package dto

type domainDTO[T, Request, Response any] interface {
	Input(*Request) *T
	Output(*T) *Response
}

type infraDTO[T, Request, Response any] interface {
	Input(*T) *Request
	Output(*Response) *T
}
