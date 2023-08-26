package interceptor

import (
	"context"
	"reflect"

	"google.golang.org/grpc"
)

// ValidateRule バリデーションルールを管理
type ValidateRule struct {
	Type     interface{}
	Validate func(ctx context.Context, val interface{}) error
}

// Validator バリデータ
type Validator struct {
	validators []*ValidateRule
}

// NewValidator Validator のコンストラクタ
func NewValidator() *Validator {
	return &Validator{
		validators: []*ValidateRule{},
	}
}

// AddRule ルールを追加
func (x *Validator) AddRule(validate *ValidateRule) {
	x.validators = append(x.validators, validate)
}

// AddRule ルールを複数追加
func (x *Validator) AddRules(rules []*ValidateRule) {
	x.validators = append(x.validators, rules...)
}

// WithContext 登録したルールにそってバリデーションを実施
func (x *Validator) WithContext(ctx context.Context, val interface{}) error {
	for _, v := range x.validators {
		if reflect.TypeOf(v.Type) == reflect.TypeOf(val) {
			return v.Validate(ctx, val)
		}
	}
	return nil
}

// gRPC バリデーションインターセプタ
func ValidateInterceptor(validate *Validator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := validate.WithContext(ctx, req); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

// interface{} から T へ変換
func validateGateway[T any](vfunc func(ctx context.Context, req *T) error) func(ctx context.Context, req interface{}) error {
	return func(ctx context.Context, req interface{}) error {
		switch r := req.(type) {
		case *T:
			return vfunc(ctx, r)
		}
		return nil
	}
}

// Validate バリデーションルールを実装する
func Validate[T any](t *T, call func(ctx context.Context, req *T) error) *ValidateRule {
	return &ValidateRule{
		Type:     t,
		Validate: validateGateway(call),
	}
}
