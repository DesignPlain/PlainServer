package dataloss

import types "DesignSphere_Server/src/resource/gcp/types"

type PreventionDeidentifyTemplate struct {
	/*
	   Configuration of the deidentify template
	   Structure is documented below.
	*/
	DeidentifyConfig types.Dataloss_PreventionDeidentifyTemplateDeidentifyConfig `json:"deidentifyConfig,omitempty" yaml:"deidentifyConfig,omitempty"`

	// A description of the template.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// User set display name of the template.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The parent of the template in any of the following formats:
	   - `projects/{{project}}`
	   - `projects/{{project}}/locations/{{location}}`
	   - `organizations/{{organization_id}}`
	   - `organizations/{{organization_id}}/locations/{{location}}`
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   The template id can contain uppercase and lowercase letters, numbers, and hyphens;
	   that is, it must match the regular expression: [a-zA-Z\d-_]+. The maximum length is
	   100 characters. Can be empty to allow the system to generate one.
	*/
	TemplateId string `json:"templateId,omitempty" yaml:"templateId,omitempty"`
}
