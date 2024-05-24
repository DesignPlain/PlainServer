package quicksight

import types "DesignSphere_Server/src/resource/aws/types"

type DataSource struct {
	// An identifier for the data source.
	DataSourceId string `json:"dataSourceId,omitempty" yaml:"dataSourceId,omitempty"`

	// Secure Socket Layer (SSL) properties that apply when Amazon QuickSight connects to your underlying source. See SSL Properties below for more details.
	SslProperties types.Quicksight_DataSourceSslProperties `json:"sslProperties,omitempty" yaml:"sslProperties,omitempty"`

	/*
	   The type of the data source. See the [AWS Documentation](https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateDataSource.html#QS-CreateDataSource-request-Type) for the complete list of valid values.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Use this parameter only when you want Amazon QuickSight to use a VPC connection when connecting to your underlying source. See VPC Connection Properties below for more details.
	VpcConnectionProperties types.Quicksight_DataSourceVpcConnectionProperties `json:"vpcConnectionProperties,omitempty" yaml:"vpcConnectionProperties,omitempty"`

	// A set of resource permissions on the data source. Maximum of 64 items. See Permission below for more details.
	Permissions []types.Quicksight_DataSourcePermission `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID for the AWS account that the data source is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`

	// The credentials Amazon QuickSight uses to connect to your underlying source. Currently, only credentials based on user name and password are supported. See Credentials below for more details.
	Credentials types.Quicksight_DataSourceCredentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`

	// A name for the data source, maximum of 128 characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The parameters used to connect to this data source (exactly one).
	Parameters types.Quicksight_DataSourceParameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
