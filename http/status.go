package http

const (
	StatusOK         int = 200
	StatusNotFound   int = 404
	StatusNotAllowed int = 405
)

func StatusText(code int) string {
	switch code {
	case StatusOK:
		return "OK"
	case StatusNotFound:
		return "Not Found"
	case StatusNotAllowed:
		return "Method Not Allowed"
	default:
		return ""
	}
}
