package controllers

import (
	"goschool/views"
)

type Static struct {
	Header  *views.View
	Footer  *views.View
	Welcome *views.View
}

func NewStatic() *Static {
	return &Static{
		Welcome: views.NewView("main", "static/welcome"),
	}
}
