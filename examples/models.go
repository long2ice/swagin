package main

import "mime/multipart"

type QueryModel struct {
	Name  string `query:"name" binding:"required" json:"name" description:"name of model" default:"test"`
	Token string `header:"token" binding:"required" json:"token" default:"test"`
}

type QueryPathModel struct {
	Name  string `query:"name" binding:"required" json:"name" description:"name of model" default:"test"`
	ID    int    `uri:"id" binding:"required" json:"id" description:"id of model" default:"1"`
	Token string `header:"token" binding:"required" json:"token" default:"test"`
}

type FormModel struct {
	ID   int    `query:"id" binding:"required" json:"id" description:"id of model" default:"1"`
	Name string `form:"name" binding:"required" json:"name" description:"name of model" default:"test"`
	List []int  `form:"list" binding:"required" json:"list" description:"list of model" default:"[1]"`
}

type FormFileModel struct {
	File *multipart.FileHeader `form:"file" binding:"required" description:"file upload"`
}

type QueryModelList []*QueryModel
