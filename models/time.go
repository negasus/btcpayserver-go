package models

import (
	"strconv"
	"time"
)

type Time struct {
	time.Time
}

func (s *Time) UnmarshalJSON(data []byte) error {
	v, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	s.Time = time.Unix(v, 0)

	return nil
}
