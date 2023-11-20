package upload 

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
  file, err := c.FormFile("file")  
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })
    return
  }

  filename := "uploads/" + file.Filename
  err = c.SaveUploadedFile(file,filename)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "File uploaded successfully",
    "filename": filename,
  })
}
