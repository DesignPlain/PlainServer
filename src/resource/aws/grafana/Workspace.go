package grafana

import types "DesignSphere_Server/src/resource/aws/types"

type Workspace struct {
	// The workspace description.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The notification destinations. If a data source is specified here, Amazon Managed Grafana will create IAM roles and permissions needed to use these destinations. Must be set to `SNS`.
	NotificationDestinations []string `json:"notificationDestinations,omitempty" yaml:"notificationDestinations,omitempty"`

	// The data sources for the workspace. Valid values are `AMAZON_OPENSEARCH_SERVICE`, `ATHENA`, `CLOUDWATCH`, `PROMETHEUS`, `REDSHIFT`, `SITEWISE`, `TIMESTREAM`, `XRAY`
	DataSources []string `json:"dataSources,omitempty" yaml:"dataSources,omitempty"`

	// Specifies the version of Grafana to support in the new workspace. Supported values are `8.4` and `9.4`. If not specified, defaults to `8.4`.
	GrafanaVersion string `json:"grafanaVersion,omitempty" yaml:"grafanaVersion,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of account access for the workspace. Valid values are `CURRENT_ACCOUNT` and `ORGANIZATION`. If `ORGANIZATION` is specified, then `organizational_units` must also be present.
	AccountAccessType string `json:"accountAccessType,omitempty" yaml:"accountAccessType,omitempty"`

	// The configuration string for the workspace that you create. For more information about the format and configuration options available, see [Working in your Grafana workspace](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-configure-workspace.html).
	Configuration string `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// The Amazon Organizations organizational units that the workspace is authorized to use data sources from.
	OrganizationalUnits []string `json:"organizationalUnits,omitempty" yaml:"organizationalUnits,omitempty"`

	/*
	   The permission type of the workspace. If `SERVICE_MANAGED` is specified, the IAM roles and IAM policy attachments are generated automatically. If `CUSTOMER_MANAGED` is specified, the IAM roles and IAM policy attachments will not be created.

	   The following arguments are optional:
	*/
	PermissionType string `json:"permissionType,omitempty" yaml:"permissionType,omitempty"`

	// The AWS CloudFormation stack set name that provisions IAM roles to be used by the workspace.
	StackSetName string `json:"stackSetName,omitempty" yaml:"stackSetName,omitempty"`

	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. See VPC Configuration below.
	VpcConfiguration types.Grafana_WorkspaceVpcConfiguration `json:"vpcConfiguration,omitempty" yaml:"vpcConfiguration,omitempty"`

	// The authentication providers for the workspace. Valid values are `AWS_SSO`, `SAML`, or both.
	AuthenticationProviders []string `json:"authenticationProviders,omitempty" yaml:"authenticationProviders,omitempty"`

	// The Grafana workspace name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration for network access to your workspace.See Network Access Control below.
	NetworkAccessControl types.Grafana_WorkspaceNetworkAccessControl `json:"networkAccessControl,omitempty" yaml:"networkAccessControl,omitempty"`

	// The role name that the workspace uses to access resources through Amazon Organizations.
	OrganizationRoleName string `json:"organizationRoleName,omitempty" yaml:"organizationRoleName,omitempty"`

	// The IAM role ARN that the workspace assumes.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
