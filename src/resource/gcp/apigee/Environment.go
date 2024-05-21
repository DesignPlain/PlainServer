package apigee

import types "DesignSphere_Server/src/resource/gcp/types"

type Environment struct {
	/*
	   Optional. API Proxy type supported by the environment. The type can be set when creating
	   the Environment and cannot be changed.
	   Possible values are: `API_PROXY_TYPE_UNSPECIFIED`, `PROGRAMMABLE`, `CONFIGURABLE`.
	*/
	ApiProxyType string `json:"apiProxyType,omitempty" yaml:"apiProxyType,omitempty"`

	/*
	   Optional. Deployment type supported by the environment. The deployment type can be
	   set when creating the environment and cannot be changed. When you enable archive
	   deployment, you will be prevented from performing a subset of actions within the
	   environment, including:
	   Managing the deployment of API proxy or shared flow revisions;
	   Creating, updating, or deleting resource files;
	   Creating, updating, or deleting target servers.
	   Possible values are: `DEPLOYMENT_TYPE_UNSPECIFIED`, `PROXY`, `ARCHIVE`.
	*/
	DeploymentType string `json:"deploymentType,omitempty" yaml:"deploymentType,omitempty"`

	// Description of the environment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Display name of the environment.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The resource ID of the environment.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   NodeConfig for setting the min/max number of nodes associated with the environment.
	   Structure is documented below.
	*/
	NodeConfig types.Apigee_EnvironmentNodeConfig `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`

	/*
	   The Apigee Organization associated with the Apigee environment,
	   in the format `organizations/{{org_name}}`.


	   - - -
	*/
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	/*
	   Types that can be selected for an Environment. Each of the types are
	   limited by capability and capacity. Refer to Apigee's public documentation
	   to understand about each of these types in details.
	   An Apigee org can support heterogeneous Environments.
	   Possible values are: `ENVIRONMENT_TYPE_UNSPECIFIED`, `BASE`, `INTERMEDIATE`, `COMPREHENSIVE`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
