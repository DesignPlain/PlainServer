package types

type Index_ProviderDefaultTags struct {
	// Resource tags to default across all resources. Can also be configured with environment variables like `TF_AWS_DEFAULT_TAGS_<tag_name>`.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
