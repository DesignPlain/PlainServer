package firebase

import types "libds/gcp/types"

type HostingVersion struct {
	/*
	   The configuration for the behavior of the site. This configuration exists in the `firebase.json` file.
	   Structure is documented below.
	*/
	Config types.Firebase_HostingVersionConfig `json:"config,omitempty" yaml:"config,omitempty"`

	/*
	   Required. The ID of the site in which to create this Version.


	   - - -
	*/
	SiteId string `json:"siteId,omitempty" yaml:"siteId,omitempty"`
}
