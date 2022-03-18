package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"todo_list/pkg/utils"
	"todo_list/service"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateUpdateTaskService
	cliams, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(cliams.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	cliams, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(cliams.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	cliams, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(cliams.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func UpdateTask(c *gin.Context) {
	var updateTask service.CreateUpdateTaskService
	cliams, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateTask); err == nil {
		res := updateTask.Update(cliams.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(claims.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(claims.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}
