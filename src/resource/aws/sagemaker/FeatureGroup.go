package sagemaker

import types "DesignSphere_Server/src/resource/aws/types"

type FeatureGroup struct {
	// The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the Offline Store if an `offline_store_config` is provided.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Map of resource tags for the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A list of Feature names and types. See Feature Definition Below.
	FeatureDefinitions []types.Sagemaker_FeatureGroupFeatureDefinition `json:"featureDefinitions,omitempty" yaml:"featureDefinitions,omitempty"`

	// The name of the Feature Group. The name must be unique within an AWS Region in an AWS account.
	FeatureGroupName string `json:"featureGroupName,omitempty" yaml:"featureGroupName,omitempty"`

	// The Offline Feature Store Configuration. See Offline Store Config Below.
	OfflineStoreConfig types.Sagemaker_FeatureGroupOfflineStoreConfig `json:"offlineStoreConfig,omitempty" yaml:"offlineStoreConfig,omitempty"`

	// The Online Feature Store Configuration. See Online Store Config Below.
	OnlineStoreConfig types.Sagemaker_FeatureGroupOnlineStoreConfig `json:"onlineStoreConfig,omitempty" yaml:"onlineStoreConfig,omitempty"`

	// A free-form description of a Feature Group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the feature that stores the EventTime of a Record in a Feature Group.
	EventTimeFeatureName string `json:"eventTimeFeatureName,omitempty" yaml:"eventTimeFeatureName,omitempty"`

	// The name of the Feature whose value uniquely identifies a Record defined in the Feature Store. Only the latest record per identifier value will be stored in the Online Store.
	RecordIdentifierFeatureName string `json:"recordIdentifierFeatureName,omitempty" yaml:"recordIdentifierFeatureName,omitempty"`
}
