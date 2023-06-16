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

	csv := NewCSVSheet()
	{
		fp, _ := os.Open("./testdata/field-test.csv")
		defer fp.Close()
		csv.Read(fp)
	}

	fieldsOK := []*Field{
		NewField("AAA", []validation.Rule{validation.Required, is.Alphanumeric}),
		NewField("BBB", []validation.Rule{validation.Required, is.Alphanumeric}),
		NewField("CCC", []validation.Rule{validation.Required, is.Alphanumeric}),
	}
	assert.Equal(t, true, csv.Validate(fieldsOK))

	fieldsNG1 := []*Field{
		NewField("BBB", []validation.Rule{validation.Required, is.Alphanumeric}),
		NewField("AAA", []validation.Rule{validation.Required, is.Alphanumeric}),
		NewField("CCC", []validation.Rule{validation.Required, is.Alphanumeric}),
	}
	assert.Equal(t, false, csv.Validate(fieldsNG1))
	logger.Warn(fieldsNG1[0].Title, zap.Error(fieldsNG1[0].Errors[0]))

	fieldsNG2 := []*Field{
		NewField("AAA", []validation.Rule{validation.Required, is.Alpha.Error("数字のみです")}),
		NewField("BBB", []validation.Rule{validation.Required, is.Alphanumeric}),
		NewField("CCC", []validation.Rule{validation.Required, is.Alphanumeric}),
	}
	assert.Equal(t, false, csv.Validate(fieldsNG2))
	logger.Warn(fieldsNG2[0].Title, zap.Error(fieldsNG2[0].Errors[0]))
}
