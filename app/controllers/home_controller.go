package controllers

import (
	"net/http"

	"github.com/unrolled/render"
<<<<<<< HEAD
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
=======

)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "wellcome to home page gotoko")
>>>>>>> d7c111a112f70f719ca912edeb2d438d4d53ae98

	render := render.New(render.Options{
		Layout: "layout",
	})

	_ = render.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"title": "Home Title",
		"body":  "Home Description",
	})

}
