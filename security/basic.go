package security

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type HttpBasic struct {
}

type User struct {
	Username string
	Password string
}

func (h *HttpBasic) Authorize(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		h.Callback(c, nil, errors.New("parse authentication error"))
	} else {
		h.Callback(c, &User{
			Username: username,
			Password: password,
		}, nil)
	}

}
func (h *HttpBasic) Callback(c *gin.Context, credentials interface{}, err error) {

}
