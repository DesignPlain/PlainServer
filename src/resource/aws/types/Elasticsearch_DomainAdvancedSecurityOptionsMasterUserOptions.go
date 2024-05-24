package types

type Elasticsearch_DomainAdvancedSecurityOptionsMasterUserOptions struct {
	// Main user's password, which is stored in the Amazon Elasticsearch Service domain's internal database. Only specify if `internal_user_database_enabled` is set to `true`.
	MasterUserPassword string `json:"masterUserPassword,omitempty" yaml:"masterUserPassword,omitempty"`

	// ARN for the main user. Only specify if `internal_user_database_enabled` is not set or set to `false`.
	MasterUserArn string `json:"masterUserArn,omitempty" yaml:"masterUserArn,omitempty"`

	// Main user's username, which is stored in the Amazon Elasticsearch Service domain's internal database. Only specify if `internal_user_database_enabled` is set to `true`.
	MasterUserName string `json:"masterUserName,omitempty" yaml:"masterUserName,omitempty"`
}
