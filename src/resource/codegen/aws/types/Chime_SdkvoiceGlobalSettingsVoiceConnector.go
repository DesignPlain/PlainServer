package types

type Chime_SdkvoiceGlobalSettingsVoiceConnector struct {
	// The S3 bucket that stores the Voice Connector's call detail records.
	CdrBucket string `json:"cdrBucket,omitempty" yaml:"cdrBucket,omitempty"`
}
