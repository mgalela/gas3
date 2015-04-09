package main

import (
  "flag"

  "github.com/mgalela/gas3/db"
  "github.com/mgalela/gas3/web"

  "github.com/gin-gonic/gin"
  "github.com/Sirupsen/logrus"
)

var (
  DBName = flag.String("db_name", "./gas3.db", "DB_Name")

  Log *logrus.Logger
)

func InitApp(dbname *string) error {
  var err error
  Log = logrus.New()
  Log.Level = logrus.Level(logrus.DebugLevel)
  err = db.Init(dbname)
  return err
}

func main() {
  flag.Parse()

  err := InitApp(DBName)
  if err != nil {
    Log.Fatal("InitApp failed:",err)
  }
  defer  db.DB.Close()

  m := gin.Default()

  v1 := m.Group("/api/v1")
  {
    v1.GET("/device", web.Devices)
    v1.GET("/device/:id", web.Device)
    v1.POST("/device", web.DeviceNew)
    v1.PUT("/device/:id", web.DeviceUpdate)
    v1.DELETE("/device/:id", web.DeviceDel)
    v1.POST("/devreg", web.DevReg)
  }

  Log.Info("starting server")
  m.Run("0.0.0.0:8001")
}
