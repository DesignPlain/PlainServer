package types

type Auditmanager_getControlControlMappingSource struct {
	//
	SourceName string `json:"sourceName,omitempty" yaml:"sourceName,omitempty"`

	//
	SourceSetUpOption string `json:"sourceSetUpOption,omitempty" yaml:"sourceSetUpOption,omitempty"`

	//
	SourceType string `json:"sourceType,omitempty" yaml:"sourceType,omitempty"`

	//
	TroubleshootingText string `json:"troubleshootingText,omitempty" yaml:"troubleshootingText,omitempty"`

	//
	SourceDescription string `json:"sourceDescription,omitempty" yaml:"sourceDescription,omitempty"`

	//
	SourceFrequency string `json:"sourceFrequency,omitempty" yaml:"sourceFrequency,omitempty"`

	//
	SourceId string `json:"sourceId,omitempty" yaml:"sourceId,omitempty"`

	//
	SourceKeyword Auditmanager_getControlControlMappingSourceSourceKeyword `json:"sourceKeyword,omitempty" yaml:"sourceKeyword,omitempty"`
}
