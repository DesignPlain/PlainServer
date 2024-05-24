package types

type Fis_ExperimentTemplateTargetFilter struct {
	// Attribute path for the filter.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	/*
	   Set of attribute values for the filter.

	   > --NOTE:-- Values specified in a `filter` are joined with an `OR` clause, while values across multiple `filter` blocks are joined with an `AND` clause. For more information, see [Targets for AWS FIS](https://docs.aws.amazon.com/fis/latest/userguide/targets.html#target-filters).
	*/
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
