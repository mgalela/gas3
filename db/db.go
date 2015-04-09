package db

import (
  "time"

  "github.com/jinzhu/gorm"
  _"github.com/mattn/go-sqlite3"
//  "golang.org/x/crypto/bcrypt"
)

var (
  DB  gorm.DB
)

type User struct {
  Id	int   `json:"id"`
  Username  string  `json:"username"; unique`
  Password  string  `json:"password"`
  Created  time.Time  `json:"created_at"`
}

type Device struct {
  Id	int   `json:"id"`
  Wifimode  string  `json:"wifimode"`
  Wifissid  string  `json:"wifissid"`
  Wifipwd   string  `json:"wifipwd"`
  Wifiip    string  `json:"wifiip"`
  Wifinetmask   string  `json:"wifinetmask"`
  Wifigateway	string	`json:"gateway"`
  Device  string  `json:"device"`
  Land	  string  `json:"land"`
  Serial  string  `json:"serial"`
  Mac   string  `json:"mac"`
  Ip   string  `json:"ip"`
  Landhq   string  `json:"landhq"`
  Swstate   string  `json:"swstate"`
  Created   time.Time	`json:"created_at"`
}
/*
type Switch struct {
  Id	int   `json:"id"`
  DeviceId  int
  Swstate   string  `json:"swstate"`
}
*/
func Init(dbname *string) error {
  var err error
//  var h []byte
//  var count int
//  user := User{}

  DB, err = gorm.Open("sqlite3", *dbname)
  if err != nil {
    return err
  }

  DB.AutoMigrate(&User{}, &Device{})
/*
  DB.Where("username = ?", "admin").Find(&user).Count(&count)
  if count < 1 {
    h,err = bcrypt.GenerateFromPassword([]byte("default"), bcrypt.DefaultCost)

    DB.DropTable(&User{})
    DB.CreateTable(&User{})
    user = User{Username: "admin", Password: string(h), Created: time.Now()}
    DB.Create(&user)
  }
*/
  return nil
}
