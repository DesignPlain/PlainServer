package deploymentmanager

import types "DesignSphere_Server/src/resource/gcp/types"

type Deployment struct {
	/*
	   Parameters that define your deployment, including the deployment
	   configuration and relevant templates.
	   Structure is documented below.
	*/
	Target types.Deploymentmanager_DeploymentTarget `json:"target,omitempty" yaml:"target,omitempty"`

	/*
	   Set the policy to use for creating new resources. Only used on
	   create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
	   `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
	   the deployment will fail. Note that updating this field does not
	   actually affect the deployment, just how it is updated.
	   Default value is `CREATE_OR_ACQUIRE`.
	   Possible values are: `ACQUIRE`, `CREATE_OR_ACQUIRE`.
	*/
	CreatePolicy string `json:"createPolicy,omitempty" yaml:"createPolicy,omitempty"`

	/*
	   Set the policy to use for deleting new resources on update/delete.
	   Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
	   resource is deleted after removal from Deployment Manager. If
	   `ABANDON`, the resource is only removed from Deployment Manager
	   and is not actually deleted. Note that updating this field does not
	   actually change the deployment, just how it is updated.
	   Default value is `DELETE`.
	   Possible values are: `ABANDON`, `DELETE`.
	*/
	DeletePolicy string `json:"deletePolicy,omitempty" yaml:"deletePolicy,omitempty"`

	// Optional user-provided description of deployment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Key-value pairs to apply to this labels.
	   Structure is documented below.
	*/
	Labels []types.Deploymentmanager_DeploymentLabel `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Unique name for the deployment
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   If set to true, a deployment is created with "shell" resources
	   that are not actually instantiated. This allows you to preview a
	   deployment. It can be updated to false to actually deploy
	   with real resources.
	   ~>--NOTE:-- Deployment Manager does not allow update
	   of a deployment in preview (unless updating to preview=false). Thus,
	   the provider will force-recreate deployments if either preview is updated
	   to true or if other fields are updated while preview is true.
	*/
	Preview bool `json:"preview,omitempty" yaml:"preview,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
