package model

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"testing"
)

func TestJsonWrite(t *testing.T) {
	f := &File{
		Filename: "filename",
		FValue:   math.NaN(),
		Tables: []*Table{
			{
				Title: "title1",
			},
			{
				Title: "title2",
			},
		},
	}
	bytes, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
