package truevieweventserver

import "strconv"
import "time"
import "strings"
import "log"

type PluginRunResult struct {
	Config            Configuration
	StartThresholdMet bool
	EndThresholdMet   bool
	FoundActualValue  float64
}

func RunConfiguration(configuration Configuration) {

	//sentMessage := false
	configDisplayString, configDisplayStringErr := ConfigToDisplayStr(configuration)

	if configDisplayStringErr != nil {
		log.Println("Error occured during config load", configDisplayStringErr)
		return
	}

	pollFrequency, freqErr := strconv.ParseInt(configuration.PollFrequency, 10, 64)

	if freqErr != nil || pollFrequency < 1 {

		log.Println("(", configuration.Name, ")Abandoning configuraiton due to invalid pollFrequency/messageSentWait config", configDisplayString)
		return
	}

	log.Println("=== Started configuration:", configDisplayString)

	for {

		time.Sleep(time.Duration(pollFrequency) * time.Millisecond)

		if strings.Compare(configuration.RunnerPlugin, INFLUX_REST_EVENT_PLUGIN) == 0 {
			RunInfluxEventPlugin(configuration)

		} else if strings.Compare(configuration.RunnerPlugin, HTTP_PROXY_POST_PLUGIN) == 0 {
			HttpProxyPostPlugin(configuration)

		} else if strings.Compare(configuration.RunnerPlugin, PER_HOST_HTTP_PROXY_POST_PLUGIN) == 0 {
			PerHostHttpProxyPostPlugin(configuration)

		} else if strings.Compare(configuration.RunnerPlugin, TRUEVIEW_CONSUMER_PLUGIN) == 0 {
			TrueviewConsumerPlugin(configuration)

		} else if strings.Compare(configuration.RunnerPlugin, EMAIL_PROXY_POST_PLUGIN) == 0 {
			EmailProxyPostPlugin(configuration)

		} else {
			log.Println("(", configuration.Name, ")Unknown plugin type abandoning configuraiton", configDisplayString)
			return
		}

	}
}
