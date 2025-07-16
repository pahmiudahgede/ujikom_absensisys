package models

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type TimeOnly struct {
	time.Time
}

func NewTimeOnly(hour, min, sec int) TimeOnly {
	t := time.Date(0, time.January, 1, hour, min, sec, 0, time.UTC)
	return TimeOnly{t}
}

func NewTimeOnlyFromString(timeStr string) (TimeOnly, error) {
	t, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		return TimeOnly{}, err
	}
	return TimeOnly{t}, nil
}

func (t TimeOnly) String() string {
	if t.IsZero() {
		return ""
	}
	return t.Time.Format("15:04:05")
}

func (t TimeOnly) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + t.String() + `"`), nil
}

func (t *TimeOnly) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" || str == `""` {
		*t = TimeOnly{}
		return nil
	}

	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}

	parsed, err := NewTimeOnlyFromString(str)
	if err != nil {
		return err
	}
	*t = parsed
	return nil
}

func (t *TimeOnly) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return t.UnmarshalText(v)
	case string:
		return t.UnmarshalText([]byte(v))
	case time.Time:
		*t = TimeOnly{v}
	case nil:
		*t = TimeOnly{}
	default:
		return fmt.Errorf("cannot scan %T into TimeOnly", v)
	}
	return nil
}

func (t TimeOnly) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.String(), nil
}

func (t *TimeOnly) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" {
		*t = TimeOnly{}
		return nil
	}

	parsed, err := time.Parse("15:04:05", str)
	if err != nil {
		return err
	}
	*t = TimeOnly{parsed}
	return nil
}

func (t TimeOnly) MarshalText() ([]byte, error) {
	if t.IsZero() {
		return []byte(""), nil
	}
	return []byte(t.String()), nil
}

func (TimeOnly) GormDataType() string {
	return "TIME"
}

func (TimeOnly) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "TIME"
}

func (t TimeOnly) IsZero() bool {
	return t.Time.IsZero()
}

func (t TimeOnly) GetTime() time.Time {
	return t.Time
}

func (t TimeOnly) Before(other TimeOnly) bool {
	return t.Time.Before(other.Time)
}

func (t TimeOnly) After(other TimeOnly) bool {
	return t.Time.After(other.Time)
}

func (t TimeOnly) Equal(other TimeOnly) bool {
	return t.Time.Equal(other.Time)
}
