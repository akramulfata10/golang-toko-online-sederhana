package controllers

import (
	"fmt"
	"net/http"

	"github.com/unrolled/render"

)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "wellcome to home page gotoko")

	render := render.New(render.Options{
		Layout: "layout",
	})

	_ = render.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"title": "Home Title",
		"body":  "Home Description",
	})

}
