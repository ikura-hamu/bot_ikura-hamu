package traq

import (
	"context"

	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"
	traq "github.com/traPtitech/go-traq"
	"go.uber.org/zap"
)

type TraqClient struct {
	client *traq.APIClient
	logger *zap.Logger
}

func NewTraqClient(l *zap.Logger) *TraqClient {
	token, ok := conf.GetTraqClientConf()
	if !ok {
		l.Panic("failed to get token")
	}
	clientConf := traq.NewConfiguration()
	clientConf.AddDefaultHeader("Authorization", "Bearer "+token)
	client := traq.NewAPIClient(clientConf)

	me, _, err := client.MeApi.GetMe(context.Background()).Execute()
	if err != nil {
		l.Panic("invalid client", zap.Error(err))
	}
	l.Info("new client", zap.String("name", me.Name))

	return &TraqClient{
		client: client,
		logger: l.Named("traQ Client"),
	}
}
