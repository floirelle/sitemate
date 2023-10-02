package controller

import (
	"fmt"
	"server/custom_errors"
	"server/models"
	"server/response"
	"server/service"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetIssues(ctx *gin.Context)
	GetIssue(ctx *gin.Context)
	AddIssue(ctx *gin.Context)
	UpdateIssue(ctx *gin.Context)
	DeleteIssue(ctx *gin.Context)
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) GetIssues(ctx *gin.Context) {
	result, _ := c.service.GetIssues()

	ctx.JSON(200, response.Json(200, result))
}

func (c *controller) GetIssue(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.service.GetIssue(id)

	if err != nil {
		tempError, _ := err.(*custom_errors.BaseError)
		ctx.JSON(tempError.StatusCode, tempError)
		return
	}

	ctx.JSON(200, response.Json(200, result))
}

func (c *controller) AddIssue(ctx *gin.Context) {
	issue := &models.Issue{}
	err := ctx.BindJSON(&issue)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, custom_errors.BadRequestError())
		return
	}

	result, err := c.service.AddIssue(*issue)

	ctx.JSON(201, response.Json(201, result))
}

func (c *controller) UpdateIssue(ctx *gin.Context) {
	issue := &models.Issue{}
	err := ctx.BindJSON(&issue)
	if err != nil {
		ctx.JSON(400, custom_errors.BadRequestError())
		return
	}

	id := ctx.Param("id")

	result, err := c.service.UpdateIssue(id, *issue)
	if err != nil {
		tempError, _ := err.(*custom_errors.BaseError)

		ctx.JSON(tempError.StatusCode, tempError)
		return
	}

	ctx.JSON(200, response.Json(200, result))
}

func (c *controller) DeleteIssue(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := c.service.DeleteIssue(id)
	if err != nil {
		tempError, _ := err.(*custom_errors.BaseError)
		ctx.JSON(tempError.StatusCode, tempError)
		return
	}

	ctx.JSON(200, response.Json(200, result))
}
