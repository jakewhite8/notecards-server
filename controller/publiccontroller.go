package controller
import (
  "net/http"
  "github.com/gin-gonic/gin"
)
func PublicResponse(context *gin.Context) {
  context.JSON(http.StatusOK, gin.H{"message": "Notecards Server"})
}