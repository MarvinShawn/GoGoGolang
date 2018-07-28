package user

import (
	"../../handle"
	"../../model"
	"../../pkg/errno"
	"github.com/gin-gonic/gin"
)

//username 关键字查询
func Get(c *gin.Context) {
	username := c.Param("username")
	//
	user, err := model.GetUser(username)

	if err != nil {

		handle.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	handle.SendResponse(c, nil, user)
}
