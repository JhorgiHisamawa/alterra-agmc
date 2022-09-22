package util

const (
	UserNotFound      = "user is not found"
	BookNotFound      = "book is not found"
	EmptyData         = "empty data"
	ErrorUnauthorized = "not authorized"
	Success           = "success"
	ErrorInput        = "error on input"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
