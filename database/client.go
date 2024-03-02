package database
import (
  "fmt"
  "notecards-api/model"
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

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)

  Instance, dbError = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if dbError != nil {
    log.Fatal(dbError)
    panic("Cannot connect to DB")
  }
  log.Println("Connected to Database!")
}
func Migrate() {
  Instance.AutoMigrate(&model.User{})
  log.Println("Database Migration Completed!")
}