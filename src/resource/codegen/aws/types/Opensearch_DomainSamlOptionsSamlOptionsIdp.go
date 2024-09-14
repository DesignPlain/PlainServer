package types

type Opensearch_DomainSamlOptionsSamlOptionsIdp struct {
	// Metadata of the SAML application in xml format.
	MetadataContent string `json:"metadataContent,omitempty" yaml:"metadataContent,omitempty"`

	// Unique Entity ID of the application in SAML Identity Provider.
	EntityId string `json:"entityId,omitempty" yaml:"entityId,omitempty"`
}
