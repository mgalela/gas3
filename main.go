package main

import (
  "flag"

  "github.com/mgalela/gas3/db"

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

  slash := m.Group("/")
  slash.Static("/static","static")
  m.LoadHTMLFiles("templates/index.tmpl")

  m.GET("/", Home)

  v1 := m.Group("/api/v1")
  {
    v1.GET("/user", UserApi)
  }

  Log.Info("starting server")
  m.Run(":8001")
}

func Home(c *gin.Context) {
  c.HTML(200, "index.tmpl", nil)
}

func UserApi(c *gin.Context){
  c.String(200, "User page")
}
