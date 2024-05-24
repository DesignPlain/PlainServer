package pinpoint

type ApnsSandboxChannel struct {
	// The application ID.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	/*
	   The default authentication method used for APNs Sandbox.
	   __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	   You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	   If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.

	   One of the following sets of credentials is also required.

	   If you choose to use __Certificate credentials__ you will have to provide:
	*/
	DefaultAuthenticationMethod string `json:"defaultAuthenticationMethod,omitempty" yaml:"defaultAuthenticationMethod,omitempty"`

	/*
	   The Certificate Private Key file (ie. `.key` file).

	   If you choose to use __Key credentials__ you will have to provide:
	*/
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`

	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId string `json:"teamId,omitempty" yaml:"teamId,omitempty"`

	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey string `json:"tokenKey,omitempty" yaml:"tokenKey,omitempty"`

	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId string `json:"tokenKeyId,omitempty" yaml:"tokenKeyId,omitempty"`

	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId string `json:"bundleId,omitempty" yaml:"bundleId,omitempty"`

	// The pem encoded TLS Certificate from Apple.
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`

	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
