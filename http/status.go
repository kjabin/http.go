package http

const (
	StatusOK                  int = 200
	StatusCreated             int = 201
	StatusNotFound            int = 404
	StatusNotAllowed          int = 405
	StatusBadRequest          int = 400
	StatusInternalServerError int = 501
)

func StatusText(code int) string {
	switch code {
	case StatusOK:
		return "OK"
	case StatusCreated:
		return "Created"
	case StatusNotFound:
		return "Not Found"
	case StatusNotAllowed:
		return "Method Not Allowed"
	case StatusBadRequest:
		return "Bad Request"
	case StatusInternalServerError:
		return "Internal Server Error"
	default:
		return ""
	}
}
