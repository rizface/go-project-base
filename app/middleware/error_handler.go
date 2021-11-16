package middleware

import (
	"base-project/app/exception"
	"base-project/helper"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				err,errOK := err.(exception.Exception);
				if errOK {
					writer.WriteHeader(err.DefaultStatusCode())
				} else {
					writer.WriteHeader(http.StatusInternalServerError)
				}
				helper.JsonWriter(writer,err.DefaultStatusCode(),err.Error(),nil)
			}
		}()
		next.ServeHTTP(writer,request)
	})
}
