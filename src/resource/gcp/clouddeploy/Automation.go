package clouddeploy

import types "DesignSphere_Server/src/resource/gcp/types"

type Automation struct {
	// The delivery_pipeline for the resource
	DeliveryPipeline string `json:"deliveryPipeline,omitempty" yaml:"deliveryPipeline,omitempty"`

	// Optional. Description of the `Automation`. Max length is 255 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Required. List of Automation rules associated with the Automation resource. Must have at least one rule and limited to 250 rules per Delivery Pipeline. Note: the order of the rules here is not the same as the order of execution.
	   Structure is documented below.
	*/
	Rules []types.Clouddeploy_AutomationRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// Required. Email address of the user-managed IAM service account that creates Cloud Deploy release and rollout resources.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	/*
	   Optional. User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. Annotations must meet the following constraints: - Annotations are key/value pairs. - Valid annotation keys have two segments: an optional prefix and name, separated by a slash (`/`). - The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between. - The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots(`.`), not longer than 253 characters in total, followed by a slash (`/`). See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/#syntax-and-character-set for more details.
	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Name of the `Automation`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Required. Selected resources to which the automation will be applied.
	   Structure is documented below.
	*/
	Selector types.Clouddeploy_AutomationSelector `json:"selector,omitempty" yaml:"selector,omitempty"`

	// Optional. When Suspended, automation is deactivated from execution.
	Suspended bool `json:"suspended,omitempty" yaml:"suspended,omitempty"`

	/*
	   Optional. Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: - Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. - All characters must use UTF-8 encoding, and international characters are allowed. - Keys must start with a lowercase letter or international character. - Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 63 characters.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
