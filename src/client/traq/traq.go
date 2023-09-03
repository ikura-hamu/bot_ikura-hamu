package traq

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ikura-hamu/bot_ikura-hamu/src/client"
	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"
	"github.com/motoki317/sc"
	traq "github.com/traPtitech/go-traq"
	"go.uber.org/zap"
)

var _ client.Client = &TraqClient{} //clientを実装しているか確認

type TraqClient struct {
	client     *traq.APIClient
	stampCache *sc.Cache[string, uuid.UUID]
	logger     *zap.Logger
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

	tc := TraqClient{client: client, logger: l.Named("traQ Client")}

	tc.stampCache, err = sc.New[string, uuid.UUID](tc.GetStampIdByName, 24*time.Hour, 48*time.Hour)
	if err != nil {
		l.Panic("failed to create stamp cache", zap.Error(err))
	}

	return &tc
}
