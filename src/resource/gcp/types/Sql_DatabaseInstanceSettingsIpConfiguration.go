package types

type Sql_DatabaseInstanceSettingsIpConfiguration struct {
	/*
	   Whether this Cloud SQL instance should be assigned
	   a public IPV4 address. At least `ipv4_enabled` must be enabled or a
	   `private_network` must be configured.
	*/
	Ipv4Enabled bool `json:"ipv4Enabled,omitempty" yaml:"ipv4Enabled,omitempty"`

	/*
	   The VPC network from which the Cloud SQL
	   instance is accessible for private IP. For example, projects/myProject/global/networks/default.
	   Specifying a network enables private IP.
	   At least `ipv4_enabled` must be enabled or a `private_network` must be configured.
	   This setting can be updated, but it cannot be removed after it is set.
	*/
	PrivateNetwork string `json:"privateNetwork,omitempty" yaml:"privateNetwork,omitempty"`

	// PSC settings for a Cloud SQL instance.
	PscConfigs []Sql_DatabaseInstanceSettingsIpConfigurationPscConfig `json:"pscConfigs,omitempty" yaml:"pscConfigs,omitempty"`

	// Whether SSL connections over IP are enforced or not. To change this field, also set the corresponding value in `ssl_mode`.
	RequireSsl bool `json:"requireSsl,omitempty" yaml:"requireSsl,omitempty"`

	/*
	   Specify how SSL connection should be enforced in DB connections. This field provides more SSL enforcment options compared to `require_ssl`. To change this field, also set the correspoding value in `require_ssl`.
	   - For PostgreSQL instances, the value pairs are listed in the [API reference doc](https://cloud.google.com/sql/docs/postgres/admin-api/rest/v1beta4/instances#ipconfiguration) for `ssl_mode` field.
	   - For MySQL instances, use the same value pairs as the PostgreSQL instances.
	   - For SQL Server instances, set it to `ALLOW_UNENCRYPTED_AND_ENCRYPTED` when `require_ssl=false` and `ENCRYPTED_ONLY` otherwise.
	*/
	SslMode string `json:"sslMode,omitempty" yaml:"sslMode,omitempty"`

	// The name of the allocated ip range for the private ip CloudSQL instance. For example: "google-managed-services-default". If set, the instance ip will be created in the allocated range. The range name must comply with [RFC 1035](https://datatracker.ietf.org/doc/html/rfc1035). Specifically, the name must be 1-63 characters long and match the regular expression a-z?.
	AllocatedIpRange string `json:"allocatedIpRange,omitempty" yaml:"allocatedIpRange,omitempty"`

	//
	AuthorizedNetworks []Sql_DatabaseInstanceSettingsIpConfigurationAuthorizedNetwork `json:"authorizedNetworks,omitempty" yaml:"authorizedNetworks,omitempty"`

	// Whether Google Cloud services such as BigQuery are allowed to access data in this Cloud SQL instance over a private IP connection. SQLSERVER database type is not supported.
	EnablePrivatePathForGoogleCloudServices bool `json:"enablePrivatePathForGoogleCloudServices,omitempty" yaml:"enablePrivatePathForGoogleCloudServices,omitempty"`
}
