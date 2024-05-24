package types

type Mq_BrokerConfiguration struct {
	// The Configuration ID.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Revision of the Configuration.
	Revision int `json:"revision,omitempty" yaml:"revision,omitempty"`
}
