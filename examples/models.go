package main

type QueryModel struct {
	Name  string `query:"name" binding:"required" json:"name" description:"name of model" default:"test"`
	Token string `header:"token" binding:"required" json:"token" default:"test"`
}

type QueryPathModel struct {
	Name  string `query:"name" binding:"required" json:"name" description:"name of model" default:"test"`
	ID    int    `uri:"id" binding:"required" json:"id" description:"id of model"`
	Token string `header:"token" binding:"required" json:"token" default:"test"`
}

type FormModel struct {
	Name string `form:"name" binding:"required" json:"name" description:"name of model" default:"test"`
}

type BodyModel struct {
	Name string `body:"name" binding:"required" json:"name" description:"name of model" default:"test"`
}
