package truevieweventserver

import "github.com/go-resty/resty"
import "strconv"
import "strings"
import "log"

import "time"

const INFLUX_REST_EVENT_PLUGIN string = "influxrest_event_alert"

/*  Json Looks like:

	{ "results":[
					{"series": [{"name":"cpu_value",
					             "columns":["time","value"],
					             "values":[["2016-05-15T14:16:55Z",2.0258379e+07]]
					             }]
					 }
		  	    ]
    }
*/

type SeriesObject struct {
	Name    string          `json:"name"`
	Columns []string        `json:"columns"`
	Values  [][]interface{} `json:"values"`
}

type Series struct {
	Series []SeriesObject `json:"series"`
}

type InfluxResponse struct {
	Results []Series `json:"results"`
}

type InfluxNResponseNSQContainer struct {
	Name string           `json:"name"`
	Data []InfluxResponse `json:"data"`
}

func influxRestGet(url string) (*InfluxResponse, error) {

	resp, err := resty.R().Get(url)
	if err != nil {

		return nil, err
	}

	influxResponse := InfluxResponse{}
	err = JsonStringStruct(resp.String(), &influxResponse)

	return &influxResponse, err
}

func thresholdCompare(value float64, compareValue float64, operand string) bool {
	returnVal := false

	if strings.Compare(operand, OPERATOR_EQUALS) == 0 && value == compareValue {
		returnVal = true
	} else if strings.Compare(operand, OPERATOR_GREATER_THEN) == 0 && value > compareValue {
		returnVal = true
	} else if strings.Compare(operand, OPERATOR_GREATER_THEN_EQUALS) == 0 && value >= compareValue {
		returnVal = true
	} else if strings.Compare(operand, OPERATOR_LESS_THEN) == 0 && value < compareValue {
		returnVal = true
	} else if strings.Compare(operand, OPERATOR_LESS_THEN_EQUALS) == 0 && value <= compareValue {
		returnVal = true
	}

	return returnVal

}

func calculateThresholds(value float64, configuration Configuration) (bool, bool) {

	startValue, _ := strconv.ParseFloat(configuration.AlertEventConfig.StartValue, 64)
	endValue, _ := strconv.ParseFloat(configuration.AlertEventConfig.EndValue, 64)
	foundThresholdStart := thresholdCompare(value, startValue, configuration.AlertEventConfig.StartOperand)
	foundThresholdEnd := thresholdCompare(value, endValue, configuration.AlertEventConfig.EndOperand)

	return foundThresholdStart, foundThresholdEnd
}

/**
 * Main exported function.  Runs a configuration against values in an influx Database.
 */
func RunInfluxEventPlugin(configuration Configuration) error {

	restUrl := configuration.AlertEventConfig.RestUrl + "?" + configuration.AlertEventConfig.RestQueryString
	influxResponse, err := influxRestGet(restUrl)
	messageSentWait, msgErr := strconv.ParseInt(configuration.AlertEventConfig.MessageSentWait, 10, 64)

	if err != nil ||
		msgErr != nil ||
		influxResponse.Results == nil || len(influxResponse.Results) < 1 ||
		influxResponse.Results[0].Series == nil || len(influxResponse.Results[0].Series) < 1 ||
		influxResponse.Results[0].Series[0].Values == nil || len(influxResponse.Results[0].Series[0].Values) < 1 ||
		influxResponse.Results[0].Series[0].Values[0] == nil || len(influxResponse.Results[0].Series[0].Values[0]) < 2 {
		log.Println("err when calling influx", err, influxResponse)
		return err
	}

	seriesObject := influxResponse.Results[0].Series[0]
	value := seriesObject.Values[0][1].(float64)
	foundThresholdStart, foundThresholdEnd := calculateThresholds(value, configuration)
	pluginRunResult := PluginRunResult{configuration, foundThresholdStart, foundThresholdEnd, value}

	if pluginRunResult.EndThresholdMet || pluginRunResult.StartThresholdMet {
		log.Println("Message to be Sent(", configuration.Name, ") pluginRunResult", pluginRunResult)

		if configuration.AlertEventConfig.EmailSender.Active == true {
			RunEmailSender(configuration, pluginRunResult)
		}

		if configuration.AlertEventConfig.WebhookSender.Active == true {
			RunWebhookSender(configuration, pluginRunResult)
		}

		if configuration.AlertEventConfig.TrueviewApiSender.Active == true {
			RunTrueviewApiSender(configuration, pluginRunResult)
		}

		time.Sleep(time.Duration(messageSentWait) * time.Millisecond)
		log.Println("(", configuration.Name, ") Messasge wait time complete - Restarting check")

	}

	return nil
}
