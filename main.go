package main
import (
  "github.com/jakewhite8/notecards-server/controller"
  "github.com/jakewhite8/notecards-server/database"
  "github.com/jakewhite8/notecards-server/middleware"
  "github.com/gin-gonic/gin"
)
func main() {
  database.Connect()
  database.Migrate()

  router := initRouter()
  router.Run(":8080")
}
func initRouter() *gin.Engine {
  router := gin.Default()
  router.GET("/", controller.PublicResponse)
  api := router.Group("/api")
  {
    api.POST("/token", controller.GenerateToken)
    api.POST("/user/register", controller.RegisterUser)
    secured := api.Group("/secured").Use(middleware.Auth())
    {
      // Secured area PoC
      secured.GET("/ping", controller.Ping)
      secured.POST("/create_notecard_set", controller.CreateNotecardSet)
      secured.GET("/get_notecard_sets", controller.GetNotecardSets)
      secured.GET("/get_notecards/:id", controller.GetNotecards)
      secured.DELETE("/delete_notecard_set/:id", controller.DeleteNotecardSet)

    }
  }
  return router
}