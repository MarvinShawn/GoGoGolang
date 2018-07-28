package user

import (
	"../../handle"
	"../../model"
	"../../pkg/errno"
	"../../util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
)

func Update(c *gin.Context) {

	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var u model.UserModel
	u.Id = uint64(userId)
	if err := c.Bind(&u); err != nil {
		handle.SendResponse(c, errno.ErrBind, nil)
		return
	}

	var queryUser model.UserModel
	if err := model.GetUserById(u.Id, &queryUser); err != nil {
		handle.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	u.CreatedAt = queryUser.CreatedAt

	// Validate the data.
	if err := u.Validate(); err != nil {
		handle.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		handle.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		log.Errorf(err, "Get an error")
		handle.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handle.SendResponse(c, nil, nil)

}
