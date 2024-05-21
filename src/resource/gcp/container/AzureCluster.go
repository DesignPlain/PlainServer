package container

import types "DesignSphere_Server/src/resource/gcp/types"

type AzureCluster struct {
	// Logging configuration.
	LoggingConfig types.Container_AzureClusterLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// The name of this resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration related to the cluster RBAC settings.
	Authorization types.Container_AzureClusterAuthorization `json:"authorization,omitempty" yaml:"authorization,omitempty"`

	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion string `json:"azureRegion,omitempty" yaml:"azureRegion,omitempty"`

	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the `AzureCluster`. `AzureClient` names are formatted as `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client string `json:"client,omitempty" yaml:"client,omitempty"`

	// Configuration related to the cluster control plane.
	ControlPlane types.Container_AzureClusterControlPlane `json:"controlPlane,omitempty" yaml:"controlPlane,omitempty"`

	/*
	   Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.

	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// Cluster-wide networking configuration.
	Networking types.Container_AzureClusterNetworking `json:"networking,omitempty" yaml:"networking,omitempty"`

	// The ARM ID of the resource group where the cluster resources are deployed. For example: `/subscriptions/-/resourceGroups/-`
	ResourceGroupId string `json:"resourceGroupId,omitempty" yaml:"resourceGroupId,omitempty"`

	// Azure authentication configuration for management of Azure resources
	AzureServicesAuthentication types.Container_AzureClusterAzureServicesAuthentication `json:"azureServicesAuthentication,omitempty" yaml:"azureServicesAuthentication,omitempty"`

	// Fleet configuration.
	Fleet types.Container_AzureClusterFleet `json:"fleet,omitempty" yaml:"fleet,omitempty"`

	// The number of the Fleet host project where this cluster will be registered.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
