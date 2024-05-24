package ssm

type ServiceSetting struct {
	// ID of the service setting.
	SettingId string `json:"settingId,omitempty" yaml:"settingId,omitempty"`

	// Value of the service setting.
	SettingValue string `json:"settingValue,omitempty" yaml:"settingValue,omitempty"`
}
