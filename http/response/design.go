package response

import (
  "github.com/gin-gonic/gin"
  . "gin-sample/domain"
)

func ResponseDesignIndex(context *gin.Context, designs []*Design, err error) {
  if err != nil {
    context.JSON(400, gin.H{
      "error": err.Error(),
    })
    return
  }

  context.JSON(200, gin.H{
    "message": designs,
  })
}
