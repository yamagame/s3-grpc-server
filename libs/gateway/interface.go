package gateway

type domainGateway[T, Request, Response any] interface {
	Input(*Request) *T
	Output(*T) *Response
}

type infraGateway[T, Request, Response any] interface {
	Input(*T) *Request
	Output(*Response) *T
}
