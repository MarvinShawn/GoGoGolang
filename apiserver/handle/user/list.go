package user

import (
	"../../handle"
	"../../pkg/errno"
	"../../service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {

	var r ListRequest

	if err := c.Bind(&r); err != nil {
		handle.SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)

	if err != nil {
		handle.SendResponse(c, err, nil)
		return
	}

	handle.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})

}
