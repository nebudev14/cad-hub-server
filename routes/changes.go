package routes

import (
	"github.com/gin-gonic/gin"
)

func StartChanges(group *gin.RouterGroup) {
	group.GET("", GetAllChanges)
	group.GET("/check", CheckChange)
}

func GetAllChanges(c *gin.Context) {

}

func CheckChange(c *gin.Context) {

}
