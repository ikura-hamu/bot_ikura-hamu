package traq

import (
	"context"

	"github.com/google/uuid"
	"github.com/ikura-hamu/bot_ikura-hamu/src/model"
)

func (tc *TraqClient) GetAllUserIds(ctx context.Context) ([]uuid.UUID, error) {
	users, _, err := tc.client.UserApi.GetUsers(ctx).Execute()
	if err != nil {
		return nil, handleError(err)
	}
	userIds := make([]uuid.UUID, 0, len(users))
	for i := range users {
		userIds = append(userIds, uuid.MustParse(users[i].Id))
	}
	return userIds, nil
}

func (tc *TraqClient) GetUserInfo(ctx context.Context, userId uuid.UUID) (*model.TraqUser, error) {
	user, _, err := tc.client.UserApi.GetUser(ctx, userId.String()).Execute()
	if err != nil {
		return nil, handleError(err)
	}
	return model.NewTraqUser(uuid.MustParse(user.Id), user.Name, user.DisplayName, user.Bio, user.GetLastOnline()), nil
}
