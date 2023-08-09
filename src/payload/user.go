package payload

import (
	"github.com/google/uuid"
)

type UserPayload struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"displayName"`
	IconID      uuid.UUID `json:"iconId"`
	Bot         bool      `json:"bot"`
}
