package cloudasset

import types "libds/gcp/types"

type ProjectFeed struct {
	/*
	   A condition which determines whether an asset update should be published. If specified, an asset
	   will be returned only when the expression evaluates to true. When set, expression field
	   must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
	   expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
	   condition are optional.
	   Structure is documented below.
	*/
	Condition types.Cloudasset_ProjectFeedCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	/*
	   Asset content type. If not specified, no content but the asset name and type will be returned.
	   Possible values are: `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, `OS_INVENTORY`, `ACCESS_POLICY`.
	*/
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	FeedId string `json:"feedId,omitempty" yaml:"feedId,omitempty"`

	/*
	   Output configuration for asset feed destination.
	   Structure is documented below.
	*/
	FeedOutputConfig types.Cloudasset_ProjectFeedFeedOutputConfig `json:"feedOutputConfig,omitempty" yaml:"feedOutputConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A list of the full names of the assets to receive updates. You must specify either or both of
	   assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	   exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	   See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	*/
	AssetNames []string `json:"assetNames,omitempty" yaml:"assetNames,omitempty"`

	/*
	   A list of types of the assets to receive updates. You must specify either or both of assetNames
	   and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	   the feed. For example: "compute.googleapis.com/Disk"
	   See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	   supported asset types.
	*/
	AssetTypes []string `json:"assetTypes,omitempty" yaml:"assetTypes,omitempty"`

	/*
	   The project whose identity will be used when sending messages to the
	   destination pubsub topic. It also specifies the project for API
	   enablement check, quota, and billing. If not specified, the resource's
	   project will be used.
	*/
	BillingProject string `json:"billingProject,omitempty" yaml:"billingProject,omitempty"`
}
