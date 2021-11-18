package truevieweventserver

import "github.com/go-resty/resty"
import "log"

type SlackMessage struct {
	Text string `json:"text"`
}

// curl -i -X POST -d 'payload={"text": "TrueVIEW ALERT:\nThis is an alert\nNext Line."}' https://slackers.freehive.io/hooks/irwbk3y8ciy1ukma398w7n566r

func RunWebhookSender(configuration Configuration, pluginRunResult PluginRunResult) {

	log.Println("About to send webhook", configuration.Name)
	webhookSender := configuration.AlertEventConfig.WebhookSender
	outgoingMessage := webhookSender.Message

	outgoingMessage += AppendPluginRunResultData(pluginRunResult)

	slackMessage := SlackMessage{outgoingMessage}
	slackJsonValue, slackJsonErr := StructToJson(slackMessage)

	if slackJsonErr != nil {
		log.Println("Error converting stuctToJson:", slackJsonErr)
	}

	response, respErr := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(slackJsonValue).
		SetResult(&AuthSuccess{}).
		Post(webhookSender.PostUrl)

	log.Println("Slack/Webhook messagesent", response, respErr)

	if respErr != nil {
		log.Println(respErr)
	}

	log.Println("Finished sending webhook", configuration.Name, configuration.AlertEventConfig.WebhookSender.PostUrl)

}
