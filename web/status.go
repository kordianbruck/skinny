package web

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"log"
	"net/http"
)

var Box = packr.NewBox("./static")

func ServeStatusPage(w http.ResponseWriter, request *http.Request) {
	s, err := Box.FindString("index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s", s)
}
