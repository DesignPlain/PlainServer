package types

type Firebase_HostingVersionConfigRewriteRun struct {
	// User-defined ID of the Cloud Run service.
	ServiceId string `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`

	// Optional. User-provided region where the Cloud Run service is hosted. Defaults to `us-central1` if not supplied.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
