package payload

type EventMessagePayload struct {
	BasePayload
	MessagePayload `json:"message"`
}
