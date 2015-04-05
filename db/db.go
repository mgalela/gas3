package db

import (
  "time"

  "github.com/jinzhu/gorm"
  _"github.com/mattn/go-sqlite3"
  "golang.org/x/crypto/bcrypt"
)

var (
  DB  gorm.DB
)

type User struct {
  Username  string  `json:"username"; unique`
  Password  string  `json:"password"`
  Created  time.Time  `json:"created_at"`
}

func Init(dbname *string) error {
  var err error
  var h []byte
  var count int
  user := User{}

  DB, err = gorm.Open("sqlite3", *dbname)
  if err != nil {
    return err
  }

  DB.Where("username = ?", "admin").Find(&user).Count(&count)
  if count < 1 {
    h,err = bcrypt.GenerateFromPassword([]byte("default"), bcrypt.DefaultCost)

    DB.DropTable(&User{})
    DB.CreateTable(&User{})
    user = User{Username: "admin", Password: string(h), Created: time.Now()}
    DB.Create(&user)
  }

  return nil
}
