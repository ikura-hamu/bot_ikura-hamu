package payload

import (
	"time"

	"github.com/google/uuid"
)

type MessagePayload struct {
	ID        uuid.UUID             `json:"id"`
	User      UserPayload           `json:"user"`
	ChannelID uuid.UUID             `json:"channelId"`
	Text      string                `json:"text"`
	PlainText string                `json:"plainText"`
	Embedded  []EmbeddedInfoPayload `json:"embedded"`
	CreatedAt time.Time             `json:"createdAt"`
	UpdatedAt time.Time             `json:"updatedAt"`
}

type EmbeddedInfoPayload struct {
	Raw  string `json:"raw"`
	Type string `json:"type"`
	ID   string `json:"id"`
}
