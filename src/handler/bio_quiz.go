package handler

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/ikura-hamu/bot_ikura-hamu/pkg/payload"
	"github.com/ikura-hamu/bot_ikura-hamu/src/repository"
)

var (
	bioQuizQuestionReg = regexp.MustCompile(`ひとことクイズ$`)
)

func (bh *BotHandler) BioQuiz(ctx context.Context, payload payload.EventMessagePayload) error {
	message := payload.MessagePayload.PlainText
	if bioQuizQuestionReg.MatchString(message) {

	}
	return nil
}

func (bh *BotHandler) bioQuizQuestion(ctx context.Context, channelId uuid.UUID) error {
	quiz, err := bh.br.GetNotAnsweredBioQuiz(ctx, channelId)
	if !errors.Is(err, repository.ErrBioQuizNotFound) {
		message := fmt.Sprintf(
			"このチャンネルではすでにひとことクイズが出題されています。やめる場合は`@BOT_ikura-hamu ひとことクイズ あきらめる`を、解答する場合は`@BOT_ikura-hamu ひとことクイズ {答え}`と送ってください！\n%s",
			quiz.ChannelId.String(),
		)
		err := bh.cl.SendMessage(ctx, channelId, message, false)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}

	userIds, err := bh.cl.GetAllUserIds(ctx)
	if err != nil {
		return err
	}
	name, bio, err := bh.makeBioQuiz(ctx, userIds)
	if err != nil {
		return err
	}

	message := fmt.Sprintf(
		"ひとことが\n> %s \nの人は誰でしょう?\n\n答えは`@BOT_ikura-hamu ひとことクイズ {答え}`と送ってください！",
		digestBio(bio),
	)
	err = bh.cl.SendMessage(ctx, channelId, message, false)
	if err != nil {
		return err
	}

	err = bh.br.CreateBioQuiz(ctx, channelId, uuid.New(), name)
	if err != nil {
		return err
	}
	return nil
}

func (bh *BotHandler) bioQuizAnswer(ctx context.Context, channelId uuid.UUID, userName string, answer string) error {
	quiz, err := bh.br.GetNotAnsweredBioQuiz(ctx, channelId)
	if errors.Is(err, repository.ErrBioQuizNotFound) {
		message := "このチャンネルではひとことクイズが出題されていません。出題する場合は`@BOT_ikura-hamu ひとことクイズ`と送ってください！"
		err := bh.cl.SendMessage(ctx, channelId, message, false)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}

	if regexp.MustCompile(fmt.Sprintf(`(?i)^%s$`, quiz.Answer)).MatchString(answer) {
		message := fmt.Sprintf("@%s :accepted.pyon:", userName)
		err := bh.cl.SendMessage(ctx, channelId, message, true)
		if err != nil {
			return err
		}
		err = bh.br.AnswerBioQuiz(ctx, quiz.ChannelId)
		if err != nil {
			return err
		}
		return nil
	}

	err = bh.cl.SendMessage(ctx, channelId, ":wrong_answer:", false)
	if err != nil {
		return err
	}
	return nil
}

func (bh *BotHandler) giveUpBioQuiz(ctx context.Context, channelId uuid.UUID) error {
	quiz, err := bh.br.GetNotAnsweredBioQuiz(ctx, channelId)
	if errors.Is(err, repository.ErrBioQuizNotFound) {
		message := "このチャンネルではひとことクイズが出題されていません。出題する場合は`@BOT_ikura-hamu ひとことクイズ`と送ってください！"
		err := bh.cl.SendMessage(ctx, channelId, message, false)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}

	err = bh.cl.SendMessage(ctx, channelId, fmt.Sprintf("正解は :@%s: %s さんでした！また遊んでください！！", quiz.Answer, quiz.Answer), false)
	if err != nil {
		return err
	}
	err = bh.br.AnswerBioQuiz(ctx, quiz.ChannelId)
	if err != nil {
		return err
	}
	return nil
}

func (bh *BotHandler) makeBioQuiz(ctx context.Context, userIds []uuid.UUID) (string, string, error) {
	for {
		userId := userIds[rand.Intn(len(userIds))]
		user, err := bh.cl.GetUserInfo(ctx, userId)
		if err != nil {
			return "", "", err
		}
		if time.Since(user.GetLastOnline()) > 31*24*time.Hour || user.GetBio() == "" {
			continue
		}
		return user.GetName(), user.GetBio(), nil
	}
}

func digestBio(bio string) string {
	bioByLines := regexp.MustCompile("\r\n|\n").Split(bio, -1)
	var newBio string
	for _, bioByLine := range bioByLines {
		newBio += ">"
		newBio += bioByLine
	}
	return newBio
}
