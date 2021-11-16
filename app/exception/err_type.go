package exception

type BadRequest struct {
	Code int
	Msg  string
}

type Unauthorized struct {
	Code int
	Msg string
}

type Forbidden struct {
	Code int
	Msg string
}

type NotFound struct {
	Code int
	Msg string
}

type MethodNotAllowed struct {
	Code int
	Msg string
}

type Conflict struct {
	Code int
	Msg string
}

type InternalServerError struct {
	Code int
	Msg string
}

