package sheet

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.uber.org/zap"
)

// FieldError FieldError 構造体
type FieldError struct {
	err    error
	opts   map[string]string
	Row    int
	Col    int
	Actual string
	Expect string
}

// NewFieldError FieldError 構造体のコンストラクタ
func NewFieldError(err error, opts map[string]string) *FieldError {
	return &FieldError{
		err:  err,
		opts: opts,
	}
}

// Error エラー文字列生成
func (e *FieldError) Error() string {
	return fmt.Sprintf("%v", e.err)
}

// Unwrap error に変換する
func (e *FieldError) Unwrap() error {
	return e.err
}

// ZapFields Zap用Fieldを返す
func (e *FieldError) ZapFields() []zap.Field {
	fields := []zap.Field{
		zap.String("reason", e.err.Error()),
		zap.String("row", fmt.Sprintf("%d", e.Row)),
		zap.String("col", fmt.Sprintf("%d", e.Col)),
		zap.String("actual", e.Actual),
	}
	if e.Expect != "" {
		fields = append(fields, zap.String("expect", e.Expect))
	}
	for key, val := range e.opts {
		fields = append(fields, zap.String(key, val))
	}
	return fields
}

// Field Field 構造体
type Field struct {
	Title   string
	Rules   []validation.Rule
	Errors  []*FieldError
	Options map[string]string
}

// NewFieldFactory Field コンストラクタを生成
func NewFieldFactory(opts map[string]string) func(title string, rules []validation.Rule) *Field {
	return func(title string, rules []validation.Rule) *Field {
		return &Field{
			Title:   title,
			Rules:   rules,
			Errors:  []*FieldError{},
			Options: opts,
		}
	}
}

func (x *Field) appendError(err *FieldError) {
	x.Errors = append(x.Errors, err)
}

// IsValidLine 有効な行か？
func (x Sheet) IsValidLine(row int) bool {
	if len(x.Cells) > row && row >= 0 {
		c := 0
		for _, v := range x.Cells[row] {
			if v != "" {
				c++
			}
		}
		return c > 0 // 一つでも値があれば有効
	}
	return false
}

func (x Sheet) validateBody(validlast, column int, field *Field) bool {
	retval := true
	for row := 1; row <= validlast; row++ {
		val := x.Cell(row, column)
		err := validation.Validate(&val, field.Rules...)
		if err != nil {
			err := NewFieldError(err, field.Options)
			err.Row = row
			err.Col = column
			err.Actual = val
			err.Expect = ""
			field.appendError(err)
			retval = false
		}
	}
	return retval
}

// Validate シートの値を検証する
func (x Sheet) Validate(fields []*Field) bool {
	retval := true
	validlast := 0
	// 有効な最後の行を検出
	for row, _ := range x.Cells {
		if x.IsValidLine(row) {
			validlast = row
		} else if validlast > 0 {
			break
		}
	}
	// カラムタイトルと値を検証
	for i, field := range fields {
		if x.Cell(0, i) != field.Title {
			err := NewFieldError(fmt.Errorf("カラムタイトルが不正です"), field.Options)
			err.Row = 0
			err.Col = i
			err.Actual = x.Cell(0, i)
			err.Expect = field.Title
			field.appendError(err)
			retval = false
		} else if len(field.Rules) > 0 {
			if !x.validateBody(validlast, i, field) {
				retval = false
			}
		}
	}
	return retval
}
