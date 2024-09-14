package cognito

type UserPoolUICustomization struct {
	// The CSS values in the UI customization, provided as a String. At least one of `css` or `image_file` is required.
	Css string `json:"css,omitempty" yaml:"css,omitempty"`

	// The uploaded logo image for the UI customization, provided as a base64-encoded String. Drift detection is not possible for this argument. At least one of `css` or `image_file` is required.
	ImageFile string `json:"imageFile,omitempty" yaml:"imageFile,omitempty"`

	// The user pool ID for the user pool.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`

	// The client ID for the client app. Defaults to `ALL`. If `ALL` is specified, the `css` and/or `image_file` settings will be used for every client that has no UI customization set previously.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`
}
