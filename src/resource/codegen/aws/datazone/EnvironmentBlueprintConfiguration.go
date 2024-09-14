package datazone

type EnvironmentBlueprintConfiguration struct {
	// Parameters for each region in which the blueprint is enabled
	RegionalParameters map[string]map[string]string `json:"regionalParameters,omitempty" yaml:"regionalParameters,omitempty"`

	// ID of the Domain.
	DomainId string `json:"domainId,omitempty" yaml:"domainId,omitempty"`

	/*
	   Regions in which the blueprint is enabled

	   The following arguments are optional:
	*/
	EnabledRegions []string `json:"enabledRegions,omitempty" yaml:"enabledRegions,omitempty"`

	// ID of the Environment Blueprint
	EnvironmentBlueprintId string `json:"environmentBlueprintId,omitempty" yaml:"environmentBlueprintId,omitempty"`

	// ARN of the manage access role with which this blueprint is created.
	ManageAccessRoleArn string `json:"manageAccessRoleArn,omitempty" yaml:"manageAccessRoleArn,omitempty"`

	// ARN of the provisioning role with which this blueprint is created.
	ProvisioningRoleArn string `json:"provisioningRoleArn,omitempty" yaml:"provisioningRoleArn,omitempty"`
}
