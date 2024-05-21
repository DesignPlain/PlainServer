package appengine

import types "DesignSphere_Server/src/resource/gcp/types"

type StandardAppVersion struct {
	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy bool `json:"deleteServiceOnDestroy,omitempty" yaml:"deleteServiceOnDestroy,omitempty"`

	/*
	   The entrypoint for the application.
	   Structure is documented below.
	*/
	Entrypoint types.Appengine_StandardAppVersionEntrypoint `json:"entrypoint,omitempty" yaml:"entrypoint,omitempty"`

	/*
	   A list of the types of messages that this application is able to receive.
	   Each value may be one of: `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, `INBOUND_SERVICE_WARMUP`.
	*/
	InboundServices []string `json:"inboundServices,omitempty" yaml:"inboundServices,omitempty"`

	/*
	   A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	   Structure is documented below.
	*/
	ManualScaling types.Appengine_StandardAppVersionManualScaling `json:"manualScaling,omitempty" yaml:"manualScaling,omitempty"`

	// Desired runtime. Example python27.
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`

	/*
	   The version of the API in the given runtime environment.
	   Please see the app.yaml reference for valid values at `https://cloud.google.com/appengine/docs/standard/<language>/config/appref`\
	   Substitute `<language>` with `python`, `java`, `php`, `ruby`, `go` or `nodejs`.
	*/
	RuntimeApiVersion string `json:"runtimeApiVersion,omitempty" yaml:"runtimeApiVersion,omitempty"`

	// AppEngine service resource
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	/*
	   Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
	   Structure is documented below.
	*/
	BasicScaling types.Appengine_StandardAppVersionBasicScaling `json:"basicScaling,omitempty" yaml:"basicScaling,omitempty"`

	/*
	   Code and application artifacts that make up this version.
	   Structure is documented below.
	*/
	Deployment types.Appengine_StandardAppVersionDeployment `json:"deployment,omitempty" yaml:"deployment,omitempty"`

	// Environment variables available to the application.
	EnvVariables map[string]string `json:"envVariables,omitempty" yaml:"envVariables,omitempty"`

	// The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	/*
	   Enables VPC connectivity for standard apps.
	   Structure is documented below.
	*/
	VpcAccessConnector types.Appengine_StandardAppVersionVpcAccessConnector `json:"vpcAccessConnector,omitempty" yaml:"vpcAccessConnector,omitempty"`

	// Allows App Engine second generation runtimes to access the legacy bundled services.
	AppEngineApis bool `json:"appEngineApis,omitempty" yaml:"appEngineApis,omitempty"`

	/*
	   An ordered list of URL-matching patterns that should be applied to incoming requests.
	   The first matching URL handles the request and other request handlers are not attempted.
	   Structure is documented below.
	*/
	Handlers []types.Appengine_StandardAppVersionHandler `json:"handlers,omitempty" yaml:"handlers,omitempty"`

	/*
	   Instance class that is used to run this version. Valid values are
	   AutomaticScaling: F1, F2, F4, F4_1G
	   BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
	   Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
	*/
	InstanceClass string `json:"instanceClass,omitempty" yaml:"instanceClass,omitempty"`

	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy bool `json:"noopOnDestroy,omitempty" yaml:"noopOnDestroy,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Whether multiple requests can be dispatched to this version at once.
	Threadsafe bool `json:"threadsafe,omitempty" yaml:"threadsafe,omitempty"`

	/*
	   Automatic scaling is based on request rate, response latencies, and other application metrics.
	   Structure is documented below.
	*/
	AutomaticScaling types.Appengine_StandardAppVersionAutomaticScaling `json:"automaticScaling,omitempty" yaml:"automaticScaling,omitempty"`

	/*
	   Configuration for third-party Python runtime libraries that are required by the application.
	   Structure is documented below.
	*/
	Libraries []types.Appengine_StandardAppVersionLibrary `json:"libraries,omitempty" yaml:"libraries,omitempty"`

	// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
	VersionId string `json:"versionId,omitempty" yaml:"versionId,omitempty"`
}
