package sagemaker

import types "libds/aws/types"

type Project struct {
	// A description for the project.
	ProjectDescription string `json:"projectDescription,omitempty" yaml:"projectDescription,omitempty"`

	// The name of the Project.
	ProjectName string `json:"projectName,omitempty" yaml:"projectName,omitempty"`

	// The product ID and provisioning artifact ID to provision a service catalog. See Service Catalog Provisioning Details below.
	ServiceCatalogProvisioningDetails types.Sagemaker_ProjectServiceCatalogProvisioningDetails `json:"serviceCatalogProvisioningDetails,omitempty" yaml:"serviceCatalogProvisioningDetails,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
