package model

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Dwarf struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Dwarf    string `json:"dwarf" gorm:"uniquel; not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", dbHost, dbPort, username, password)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Dwarf{})
	if err != nil {
		fmt.Println(err)
	}

}
