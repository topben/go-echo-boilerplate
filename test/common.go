package test

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	BlogModels "github.com/topben/go-echo-boilerplate/blogs/models"
	"github.com/topben/go-echo-boilerplate/database"
	UserModels "github.com/topben/go-echo-boilerplate/users/models"
)

func LoadTestEnv() error {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/topben/go-echo-boilerplate/test.env"))
	if err != nil {
		log.Fatal("failed to load test env config: ", err)
	}
	return err
}

func InitTest() {
	err := LoadTestEnv()
	db := database.GetInstance()
	db.DropTable("migrations")
	db.DropTableIfExists(&UserModels.User{})
	db.DropTableIfExists(&BlogModels.Blog{})
	m := database.GetMigrations(db)
	err = m.Migrate()
	if err != nil {
		log.Fatal("failed to run db migration: ", err)
	}
}
