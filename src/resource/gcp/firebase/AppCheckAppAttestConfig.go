package firebase

type AppCheckAppAttestConfig struct {
	/*
	   The ID of an
	   [Apple App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).


	   - - -
	*/
	AppId string `json:"appId,omitempty" yaml:"appId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Specifies the duration for which App Check tokens exchanged from App Attest artifacts will be valid.
	   If unset, a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive.
	   A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	*/
	TokenTtl string `json:"tokenTtl,omitempty" yaml:"tokenTtl,omitempty"`
}
