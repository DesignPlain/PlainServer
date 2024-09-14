package servicecatalog

import types "libds/aws/types"

type ProvisionedProduct struct {
	// Configuration block with parameters specified by the administrator that are required for provisioning the product. See `provisioning_parameters` Block for details.
	ProvisioningParameters []types.Servicecatalog_ProvisionedProductProvisioningParameter `json:"provisioningParameters,omitempty" yaml:"provisioningParameters,omitempty"`

	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related events.
	NotificationArns []string `json:"notificationArns,omitempty" yaml:"notificationArns,omitempty"`

	// Name of the provisioning artifact. You must provide the `provisioning_artifact_id` or `provisioning_artifact_name`, but not both.
	ProvisioningArtifactName string `json:"provisioningArtifactName,omitempty" yaml:"provisioningArtifactName,omitempty"`

	// _Only applies to deleting._ Whether to delete the Service Catalog provisioned product but leave the CloudFormation stack, stack set, or the underlying resources of the deleted provisioned product. The default value is `false`.
	RetainPhysicalResources bool `json:"retainPhysicalResources,omitempty" yaml:"retainPhysicalResources,omitempty"`

	/*
	   User-friendly name of the provisioned product.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Name of the product. You must provide `product_id` or `product_name`, but not both.
	ProductName string `json:"productName,omitempty" yaml:"productName,omitempty"`

	// Tags to apply to the provisioned product. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Language code. Valid values: `en` (English), `jp` (Japanese), `zh` (Chinese). Default value is `en`.
	AcceptLanguage string `json:"acceptLanguage,omitempty" yaml:"acceptLanguage,omitempty"`

	// Path identifier of the product. This value is optional if the product has a default path, and required if the product has more than one path. To list the paths for a product, use `aws.servicecatalog.getLaunchPaths`. When required, you must provide `path_id` or `path_name`, but not both.
	PathId string `json:"pathId,omitempty" yaml:"pathId,omitempty"`

	// Product identifier. For example, `prod-abcdzk7xy33qa`. You must provide `product_id` or `product_name`, but not both.
	ProductId string `json:"productId,omitempty" yaml:"productId,omitempty"`

	// Identifier of the provisioning artifact. For example, `pa-4abcdjnxjj6ne`. You must provide the `provisioning_artifact_id` or `provisioning_artifact_name`, but not both.
	ProvisioningArtifactId string `json:"provisioningArtifactId,omitempty" yaml:"provisioningArtifactId,omitempty"`

	// Configuration block with information about the provisioning preferences for a stack set. See `stack_set_provisioning_preferences` Block for details.
	StackSetProvisioningPreferences types.Servicecatalog_ProvisionedProductStackSetProvisioningPreferences `json:"stackSetProvisioningPreferences,omitempty" yaml:"stackSetProvisioningPreferences,omitempty"`

	// _Only applies to deleting._ If set to `true`, AWS Service Catalog stops managing the specified provisioned product even if it cannot delete the underlying resources. The default value is `false`.
	IgnoreErrors bool `json:"ignoreErrors,omitempty" yaml:"ignoreErrors,omitempty"`

	// Name of the path. You must provide `path_id` or `path_name`, but not both.
	PathName string `json:"pathName,omitempty" yaml:"pathName,omitempty"`
}
