package truevieweventserver

import "github.com/go-resty/resty"
import "log"
import "strconv"
import "time"

type TrueviewMessage struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Time     string `json:"time"`
}

//  ```message api: POST /api/v1/message/add?api_token=
//  (eg: POST http://10.248.1.226/api/v1/message/add?api_token=aa3c5Wucvq1Ru6v1k3foxUbPqgeK9Pv5tprJazq98sB2I95Zm9bi6t64aS33)
//  data format:
//  {
//  "name”:”alerts info”,
//  ”hostname":"stinger_freehive_io",
//  "time":"2010-10-10 10:00:00"
//  }
//  where to get api_token: check table users```
// current token in trueview-demo mysql/trueview is
// http://trueview-demo.freehive.io/api/v1/message/add?api_token=20ZyPBCIDPQjVkXajnyuiFslW9WJPX7WfyQTYBVzVotu9Q6fzLK5jsLnhabJ

func RunTrueviewApiSender(configuration Configuration, pluginRunResult PluginRunResult) {

	log.Println("About to send webhook", configuration.Name)
	trueviewApiSender := configuration.AlertEventConfig.TrueviewApiSender
	outgoingMessage := trueviewApiSender.Message
	outgoingMessage += " value was: " + strconv.FormatFloat(pluginRunResult.FoundActualValue, 'f', 6, 64)
	t := time.Now()
	timeNowString := t.String()

	trueviewMessage := TrueviewMessage{outgoingMessage, trueviewApiSender.Hostname, timeNowString}
	trueviewJsonValue, trueviewJsonErr := StructToJson(trueviewMessage)

	if trueviewJsonErr != nil {
		log.Println("Error converting stuctToJson:", trueviewJsonErr)
	}

	response, respErr := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(trueviewJsonValue).
		SetResult(&AuthSuccess{}).
		Post(trueviewApiSender.PostUrl)

	log.Println("trueview/Webhook message sent", response, respErr)

	if respErr != nil {
		log.Println(respErr)
	}

	log.Println("Finished sending trueview/Webhook message sent", configuration.Name, trueviewApiSender.PostUrl)

}
