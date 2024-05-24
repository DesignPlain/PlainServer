package types

type Fis_ExperimentTemplateActionParameter struct {
	// Parameter name.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   Parameter value.

	   For a list of parameters supported by each action, see [AWS FIS actions reference](https://docs.aws.amazon.com/fis/latest/userguide/fis-actions-reference.html).
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
