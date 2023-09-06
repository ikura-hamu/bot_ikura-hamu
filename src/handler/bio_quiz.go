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
	bioQuizQuestionReg = regexp.MustCompile(`ひとことクイズ\s*$`)
	bioQuizAnswerReg   = regexp.MustCompile(`[0-9a-zA-Z-_]+\s*$`)
	giveUpBioQuizReg   = regexp.MustCompile(`あきらめる\s*$`)
)

func (bh *BotHandler) bioQuiz(ctx context.Context, payload payload.EventMessagePayload) error {
	message := payload.MessagePayload.PlainText
	if bioQuizQuestionReg.MatchString(message) {
		return bh.bioQuizQuestion(ctx, payload.MessagePayload.ChannelID)
	}
	if bioQuizAnswerReg.MatchString(message) {
		answer := bioQuizAnswerReg.FindString(message)
		return bh.bioQuizAnswer(ctx, payload.MessagePayload.ChannelID, payload.MessagePayload.User.Name, answer)
	}
	if giveUpBioQuizReg.MatchString(message) {
		return bh.giveUpBioQuiz(ctx, payload.MessagePayload.ChannelID)
	}
	return nil
}

func (bh *BotHandler) bioQuizQuestion(ctx context.Context, channelId uuid.UUID) error {
	quiz, err := bh.br.GetNotAnsweredBioQuiz(ctx, channelId)
	if err == nil {
		message := fmt.Sprintf(
			"このチャンネルではすでにひとことクイズが出題されています。やめる場合は`@BOT_ikura-hamu ひとことクイズ あきらめる`を、解答する場合は`@BOT_ikura-hamu ひとことクイズ {答え}`と送ってください！\nhttps://q.trap.jp/messages/%s",
			quiz.MessageId.String(),
		)
		_, err := bh.cl.SendMessage(ctx, channelId, message, false)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil && !errors.Is(err, repository.ErrBioQuizNotFound) {
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
	messageId, err := bh.cl.SendMessage(ctx, channelId, message, false)
	if err != nil {
		return err
	}

	err = bh.br.CreateBioQuiz(ctx, channelId, messageId, name)
	if err != nil {
		return err
	}
	return nil
}

func (bh *BotHandler) bioQuizAnswer(ctx context.Context, channelId uuid.UUID, userName string, answer string) error {
	quiz, err := bh.br.GetNotAnsweredBioQuiz(ctx, channelId)
	if errors.Is(err, repository.ErrBioQuizNotFound) {
		message := "このチャンネルではひとことクイズが出題されていません。出題する場合は`@BOT_ikura-hamu ひとことクイズ`と送ってください！"
		_, err := bh.cl.SendMessage(ctx, channelId, message, false)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}

	if regexp.MustCompile(fmt.Sprintf(`(?i)^%s\s*$`, quiz.Answer)).MatchString(answer) {
		message := fmt.Sprintf("@%s :accepted.pyon:", userName)
		_, err := bh.cl.SendMessage(ctx, channelId, message, true)
		if err != nil {
			return err
		}
		err = bh.br.AnswerBioQuiz(ctx, quiz.Id)
		if err != nil {
			return err
		}
		return nil
	}

	message := fmt.Sprintf("@%s :wrong_answer:", userName)
	_, err = bh.cl.SendMessage(ctx, channelId, message, false)
	if err != nil {
		return err
	}
	return nil
}

func (bh *BotHandler) giveUpBioQuiz(ctx context.Context, channelId uuid.UUID) error {
	quiz, err := bh.br.GetNotAnsweredBioQuiz(ctx, channelId)
	if errors.Is(err, repository.ErrBioQuizNotFound) {
		message := "このチャンネルではひとことクイズが出題されていません。出題する場合は`@BOT_ikura-hamu ひとことクイズ`と送ってください！"
		_, err := bh.cl.SendMessage(ctx, channelId, message, false)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}

	_, err = bh.cl.SendMessage(ctx, channelId, fmt.Sprintf("正解は :@%s: %s さんでした！また遊んでください！！", quiz.Answer, quiz.Answer), false)
	if err != nil {
		return err
	}
	err = bh.br.AnswerBioQuiz(ctx, quiz.Id)
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
