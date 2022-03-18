package service

import (
	"time"
	"todo_list/model"
	"todo_list/pkg/e"
	"todo_list/serializer"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做 1是已做
}

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := e.SUCCESS
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "创建成功",
	}
}

type ShowTaskService struct {
}

func (service *ShowTaskService) Show(uid uint, tid string) serializer.Response {
	code := e.SUCCESS
	var user model.User
	model.DB.First(&user, uid)
	var task model.Task
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	//todo: 判断备忘录是否属于当前用户
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
	}
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	var count uint
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), count)
}
