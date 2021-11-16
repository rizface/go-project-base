package test

import (
	"base-project/app/exception"
	"base-project/app/middleware"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

func TestHTTP(t *testing.T) {
	t.Skip("this test is skipped")
	r := mux.NewRouter()

	r.Use(middleware.ErrorHandler)

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		panic(exception.BadRequest{Msg: "Bad Request"})
	})

	server := http.Server{
		Addr:              ":8080",
		Handler:           r,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
