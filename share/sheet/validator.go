package sheet

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Field struct {
	Title  string
	Rules  []validation.Rule
	Errors []error
}

func NewField(title string, rules []validation.Rule) *Field {
	return &Field{
		Title:  title,
		Rules:  rules,
		Errors: []error{},
	}
}

func (x *Field) AppendError(err error) {
	x.Errors = append(x.Errors, err)
}

func (x Sheet) Validate(fields []*Field) bool {
	retval := true
	for i, field := range fields {
		if x.Cell(0, i) != field.Title {
			field.AppendError(fmt.Errorf("カラムタイトルが不正です"))
			retval = false
		} else if len(field.Rules) > 0 {
			for y := 1; y < 3; y++ {
				val := x.Cell(y, i)
				err := validation.Validate(&val, field.Rules...)
				if err != nil {
					field.AppendError(err)
					retval = false
				}
			}
		}
	}
	return retval
}
