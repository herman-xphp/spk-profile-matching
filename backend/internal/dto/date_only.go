package dto

import (
	"strings"
	"time"
)

type DateOnly time.Time

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" {
		return nil
	}

	// coba format YYYY-MM-DD
	t, err := time.Parse("2006-01-02", s)
	if err == nil {
		*d = DateOnly(t)
		return nil
	}

	// fallback ke RFC3339 (biar kompatibel juga)
	t, err = time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}

	*d = DateOnly(t)
	return nil
}

func (d DateOnly) Time() time.Time {
	return time.Time(d)
}
