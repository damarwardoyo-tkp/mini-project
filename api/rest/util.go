package rest

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ParseBody(body io.ReadCloser, dest interface{}) error {
	defer body.Close()
	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(raw, &dest)
	if err != nil {
		return err
	}
	return nil
}
