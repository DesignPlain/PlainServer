package integrationconnectors

import types "libds/gcp/types"

type Connection struct {
	/*
	   Log configuration for the connection.
	   Structure is documented below.
	*/
	LogConfig types.Integrationconnectors_ConnectionLogConfig `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`

	// Service account needed for runtime plane to access Google Cloud resources.
	ServiceAccount string `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`

	// An arbitrary description for the Conection.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Eventing enablement type. Will be nil if eventing is not enabled.
	   Possible values are: `EVENTING_AND_CONNECTION`, `ONLY_EVENTING`.
	*/
	EventingEnablementType string `json:"eventingEnablementType,omitempty" yaml:"eventingEnablementType,omitempty"`

	/*
	   Resource labels to represent user provided metadata.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Determines whether or no a connection is locked. If locked, a reason must be specified.
	   Structure is documented below.
	*/
	LockConfig types.Integrationconnectors_ConnectionLockConfig `json:"lockConfig,omitempty" yaml:"lockConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Node configuration for the connection.
	   Structure is documented below.
	*/
	NodeConfig types.Integrationconnectors_ConnectionNodeConfig `json:"nodeConfig,omitempty" yaml:"nodeConfig,omitempty"`

	/*
	   SSL Configuration of a connection
	   Structure is documented below.
	*/
	SslConfig types.Integrationconnectors_ConnectionSslConfig `json:"sslConfig,omitempty" yaml:"sslConfig,omitempty"`

	// Suspended indicates if a user has suspended a connection or not.
	Suspended bool `json:"suspended,omitempty" yaml:"suspended,omitempty"`

	/*
	   authConfig for the connection.
	   Structure is documented below.
	*/
	AuthConfig types.Integrationconnectors_ConnectionAuthConfig `json:"authConfig,omitempty" yaml:"authConfig,omitempty"`

	/*
	   Config Variables for the connection.
	   Structure is documented below.
	*/
	ConfigVariables []types.Integrationconnectors_ConnectionConfigVariable `json:"configVariables,omitempty" yaml:"configVariables,omitempty"`

	// connectorVersion of the Connector.
	ConnectorVersion string `json:"connectorVersion,omitempty" yaml:"connectorVersion,omitempty"`

	/*
	   Define the Connectors target endpoint.
	   Structure is documented below.
	*/
	DestinationConfigs []types.Integrationconnectors_ConnectionDestinationConfig `json:"destinationConfigs,omitempty" yaml:"destinationConfigs,omitempty"`

	/*
	   Eventing Configuration of a connection
	   Structure is documented below.
	*/
	EventingConfig types.Integrationconnectors_ConnectionEventingConfig `json:"eventingConfig,omitempty" yaml:"eventingConfig,omitempty"`

	// Location in which Connection needs to be created.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Name of Connection needs to be created.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
