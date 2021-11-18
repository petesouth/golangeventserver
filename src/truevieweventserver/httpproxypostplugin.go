package truevieweventserver

import "github.com/go-resty/resty"
import "log"
import "bytes"

const HTTP_PROXY_POST_PLUGIN string = "http_proxy_post_plugin"

/**
 * Main exported function.  Runs a configuration against values in an influx Database.
 */
func HttpProxyPostPlugin(configuration Configuration) error {

	urlsArrayLength := len(configuration.ProxyPluginConfig.GetUrls)
	var buffer bytes.Buffer

	buffer.WriteString("[")

	// Sping it all off using goroutines
	for i := 0; i < urlsArrayLength; i += 1 {
		restUrl := configuration.ProxyPluginConfig.GetUrls[i]

		resp, err := resty.R().Get(restUrl)

		if err != nil {

			return err
		}

		buffer.Write(resp.Body())

		if i != (urlsArrayLength - 1) {
			buffer.WriteString(",")
		}

	}

	buffer.WriteString("]")
	data := buffer.String()

	bodyStr := "{ \"name\": \"" + configuration.Name + "\", \"data\":" + data + " }"

	response, respErr := resty.R().
		SetHeader("Content-Type", "application/text").
		SetBody(bodyStr).
		SetResult(&AuthSuccess{}).
		Post(configuration.ProxyPluginConfig.PostUrl)

	log.Println("HttpProxy send message to post url:", configuration.ProxyPluginConfig.PostUrl, data, response, respErr)

	if respErr != nil {
		return respErr
	}

	return nil
}
