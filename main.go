package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

//main func
func main() {
	http.Handle("/", loggingMiddlware(http.HandlerFunc(handler)))
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "package main #14")
}

func loggingMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("uri:%s", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
