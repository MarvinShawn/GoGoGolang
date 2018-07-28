package user

import (
	"../../handle"
	"../../model"
	"../../pkg/errno"
	"../../util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {

	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		handle.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// 校验数据
	if err := u.Validate(); err != nil {
		handle.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//加密用户密码
	if err := u.Encrypt(); err != nil {

		handle.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	//插进数据库
	if err := u.Create(); err != nil {

		handle.SendResponse(c, errno.ErrDatabase, nil)
		return

	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	handle.SendResponse(c, nil, rsp)

}
