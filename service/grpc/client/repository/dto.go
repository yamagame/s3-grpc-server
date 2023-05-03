package repository

type RepositoryDTO[T, Request, Response any] interface {
	Input(*T) *Request
	Output(*Response) *T
}

func Domain[T, Request, Response any](dto RepositoryDTO[T, Request, Response], req *T, call func(req *Request) (*Response, error)) (*T, error) {
	file, err := call(dto.Input(req))
	if err != nil {
		return nil, err
	}
	return dto.Output(file), nil
}
