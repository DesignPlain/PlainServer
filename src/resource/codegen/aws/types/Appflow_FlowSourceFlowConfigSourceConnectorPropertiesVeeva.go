package types

type Appflow_FlowSourceFlowConfigSourceConnectorPropertiesVeeva struct {
	// Document type specified in the Veeva document extract flow.
	DocumentType string `json:"documentType,omitempty" yaml:"documentType,omitempty"`

	// Boolean value to include All Versions of files in Veeva document extract flow.
	IncludeAllVersions bool `json:"includeAllVersions,omitempty" yaml:"includeAllVersions,omitempty"`

	// Boolean value to include file renditions in Veeva document extract flow.
	IncludeRenditions bool `json:"includeRenditions,omitempty" yaml:"includeRenditions,omitempty"`

	// Boolean value to include source files in Veeva document extract flow.
	IncludeSourceFiles bool `json:"includeSourceFiles,omitempty" yaml:"includeSourceFiles,omitempty"`

	//
	Object string `json:"object,omitempty" yaml:"object,omitempty"`
}
