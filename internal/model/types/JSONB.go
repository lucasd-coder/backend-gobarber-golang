package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type JSON map[string]string

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j JSON) Scan(v interface{}) error {
	if v == nil {
		return nil
	}
	switch data := v.(type) {
	case string:
		return json.Unmarshal([]byte(data), &j)
	case []byte:
		return json.Unmarshal(data, &j)
	default:
		return errors.New("cannot scan type %t into Map")
	}
}

func (JSON) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// use field.Tag, field.TagSettings gets field's tags
	// checkout https://github.com/go-gorm/gorm/blob/master/schema/field.go for all options

	// returns different database type based on driver name
	switch db.Dialector.Name() {
	case "mysql", "sqlite":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
