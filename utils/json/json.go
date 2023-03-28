package json

import (
	"encoding/json"
	"io"
	"log"
)

func Decode(r io.Reader, v interface{}) error {
	body, err := io.ReadAll(r)
	if err != nil {
		log.Println("Cannot read data")
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		log.Println("Error when unmarshal data")
	}

	return err
}
