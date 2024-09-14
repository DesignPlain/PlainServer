package types

type Mq_getBrokerInstance struct {
	//
	Endpoints []string `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`

	//
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	//
	ConsoleUrl string `json:"consoleUrl,omitempty" yaml:"consoleUrl,omitempty"`
}
