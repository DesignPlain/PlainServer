package types

type Mskconnect_ConnectorWorkerConfiguration struct {
	// The Amazon Resource Name (ARN) of the worker configuration.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The revision of the worker configuration.
	Revision int `json:"revision,omitempty" yaml:"revision,omitempty"`
}
