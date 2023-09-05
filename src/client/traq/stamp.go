package traq

import (
	"context"

	"github.com/google/uuid"
	"github.com/ikura-hamu/bot_ikura-hamu/src/client"
)

func (tc *TraqClient) AddStamp(ctx context.Context, messageId uuid.UUID, stampID uuid.UUID, count int) error {
	for i := 0; i < count; i++ {
		_, err := tc.client.StampApi.AddMessageStamp(ctx, messageId.String(), stampID.String()).Execute()
		if err != nil {
			return handleError(err)
		}
	}
	return nil
}

func (tc *TraqClient) GetAllStamps(ctx context.Context) (map[string]uuid.UUID, error) {
	stamps, _, err := tc.client.StampApi.GetStamps(ctx).Execute()
	if err != nil {
		return nil, handleError(err)
	}
	stampsMap := make(map[string]uuid.UUID, len(stamps))
	for s := range stamps {
		stampsMap[stamps[s].Name] = uuid.MustParse(stamps[s].Id)
	}

	return stampsMap, nil
}

func (tc *TraqClient) GetStampIdByName(ctx context.Context, name string) (uuid.UUID, error) {
	return tc.stampCache.Get(ctx, name)
}

func (tc *TraqClient) getStampIdByName(ctx context.Context, name string) (uuid.UUID, error) {
	stamps, err := tc.GetAllStamps(ctx)
	if err != nil {
		return uuid.Nil, handleError(err)
	}

	stampId, ok := stamps[name]
	if !ok {
		return uuid.Nil, client.ErrInvalidStampName
	}
	return stampId, nil
}
