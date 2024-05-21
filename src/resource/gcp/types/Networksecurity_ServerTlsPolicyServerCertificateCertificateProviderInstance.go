package types

type Networksecurity_ServerTlsPolicyServerCertificateCertificateProviderInstance struct {
	// Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.
	PluginInstance string `json:"pluginInstance,omitempty" yaml:"pluginInstance,omitempty"`
}
