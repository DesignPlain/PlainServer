package types

type Appsync_DataSourceHttpConfig struct {
	// Authorization configuration in case the HTTP endpoint requires authorization. See Authorization Config.
	AuthorizationConfig Appsync_DataSourceHttpConfigAuthorizationConfig `json:"authorizationConfig,omitempty" yaml:"authorizationConfig,omitempty"`

	// HTTP URL.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
}
