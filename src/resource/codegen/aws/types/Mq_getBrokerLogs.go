package types

type Mq_getBrokerLogs struct {
	//
	Audit bool `json:"audit,omitempty" yaml:"audit,omitempty"`

	//
	General bool `json:"general,omitempty" yaml:"general,omitempty"`
}
