package traq

import (
	"context"

	"github.com/google/uuid"
)

func (tc *TraqClient) AddStamp(ctx context.Context, messageId uuid.UUID, stampID uuid.UUID, count int) error {
	for i := 0; i < count; i++ {
		_, err := tc.client.StampApi.AddMessageStamp(ctx, messageId.String(), stampID.String()).Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

func (tc *TraqClient) GetAllStamps(ctx context.Context) (map[string]uuid.UUID, error) {
	stamps, _, err := tc.client.StampApi.GetStamps(ctx).Execute()
	if err != nil {
		return nil, err
	}
	stampsMap := make(map[string]uuid.UUID)
	for s := range stamps {
		stampsMap[stamps[s].Name] = stampsMap[stamps[s].Id]
	}

	return stampsMap, nil
}
