package types

type Datacatalog_TagTemplateFieldType struct {
	/*
	   Represents an enum type.
	   Exactly one of `primitive_type` or `enum_type` must be set
	   Structure is documented below.
	*/
	EnumType Datacatalog_TagTemplateFieldTypeEnumType `json:"enumType,omitempty" yaml:"enumType,omitempty"`

	/*
	   Represents primitive types - string, bool etc.
	   Exactly one of `primitive_type` or `enum_type` must be set
	   Possible values are: `DOUBLE`, `STRING`, `BOOL`, `TIMESTAMP`.
	*/
	PrimitiveType string `json:"primitiveType,omitempty" yaml:"primitiveType,omitempty"`
}
