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
	PublicField  int `tempest:"mrots"`
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
	PublicField   int       `tempest:"unique"`
	privateField  string    `tempest:"unique"`
	privateField2 string    `tempest:"unique"`
	Date          time.Time `tempest:"unique"`
	InlineStruct  struct {
		A float32
		B float64
	} `tempest:"unique"`
	Interf io.Writer `tempest:"unique"`
}

type ClassicIndex struct {
	ID           string
	PublicField  int       `tempest:"index"`
	privateField string    `tempest:"index"`
	Date         time.Time `tempest:"index"`
	InlineStruct struct {
		a float32
		B float64
	} `tempest:"index"`
	InlineStructPtr *UserWithNoID `tempest:"index"`
	Interf          io.Writer     `tempest:"index"`
}

type ClassicInline struct {
	PublicField  int `tempest:"unique"`
	ClassicIndex `tempest:"inline"`
	*ToEmbed     `tempest:"inline"`
	Date         time.Time `tempest:"unique"`
}

type User struct {
	ID              int       `tempest:"id,increment"`
	Name            string    `tempest:"index"`
	Age             int       `tempest:"index,increment"`
	DateOfBirth     time.Time `tempest:"index"`
	Group           string
	unexportedField int
	Slug            string `tempest:"unique"`
}

type ToEmbed struct {
	ID string
}

type NestedID struct {
	ToEmbed `tempest:"inline"`
	Name    string
}

type SimpleUser struct {
	ID   int `tempest:"id"`
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
	UserWithIDField `tempest:"inline"`
	Age             int
}

type UserWithEmbeddedField struct {
	UserWithNoID `tempest:"inline"`
	ID           uint64
}

type UserWithIncrementField struct {
	ID   int
	Name string
	Age  int `tempest:"unique,increment"`
}

type IndexedNameUser struct {
	ID          int    `tempest:"id"`
	Name        string `tempest:"index"`
	Score       int    `tempest:"index,increment"`
	age         int
	DateOfBirth time.Time `tempest:"index"`
	Group       string
}

type UniqueNameUser struct {
	ID   int    `tempest:"id"`
	Name string `tempest:"unique"`
	Age  int    `tempest:"index,increment"`
}
