package controllers

import (
	"github.com/andrqxa/goschool2/views"
)

type Static struct {
	Header  *views.View
	Footer  *views.View
	Welcome *views.View
}

func NewStatic() *Static {
	return &Static{
		Header:  views.NewView("main", "static/header"),
		Footer:  views.NewView("main", "static/footer"),
		Welcome: views.NewView("main", "static/welcome"),
	}
}
