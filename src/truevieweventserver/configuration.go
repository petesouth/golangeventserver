package truevieweventserver

const OPERATOR_EQUALS string = "=="
const OPERATOR_GREATER_THEN string = ">"
const OPERATOR_GREATER_THEN_EQUALS string = ">="
const OPERATOR_LESS_THEN string = "<"
const OPERATOR_LESS_THEN_EQUALS string = "<="

type TrueviewApiPlugin struct {
	Active   bool   `json:"active"`
	PostUrl  string `json:"postUrl"`
	Hostname string `json:"hostname"`
	Message  string `json:"message"`
}

type PerHostProxyPlugin struct {
	InfluxHostsUrl string   `json:"influxHostsUrl"`
	PostUrl        string   `json:"postUrl"`
	GetUrls        []string `json:"getUrls"`
}

type ProxyPlugin struct {
	PostUrl string   `json:"postUrl"`
	GetUrls []string `json:"getUrls"`
}

type EmailProxyPlugin struct {
	PostUrl        string `json:"postUrl"`
	EmailServerUrl string `json:"emailServerUrl"`
}

type ConsumerPlugin struct {
	NsqdTcpAddress         string   `json:"nsqdTcpAddress"`
	Topic                  string   `json:"topic"`
	OutputDir              []string `json:"outputDir"`
	MessageBatchRemoveSize int      `json:"messageBatchRemoveSize"`
}

type WebhookPlugin struct {
	Active  bool   `json:"active"`
	PostUrl string `json:"postUrl"`
	Message string `json:"message"`
}

type EmailPlugin struct {
	Active   bool     `json:"active"`
	Identity string   `json:"identity"`
	Username string   `json:"username"`
	Sender   string   `json:"sender"`
	Password string   `json:"password"`
	Host     string   `json:"host"`
	Port     int      `json:"port"`
	Targets  []string `json:"targets"`
	Message  string   `json:"message"`
	Body     string   `json:"body"`
}

type AlertEvent struct {
	MessageSentWait string `json:"messageSentWait"`
	StartValue      string `json:"startValue"`
	StartOperand    string `json:"startOperand"`
	EndValue        string `json:"endValue"`
	EndOperand      string `json:"endOperand"`
	RestUrl         string `json:"restUrl"`
	RestQueryString string `json:"restQueryString"`

	EmailSender       EmailPlugin       `json:"emailSender"`
	WebhookSender     WebhookPlugin     `json:"webhookSender"`
	TrueviewApiSender TrueviewApiPlugin `json:"trueviewApiSender"`
}

type Configuration struct {
	RunnerPlugin             string             `json:"runnerPlugin"`
	Name                     string             `json:"name"`
	PollFrequency            string             `json:"pollFrequency"`
	AlertEventConfig         AlertEvent         `json:"alertEventConfig"`
	ProxyPluginConfig        ProxyPlugin        `json:"proxyPluginConfig"`
	EmailProxyPluginConfig   EmailProxyPlugin   `json:"emailProxyPlufingConfig"`
	ConsumerPluginConfig     ConsumerPlugin     `json:"consumerPluginConfig"`
	PerHostProxyPluginConfig PerHostProxyPlugin `json:"perHostProxyPluginConfig"`
}

func ConfigToDisplayStr(config Configuration) (string, error) {
	return StructToJson(config)
}
