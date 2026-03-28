package utils

// Cấu trúc Enum cho Error Codes
type ErrorCode int

const (
	SUCCESS               ErrorCode = 200
	VALIDATION_ERROR      ErrorCode = 400
	UNAUTHORIZED          ErrorCode = 401
	NOT_FOUND             ErrorCode = 404
	INTERNAL_SERVER_ERROR ErrorCode = 500
)

// String() function to convert ErrorCode to string if needed
func (e ErrorCode) String() string {
	switch e {
	case SUCCESS:
		return "Success"
	case VALIDATION_ERROR:
		return "Validation Error"
	case UNAUTHORIZED:
		return "Unauthorized"
	case NOT_FOUND:
		return "Not Found"
	case INTERNAL_SERVER_ERROR:
		return "Internal Server Error"
	default:
		return "Unknown Error"
	}
}
