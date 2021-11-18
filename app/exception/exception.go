package exception

import "net/http"

func(b BadRequest) DefaultStatusCode() int {
	return http.StatusBadRequest
}

func(b BadRequest) Error() string {
	return b.Msg
}

func(u Unauthorized) DefaultStatusCode() int {
	return http.StatusUnauthorized
}

func(u Unauthorized) Error() string {
	return u.Msg
}

func(f Forbidden) DefaultStatusCode() int {
	return http.StatusForbidden
}

func(f Forbidden) Error() string {
	return f.Msg
}

func(n NotFound) DefaultStatusCode() int {
	return http.StatusNotFound
}

func(n NotFound) Error() string {
	return n.Msg
}

func(m MethodNotAllowed) DefaultStatusCode() int {
	return http.StatusMethodNotAllowed
}

func(m MethodNotAllowed) Error() string {
	return m.Msg
}

func(c Conflict) DefaultStatusCode() int {
	return http.StatusConflict
}

func(c Conflict) Error() string {
	return c.Msg
}

func(i InternalServerError) DefaultStatusCode() int {
	return http.StatusInternalServerError
}

func(i InternalServerError) Error() string {
	return i.Msg
}



