package firebase

type AppCheckDebugToken struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The secret token itself. Must be provided during creation, and must be a UUID4,
	   case insensitive. You may use a method of your choice such as random/random_uuid
	   to generate the token.
	   This field is immutable once set, and cannot be updated. You can, however, delete
	   this debug token to revoke it.
	   For security reasons, this field will never be populated in any response.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Token string `json:"token,omitempty" yaml:"token,omitempty"`

	/*
	   The ID of a
	   [Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id),
	   [Apple App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id),
	   or [Android App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.androidApps#AndroidApp.FIELDS.app_id)


	   - - -
	*/
	AppId string `json:"appId,omitempty" yaml:"appId,omitempty"`

	// A human readable display name used to identify this debug token.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
