package types

type Mq_BrokerLogs struct {
	// Enables audit logging. Auditing is only possible for `engine_type` of `ActiveMQ`. User management action made using JMX or the ActiveMQ Web Console is logged. Defaults to `false`.
	Audit bool `json:"audit,omitempty" yaml:"audit,omitempty"`

	// Enables general logging via CloudWatch. Defaults to `false`.
	General bool `json:"general,omitempty" yaml:"general,omitempty"`
}
