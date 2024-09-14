package types

type Glue_MLTransformSchema struct {
	// The type of data in the column.
	DataType string `json:"dataType,omitempty" yaml:"dataType,omitempty"`

	// The name you assign to this ML Transform. It must be unique in your account.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
