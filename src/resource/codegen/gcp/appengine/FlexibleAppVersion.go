package appengine

import types "libds/gcp/types"

type FlexibleAppVersion struct {
	/*
	   Duration that static files should be cached by web proxies and browsers.
	   Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
	*/
	DefaultExpiration string `json:"defaultExpiration,omitempty" yaml:"defaultExpiration,omitempty"`

	// If set to `true`, the service will be deleted if it is the last version.
	DeleteServiceOnDestroy bool `json:"deleteServiceOnDestroy,omitempty" yaml:"deleteServiceOnDestroy,omitempty"`

	/*
	   Code and application artifacts that make up this version.
	   Structure is documented below.
	*/
	Deployment types.Appengine_FlexibleAppVersionDeployment `json:"deployment,omitempty" yaml:"deployment,omitempty"`

	/*
	   An ordered list of URL-matching patterns that should be applied to incoming requests.
	   The first matching URL handles the request and other request handlers are not attempted.
	   Structure is documented below.
	*/
	Handlers []types.Appengine_FlexibleAppVersionHandler `json:"handlers,omitempty" yaml:"handlers,omitempty"`

	/*
	   Instance class that is used to run this version. Valid values are
	   AutomaticScaling: F1, F2, F4, F4_1G
	   ManualScaling: B1, B2, B4, B8, B4_1G
	   Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
	*/
	InstanceClass string `json:"instanceClass,omitempty" yaml:"instanceClass,omitempty"`

	/*
	   A list of the types of messages that this application is able to receive.
	   Each value may be one of: `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, `INBOUND_SERVICE_WARMUP`.
	*/
	InboundServices []string `json:"inboundServices,omitempty" yaml:"inboundServices,omitempty"`

	/*
	   A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
	   Structure is documented below.
	*/
	ManualScaling types.Appengine_FlexibleAppVersionManualScaling `json:"manualScaling,omitempty" yaml:"manualScaling,omitempty"`

	/*
	   Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.
	   Structure is documented below.
	*/
	ReadinessCheck types.Appengine_FlexibleAppVersionReadinessCheck `json:"readinessCheck,omitempty" yaml:"readinessCheck,omitempty"`

	/*
	   The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as
	   default if this field is neither provided in app.yaml file nor through CLI flag.
	*/
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	/*
	   Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens.
	   Reserved names,"default", "latest", and any name with the prefix "ah-".
	*/
	VersionId string `json:"versionId,omitempty" yaml:"versionId,omitempty"`

	/*
	   Extra network settings
	   Structure is documented below.
	*/
	Network types.Appengine_FlexibleAppVersionNetwork `json:"network,omitempty" yaml:"network,omitempty"`

	// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
	NobuildFilesRegex string `json:"nobuildFilesRegex,omitempty" yaml:"nobuildFilesRegex,omitempty"`

	// The channel of the runtime to use. Only available for some runtimes.
	RuntimeChannel string `json:"runtimeChannel,omitempty" yaml:"runtimeChannel,omitempty"`

	// AppEngine service resource. Can contain numbers, letters, and hyphens.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	// If set to `true`, the application version will not be deleted.
	NoopOnDestroy bool `json:"noopOnDestroy,omitempty" yaml:"noopOnDestroy,omitempty"`

	/*
	   The version of the API in the given runtime environment.
	   Please see the app.yaml reference for valid values at `https://cloud.google.com/appengine/docs/standard/<language>/config/appref`\
	   Substitute `<language>` with `python`, `java`, `php`, `ruby`, `go` or `nodejs`.
	*/
	RuntimeApiVersion string `json:"runtimeApiVersion,omitempty" yaml:"runtimeApiVersion,omitempty"`

	/*
	   The entrypoint for the application.
	   Structure is documented below.
	*/
	Entrypoint types.Appengine_FlexibleAppVersionEntrypoint `json:"entrypoint,omitempty" yaml:"entrypoint,omitempty"`

	/*
	   Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
	   Default value is `SERVING`.
	   Possible values are: `SERVING`, `STOPPED`.
	*/
	ServingStatus string `json:"servingStatus,omitempty" yaml:"servingStatus,omitempty"`

	// Metadata settings that are supplied to this version to enable beta runtime features.
	BetaSettings map[string]string `json:"betaSettings,omitempty" yaml:"betaSettings,omitempty"`

	/*
	   Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.
	   Structure is documented below.
	*/
	LivenessCheck types.Appengine_FlexibleAppVersionLivenessCheck `json:"livenessCheck,omitempty" yaml:"livenessCheck,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Machine resources for a version.
	   Structure is documented below.
	*/
	Resources types.Appengine_FlexibleAppVersionResources `json:"resources,omitempty" yaml:"resources,omitempty"`

	// Desired runtime. Example python27.
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`

	/*
	   Enables VPC connectivity for standard apps.
	   Structure is documented below.
	*/
	VpcAccessConnector types.Appengine_FlexibleAppVersionVpcAccessConnector `json:"vpcAccessConnector,omitempty" yaml:"vpcAccessConnector,omitempty"`

	/*
	   Serving configuration for Google Cloud Endpoints.
	   Structure is documented below.
	*/
	ApiConfig types.Appengine_FlexibleAppVersionApiConfig `json:"apiConfig,omitempty" yaml:"apiConfig,omitempty"`

	/*
	   Automatic scaling is based on request rate, response latencies, and other application metrics.
	   Structure is documented below.
	*/
	AutomaticScaling types.Appengine_FlexibleAppVersionAutomaticScaling `json:"automaticScaling,omitempty" yaml:"automaticScaling,omitempty"`

	/*
	   Code and application artifacts that make up this version.
	   Structure is documented below.
	*/
	EndpointsApiService types.Appengine_FlexibleAppVersionEndpointsApiService `json:"endpointsApiService,omitempty" yaml:"endpointsApiService,omitempty"`

	// Environment variables available to the application.  As these are not returned in the API request, the provider will not detect any changes made outside of the config.
	EnvVariables map[string]string `json:"envVariables,omitempty" yaml:"envVariables,omitempty"`

	// The path or name of the app's main executable.
	RuntimeMainExecutablePath string `json:"runtimeMainExecutablePath,omitempty" yaml:"runtimeMainExecutablePath,omitempty"`
}
