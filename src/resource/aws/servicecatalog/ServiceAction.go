package servicecatalog

import types "DesignSphere_Server/src/resource/aws/types"

type ServiceAction struct {
	// Self-service action definition configuration block. Detailed below.
	Definition types.Servicecatalog_ServiceActionDefinition `json:"definition,omitempty" yaml:"definition,omitempty"`

	// Self-service action description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Self-service action name.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Language code. Valid values are `en` (English), `jp` (Japanese), and `zh` (Chinese). Default is `en`.
	AcceptLanguage string `json:"acceptLanguage,omitempty" yaml:"acceptLanguage,omitempty"`
}
