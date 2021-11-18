package truevieweventserver

import "strings"
import "log"

const PER_HOST_HTTP_PROXY_POST_PLUGIN string = "per_host_http_proxy_plugin"

/**
 * Main exported function.  Runs a configuratiosn against values in an influx Database.
 */
func PerHostHttpProxyPostPlugin(configuration Configuration) error {

	hostsInfluxRestUrl := configuration.PerHostProxyPluginConfig.InfluxHostsUrl
	influxResp, influxRespError := influxRestGet(hostsInfluxRestUrl)

	if influxRespError != nil {
		return influxRespError
	}

	influxRespSeriesLen := len(influxResp.Results)
	var hostNames []string

	for iResult := 0; iResult < influxRespSeriesLen; iResult += 1 {

		resultSeries := influxResp.Results[iResult].Series
		resultSeriesLen := len(resultSeries)

		for iResultSeries := 0; iResultSeries < resultSeriesLen; iResultSeries += 1 {

			values := resultSeries[iResultSeries].Values
			valuesLen := len(values)

			for iValues := 0; iValues < valuesLen; iValues += 1 {

				stringSplitValues := strings.Split(values[iValues][0].(string), ",")

				host := strings.Split(stringSplitValues[1], "=")[1]
				hostNames = append(hostNames, host)
			}

		}

	}

	log.Println("------->Found hostnames: %s", hostNames)

	hostNamesLen := len(hostNames)

	for iHostNames := 0; iHostNames < hostNamesLen; iHostNames += 1 {
		hostName := hostNames[iHostNames]
		name := hostName + "_" + configuration.Name

		hostConfiguration := Configuration{}
		hostConfiguration.Name = name

		hostConfiguration.ProxyPluginConfig.PostUrl = configuration.PerHostProxyPluginConfig.PostUrl

		getUrlsLen := len(configuration.PerHostProxyPluginConfig.GetUrls)
		hostConfiguration.ProxyPluginConfig.GetUrls = []string{}

		for iGetUrls := 0; iGetUrls < getUrlsLen; iGetUrls += 1 {
			getUrl := configuration.PerHostProxyPluginConfig.GetUrls[iGetUrls]
			getUrl = strings.Replace(getUrl, "{{host_name}}", hostName, -1)
			hostConfiguration.ProxyPluginConfig.GetUrls = append(hostConfiguration.ProxyPluginConfig.GetUrls, getUrl)
		}

		log.Println("======Running a HttpProxyPostPlugin for urls:", hostConfiguration)

		err := HttpProxyPostPlugin(hostConfiguration)

		if err != nil {
			log.Println("Error calling HttpProxyPostPlugin from PerHostHttpProxyPostPlugin", err)
		}

	}

	return nil
}
