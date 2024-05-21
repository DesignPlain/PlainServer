package firebase

type AppCheckRecaptchaEnterpriseConfig struct {
	/*
	   The ID of an
	   [Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id).


	   - - -
	*/
	AppId string `json:"appId,omitempty" yaml:"appId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The score-based site key created in reCAPTCHA Enterprise used to invoke reCAPTCHA and generate the reCAPTCHA tokens for your application.
	   --Important--: This is not the siteSecret (as it is in reCAPTCHA v3), but rather your score-based reCAPTCHA Enterprise site key.
	*/
	SiteKey string `json:"siteKey,omitempty" yaml:"siteKey,omitempty"`

	/*
	   Specifies the duration for which App Check tokens exchanged from reCAPTCHA Enterprise artifacts will be valid.
	   If unset, a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive.
	   A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	*/
	TokenTtl string `json:"tokenTtl,omitempty" yaml:"tokenTtl,omitempty"`
}
