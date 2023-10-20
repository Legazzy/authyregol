package response

func NewSuccessCreated(detail string) Response {
	res := Response{}

	res.Status = 201

	res.Header = "Created"
	res.Extern = "https://www.rfc-editor.org/rfc/rfc7231.html#section-6.3.2"

	res.Detail = detail

	return res
}

func NewSuccessNoContent(detail string) Response {
	res := Response{}

	res.Status = 204

	res.Header = "No Content"
	res.Extern = "https://www.rfc-editor.org/rfc/rfc7231.html#section-6.3.5"

	res.Detail = detail

	return res
}

func NewSuccessOK(detail string) Response {
	res := Response{}

	res.Status = 200

	res.Header = "OK"
	res.Extern = "https://www.rfc-editor.org/rfc/rfc7231.html#section-6.3.1"

	res.Detail = detail

	return res
}

func NewClientBadRequest(detail string) Response {
	res := Response{}

	res.Status = 400

	res.Header = "Bad Request"
	res.Extern = "https://www.rfc-editor.org/rfc/rfc7231.html#section-6.5.1"

	res.Detail = detail

	return res
}

func NewClientConflict(detail string) Response {
	res := Response{}

	res.Status = 409

	res.Header = "Conflict"
	res.Extern = "https://www.rfc-editor.org/rfc/rfc7231.html#section-6.5.8"

	res.Detail = detail

	return res
}

func NewClientForbidden(detail string) Response {
	res := Response{}

	res.Status = 403

	res.Header = "Forbidden"
	res.Extern = "https://www.rfc-editor.org/rfc/rfc7231.html#section-6.5.3"

	res.Detail = detail

	return res
}

func NewClientMethodNotAllowed(detail string) Response {
	res := Response{}

	res.Status = 405

	res.Header = "Method Not Allowed"
	res.Extern = "https://www.rfc-editor.org/rfc/rfc7231.html#section-6.5.5"

	res.Detail = detail

	return res
}

func NewClientNotFound(detail string) Response {
	res := Response{}

	res.Status = 404

	res.Header = "Not Found"
	res.Extern = "https://www.rfc-editor.org/rfc/rfc7231.html#section-6.5.4"

	res.Detail = detail

	return res
}

func NewClientUnauthorized(detail string) Response {
	res := Response{}

	res.Status = 401

	res.Header = "Unauthorized"
	res.Extern = "https://www.rfc-editor.org/rfc/rfc7235.html#section-3.1"

	res.Detail = detail

	return res
}

func NewServerInternalServerError(detail string) Response {
	res := Response{}

	res.Status = 500

	res.Header = "Internal Server Error"
	res.Extern = "https://www.rfc-editor.org/rfc/rfc7231.html#section-6.6.1"

	res.Detail = detail

	return res
}
