package tempest

import (
	"bytes"
	"encoding/binary"
	"time"

	"github.com/malivvan/tempest/codec"
	"github.com/malivvan/tempest/codec/json"
	bolt "go.etcd.io/bbolt"
)

var _metaBucket = []byte{0x00}

// Defaults to json
var defaultCodec = json.Codec

// Open opens a database at the given path with optional Tempest options.
func Open(path string, tempestOptions ...func(*Options) error) (*DB, error) {
	var err error

	var opts Options
	for _, option := range tempestOptions {
		if err = option(&opts); err != nil {
			return nil, err
		}
	}

	s := DB{
		bolt: opts.bolt,
	}

	n := node{
		s:          &s,
		codec:      opts.codec,
		batchMode:  opts.batchMode,
		rootBucket: opts.rootBucket,
	}

	if n.codec == nil {
		n.codec = defaultCodec
	}

	if opts.boltMode == 0 {
		opts.boltMode = 0600
	}

	if opts.boltOptions == nil {
		opts.boltOptions = &bolt.Options{Timeout: 1 * time.Second}
	}

	s.Node = &n

	// skip if UseDB option is used
	if s.bolt == nil {
		s.bolt, err = bolt.Open(path, opts.boltMode, opts.boltOptions)
		if err != nil {
			return nil, err
		}
	}

	err = s.checkVersion()
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// DB is the wrapper around BoltDB. It contains an instance of BoltDB and uses it to perform all the
// needed operations
type DB struct {
	Node
	bolt *bolt.DB
}

// Bolt returns the underlying BoltDB instance.
func (s *DB) Bolt() *bolt.DB {
	return s.bolt
}

// Close the database
func (s *DB) Close() error {
	return s.Bolt().Close()
}

func (s *DB) checkVersion() error {
	var v string
	b, err := s.getMeta("version")
	if err != nil && err != ErrNotFound {
		return err
	}
	v = string(b)

	// for now, we only set the current version if it doesn't exist.
	// v1 and v2 database files are compatible.
	if v == "" {
		return s.setMeta("version", []byte(Version))
	}

	return nil
}

func (s *DB) getMeta(key string) (val []byte, err error) {
	err = s.Bolt().View(func(tx *bolt.Tx) error {
		b := tx.Bucket(_metaBucket)
		if b == nil {
			return ErrNotFound
		}
		v := b.Get([]byte(key))
		if v == nil {
			return ErrNotFound
		}
		val = make([]byte, len(v))
		copy(val, v)
		return nil
	})
	return
}

func (s *DB) setMeta(key string, val []byte) error {
	return s.Bolt().Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(_metaBucket)
		if err != nil {
			return err
		}
		return b.Put([]byte(key), val)
	})
}

// toBytes turns an interface into a slice of bytes
func toBytes(key interface{}, codec codec.MarshalUnmarshaler) ([]byte, error) {
	if key == nil {
		return nil, nil
	}
	switch t := key.(type) {
	case []byte:
		return t, nil
	case string:
		return []byte(t), nil
	case int:
		return numbertob(int64(t))
	case uint:
		return numbertob(uint64(t))
	case int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		return numbertob(t)
	default:
		return codec.Marshal(key)
	}
}

func numbertob(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func numberfromb(raw []byte) (int64, error) {
	r := bytes.NewReader(raw)
	var to int64
	err := binary.Read(r, binary.BigEndian, &to)
	if err != nil {
		return 0, err
	}
	return to, nil
}
