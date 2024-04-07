package controller
import (
  "github.com/jakewhite8/notecards-server/database"
  "github.com/jakewhite8/notecards-server/model"
  "net/http"
  "github.com/gin-gonic/gin"
  "sync"
)

func insertNotecards(wg *sync.WaitGroup, rawNotecards [][]string, notecardSetID uint, context *gin.Context) {
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
    context.JSON(http.StatusInternalServerError, gin.H{"error": notecardsRecord.Error.Error()})
    context.Abort()
    return
  }
}

func insertUserNotecards(wg *sync.WaitGroup, notecardSetID uint, userID uint, context *gin.Context) {
  defer wg.Done()
  userNotecards := model.UserNotecards{
    UserID: userID,
    NotecardSetID: notecardSetID,
  }
  userNotecardsRecord := database.Instance.Create(&userNotecards)
  if userNotecardsRecord.Error != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": userNotecardsRecord.Error.Error()})
    context.Abort()
    return
  }
}


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
  // Goroutine
  go insertNotecards(wg, newNotecardSet.Notecards, notecardSet.ID, context)
  go insertUserNotecards(wg, notecardSet.ID, userID, context)

  wg.Wait()

  // Send client success response
  context.JSON(http.StatusCreated, gin.H{"success": true})
}