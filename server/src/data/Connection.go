package data

import (
  "github.com/jinzhu/gorm"
  "github.com/gin-gonic/gin"
  eh "github.com/estebanborai/songs-share-server/lib/error_handlers"
)

func Connection(c *gin.Context) (db *gorm.DB) {
  db, err := gorm.Open("mysql", "root:root@tcp(songs-share-db)/songs-share?charset=utf8mb4&parseTime=True&loc=Local")

  if err != nil {
    eh.ResponseWithError(c, 500, "Unable to connect to the database")
  }

  return db
}