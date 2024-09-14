package types

type Sql_getDatabaseInstanceSettingIpConfiguration struct {
	// Whether this Cloud SQL instance should be assigned a public IPV4 address. At least ipv4_enabled must be enabled or a private_network must be configured.
	Ipv4Enabled bool `json:"ipv4Enabled,omitempty" yaml:"ipv4Enabled,omitempty"`

	// The VPC network from which the Cloud SQL instance is accessible for private IP. For example, projects/myProject/global/networks/default. Specifying a network enables private IP. At least ipv4_enabled must be enabled or a private_network must be configured. This setting can be updated, but it cannot be removed after it is set.
	PrivateNetwork string `json:"privateNetwork,omitempty" yaml:"privateNetwork,omitempty"`

	// PSC settings for a Cloud SQL instance.
	PscConfigs []Sql_getDatabaseInstanceSettingIpConfigurationPscConfig `json:"pscConfigs,omitempty" yaml:"pscConfigs,omitempty"`

	// Whether SSL connections over IP are enforced or not. To change this field, also set the corresponding value in ssl_mode if it has been set too.
	RequireSsl bool `json:"requireSsl,omitempty" yaml:"requireSsl,omitempty"`

	// Specify how SSL connection should be enforced in DB connections. This field provides more SSL enforcment options compared to require_ssl. To change this field, also set the correspoding value in require_ssl.
	SslMode string `json:"sslMode,omitempty" yaml:"sslMode,omitempty"`

	// The name of the allocated ip range for the private ip CloudSQL instance. For example: "google-managed-services-default". If set, the instance ip will be created in the allocated range. The range name must comply with RFC 1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z?.
	AllocatedIpRange string `json:"allocatedIpRange,omitempty" yaml:"allocatedIpRange,omitempty"`

	//
	AuthorizedNetworks []Sql_getDatabaseInstanceSettingIpConfigurationAuthorizedNetwork `json:"authorizedNetworks,omitempty" yaml:"authorizedNetworks,omitempty"`

	// Whether Google Cloud services such as BigQuery are allowed to access data in this Cloud SQL instance over a private IP connection. SQLSERVER database type is not supported.
	EnablePrivatePathForGoogleCloudServices bool `json:"enablePrivatePathForGoogleCloudServices,omitempty" yaml:"enablePrivatePathForGoogleCloudServices,omitempty"`
}
