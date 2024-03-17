package main
import (
  "notecards-api/controller"
  "notecards-api/database"
  "notecards-api/middleware"
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
    }
  }
  return router
}