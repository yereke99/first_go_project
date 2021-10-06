package main

import (
  "github.com/gin-gonic/gin"
	//"net/http"
  "restAPI/models"
  "restAPI/controllers"
)

func main(){
  r := gin.Default()

  models.ConnectDB()

  r.GET("/tracks", controllers.GetAllTracks)
  r.POST("/tracks", controllers.CreateTrack)
  r.GET("/tracks/:id", controllers.GetTrack)
  r.PATCH("/tracks/:id", controllers.UpdateTrack)
  r.DELETE("/tracks/:id",controllers.DeleteTrack)
  r.Run()
}
