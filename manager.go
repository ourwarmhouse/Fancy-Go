package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/BetaGoRobot/BetaGo/betagovar"
	"github.com/BetaGoRobot/BetaGo/commandHandler/admin"
	"github.com/BetaGoRobot/BetaGo/commandHandler/cal"
	errorsender "github.com/BetaGoRobot/BetaGo/commandHandler/error_sender"
	"github.com/BetaGoRobot/BetaGo/commandHandler/helper"
	"github.com/BetaGoRobot/BetaGo/commandHandler/music"
	"github.com/BetaGoRobot/BetaGo/commandHandler/roll"
	"github.com/BetaGoRobot/BetaGo/dbpack"
	"github.com/BetaGoRobot/BetaGo/utility"
	"github.com/enescakir/emoji"
	jsoniter "github.com/json-iterator/go"
	"github.com/lonelyevil/khl"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func clickEventHandler(ctx *khl.MessageButtonClickContext) {
	var err error
	clickValue := ctx.Extra.Value
	isAdmin := dbpack.CheckIsAdmin(ctx.Extra.UserID)
	switch clickValue {
	case "SHOWADMIN":
		if isAdmin {
			err = admin.ShowAdminHandler(ctx.Extra.TargetID, "")
		}
	case "HELP":
		if isAdmin {
			err = helper.AdminCommandHelperHandler(ctx.Extra.TargetID, "", ctx.Extra.UserID)
		} else {
			err = helper.UserCommandHelperHandler(ctx.Extra.TargetID, "", ctx.Extra.UserID)
		}
	case "ROLL":
		err = roll.RandRollHandler(ctx.Extra.TargetID, "", ctx.Extra.UserID)
	case "ONEWORD":
		err = roll.OneWordHandler(ctx.Extra.TargetID, "", ctx.Extra.UserID)
	case "PING":
		helper.PingHandler(ctx.Extra.TargetID, "", ctx.Extra.UserID)
	case "SHOWCAL":
		err = cal.ShowCalHandler(ctx.Extra.TargetID, "", ctx.Extra.UserID, ctx.Extra.GuildID)
	default:
		err = fmt.Errorf("非法操作" + emoji.Warning.String())
	}
	if err != nil {
		errorsender.SendErrorInfo(ctx.Extra.TargetID, "", ctx.Extra.UserID, err)
	}
}

func commandHandler(ctx *khl.KmarkdownMessageContext) {
	// 判断是否被at到,且消息不是引用/回复
	if !utility.IsInSlice(betagovar.RobotID, ctx.Extra.Mention) {
		return
	}
	// 示例中，由于用户发送的命令的Content格式为(met)id(met) <command> <parameters>
	// 针对解析的判断逻辑，首先判断是否为空字符串，若为空发送help信息
	// ? 解析出不包含at信息的实际内容
	trueContent := strings.Trim(ctx.Common.Content[strings.LastIndex(ctx.Common.Content, `)`)+1:], " ")
	if trueContent != "" {
		// 内容非空，解析命令
		var (
			command     string
			parameters  []string
			adminSolved bool
			slice       = strings.Split(strings.Trim(trueContent, " "), " ")
		)
		// 判断指令类型
		if len(slice) == 1 {
			command = slice[0]
		} else {
			command = slice[0]
			parameters = slice[1:]
		}
		command = strings.ToUpper(command)
		var err error
		// 进入指令执行逻辑,首先判断是否为Admin
		if dbpack.CheckIsAdmin(ctx.Common.AuthorID) {
			// 是Admin，判断指令
			switch command {
			case "HELP":
				err = helper.AdminCommandHelperHandler(ctx.Common.TargetID, ctx.Common.MsgID, ctx.Common.AuthorID)
				adminSolved = true
			case "ADDADMIN":
				if len(parameters) == 2 {
					userID, userName := parameters[0], parameters[1]
					err = admin.AddAdminHandler(userID, userName, ctx.Common.MsgID, ctx.Common.TargetID)
				} else {
					err = fmt.Errorf("请输入正确的用户ID和用户名格式")
				}
				adminSolved = true
			case "REMOVEADMIN":
				if len(parameters) == 1 {
					targetUserID := parameters[0]
					err = admin.RemoveAdminHandler(ctx.Common.AuthorID, targetUserID, ctx.Common.MsgID, ctx.Common.TargetID)
				} else {
					err = fmt.Errorf("请输入正确的用户ID和用户名格式")
				}
				adminSolved = true
			case "SHOWADMIN":
				err = admin.ShowAdminHandler(ctx.Common.TargetID, ctx.Common.MsgID)
				adminSolved = true
			default:
				adminSolved = false
			}
			if err != nil {
				errorsender.SendErrorInfo(ctx.Common.TargetID, ctx.Common.MsgID, ctx.Common.AuthorID, err)
			}
		}
		// 非管理员命令
		if !adminSolved {
			switch command {
			case "HELP":
				err = helper.UserCommandHelperHandler(ctx.Common.TargetID, ctx.Common.MsgID, ctx.Common.AuthorID)
			case "ROLL":
				err = roll.RandRollHandler(ctx.Common.TargetID, ctx.Common.MsgID, ctx.Common.AuthorID)
			case "PING":
				helper.PingHandler(ctx.Common.TargetID, ctx.Common.MsgID, ctx.Common.AuthorID)
			case "ONEWORD":
				err = roll.OneWordHandler(ctx.Common.TargetID, ctx.Common.MsgID, ctx.Common.AuthorID)
			case "SEARCHMUSIC":
				err = music.SearchMusicByRobot(ctx.Common.TargetID, ctx.Common.MsgID, ctx.Common.AuthorID, parameters)
			case "GETUSER":
				err = helper.GetUserInfoHandler(parameters[0], ctx.Extra.GuildID, ctx.Common.TargetID, ctx.Common.MsgID)
			case "SHOWCAL":
				err = cal.ShowCalHandler(ctx.Common.TargetID, ctx.Common.MsgID, ctx.Common.AuthorID, ctx.Extra.GuildID)
			default:
				err = fmt.Errorf("未知的指令`%s`, 尝试使用command `help`来获取可用的命令列表", command)
			}
			if err != nil {
				errorsender.SendErrorInfo(ctx.Common.TargetID, ctx.Common.MsgID, ctx.Common.AuthorID, err)
			}
		}
	}

}

func channelJoinedHandler(ctx *khl.GuildChannelMemberAddContext) {

	userInfo, err := utility.GetUserInfo(ctx.Extra.UserID, ctx.Common.TargetID)
	if err != nil {
		errorsender.SendErrorInfo(betagovar.NotifierChanID, "", userInfo.ID, err)
		return
	}
	channelInfo, err := utility.GetChannnelInfo(ctx.Extra.ChannelID)
	if err != nil {
		errorsender.SendErrorInfo(betagovar.NotifierChanID, "", userInfo.ID, err)
		return
	}
	// !频道日志记录
	newChanLog := &dbpack.ChannelLog{
		UserID:      userInfo.ID,
		UserName:    userInfo.Username,
		ChannelID:   channelInfo.ID,
		ChannelName: channelInfo.Name,
		JoinedTime:  ctx.Extra.JoinedAt.ToTime(),
		LeftTime:    time.Time{},
	}
	if err = newChanLog.AddJoinedRecord(); err != nil {
		errorsender.SendErrorInfo(betagovar.NotifierChanID, "", userInfo.ID, err)
	}

	cardMessageStr, err := khl.CardMessage{&khl.CardMessageCard{
		Theme: khl.CardThemeInfo,
		Size:  khl.CardSizeLg,
		Modules: []interface{}{
			khl.CardMessageSection{
				Text: khl.CardMessageElementKMarkdown{
					Content: "`" + userInfo.Nickname + "`悄悄加入了语音频道`" + channelInfo.Name + "`",
				},
			},
		},
	}}.BuildMessage()
	if err != nil {
		errorsender.SendErrorInfo(ctx.Common.TargetID, "", "", err)
		return
	}
	_, err = betagovar.GlobalSession.MessageCreate(&khl.MessageCreate{
		MessageCreateBase: khl.MessageCreateBase{
			Type:     khl.MessageTypeCard,
			TargetID: betagovar.NotifierChanID,
			Content:  cardMessageStr,
		},
	})
	if err != nil {
		errorsender.SendErrorInfo(betagovar.NotifierChanID, "", "", err)
	}
}

func channelLeftHandler(ctx *khl.GuildChannelMemberDeleteContext) {
	// 离开频道时，记录频道信息
	userInfo, err := utility.GetUserInfo(ctx.Extra.UserID, ctx.Common.TargetID)
	if err != nil {
		errorsender.SendErrorInfo(betagovar.NotifierChanID, "", userInfo.ID, err)
		return
	}
	channelInfo, err := utility.GetChannnelInfo(ctx.Extra.ChannelID)
	if err != nil {
		errorsender.SendErrorInfo(betagovar.NotifierChanID, "", userInfo.ID, err)
		return
	}

	// !频道日志记录
	newChanLog := &dbpack.ChannelLog{
		UserID:      userInfo.ID,
		UserName:    userInfo.Username,
		ChannelID:   channelInfo.ID,
		ChannelName: channelInfo.Name,
		JoinedTime:  time.Time{},
		LeftTime:    ctx.Extra.ExitedAt.ToTime(),
	}
	if err = newChanLog.UpdateLeftTime(); err != nil {
		errorsender.SendErrorInfo(betagovar.NotifierChanID, "", userInfo.ID, err)
	}
}

func sendMessageToTestChannel(session *khl.Session, content string) {

	session.MessageCreate(&khl.MessageCreate{
		MessageCreateBase: khl.MessageCreateBase{
			Type:     9,
			TargetID: betagovar.TestChanID,
			Content:  content,
		}})
}

func receiveDirectMessage(ctx *khl.DirectMessageReactionAddContext) {
	log.Println("-----------Test")
}
