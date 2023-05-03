package repository

type RepositoryDTO[T, Request, Response any] interface {
	Input(*Request) *T
	Output(*T) *Response
}

func Domain[T, Request, Response any](dto RepositoryDTO[T, Request, Response], req *Request, call func(req *T) (*T, error)) (*Response, error) {
	file, err := call(dto.Input(req))
	if err != nil {
		return nil, err
	}
	return dto.Output(file), nil
}
