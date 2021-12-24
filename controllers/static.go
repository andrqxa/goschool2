package controllers

import (
	"goschool/views"
)

type Static struct {
	Welcome *views.View
}

func NewStatic() *Static {
	return &Static{
		Welcome: views.NewView("main", "static/welcome"),
	}
}
