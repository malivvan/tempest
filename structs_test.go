package tempest

import (
	"io"
	"time"
)

type ClassicNoTags struct {
	PublicField  int
	privateField string
	Date         time.Time
	InlineStruct struct {
		a float32
		B float64
	}
	Interf io.Writer
}

type ClassicBadTags struct {
	ID           string
	PublicField  int `db:"mrots"`
	privateField string
	Date         time.Time
	InlineStruct struct {
		a float32
		B float64
	}
	Interf io.Writer
}

type ClassicUnique struct {
	ID            string
	PublicField   int       `db:"unique"`
	privateField  string    `db:"unique"`
	privateField2 string    `db:"unique"`
	Date          time.Time `db:"unique"`
	InlineStruct  struct {
		A float32
		B float64
	} `db:"unique"`
	Interf io.Writer `db:"unique"`
}

type ClassicIndex struct {
	ID           string
	PublicField  int       `db:"index"`
	privateField string    `db:"index"`
	Date         time.Time `db:"index"`
	InlineStruct struct {
		a float32
		B float64
	} `db:"index"`
	InlineStructPtr *UserWithNoID `db:"index"`
	Interf          io.Writer     `db:"index"`
}

type ClassicInline struct {
	PublicField  int `db:"unique"`
	ClassicIndex `db:"inline"`
	*ToEmbed     `db:"inline"`
	Date         time.Time `db:"unique"`
}

type User struct {
	ID              int       `db:"id,increment"`
	Name            string    `db:"index"`
	Age             int       `db:"index,increment"`
	DateOfBirth     time.Time `db:"index"`
	Group           string
	unexportedField int
	Slug            string `db:"unique"`
}

type ToEmbed struct {
	ID string
}

type NestedID struct {
	ToEmbed `db:"inline"`
	Name    string
}

type SimpleUser struct {
	ID   int `db:"id"`
	Name string
	age  int
}

type UserWithNoID struct {
	Name string
}

type UserWithIDField struct {
	ID   int
	Name string
}

type UserWithUint64IDField struct {
	ID   uint64
	Name string
}

type UserWithStringIDField struct {
	ID   string
	Name string
}

type UserWithEmbeddedIDField struct {
	UserWithIDField `db:"inline"`
	Age             int
}

type UserWithEmbeddedField struct {
	UserWithNoID `db:"inline"`
	ID           uint64
}

type UserWithIncrementField struct {
	ID   int
	Name string
	Age  int `db:"unique,increment"`
}

type IndexedNameUser struct {
	ID          int    `db:"id"`
	Name        string `db:"index"`
	Score       int    `db:"index,increment"`
	age         int
	DateOfBirth time.Time `db:"index"`
	Group       string
}

type UniqueNameUser struct {
	ID   int    `db:"id"`
	Name string `db:"unique"`
	Age  int    `db:"index,increment"`
}
