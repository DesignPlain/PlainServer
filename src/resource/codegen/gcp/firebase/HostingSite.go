package firebase

type HostingSite struct {
	/*
	   Optional. The [ID of a Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id)
	   associated with the Hosting site.
	*/
	AppId string `json:"appId,omitempty" yaml:"appId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Required. Immutable. A globally unique identifier for the Hosting site. This identifier is
	   used to construct the Firebase-provisioned subdomains for the site, so it must also be a valid
	   domain name label.
	*/
	SiteId string `json:"siteId,omitempty" yaml:"siteId,omitempty"`
}
