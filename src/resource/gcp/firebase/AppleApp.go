package firebase

type AppleApp struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The Apple Developer Team ID associated with the App in the App Store.
	TeamId string `json:"teamId,omitempty" yaml:"teamId,omitempty"`

	/*
	   The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AppleApp.
	   If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AppleApp.
	   This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	*/
	ApiKeyId string `json:"apiKeyId,omitempty" yaml:"apiKeyId,omitempty"`

	// The automatically generated Apple ID assigned to the Apple app by Apple in the Apple App Store.
	AppStoreId string `json:"appStoreId,omitempty" yaml:"appStoreId,omitempty"`

	/*
	   The canonical bundle ID of the Apple app as it would appear in the Apple AppStore.


	   - - -
	*/
	BundleId string `json:"bundleId,omitempty" yaml:"bundleId,omitempty"`

	/*
	   (Optional) Set to 'ABANDON' to allow the Apple to be untracked from terraform state rather than deleted upon 'terraform
	   destroy'. This is useful because the Apple may be serving traffic. Set to 'DELETE' to delete the Apple. Defaults to
	   'DELETE'.
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	// The user-assigned display name of the App.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
