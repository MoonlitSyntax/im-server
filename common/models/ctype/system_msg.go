package ctype

import (
	"database/sql/driver"
	"encoding/json"
)

type SystemMsg struct {
	Type int8 `json:"type"` // 1 涉黄 2 涉恐 3 涉政 4 不正当言论
}

func (c *SystemMsg) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}

func (c SystemMsg) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
