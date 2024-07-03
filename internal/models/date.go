package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

const dateFormat = "2006-01-02"

// UnmarshalJSON парсит дату из строки в формате JSON
func (d *Date) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	t, err := time.Parse(`"`+dateFormat+`"`, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// MarshalJSON форматирует дату в строку в формате JSON
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format(dateFormat))), nil
}

// String возвращает дату в виде строки
func (d Date) String() string {
	return d.Format(dateFormat)
}

// Scan реализует интерфейс Scanner для чтения из базы данных
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		d.Time = v
	case string:
		t, err := time.Parse(dateFormat, v)
		if err != nil {
			return err
		}
		d.Time = t
	default:
		return fmt.Errorf("cannot scan type %T into Date", value)
	}
	return nil
}

// Value реализует интерфейс Valuer для записи в базу данных
func (d Date) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil
	}
	return d.Format(dateFormat), nil
}
