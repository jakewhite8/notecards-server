package database
import (
  "fmt"
  "github.com/jakewhite8/notecards-server/model"
  "log"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)
var Instance *gorm.DB
var dbError error
func Connect() () {
  host := "localhost"
  user := "admin"
  password := "password"
  dbName := "go_rest_api"
  port := 5432

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port)

  Instance, dbError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if dbError != nil {
    log.Fatal(dbError)
    panic("Cannot connect to DB")
  }
  log.Println("Connected to Database!")
}
func Migrate() {
  Instance.AutoMigrate(&model.User{}, &model.NotecardSet{}, &model.UserNotecards{}, &model.Notecards{})
  log.Println("Database Migration Completed!")
}