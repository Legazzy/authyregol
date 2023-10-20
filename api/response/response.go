package response

type Response struct {
	Status int

	Header string
	Detail string
	Extern string

	Return interface{}
}
