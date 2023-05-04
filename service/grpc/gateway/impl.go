package gateway

func ToDomain[T, Request, Response any](gateway domainGateway[T, Request, Response], req *Request, call func(req *T) (*T, error)) (*Response, error) {
	file, err := call(gateway.Input(req))
	if err != nil {
		return nil, err
	}
	return gateway.Output(file), nil
}

func ToInfra[T, Request, Response any](gateway infraGateway[T, Request, Response], req *T, call func(req *Request) (*Response, error)) (*T, error) {
	file, err := call(gateway.Input(req))
	if err != nil {
		return nil, err
	}
	return gateway.Output(file), nil
}
