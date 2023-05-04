package dto

func ToDomain[T, Request, Response any](dto domainDTO[T, Request, Response], req *Request, call func(req *T) (*T, error)) (*Response, error) {
	file, err := call(dto.Input(req))
	if err != nil {
		return nil, err
	}
	return dto.Output(file), nil
}

func ToInfra[T, Request, Response any](dto infraDTO[T, Request, Response], req *T, call func(req *Request) (*Response, error)) (*T, error) {
	file, err := call(dto.Input(req))
	if err != nil {
		return nil, err
	}
	return dto.Output(file), nil
}
