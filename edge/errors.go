package edge

// CodeError is a error with a code (like 404 Not Found).
type CodeError struct {
	Code int
	Text string
}

func (e CodeError) Error() string {
	return e.Text
}
