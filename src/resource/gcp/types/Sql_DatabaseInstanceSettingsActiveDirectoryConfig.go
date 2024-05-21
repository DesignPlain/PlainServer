package types

type Sql_DatabaseInstanceSettingsActiveDirectoryConfig struct {
	/*
	   The domain name for the active directory (e.g., mydomain.com).
	   Can only be used with SQL Server.
	*/
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
