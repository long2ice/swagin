package main

type QueryModel struct {
	Name string `query:"name" binding:"required" json:"name" description:"name of model" default:"test"`
}

type FormModel struct {
	Name string `form:"name" binding:"required" json:"name" description:"name of model" default:"test"`
}
