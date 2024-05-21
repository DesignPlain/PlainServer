package dataloss

import types "DesignSphere_Server/src/resource/gcp/types"

type PreventionStoredInfoType struct {
	/*
	   The parent of the info type in any of the following formats:
	   - `projects/{{project}}`
	   - `projects/{{project}}/locations/{{location}}`
	   - `organizations/{{organization_id}}`
	   - `organizations/{{organization_id}}/locations/{{location}}`


	   - - -
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   Regular expression which defines the rule.
	   Structure is documented below.
	*/
	Regex types.Dataloss_PreventionStoredInfoTypeRegex `json:"regex,omitempty" yaml:"regex,omitempty"`

	/*
	   The storedInfoType ID can contain uppercase and lowercase letters, numbers, and hyphens;
	   that is, it must match the regular expression: [a-zA-Z\d-_]+. The maximum length is 100
	   characters. Can be empty to allow the system to generate one.
	*/
	StoredInfoTypeId string `json:"storedInfoTypeId,omitempty" yaml:"storedInfoTypeId,omitempty"`

	// A description of the info type.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Dictionary which defines the rule.
	   Structure is documented below.
	*/
	Dictionary types.Dataloss_PreventionStoredInfoTypeDictionary `json:"dictionary,omitempty" yaml:"dictionary,omitempty"`

	// User set display name of the info type.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Dictionary which defines the rule.
	   Structure is documented below.
	*/
	LargeCustomDictionary types.Dataloss_PreventionStoredInfoTypeLargeCustomDictionary `json:"largeCustomDictionary,omitempty" yaml:"largeCustomDictionary,omitempty"`
}
