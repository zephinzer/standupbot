package main

import (
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	api := slack.New(
		os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN"),
		slack.OptionAppLevelToken(os.Getenv("SLACK_BOT_SOCKET_TOKEN")),
		slack.OptionLog(
			log.New(os.Stdout, "api: ", log.Lshortfile),
		),
		slack.OptionDebug(true),
	)
	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(
			log.New(os.Stdout, "socketmode: ", log.Lshortfile),
		),
	)

	go func() {
		for event := range client.Events {
			switch event.Type {
			case socketmode.EventTypeInteractive:
				log.Printf("received interactive event")
				callback, ok := event.Data.(slack.InteractionCallback)
				if !ok {
					log.Printf("ignored %+v", event)
					continue
				}
				// o, _ := json.MarshalIndent(callback.User, "", "  ")
				// fmt.Println(string(o))
				log.Printf("received action[%s] in channel[%s] from user[%v]", callback.ActionID, callback.Message.Channel, callback.User.ID)
				client.Ack(*event.Request)
			case socketmode.EventTypeSlashCommand:
				log.Printf("received slash command event")
				command, ok := event.Data.(slack.SlashCommand)
				if !ok {
					log.Printf("ignored %+v", event)
					continue
				}
				log.Printf("received command: %+v", command)
				response := map[string]interface{}{
					"blocks": []slack.Block{
						slack.NewSectionBlock(
							&slack.TextBlockObject{
								Type: slack.MarkdownType,
								Text: "Let's configure your standup schedule",
							},
							nil,
							slack.NewAccessory(
								slack.NewButtonBlockElement(
									"configure",
									"",
									&slack.TextBlockObject{
										Type: slack.PlainTextType,
										Text: "LFG",
									},
								),
							),
						),
					},
				}
				client.Ack(*event.Request, response)
			case socketmode.EventTypeEventsAPI:
				log.Printf("received events api event")
				eventsApiEvent, ok := event.Data.(slackevents.EventsAPIEvent)
				if !ok {
					log.Printf("ignored %+v", event)
					continue
				}
				client.Ack(*event.Request)
				switch eventsApiEvent.Type {
				case slackevents.CallbackEvent:
					switch ev := eventsApiEvent.InnerEvent.Data.(type) {
					case *slackevents.MessageEvent:
						log.Printf("received message in channel[%s]", ev.Channel)
						log.Printf("received message content: %s", ev.Text)
					default:
						log.Printf("received unknown message event type: %v", eventsApiEvent.InnerEvent.Data)
					}
				default:
					log.Printf("received unknown event type: %s", eventsApiEvent.Type)
				}
			}
		}
	}()

	client.Run()
}
