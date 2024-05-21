package clouddeploy

import types "DesignSphere_Server/src/resource/gcp/types"

type CustomTargetType struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	/*
	   Configures render and deploy for the `CustomTargetType` using Skaffold custom actions.
	   Structure is documented below.
	*/
	CustomActions types.Clouddeploy_CustomTargetTypeCustomActions `json:"customActions,omitempty" yaml:"customActions,omitempty"`

	// Description of the `CustomTargetType`. Max length is 255 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: - Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. - All characters must use UTF-8 encoding, and international characters are allowed. - Keys must start with a lowercase letter or international character. - Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The location of the source.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Name of the `CustomTargetType`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
