package errs

const (
	// 4XX
	TypeNotFound        = "not_found"
	TypeBadRequest      = "bad_request"
	TypeUnauthorized    = "unauthorized"
	TypeForbidden       = "forbidden"
	TypeConflict        = "conflict"
	TypeTooManyRequests = "too_many_requests"
	// 5XX
	TypeInternalError = "internal_server_error"
)
