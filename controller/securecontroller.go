package controller
import (
  "net/http"
  "github.com/gin-gonic/gin"
)
func Ping(context *gin.Context) {
  userID := context.MustGet("user_id").(uint)
  context.JSON(http.StatusOK, gin.H{"user_id": userID})
}