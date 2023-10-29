package Handlers

import (
	"Test/models"
	"Test/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandlers struct {
	Usecase usecase.UseCase
}

func NewTestHandlers() *TestHandlers {
	return &TestHandlers{}
}

func (th *TestHandlers) Sample(c *gin.Context) {
	var Name models.DataModel
	if err := c.ShouldBind(&Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Can't bind the name": err.Error()})
		return
	}
	savedName, err := th.Usecase.ExecuteAddname(Name.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Name Saved": savedName})
}

func (th *TestHandlers) ShowName(c *gin.Context) {
	Name, err := th.Usecase.ExecuteShowName()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Name Is": Name})
}
