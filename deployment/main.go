package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/codegangsta/negroni"
)

func main() {
	// Middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleware),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	n.Run(":8080")
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way there...")

	var assetUrl = regexp.MustCompile(`^/(css|fonts|js)/`)

	if assetUrl.MatchString(r.RequestURI) {
		next(rw, r)
	} else if r.URL.Query().Get("password") == "secret123" {
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorized", 401)
	}

	log.Println("Logging on the way back...")
}
