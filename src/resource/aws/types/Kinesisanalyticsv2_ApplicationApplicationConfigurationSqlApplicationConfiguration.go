package types

type Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfiguration struct {
	// The reference data source used by the application.
	ReferenceDataSource Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSource `json:"referenceDataSource,omitempty" yaml:"referenceDataSource,omitempty"`

	// The input stream used by the application.
	Input Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationInput `json:"input,omitempty" yaml:"input,omitempty"`

	// The destination streams used by the application.
	Outputs []Kinesisanalyticsv2_ApplicationApplicationConfigurationSqlApplicationConfigurationOutput `json:"outputs,omitempty" yaml:"outputs,omitempty"`
}
