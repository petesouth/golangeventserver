package truevieweventserver

import "net/smtp"

//import "encoding/base64"
import "log"
import "strconv"

//import "bytes"
import "net/mail"
import "strings"
import "fmt"

type AuthSuccess struct {
	ID, Message string
}

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func AppendPluginRunResultData(pluginRunResult PluginRunResult) string {
	message := ""
	floatValueString := strconv.FormatFloat(pluginRunResult.FoundActualValue, 'f', 6, 64)
	message += "\n\n Condition was " + floatValueString + " " + pluginRunResult.Config.AlertEventConfig.StartOperand + " " + pluginRunResult.Config.AlertEventConfig.StartValue

	if len(pluginRunResult.Config.AlertEventConfig.EndOperand) > 0 {
		message += "\n               " + floatValueString + " " + pluginRunResult.Config.AlertEventConfig.EndOperand + " " + pluginRunResult.Config.AlertEventConfig.EndValue
	}
	message += "\n\n Query was:" + pluginRunResult.Config.AlertEventConfig.RestQueryString

	return message
}

// curl -i -X POST -d 'payload={"text": "TrueVIEW ALERT:\nThis is an alert\nNext Line."}' https://slackers.freehive.io/hooks/irwbk3y8ciy1ukma398w7n566r

func RunEmailSender(configuration Configuration, pluginRunResult PluginRunResult) {

	log.Println("About to send email", configuration.Name)

	emailSender := configuration.AlertEventConfig.EmailSender

	auth := smtp.PlainAuth(emailSender.Identity,
		emailSender.Username,
		emailSender.Password,
		emailSender.Host)

	log.Println("emailServer", emailSender.Host)

	for i := 0; i < len(emailSender.Targets); i += 1 {

		targetAddr := emailSender.Targets[i]
		emailServer := emailSender.Host + ":" + strconv.Itoa(emailSender.Port)

		header := make(map[string]string)
		header["From"] = emailSender.Sender
		header["To"] = targetAddr
		header["Subject"] = encodeRFC2047(emailSender.Message)
		header["MIME-Version"] = "1.0"
		header["Content-Type"] = "text/plain; charset=\"utf-8\""
		header["Content-Transfer-Encoding"] = "bas64"

		message := ""
		for k, v := range header {
			message += fmt.Sprintf("%s: %s\r\n", k, v)
		}
		message += "\r\n" + emailSender.Body
		message += AppendPluginRunResultData(pluginRunResult)

		log.Println("Email body:", message)
		// Connect to the server, authenticate, set the sender and recipient,
		// and send the email all in one step.
		err := smtp.SendMail(
			emailServer,
			auth,
			emailSender.Sender,
			[]string{targetAddr},
			[]byte(message))

		if err != nil {
			log.Println("Error sending email:", err)
		}

		log.Println("Finished sending email", configuration.Name)

	}

}
