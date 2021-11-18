package middleware

import (
	"base-project/app/exception"
	"base-project/helper"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			gErr := recover()
			if gErr != nil {
				err,errOK := gErr.(exception.Exception);
				if errOK {
					writer.WriteHeader(err.DefaultStatusCode())
					helper.JsonWriter(writer,err.DefaultStatusCode(),err.Error(),nil)
				} else {
					err,_ := gErr.(error)
					writer.WriteHeader(http.StatusInternalServerError)
					helper.JsonWriter(writer,http.StatusInternalServerError,err.Error(),nil)
				}
			}
		}()
		next.ServeHTTP(writer,request)
	})
}
