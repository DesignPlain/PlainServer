package firebase

type AppCheckRecaptchaV3Config struct {
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
	   The site secret used to identify your service for reCAPTCHA v3 verification.
	   For security reasons, this field will never be populated in any response.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	SiteSecret string `json:"siteSecret,omitempty" yaml:"siteSecret,omitempty"`

	/*
	   Specifies the duration for which App Check tokens exchanged from reCAPTCHA V3 artifacts will be valid.
	   If unset, a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive.
	   A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	*/
	TokenTtl string `json:"tokenTtl,omitempty" yaml:"tokenTtl,omitempty"`
}
