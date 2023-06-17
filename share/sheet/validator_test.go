package sheet

import (
	"os"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestValidator(t *testing.T) {
	logger := zap.NewExample()
	filename := "./testdata/field-test.csv"

	csv := NewCSVSheet()
	{
		fp, _ := os.Open(filename)
		defer fp.Close()
		csv.Read(fp)
	}

	opts := map[string]string{
		"file": filename,
	}

	fieldFactory := NewFieldFactory(opts)

	fieldsOK := []*Field{
		fieldFactory("AAA", []validation.Rule{validation.Required, is.Alphanumeric}),
		fieldFactory("BBB", []validation.Rule{validation.Required, is.Alphanumeric}),
		fieldFactory("CCC", []validation.Rule{validation.Required, is.Alphanumeric}),
	}
	assert.Equal(t, true, csv.Validate(fieldsOK))

	fieldsNG1 := []*Field{
		fieldFactory("BBB", []validation.Rule{validation.Required, is.Alphanumeric}),
		fieldFactory("AAA", []validation.Rule{validation.Required, is.Alphanumeric}),
		fieldFactory("CCC", []validation.Rule{validation.Required, is.Alphanumeric}),
	}
	assert.Equal(t, false, csv.Validate(fieldsNG1))
	logger.Warn(fieldsNG1[0].Title, fieldsNG1[0].Errors[0].ZapFields()...)

	fieldsNG2 := []*Field{
		fieldFactory("AAA", []validation.Rule{validation.Required, is.Alpha.Error("アルファベット以外が使用されています")}),
		fieldFactory("BBB", []validation.Rule{validation.Required, is.Alphanumeric}),
		fieldFactory("CCC", []validation.Rule{validation.Required, is.Alpha.Error("アルファベット以外が使用されています")}),
	}
	assert.Equal(t, false, csv.Validate(fieldsNG2))
	logger.Warn(fieldsNG2[0].Title, fieldsNG2[0].Errors[0].ZapFields()...)
	logger.Warn(fieldsNG2[0].Title, fieldsNG2[2].Errors[0].ZapFields()...)
}
