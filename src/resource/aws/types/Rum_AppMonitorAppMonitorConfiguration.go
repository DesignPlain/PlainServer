package types

type Rum_AppMonitorAppMonitorConfiguration struct {
	// A list of pages in the CloudWatch RUM console that are to be displayed with a "favorite" icon.
	FavoritePages []string `json:"favoritePages,omitempty" yaml:"favoritePages,omitempty"`

	// If you set this to `true`, RUM enables X-Ray tracing for the user sessions  that RUM samples. RUM adds an X-Ray trace header to allowed HTTP requests. It also records an X-Ray segment for allowed HTTP requests.
	EnableXray bool `json:"enableXray,omitempty" yaml:"enableXray,omitempty"`

	// A list of URLs in your website or application to exclude from RUM data collection.
	ExcludedPages []string `json:"excludedPages,omitempty" yaml:"excludedPages,omitempty"`

	// The ID of the Amazon Cognito identity pool that is used to authorize the sending of data to RUM.
	IdentityPoolId string `json:"identityPoolId,omitempty" yaml:"identityPoolId,omitempty"`

	// If this app monitor is to collect data from only certain pages in your application, this structure lists those pages.
	IncludedPages []string `json:"includedPages,omitempty" yaml:"includedPages,omitempty"`

	// Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. Default value is `0.1`.
	SessionSampleRate float64 `json:"sessionSampleRate,omitempty" yaml:"sessionSampleRate,omitempty"`

	// An array that lists the types of telemetry data that this app monitor is to collect. Valid values are `errors`, `performance`, and `http`.
	Telemetries []string `json:"telemetries,omitempty" yaml:"telemetries,omitempty"`

	// If you set this to `true`, RUM web client sets two cookies, a session cookie  and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
	AllowCookies bool `json:"allowCookies,omitempty" yaml:"allowCookies,omitempty"`

	// The ARN of the guest IAM role that is attached to the Amazon Cognito identity pool that is used to authorize the sending of data to RUM.
	GuestRoleArn string `json:"guestRoleArn,omitempty" yaml:"guestRoleArn,omitempty"`
}
