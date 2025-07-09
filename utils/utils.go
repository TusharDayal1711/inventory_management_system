package utils

import (
	"encoding/json"
	"net/http"
)

func ParseJSONBody(r *http.Request, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(dst)
	if err != nil {
		return err
	}
	return nil
}
