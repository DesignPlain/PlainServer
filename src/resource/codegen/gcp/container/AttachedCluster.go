package container

import types "libds/gcp/types"

type AttachedCluster struct {
	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Monitoring configuration.
	   Structure is documented below.
	*/
	MonitoringConfig types.Container_AttachedClusterMonitoringConfig `json:"monitoringConfig,omitempty" yaml:"monitoringConfig,omitempty"`

	// The name of this resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Fleet configuration.
	   Structure is documented below.
	*/
	Fleet types.Container_AttachedClusterFleet `json:"fleet,omitempty" yaml:"fleet,omitempty"`

	/*
	   Logging configuration.
	   Structure is documented below.
	*/
	LoggingConfig types.Container_AttachedClusterLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	/*
	   OIDC discovery information of the target cluster.
	   Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
	   API server. This fields indicates how GCP services
	   validate KSA tokens in order to allow system workloads (such as GKE Connect
	   and telemetry agents) to authenticate back to GCP.
	   Both clusters with public and private issuer URLs are supported.
	   Clusters with public issuers only need to specify the `issuer_url` field
	   while clusters with private issuers need to provide both
	   `issuer_url` and `jwks`.
	   Structure is documented below.
	*/
	OidcConfig types.Container_AttachedClusterOidcConfig `json:"oidcConfig,omitempty" yaml:"oidcConfig,omitempty"`

	/*
	   Support for proxy configuration.
	   Structure is documented below.
	*/
	ProxyConfig types.Container_AttachedClusterProxyConfig `json:"proxyConfig,omitempty" yaml:"proxyConfig,omitempty"`

	/*
	   Binary Authorization configuration.
	   Structure is documented below.
	*/
	BinaryAuthorization types.Container_AttachedClusterBinaryAuthorization `json:"binaryAuthorization,omitempty" yaml:"binaryAuthorization,omitempty"`

	/*
	   Configuration related to the cluster RBAC settings.
	   Structure is documented below.
	*/
	Authorization types.Container_AttachedClusterAuthorization `json:"authorization,omitempty" yaml:"authorization,omitempty"`

	// Policy to determine what flags to send on delete.
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	/*
	   A human readable description of this attached cluster. Cannot be longer
	   than 255 UTF-8 encoded bytes.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The platform version for the cluster (e.g. `1.23.0-gke.1`).
	PlatformVersion string `json:"platformVersion,omitempty" yaml:"platformVersion,omitempty"`

	/*
	   Optional. Annotations on the cluster. This field has the same
	   restrictions as Kubernetes annotations. The total size of all keys and
	   values combined is limited to 256k. Key can have 2 segments: prefix (optional)
	   and name (required), separated by a slash (/). Prefix must be a DNS subdomain.
	   Name must be 63 characters or less, begin and end with alphanumerics,
	   with dashes (-), underscores (_), dots (.), and alphanumerics between.

	   --Note--: This field is non-authoritative, and will only manage the annotations present in your configuration.
	   Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	*/
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	/*
	   The Kubernetes distribution of the underlying attached cluster. Supported values:
	   "eks", "aks".
	*/
	Distribution string `json:"distribution,omitempty" yaml:"distribution,omitempty"`
}
