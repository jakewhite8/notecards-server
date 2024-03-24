package model
import (
  "gorm.io/gorm"
)

//  Table that stores User informatino
type User struct {
  gorm.Model
  Name     string `json:"name"`
  Username string `json:"username" gorm:"unique"`
  Email    string `json:"email" gorm:"unique"`
  Password string `json:"password"`
}

// Table that stores general information about a set of Notecards
type NotecardSet struct {
  gorm.Model
  Title       string    `json:"title"`
  UserID      uint      `json:"creator"`
  Description string    `json:"description"`
}

// Table that stores the Notecards a User subscribes to
type UserNotecards struct {
  gorm.Model
  UserID            uint `json:"user_id"`
  NotecardSetID     uint `json:"notecard_set_id"`
}

// Table that stores the individual notecards for a set of Notecards
type Notecards struct {
  gorm.Model
  NotecardSetId     uint    `json:"notecard_set_id"`
  Front             string  `json:"front"`
  Back              string  `json:"back"`
}

