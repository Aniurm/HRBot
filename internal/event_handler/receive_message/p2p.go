package receiveMessage

import (
	"github.com/sirupsen/logrus"
	"strings"
	"xlab-feishu-robot/internal/controller"
	"xlab-feishu-robot/internal/store"
)

func p2p(messageevent *store.MessageEvent) {
	switch strings.ToUpper(messageevent.Message.Message_type) {
	case "TEXT":
		p2pTextMessage(messageevent)
	default:
		logrus.WithFields(logrus.Fields{"message type": messageevent.Message.Message_type}).Warn("Receive p2p message, but this type is not supported")
	}
}

func p2pTextMessage(messageevent *store.MessageEvent) {
	// get the pure text message
	messageevent.Message.Content = strings.TrimSuffix(strings.TrimPrefix(messageevent.Message.Content, "{\"text\":\""), "\"}")
	content := messageevent.Message.Content
	logrus.WithFields(logrus.Fields{"message content": content}).Info("Receive p2p TEXT message")

	switch {
	case strings.Contains(content, "help"):
		controller.HelpP2P(messageevent)
	case strings.Contains(content, "批量加人"):
		controller.AddPeople(messageevent)
	case strings.Contains(content, "批量删人"):
		controller.DelPeople(messageevent)
	case strings.Contains(content, "开始使用"):
		controller.Login(messageevent)
	default:
		logrus.Errorf("Receive p2p TEXT message, but this type is not supported")
	}
}
