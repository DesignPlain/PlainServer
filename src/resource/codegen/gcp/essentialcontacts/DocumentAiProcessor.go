package essentialcontacts

type DocumentAiProcessor struct {
	/*
	   The location of the resource.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The type of processor. For possible types see the [official list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The display name. Must be unique.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName string `json:"kmsKeyName,omitempty" yaml:"kmsKeyName,omitempty"`
}
