package clouddeploy

import types "libds/gcp/types"

type DeliveryPipeline struct {
	// When suspended, no new releases or rollouts can be created, but in-progress ones will complete.
	Suspended bool `json:"suspended,omitempty" yaml:"suspended,omitempty"`

	/*
	   User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.

	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// Description of the `DeliveryPipeline`. Max length is 255 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: - Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. - All characters must use UTF-8 encoding, and international characters are allowed. - Keys must start with a lowercase letter or international character. - Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Name of the `DeliveryPipeline`. Format is [a-z][a-z0-9\-]{0,62}.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
	SerialPipeline types.Clouddeploy_DeliveryPipelineSerialPipeline `json:"serialPipeline,omitempty" yaml:"serialPipeline,omitempty"`
}
