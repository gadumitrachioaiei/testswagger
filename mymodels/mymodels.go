package mymodels

import (
	"encoding/json"

	"github.com/go-openapi/strfmt"
)

type RawMessage struct {
	json.RawMessage
}

func (*RawMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// func (m RawMessage) MarshalJSON() ([]byte, error) {
// 	return json.RawMessage(m).MarshalJSON()
// }

// // UnmarshalJSON sets *m to a copy of data.
// func (m *RawMessage) UnmarshalJSON(data []byte) error {
