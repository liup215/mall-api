package time

import (
	"database/sql/driver"
	"errors"
	"strconv"
	xtime "time"
)

const (
	timeFormart = "2006-01-02 15:04:05"
)

// Time be used to MySql timestamp converting.
type Time int64

// Scan scan time.
func (jt *Time) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case xtime.Time:
		*jt = Time(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*jt = Time(i)
	case int64:
		*jt = Time(sc)
	}
	return
}

// Value get time value.
func (jt *Time) Value() (driver.Value, error) {
	if jt == nil {
		return nil, nil
	}
	return xtime.Unix(int64(*jt), 0), nil
}

// Time get time.
func (jt Time) Time() xtime.Time {
	return xtime.Unix(int64(jt), 0)
}

func (jt *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := xtime.ParseInLocation(`"`+timeFormart+`"`, string(data), xtime.Local)
	if err != nil {
		jt = nil
		return nil
	}
	*jt = Time(now.Unix())
	//t=nil
	return
}

func (jt Time) MarshalJSON() ([]byte, error) {
	if y := jt.Time().Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	if !jt.Time().IsZero() {
		b = jt.Time().AppendFormat(b, timeFormart)
	}

	b = append(b, '"')
	return b, nil
}

func Now() Time {
	now := xtime.Now()
	return Time(now.Unix())
}
