package controller
import (
  "net/http"
  "github.com/gin-gonic/gin"
  "fmt"
)
func CreateNotecardSet(context *gin.Context) {
  fmt.Println("Create Notecard")

  type NewNotecardSet struct {
    Title   string
    Notecards [][]string
  }

  var notecardSet NewNotecardSet
  if err := context.ShouldBindJSON(&notecardSet); err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    context.Abort()
    return
  }
  fmt.Printf("Create notecard set: %v", notecardSet)
  context.JSON(http.StatusCreated, gin.H{"success": true})
}