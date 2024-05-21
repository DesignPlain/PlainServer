package types

type Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionPropertyTypeOptionsPropertyDefinition struct {
	// Whether the property can have multiple values.
	IsRepeatable bool `json:"isRepeatable,omitempty" yaml:"isRepeatable,omitempty"`

	// The name of the metadata property.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The schema source information.
	   Structure is documented below.
	*/
	SchemaSources []Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionPropertyTypeOptionsPropertyDefinitionSchemaSource `json:"schemaSources,omitempty" yaml:"schemaSources,omitempty"`

	// Whether the property can be filtered. If this is a sub-property, all the parent properties must be marked filterable.
	IsFilterable bool `json:"isFilterable,omitempty" yaml:"isFilterable,omitempty"`

	// Whether the property is user supplied metadata.
	IsMetadata bool `json:"isMetadata,omitempty" yaml:"isMetadata,omitempty"`

	// Indicates that the property should be included in a global search.
	IsSearchable bool `json:"isSearchable,omitempty" yaml:"isSearchable,omitempty"`

	// Map property.
	MapTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionPropertyTypeOptionsPropertyDefinitionMapTypeOptions `json:"mapTypeOptions,omitempty" yaml:"mapTypeOptions,omitempty"`

	// The display-name for the property, used for front-end.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Float property.
	FloatTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionPropertyTypeOptionsPropertyDefinitionFloatTypeOptions `json:"floatTypeOptions,omitempty" yaml:"floatTypeOptions,omitempty"`

	// Whether the property is mandatory.
	IsRequired bool `json:"isRequired,omitempty" yaml:"isRequired,omitempty"`

	/*
	   Stores the retrieval importance.
	   Possible values are: `HIGHEST`, `HIGHER`, `HIGH`, `MEDIUM`, `LOW`, `LOWEST`.
	*/
	RetrievalImportance string `json:"retrievalImportance,omitempty" yaml:"retrievalImportance,omitempty"`

	/*
	   Enum/categorical property.
	   Structure is documented below.
	*/
	EnumTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionPropertyTypeOptionsPropertyDefinitionEnumTypeOptions `json:"enumTypeOptions,omitempty" yaml:"enumTypeOptions,omitempty"`

	// Integer property.
	IntegerTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionPropertyTypeOptionsPropertyDefinitionIntegerTypeOptions `json:"integerTypeOptions,omitempty" yaml:"integerTypeOptions,omitempty"`

	// Text property.
	TextTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionPropertyTypeOptionsPropertyDefinitionTextTypeOptions `json:"textTypeOptions,omitempty" yaml:"textTypeOptions,omitempty"`

	// Timestamp property. Not supported by CMEK compliant deployment.
	TimestampTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionPropertyTypeOptionsPropertyDefinitionTimestampTypeOptions `json:"timestampTypeOptions,omitempty" yaml:"timestampTypeOptions,omitempty"`

	// Date time property. Not supported by CMEK compliant deployment.
	DateTimeTypeOptions Essentialcontacts_DocumentAiWarehouseDocumentSchemaPropertyDefinitionPropertyTypeOptionsPropertyDefinitionDateTimeTypeOptions `json:"dateTimeTypeOptions,omitempty" yaml:"dateTimeTypeOptions,omitempty"`
}
