package types

type Ec2clientvpn_EndpointClientLoginBannerOptions struct {
	// Customizable text that will be displayed in a banner on AWS provided clients when a VPN session is established. UTF-8 encoded characters only. Maximum of 1400 characters.
	BannerText string `json:"bannerText,omitempty" yaml:"bannerText,omitempty"`

	// Enable or disable a customizable text banner that will be displayed on AWS provided clients when a VPN session is established. The default is `false` (not enabled).
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
