package types

type Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinition struct {
	// Integer property.
	IntegerTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionIntegerTypeOptions `json:"integerTypeOptions,omitempty" yaml:"integerTypeOptions,omitempty"`

	// Whether the property is mandatory.
	IsRequired bool `json:"isRequired,omitempty" yaml:"isRequired,omitempty"`

	// The name of the metadata property.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The schema source information.
	   Structure is documented below.
	*/
	SchemaSources []Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionSchemaSource `json:"schemaSources,omitempty" yaml:"schemaSources,omitempty"`

	/*
	   Nested structured data property.
	   Structure is documented below.
	*/
	PropertyTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionPropertyTypeOptions `json:"propertyTypeOptions,omitempty" yaml:"propertyTypeOptions,omitempty"`

	// Timestamp property. Not supported by CMEK compliant deployment.
	TimestampTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionTimestampTypeOptions `json:"timestampTypeOptions,omitempty" yaml:"timestampTypeOptions,omitempty"`

	// Date time property. Not supported by CMEK compliant deployment.
	DateTimeTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionDateTimeTypeOptions `json:"dateTimeTypeOptions,omitempty" yaml:"dateTimeTypeOptions,omitempty"`

	// The display-name for the property, used for front-end.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Enum/categorical property.
	   Structure is documented below.
	*/
	EnumTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionEnumTypeOptions `json:"enumTypeOptions,omitempty" yaml:"enumTypeOptions,omitempty"`

	// Whether the property can have multiple values.
	IsRepeatable bool `json:"isRepeatable,omitempty" yaml:"isRepeatable,omitempty"`

	// Indicates that the property should be included in a global search.
	IsSearchable bool `json:"isSearchable,omitempty" yaml:"isSearchable,omitempty"`

	// Float property.
	FloatTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionFloatTypeOptions `json:"floatTypeOptions,omitempty" yaml:"floatTypeOptions,omitempty"`

	// Whether the property can be filtered. If this is a sub-property, all the parent properties must be marked filterable.
	IsFilterable bool `json:"isFilterable,omitempty" yaml:"isFilterable,omitempty"`

	/*
	   Stores the retrieval importance.
	   Possible values are: `HIGHEST`, `HIGHER`, `HIGH`, `MEDIUM`, `LOW`, `LOWEST`.
	*/
	RetrievalImportance string `json:"retrievalImportance,omitempty" yaml:"retrievalImportance,omitempty"`

	// Whether the property is user supplied metadata.
	IsMetadata bool `json:"isMetadata,omitempty" yaml:"isMetadata,omitempty"`

	// Map property.
	MapTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionMapTypeOptions `json:"mapTypeOptions,omitempty" yaml:"mapTypeOptions,omitempty"`

	// Text property.
	TextTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionTextTypeOptions `json:"textTypeOptions,omitempty" yaml:"textTypeOptions,omitempty"`
}
