package types

type Datacatalog_TagField struct {
	/*
	   Holds the value for a tag field with enum type. This value must be one of the allowed values in the definition of this enum.

	   - - -
	*/
	EnumValue string `json:"enumValue,omitempty" yaml:"enumValue,omitempty"`

	// The identifier for this object. Format specified above.
	FieldName string `json:"fieldName,omitempty" yaml:"fieldName,omitempty"`

	/*
	   (Output)
	   The order of this field with respect to other fields in this tag. For example, a higher value can indicate
	   a more important field. The value can be negative. Multiple fields can have the same order, and field orders
	   within a tag do not have to be sequential.
	*/
	Order int `json:"order,omitempty" yaml:"order,omitempty"`

	// Holds the value for a tag field with string type.
	StringValue string `json:"stringValue,omitempty" yaml:"stringValue,omitempty"`

	// Holds the value for a tag field with timestamp type.
	TimestampValue string `json:"timestampValue,omitempty" yaml:"timestampValue,omitempty"`

	// Holds the value for a tag field with boolean type.
	BoolValue bool `json:"boolValue,omitempty" yaml:"boolValue,omitempty"`

	/*
	   (Output)
	   The display name of this field
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Holds the value for a tag field with double type.
	DoubleValue float64 `json:"doubleValue,omitempty" yaml:"doubleValue,omitempty"`
}
