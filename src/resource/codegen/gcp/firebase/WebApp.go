package firebase

type WebApp struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the WebApp.
	   If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the WebApp.
	   This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned.
	*/
	ApiKeyId string `json:"apiKeyId,omitempty" yaml:"apiKeyId,omitempty"`

	/*
	   Set to 'ABANDON' to allow the WebApp to be untracked from terraform state rather than deleted upon 'terraform destroy'.
	   This is useful becaue the WebApp may be serving traffic. Set to 'DELETE' to delete the WebApp. Default to 'DELETE'
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	/*
	   The user-assigned display name of the App.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
