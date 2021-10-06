package controllers

import (
  "github.com/gin-gonic/gin"
	"net/http"
  "restAPI/models"
)

type CreateTrackInput struct {
	Artist string `json:"artist" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateTrackInput struct{
  Artist string `json:"artist"`
  Title string `json:"title"`
}

func GetAllTracks(c *gin.Context){
  var tracks []models.Track
  models.ConnectDB().Find(&tracks)

  c.JSON(http.StatusOK, gin.H{"tracks": tracks})
}

func CreateTrack(c *gin.Context){
  var input CreateTrackInput
  if err := c.ShouldBindJSON(&input); err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  track := models.Track{Artist: input.Artist, Title: input.Title}
  models.ConnectDB().Create(&track)
  c.JSON(http.StatusOK, gin.H{"tracks": track})
}

func GetTrack(c *gin.Context){
  var track models.Track
  if err := models.ConnectDB().Where("id = ?", c.Param("id")).First(&track).Error; err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"tracks": track})
 }


func UpdateTrack(c *gin.Context){
  var track models.Track
  if err := models.ConnectDB().Where("id = ?", c.Param("id")).First(&track).Error; err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
    return
  }

  var input UpdateTrackInput
  if err := c.ShouldBindJSON(&input); err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.ConnectDB().Model(&track).Update(input)
  c.JSON(http.StatusOK, gin.H{"tracks": track})
}

func DeleteTrack(c *gin.Context){
  var track models.Track
  if err := models.ConnectDB().Where("id = ?", c.Param("id")).First(&track).Error; err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error": "Запись удалена"})
    return
  }
  models.ConnectDB().Delete(track)
  c.JSON(http.StatusOK, gin.H{"tracks": true})
}
