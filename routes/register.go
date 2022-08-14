package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type DirectoryRequestBody struct {
	Directory 	string
}

func StartRegister(group *gin.RouterGroup) {
	group.POST("/directory", RegisterDirectory);
}

func RegisterDirectory(c *gin.Context) {
	var reqBody DirectoryRequestBody
	if err := c.Copy().ShouldBindJSON(&reqBody); err != nil {
		fmt.Println(err);
	}
	fmt.Println(reqBody.Directory);

}


