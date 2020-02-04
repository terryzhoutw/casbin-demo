package main

import (
	"casbindemo/middlerware"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func main() {

	e, _ := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")

	router := gin.New()
	router.Use(middlerware.NewAuthorizer(e))

	router.GET("/user/getuserbyid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getuserbyid",
		})
	})

	router.POST("/unit/getunitbyid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getunitbyid",
		})
	})

	router.Run()
}
