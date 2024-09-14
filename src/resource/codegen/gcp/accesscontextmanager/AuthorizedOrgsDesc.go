package accesscontextmanager

type AuthorizedOrgsDesc struct {
	/*
	   The type of entities that need to use the authorization relationship during
	   evaluation, such as a device. Valid values are "ASSET_TYPE_DEVICE" and
	   "ASSET_TYPE_CREDENTIAL_STRENGTH".
	   Possible values are: `ASSET_TYPE_DEVICE`, `ASSET_TYPE_CREDENTIAL_STRENGTH`.
	*/
	AssetType string `json:"assetType,omitempty" yaml:"assetType,omitempty"`

	/*
	   The direction of the authorization relationship between this organization
	   and the organizations listed in the "orgs" field. The valid values for this
	   field include the following:
	   AUTHORIZATION_DIRECTION_FROM: Allows this organization to evaluate traffic
	   in the organizations listed in the `orgs` field.
	   AUTHORIZATION_DIRECTION_TO: Allows the organizations listed in the `orgs`
	   field to evaluate the traffic in this organization.
	   For the authorization relationship to take effect, all of the organizations
	   must authorize and specify the appropriate relationship direction. For
	   example, if organization A authorized organization B and C to evaluate its
	   traffic, by specifying "AUTHORIZATION_DIRECTION_TO" as the authorization
	   direction, organizations B and C must specify
	   "AUTHORIZATION_DIRECTION_FROM" as the authorization direction in their
	   "AuthorizedOrgsDesc" resource.
	   Possible values are: `AUTHORIZATION_DIRECTION_TO`, `AUTHORIZATION_DIRECTION_FROM`.
	*/
	AuthorizationDirection string `json:"authorizationDirection,omitempty" yaml:"authorizationDirection,omitempty"`

	/*
	   A granular control type for authorization levels. Valid value is "AUTHORIZATION_TYPE_TRUST".
	   Possible values are: `AUTHORIZATION_TYPE_TRUST`.
	*/
	AuthorizationType string `json:"authorizationType,omitempty" yaml:"authorizationType,omitempty"`

	/*
	   Resource name for the `AuthorizedOrgsDesc`. Format:
	   `accessPolicies/{access_policy}/authorizedOrgsDescs/{authorized_orgs_desc}`.
	   The `authorized_orgs_desc` component must begin with a letter, followed by
	   alphanumeric characters or `_`.
	   After you create an `AuthorizedOrgsDesc`, you cannot change its `name`.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The list of organization ids in this AuthorizedOrgsDesc.
	   Format: `organizations/<org_number>`
	   Example: `organizations/123456`
	*/
	Orgs []string `json:"orgs,omitempty" yaml:"orgs,omitempty"`

	// Required. Resource name for the access policy which owns this `AuthorizedOrgsDesc`.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
