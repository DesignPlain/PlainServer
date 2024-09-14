package ssm

type ServiceSetting struct {
	// Value of the service setting.
	SettingValue string `json:"settingValue,omitempty" yaml:"settingValue,omitempty"`

	// ID of the service setting.
	SettingId string `json:"settingId,omitempty" yaml:"settingId,omitempty"`
}
