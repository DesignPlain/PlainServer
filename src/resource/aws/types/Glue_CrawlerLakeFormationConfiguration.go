package types

type Glue_CrawlerLakeFormationConfiguration struct {
	// Required for cross account crawls. For same account crawls as the target data, this can omitted.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// Specifies whether to use Lake Formation credentials for the crawler instead of the IAM role credentials.
	UseLakeFormationCredentials bool `json:"useLakeFormationCredentials,omitempty" yaml:"useLakeFormationCredentials,omitempty"`
}
