package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

func main() {
	api := slack.New(
		//#enter bot token here
	)

	fmt.Println("Hello World")

	realtimeApi := api.NewRTM()
	go realtimeApi.ManageConnection()

	for msg := range realtimeApi.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			user, _ := api.GetUserInfo(ev.User)
			status, _ := processNotification(ev)
			if strings.HasPrefix(ev.Text, "-loc") {
			//Emma's ID
				if (user.ID == "U7GLS08AJ") {
					realtimeApi.SendMessage(realtimeApi.NewOutgoingMessage(fmt.Sprintf("Today %v is %v", user.RealName, status), "CBZSE4DS6"))
					realtimeApi.SendMessage(realtimeApi.NewOutgoingMessage(fmt.Sprintf("Today %v is %v", user.RealName, status), "G6LHR73V1"))

				}
				realtimeApi.SendMessage(realtimeApi.NewOutgoingMessage(fmt.Sprintf("Today %v is %v", user.RealName, status), "CBZSE4DS6"))
			}
		}
	}

}

func processNotification(event *slack.MessageEvent) (string, error) {
	notifRegex, err := regexp.Compile("-loc (.*)")

	if err != nil {
		//return err
	}

	validSyntax := notifRegex.FindStringSubmatch(event.Text)

	fmt.Printf("%v", validSyntax)

	if validSyntax == nil || len(validSyntax) < 1 {
		return "", errors.New("invalid syntax")
	}

	return validSyntax[1], nil

}
