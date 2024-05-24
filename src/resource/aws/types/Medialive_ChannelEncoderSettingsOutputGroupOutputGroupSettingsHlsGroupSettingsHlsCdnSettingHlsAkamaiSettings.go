package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsHlsGroupSettingsHlsCdnSettingHlsAkamaiSettings struct {
	// Number of seconds to wait until a restart is initiated.
	RestartDelay int `json:"restartDelay,omitempty" yaml:"restartDelay,omitempty"`

	//
	Salt string `json:"salt,omitempty" yaml:"salt,omitempty"`

	//
	Token string `json:"token,omitempty" yaml:"token,omitempty"`

	// Number of seconds to wait before retrying connection to the flash media server if the connection is lost.
	ConnectionRetryInterval int `json:"connectionRetryInterval,omitempty" yaml:"connectionRetryInterval,omitempty"`

	//
	FilecacheDuration int `json:"filecacheDuration,omitempty" yaml:"filecacheDuration,omitempty"`

	//
	HttpTransferMode string `json:"httpTransferMode,omitempty" yaml:"httpTransferMode,omitempty"`

	// Number of retry attempts.
	NumRetries int `json:"numRetries,omitempty" yaml:"numRetries,omitempty"`
}
