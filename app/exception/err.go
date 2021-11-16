package exception

type Exception interface {
	DefaultStatusCode() int
	Error() string
}
