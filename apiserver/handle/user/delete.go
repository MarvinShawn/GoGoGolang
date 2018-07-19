package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"../../model"
	"../../handle"
	"../../pkg/errno"
)


// 通过id 删除用户
func Delete(c *gin.Context)  {
	userId,_ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(userId)); err != nil {

		handle.SendResponse(c,errno.ErrDatabase,nil)
		return

	}
	handle.SendResponse(c,nil,nil)

}