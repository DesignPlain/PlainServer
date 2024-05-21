package firebase

type HostingRelease struct {
	/*
	   The ID of the channel to which the release belongs. If not provided, the release will
	   belong to the default "live" channel
	*/
	ChannelId string `json:"channelId,omitempty" yaml:"channelId,omitempty"`

	// The deploy description when the release was created. The value can be up to 512 characters.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	/*
	   Required. The ID of the site to which the release belongs.


	   - - -
	*/
	SiteId string `json:"siteId,omitempty" yaml:"siteId,omitempty"`

	/*
	   The type of the release; indicates what happened to the content of the site. There is no need to specify
	   `DEPLOY` or `ROLLBACK` type if a `version_name` is provided.
	   DEPLOY: A version was uploaded to Firebase Hosting and released. Output only.
	   ROLLBACK: The release points back to a previously deployed version. Output only.
	   SITE_DISABLE: The release prevents the site from serving content. Firebase Hosting acts as if the site never existed
	   Possible values are: `DEPLOY`, `ROLLBACK`, `SITE_DISABLE`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   The unique identifier for a version, in the format: sites/SITE_ID/versions/VERSION_ID.
	   The content of the version specified will be actively displayed on the appropriate URL.
	   The Version must belong to the same site as in the `site_id`.
	   This parameter must be empty if the `type` of the release is `SITE_DISABLE`.
	*/
	VersionName string `json:"versionName,omitempty" yaml:"versionName,omitempty"`
}
