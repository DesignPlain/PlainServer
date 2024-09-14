package types

type Guardduty_OrganizationConfigurationDatasourcesS3Logs struct {
	// Set to `true` if you want S3 data event logs to be automatically enabled for new members of the organization. Default: `false`
	AutoEnable bool `json:"autoEnable,omitempty" yaml:"autoEnable,omitempty"`
}
