package time

import (
	"database/sql/driver"
	"errors"
	xtime "time"

	"gopkg.in/mgo.v2/bson"
)

type DateTime string

// Scan scan time.
func (jt *DateTime) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case xtime.Time:
		*jt = DateTime(sc.Format(timeFormart))
	case string:
		_, err = xtime.Parse(timeFormart, sc)
		if err == nil {
			*jt = DateTime(sc)
		}
	}
	return
}

// Value get time value.
func (jt *DateTime) Value() (driver.Value, error) {
	if jt == nil {
		return nil, nil
	}
	return xtime.Parse(timeFormart, string(*jt))
}

func (jt DateTime) GetBSON() (interface{}, error) {
	if jt == "" {
		return nil, nil
	}
	return jt.Time(), nil
}

func (t *DateTime) SetBSON(raw bson.Raw) error {
	var decode xtime.Time
	err := raw.Unmarshal(&decode)
	if err != nil {
		return err
	}

	*t = DateTime(decode.Format(timeFormart))

	return nil
}

// Time get time.
func (jt DateTime) Time() xtime.Time {
	t, _ := xtime.Parse(timeFormart, string(jt))
	return t
}

func (jt *DateTime) UnmarshalJSON(data []byte) (err error) {
	dataString := string(data)
	_, err = xtime.ParseInLocation(`"`+timeFormart+`"`, dataString, xtime.Local)
	if err != nil {
		jt = nil
		return nil
	}
	*jt = DateTime(dataString)
	//t=nil
	return
}

func (jt DateTime) MarshalJSON() ([]byte, error) {
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

func NowDateTime() DateTime {
	now := xtime.Now()
	return DateTime(now.Format(timeFormart))
}
