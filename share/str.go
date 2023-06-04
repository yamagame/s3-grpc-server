package share

import (
	"encoding/json"
	"log"
)

func JsonString(a any) string {
	bytes, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
