package firebase

type AndroidApp struct {
	/*
	   The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the AndroidApp.
	   If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the AndroidApp.
	   This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	*/
	ApiKeyId string `json:"apiKeyId,omitempty" yaml:"apiKeyId,omitempty"`

	/*
	   (Optional) Set to 'ABANDON' to allow the AndroidApp to be untracked from terraform state rather than deleted upon
	   'terraform destroy'. This is useful because the AndroidApp may be serving traffic. Set to 'DELETE' to delete the
	   AndroidApp. Defaults to 'DELETE'.
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	/*
	   The user-assigned display name of the AndroidApp.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Immutable. The canonical package name of the Android app as would appear in the Google Play
	   Developer Console.
	*/
	PackageName string `json:"packageName,omitempty" yaml:"packageName,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The SHA1 certificate hashes for the AndroidApp.
	Sha1Hashes []string `json:"sha1Hashes,omitempty" yaml:"sha1Hashes,omitempty"`

	// The SHA256 certificate hashes for the AndroidApp.
	Sha256Hashes []string `json:"sha256Hashes,omitempty" yaml:"sha256Hashes,omitempty"`
}
