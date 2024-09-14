package types

type Elasticsearch_DomainSamlOptionsSamlOptionsIdp struct {
	// The unique Entity ID of the application in SAML Identity Provider.
	EntityId string `json:"entityId,omitempty" yaml:"entityId,omitempty"`

	// The Metadata of the SAML application in xml format.
	MetadataContent string `json:"metadataContent,omitempty" yaml:"metadataContent,omitempty"`
}
