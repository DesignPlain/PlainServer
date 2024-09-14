package types

type Fis_ExperimentTemplateTarget struct {
	/*
	   The resource type parameters.

	   > --NOTE:-- The `target` configuration block requires either `resource_arns` or `resource_tag`.
	*/
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Set of ARNs of the resources to target with an action. Conflicts with `resource_tag`.
	ResourceArns []string `json:"resourceArns,omitempty" yaml:"resourceArns,omitempty"`

	// Tag(s) the resources need to have to be considered a valid target for an action. Conflicts with `resource_arns`. See below.
	ResourceTags []Fis_ExperimentTemplateTargetResourceTag `json:"resourceTags,omitempty" yaml:"resourceTags,omitempty"`

	// AWS resource type. The resource type must be supported for the specified action. To find out what resource types are supported, see [Targets for AWS FIS](https://docs.aws.amazon.com/fis/latest/userguide/targets.html#resource-types).
	ResourceType string `json:"resourceType,omitempty" yaml:"resourceType,omitempty"`

	// Scopes the identified resources. Valid values are `ALL` (all identified resources), `COUNT(n)` (randomly select `n` of the identified resources), `PERCENT(n)` (randomly select `n` percent of the identified resources).
	SelectionMode string `json:"selectionMode,omitempty" yaml:"selectionMode,omitempty"`

	// Filter(s) for the target. Filters can be used to select resources based on specific attributes returned by the respective describe action of the resource type. For more information, see [Targets for AWS FIS](https://docs.aws.amazon.com/fis/latest/userguide/targets.html#target-filters). See below.
	Filters []Fis_ExperimentTemplateTargetFilter `json:"filters,omitempty" yaml:"filters,omitempty"`

	// Friendly name given to the target.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
