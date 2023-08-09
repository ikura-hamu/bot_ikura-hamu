package payload

import "time"

type BasePayload struct {
	EventTime time.Time `json:"eventTime"`
}
