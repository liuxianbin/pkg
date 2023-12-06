package ginx

import "time"

type JsonTime time.Time

const LocalDateTimeFormat string = "2006-01-02 15:04:05"

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Format(LocalDateTimeFormat) + `"`), nil
}

func (j *JsonTime) UnmarshalJSON(b []byte) error {
	t, err := time.ParseInLocation(`"`+LocalDateTimeFormat+`"`, string(b), time.Local)
	*j = JsonTime(t)
	return err
}
