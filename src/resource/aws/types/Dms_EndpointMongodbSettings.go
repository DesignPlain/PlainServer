package types

type Dms_EndpointMongodbSettings struct {
	// Number of documents to preview to determine the document organization. Use this setting when `nesting_level` is set to `one`. Default is `1000`.
	DocsToInvestigate string `json:"docsToInvestigate,omitempty" yaml:"docsToInvestigate,omitempty"`

	// Document ID. Use this setting when `nesting_level` is set to `none`. Default is `false`.
	ExtractDocId string `json:"extractDocId,omitempty" yaml:"extractDocId,omitempty"`

	// Specifies either document or table mode. Default is `none`. Valid values are `one` (table mode) and `none` (document mode).
	NestingLevel string `json:"nestingLevel,omitempty" yaml:"nestingLevel,omitempty"`

	// Authentication mechanism to access the MongoDB source endpoint. Default is `default`.
	AuthMechanism string `json:"authMechanism,omitempty" yaml:"authMechanism,omitempty"`

	// Authentication database name. Not used when `auth_type` is `no`. Default is `admin`.
	AuthSource string `json:"authSource,omitempty" yaml:"authSource,omitempty"`

	// Authentication type to access the MongoDB source endpoint. Default is `password`.
	AuthType string `json:"authType,omitempty" yaml:"authType,omitempty"`
}
