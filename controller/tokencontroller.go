package controller
import (
  "github.com/jakewhite8/notecards-server/auth"
  "github.com/jakewhite8/notecards-server/database"
  "github.com/jakewhite8/notecards-server/model"
  "net/http"
  "github.com/gin-gonic/gin"
)
type TokenRequest struct {
  Email    string `json:"email"`
  Password string `json:"password"`
}
func GenerateToken(context *gin.Context) {
  var request TokenRequest
  var user model.User
  // Accept data from client
  if err := context.ShouldBindJSON(&request); err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    context.Abort()
    return
  }
  // Check if email exists and password is correct
  record := database.Instance.Where("LOWER(email) = LOWER(?)", request.Email).First(&user)
  if record.Error != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
    context.Abort()
    return
  }
  credentialError := user.CheckPassword(request.Password)
  if credentialError != nil {
    context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
    context.Abort()
    return
  }
  tokenString, err:= auth.GenerateJWT(user.Email, user.Username, user.ID)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    context.Abort()
    return
  }
  context.JSON(http.StatusOK, gin.H{"token": tokenString, "name": user.Name, "username": user.Username, "id": user.ID, "email": user.Email})
}