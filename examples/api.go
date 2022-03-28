package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/long2ice/swagin/security"
	"mime/multipart"
	"net/http"
)

type TestQuery struct {
	Name     string `query:"name" validate:"required" json:"name" description:"name of model" default:"test"`
	Token    string `header:"token" validate:"required" json:"token" default:"test"`
	Optional string `query:"optional" json:"optional"`
}

func (t *TestQuery) Handler(c *gin.Context) {
	user := c.MustGet(security.Credentials).(*security.User)
	fmt.Println(user)
	c.JSON(http.StatusOK, t)
}

type TestQueryList struct {
	Name  string `query:"name" validate:"required" json:"name" description:"name of model" default:"test"`
	Token string `header:"token" validate:"required" json:"token" default:"test"`
}

func (t *TestQueryList) Handler(c *gin.Context) {
	user := c.MustGet(security.Credentials).(*security.User)
	fmt.Println(user)
	c.JSON(http.StatusOK, []TestQueryList{*t})
}

type TestQueryPath struct {
	Name  string `query:"name" validate:"required" json:"name" description:"name of model" default:"test"`
	ID    int    `uri:"id" validate:"required" json:"id" description:"id of model" default:"1"`
	Token string `header:"token" validate:"required" json:"token" default:"test"`
}

func (t *TestQueryPath) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, t)
}

type TestForm struct {
	ID   int    `query:"id" validate:"required" json:"id" description:"id of model" default:"1"`
	Name string `form:"name" validate:"required" json:"name" description:"name of model" default:"test"`
	List []int  `form:"list" validate:"required" json:"list" description:"list of model"`
}

func (t *TestForm) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, t)
}

type TestNoModel struct {
	Authorization string `header:"authorization" validate:"required" json:"authorization" default:"authorization"`
	Token         string `header:"token" binding:"required" json:"token" default:"token"`
}

func (t *TestNoModel) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, t)
}

type TestFile struct {
	File *multipart.FileHeader `form:"file" validate:"required" description:"file upload"`
}

func (t *TestFile) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"file": t.File.Filename})
}
