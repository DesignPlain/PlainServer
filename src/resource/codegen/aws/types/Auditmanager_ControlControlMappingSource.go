package types

type Auditmanager_ControlControlMappingSource struct {
	// The setup option for the data source. This option reflects if the evidence collection is automated or manual. Valid values are `System_Controls_Mapping` (automated) and `Procedural_Controls_Mapping` (manual).
	SourceSetUpOption string `json:"sourceSetUpOption,omitempty" yaml:"sourceSetUpOption,omitempty"`

	/*
	   Type of data source for evidence collection. If `source_set_up_option` is manual, the only valid value is `MANUAL`. If `source_set_up_option` is automated, valid values are `AWS_Cloudtrail`, `AWS_Config`, `AWS_Security_Hub`, or `AWS_API_Call`.

	   The following arguments are optional:
	*/
	SourceType string `json:"sourceType,omitempty" yaml:"sourceType,omitempty"`

	// Instructions for troubleshooting the control.
	TroubleshootingText string `json:"troubleshootingText,omitempty" yaml:"troubleshootingText,omitempty"`

	// Description of the source.
	SourceDescription string `json:"sourceDescription,omitempty" yaml:"sourceDescription,omitempty"`

	// Frequency of evidence collection. Valid values are `DAILY`, `WEEKLY`, or `MONTHLY`.
	SourceFrequency string `json:"sourceFrequency,omitempty" yaml:"sourceFrequency,omitempty"`

	//
	SourceId string `json:"sourceId,omitempty" yaml:"sourceId,omitempty"`

	// The keyword to search for in CloudTrail logs, Config rules, Security Hub checks, and Amazon Web Services API names. See `source_keyword` below.
	SourceKeyword Auditmanager_ControlControlMappingSourceSourceKeyword `json:"sourceKeyword,omitempty" yaml:"sourceKeyword,omitempty"`

	// Name of the source.
	SourceName string `json:"sourceName,omitempty" yaml:"sourceName,omitempty"`
}
