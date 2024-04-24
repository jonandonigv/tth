package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jonandonigv/tth/view"
	"github.com/jonandonigv/tth/view/layout"
	"github.com/jonandonigv/tth/view/partial"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	c := layout.Base(view.Index())
	http.Handle("/", templ.Handler(c))

	http.Handle("/foo", templ.Handler(partial.Foo()))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
