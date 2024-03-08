package controller
import (
  "notecards-api/auth"
  "notecards-api/database"
  "notecards-api/model"
  "net/http"
  "github.com/gin-gonic/gin"
)
func RegisterUser(context *gin.Context) {
  var user model.User
  // Store data from client in user variable (of type User)
  if err := context.ShouldBindJSON(&user); err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    context.Abort()
    return
  }
  // Set user.Password (in hashed format)
  if err := user.HashPassword(user.Password); err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    context.Abort()
    return
  }
  // Store User data in database
  record := database.Instance.Create(&user)
  if record.Error != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
    context.Abort()
    return
  }
  tokenString, err:= auth.GenerateJWT(user.Email, user.Username)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    context.Abort()
    return
  }
  // Send client success response
  context.JSON(http.StatusCreated, gin.H{"id": user.ID, "email": user.Email, "username": user.Username, "name": user.Name, "token": tokenString})
}