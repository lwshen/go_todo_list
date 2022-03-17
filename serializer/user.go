package serializer

import "todo_list/model"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`                  //用户ID
	Username string `json:"username" form:"username" example:"FanOne"` //用户名
	Status   string `json:"status" form:"status"`                      //状态
	CreateAt int64  `json:"create_at" form:"create_at"`                //创建时间
}

func BuildUser(user *model.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		CreateAt: user.CreatedAt.Unix(),
	}
}
