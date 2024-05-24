package types

type Mskconnect_ConnectorPluginCustomPlugin struct {
	// The Amazon Resource Name (ARN) of the custom plugin.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The revision of the custom plugin.
	Revision int `json:"revision,omitempty" yaml:"revision,omitempty"`
}
