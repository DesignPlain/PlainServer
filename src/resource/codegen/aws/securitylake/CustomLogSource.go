package securitylake

import types "libds/aws/types"

type CustomLogSource struct {
	// The configuration for the third-party custom source.
	Configuration types.Securitylake_CustomLogSourceConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// The Open Cybersecurity Schema Framework (OCSF) event classes which describes the type of data that the custom source will send to Security Lake.
	EventClasses []string `json:"eventClasses,omitempty" yaml:"eventClasses,omitempty"`

	/*
	   Specify the name for a third-party custom source.
	   This must be a Regionally unique value.
	   Has a maximum length of 20.
	*/
	SourceName string `json:"sourceName,omitempty" yaml:"sourceName,omitempty"`

	// Specify the source version for the third-party custom source, to limit log collection to a specific version of custom data source.
	SourceVersion string `json:"sourceVersion,omitempty" yaml:"sourceVersion,omitempty"`
}
