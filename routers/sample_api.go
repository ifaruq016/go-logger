package routers

import (
	// "fmt"
	// "encoding/json"
	// "io/ioutil"
	// "bytes"
	"github.com/gin-gonic/gin"

	"logger/app/services"
	"logger/app/utils"
)

// SetupRoutes forwards QR routes to the service layer
func SetupRoutes(router *gin.Engine) {
	sample := router.Group("/sample")
	mobileGroup := sample.Group("/mobile")

	// Sample Request route
	mobileGroup.POST("/request", func(ctx *gin.Context) {
		apiIdentifier := utils.ApiCallGenerator()
		
		utils.RequestStore(ctx, apiIdentifier)
		go utils.LogInfo(ctx, apiIdentifier, utils.DataInfo{
			State 	: "REQUEST_START",
		})

		status_code, response, _ := services.RequestServices(ctx, apiIdentifier)
		go utils.LogInfo(ctx, apiIdentifier, utils.DataInfo{
			State	: "RESPONSE_END",
			Resp	: response,
		})

		ctx.JSON(status_code, response)
	})
}
