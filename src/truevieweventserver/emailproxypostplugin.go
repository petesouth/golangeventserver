package truevieweventserver

import (
	"bytes"
	"encoding/base64"
	"github.com/mhale/smtpd"
	"io/ioutil"
	"log"
	"mime/quotedprintable"
	"net"
	"net/mail"
)

const EMAIL_PROXY_POST_PLUGIN string = "email_proxy_post_plugin"

func EmailProxyEmailHandler(origin net.Addr, from string, to []string, data []byte) {
	msg, _ := mail.ReadMessage(bytes.NewReader(data))
	subject := msg.Header.Get("Subject")

	test := quotedprintable.NewReader(msg.Body)
	body, _ := ioutil.ReadAll(test) // body now contains the decoded body
	bodyTest := string(body[:])
	bodyTestDecode, _ := base64.StdEncoding.DecodeString(bodyTest)

	log.Printf("Received mail from %s for %s with subject %s, body: %s", from, to[0], subject, bodyTestDecode)
}

/**
 * Main exported function.  Runs a configuration against values in an influx Database.
 */
func EmailProxyPostPlugin(configuration Configuration) error {

	emailServerUrl := configuration.EmailProxyPluginConfig.EmailServerUrl
	log.Printf("Running emial server address/port: %s", emailServerUrl)

	smtpd.ListenAndServe(emailServerUrl, EmailProxyEmailHandler, "EmailProxyPostPlugin", "")

	return nil
}
