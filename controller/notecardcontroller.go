package controller
import (
  "github.com/jakewhite8/notecards-server/database"
  "github.com/jakewhite8/notecards-server/model"
  "net/http"
  "github.com/gin-gonic/gin"
  "sync"
  "fmt"
)

func insertNotecards(wg *sync.WaitGroup, rawNotecards [][]string, notecardSetID uint, context *gin.Context, errorChan chan string) {
  defer wg.Done()
  var notecards []model.Notecards
  for i := range rawNotecards {
    notecard := model.Notecards {
      NotecardSetID: notecardSetID,
      Front: rawNotecards[i][0],
      Back: rawNotecards[i][1],
    }
    notecards = append(notecards, notecard)
  }

  notecardsRecord := database.Instance.Create(notecards)
  if notecardsRecord.Error != nil {
    errorChan <- notecardsRecord.Error.Error()
  }
}

func insertUserNotecards(wg *sync.WaitGroup, notecardSetID uint, userID uint, context *gin.Context, errorChan chan string) {
  defer wg.Done()
  userNotecards := model.UserNotecards{
    UserID: userID,
    NotecardSetID: notecardSetID,
  }
  userNotecardsRecord := database.Instance.Create(&userNotecards)
  if userNotecardsRecord.Error != nil {
    errorChan <- userNotecardsRecord.Error.Error()
  }
}

// Accepts a Title and an array of Notecards (with a front and back) to create a Set of Notecards
func CreateNotecardSet(context *gin.Context) {
  userID := context.MustGet("user_id").(uint)

  type NewNotecardSet struct {
    Title   string
    Notecards [][]string
  }

  var newNotecardSet NewNotecardSet
  if err := context.ShouldBindJSON(&newNotecardSet); err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    context.Abort()
    return
  }

  // Create NotecardSet
  var notecardSet model.NotecardSet
  notecardSet.Title = newNotecardSet.Title
  // notecardSet.Description = "Notecard Description"
  notecardSet.UserID = userID
  notecardSetRecord := database.Instance.Create(&notecardSet)
  if notecardSetRecord.Error != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": notecardSetRecord.Error.Error()})
    context.Abort()
    return
  }

  wg := new(sync.WaitGroup)
  wg.Add(2)
  errorChan := make(chan string)
  // Goroutine
  go insertNotecards(wg, newNotecardSet.Notecards, notecardSet.ID, context, errorChan)
  go insertUserNotecards(wg, notecardSet.ID, userID, context, errorChan)

  wg.Wait()
  close(errorChan)

  success := true
  for err := range errorChan {
    success = false
    fmt.Println("Error in CreateNotecardSet:", err)
  }
  if success {
    context.JSON(http.StatusCreated, gin.H{"success": true})
  } else {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating notecards"})
  }
}

// Returns NotecardSet information (Title, ID, Creator (id), Description) for all Notecard Sets
// that belong to the logged in User
func GetNotecardSets(context *gin.Context) {
  userID := context.MustGet("user_id").(uint)

  var notecardSets []model.NotecardSet
  getNotecardSets := database.Instance.Where("user_id = ?", userID).Find(&notecardSets)
  if getNotecardSets.Error != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": getNotecardSets.Error.Error()})
    context.Abort()
    return
  }
  context.JSON(http.StatusOK, gin.H{"notecardSets": notecardSets}) 
}

// Get all the Notecards that belong to the requested Notecard Set
func GetNotecards(context *gin.Context) {
  notecardSetID := context.Param("id")
  var notecards []model.Notecards
  getNotecards := database.Instance.Where("notecard_set_id = ?", notecardSetID).Find(&notecards)
  if getNotecards.Error != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": getNotecards.Error.Error()})
    context.Abort()
    return
  }
  context.JSON(http.StatusOK, gin.H{"notecards": notecards})
}