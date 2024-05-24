package types

type Mq_getBrokerInstance struct {
	//
	ConsoleUrl string `json:"consoleUrl,omitempty" yaml:"consoleUrl,omitempty"`

	//
	Endpoints []string `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`

	//
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
}
