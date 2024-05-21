package types

type Sql_getDatabaseInstanceSettingActiveDirectoryConfig struct {
	// Domain name of the Active Directory for SQL Server (e.g., mydomain.com).
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
