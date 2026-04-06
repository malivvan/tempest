package q_test

import (
	"fmt"
	"log"

	"time"

	"os"

	"path/filepath"
	"strings"

	"github.com/malivvan/tempest"
	"github.com/malivvan/tempest/q"
)

func ExampleRe() {
	dir, db := prepareDB()
	defer os.RemoveAll(dir)
	defer db.Close()

	var users []User

	// Find all users with name that starts with the letter D.
	if err := db.Select(q.Re("Name", "^D")).Find(&users); err != nil {
		log.Println("error: Select failed:", err)
		return
	}

	// Donald and Dilbert
	fmt.Println("Found", len(users), "users.")

	// Output:
	// Found 2 users.
}

type User struct {
	ID        int    `tempest:"id,increment"`
	Group     string `tempest:"index"`
	Email     string `tempest:"unique"`
	Name      string
	Age       int       `tempest:"index"`
	CreatedAt time.Time `tempest:"index"`
}

func prepareDB() (string, *tempest.DB) {
	dir, _ := os.MkdirTemp(os.TempDir(), "tempest")
	db, _ := tempest.Open(filepath.Join(dir, "tempest.db"))

	for i, name := range []string{"John", "Norm", "Donald", "Eric", "Dilbert"} {
		email := strings.ToLower(name + "@provider.com")
		user := User{
			Group:     "staff",
			Email:     email,
			Name:      name,
			Age:       21 + i,
			CreatedAt: time.Now(),
		}
		err := db.Save(&user)

		if err != nil {
			log.Fatal(err)
		}
	}

	return dir, db
}
