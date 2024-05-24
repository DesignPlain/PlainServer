package types

type Dms_getEndpointMongodbSetting struct {
	//
	NestingLevel string `json:"nestingLevel,omitempty" yaml:"nestingLevel,omitempty"`

	//
	AuthMechanism string `json:"authMechanism,omitempty" yaml:"authMechanism,omitempty"`

	//
	AuthSource string `json:"authSource,omitempty" yaml:"authSource,omitempty"`

	//
	AuthType string `json:"authType,omitempty" yaml:"authType,omitempty"`

	//
	DocsToInvestigate string `json:"docsToInvestigate,omitempty" yaml:"docsToInvestigate,omitempty"`

	//
	ExtractDocId string `json:"extractDocId,omitempty" yaml:"extractDocId,omitempty"`
}
