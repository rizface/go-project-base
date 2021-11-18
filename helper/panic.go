package helper

import (
	"base-project/app/exception"
	"strings"
)

func Panic(err error) {
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg,"23505") {
			panic(exception.Conflict{
				Msg:  strings.Split(msg,"_")[1] + " sudah digunakan",
			})
		} else if strings.Contains(msg, "required"){
			panic(exception.BadRequest{
				Msg: "Lengkapi " + strings.Split(msg,".")[1],
			})
		} else {
			panic(err)
		}
	}
}