package models

import (
	"encoding/json"
	"strconv"
)

type Float float64

func (s *Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatFloat(float64(*s), 'f', -1, 64))
}

func (s *Float) UnmarshalJSON(data []byte) error {
	var v string
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return err
	}

	*s = Float(f)

	return nil
}
