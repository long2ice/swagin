package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/long2ice/fastgo/security"
	"net/http"
)

type TestQuery struct {
	Model *QueryModel
}

func (t *TestQuery) NewModel() interface{} {
	return &QueryModel{}
}

func (t *TestQuery) Handler(c *gin.Context) {
	user := c.MustGet(security.Credentials).(*security.User)
	fmt.Println(user)
	c.JSON(http.StatusOK, t.Model)
}

type TestQueryPath struct {
	Model *QueryPathModel
}

func (t *TestQueryPath) NewModel() interface{} {
	return &QueryPathModel{}
}

func (t *TestQueryPath) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, t.Model)
}

type TestForm struct {
	Model *FormModel
}

func (t *TestForm) NewModel() interface{} {
	return &FormModel{}
}

func (t *TestForm) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, t.Model)
}

type TestNoModel struct {
}

func (t *TestNoModel) NewModel() interface{} {
	return nil
}

func (t *TestNoModel) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

type TestFile struct {
	Model *FormFileModel
}

func (t *TestFile) NewModel() interface{} {
	return &FormFileModel{}
}

func (t *TestFile) Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"file": t.Model.File.Filename})
}
