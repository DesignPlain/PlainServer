package firebase

type HostingChannel struct {
	// Required. The ID of the site in which to create this channel.
	SiteId string `json:"siteId,omitempty" yaml:"siteId,omitempty"`

	/*
	   Input only. A time-to-live for this channel. Sets `expire_time` to the provided
	   duration past the time of the request. A duration in seconds with up to nine fractional
	   digits, terminated by 's'. Example: "86400s" (one day).
	*/
	Ttl string `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	/*
	   Required. Immutable. A unique ID within the site that identifies the channel.


	   - - -
	*/
	ChannelId string `json:"channelId,omitempty" yaml:"channelId,omitempty"`

	/*
	   The time at which the channel will be automatically deleted. If null, the channel
	   will not be automatically deleted. This field is present in the output whether it's
	   set directly or via the `ttl` field.
	*/
	ExpireTime string `json:"expireTime,omitempty" yaml:"expireTime,omitempty"`

	/*
	   Text labels used for extra metadata and/or filtering
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The number of previous releases to retain on the channel for rollback or other
	   purposes. Must be a number between 1-100. Defaults to 10 for new channels.
	*/
	RetainedReleaseCount int `json:"retainedReleaseCount,omitempty" yaml:"retainedReleaseCount,omitempty"`
}
