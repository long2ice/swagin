package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/long2ice/swagin/security"
	"mime/multipart"
	"net/http"
)

type TestQueryReq struct {
	unexported string
	Name       string `query:"name" validate:"required" json:"name" description:"name of model" default:"test"`
	Token      string `header:"token" validate:"required" json:"token" default:"test"`
	Optional   string `query:"optional" json:"optional"`
}

func TestQuery(c *gin.Context, req TestQueryReq) {
	user := c.MustGet(security.Credentials).(*security.User)
	fmt.Println(user)
	c.JSON(http.StatusOK, req)
}

type TestQueryListReq struct {
	Name  string `query:"name" validate:"required" json:"name" description:"name of model" default:"test"`
	Token string `header:"token" validate:"required" json:"token" default:"test"`
}

func TestQueryList(c *gin.Context, req TestQueryListReq) {
	user := c.MustGet(security.Credentials).(*security.User)
	fmt.Println(user)
	c.JSON(http.StatusOK, []TestQueryListReq{req})
}

type TestQueryPathReq struct {
	Name  string `query:"name" validate:"required" json:"name" description:"name of model" default:"test"`
	ID    int    `uri:"id" validate:"required" json:"id" description:"id of model" default:"1"`
	Token string `header:"token" validate:"required" json:"token" default:"test"`
}

func TestQueryPath(c *gin.Context, req TestQueryPathReq) {
	c.JSON(http.StatusOK, req)
}

type TestFormReq struct {
	ID   int               `query:"id" validate:"required" json:"id" description:"id of model" default:"1"`
	Name string            `form:"name" validate:"required" json:"name" description:"name of model" default:"test"`
	List []int             `form:"list" validate:"required" json:"list" description:"list of model"`
	Map  map[string]string `form:"map" validate:"required" json:"map" description:"a map"`
}

func TestForm(c *gin.Context, req TestFormReq) {
	c.JSON(http.StatusOK, req)
}

type TestNoModelReq struct {
	Authorization string `header:"authorization" validate:"required" json:"authorization" default:"authorization"`
	Token         string `header:"token" binding:"required" json:"token" default:"token"`
}

func TestNoModel(c *gin.Context, req TestNoModelReq) {
	c.JSON(http.StatusOK, req)
}

type TestFileReq struct {
	File *multipart.FileHeader `form:"file" validate:"required" description:"file upload"`
}

func TestFile(c *gin.Context, req TestFileReq) {
	c.JSON(http.StatusOK, gin.H{"file": req.File.Filename})
}
