package controllers

import (
	"github.com/cloudimpl/next-coder-sdk/apicontext"
	"github.com/cloudimpl/next-coder-sdk/polycode"
	"github.com/gin-gonic/gin"
	"portal/register/model"
)

func Greeting(c *gin.Context) {
	var req model.HelloRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	apiCtx, err := apicontext.FromContext(c.Request.Context())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	greetingService := apiCtx.Service("greeting-service").Get()
	res, err := greetingService.RequestReply(polycode.TaskOptions{}, "Greeting", req).GetAny()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}
