package gateway

func ToInfra[T, Request, Response any](req *T, input func(*T) *Request, call func(req *Request) (*Response, error), output func(*Response) *T) (*T, error) {
	value, err := call(input(req))
	if err != nil {
		return nil, err
	}
	return output(value), nil
}

func ToInfraList[T, Request, Response any](req *T, input func(*T) *Request, call func(req *Request) (*Response, error), output func(*Response) []*T) ([]*T, error) {
	value, err := call(input(req))
	if err != nil {
		return nil, err
	}
	return output(value), nil
}

func ToDomain[T, Request, Response any](req *Request, input func(*Request) *T, call func(req *T) (*T, error), output func(*T) *Response) (*Response, error) {
	value, err := call(input(req))
	if err != nil {
		return nil, err
	}
	return output(value), nil
}

func ToDomainList[T, Request, Response any](req *Request, input func(*Request) *T, call func(req *T) ([]*T, error), output func([]*T) *Response) (*Response, error) {
	value, err := call(input(req))
	if err != nil {
		return nil, err
	}
	return output(value), nil
}
