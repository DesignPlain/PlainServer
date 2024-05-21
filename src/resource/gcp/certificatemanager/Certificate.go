package certificatemanager

import types "DesignSphere_Server/src/resource/gcp/types"

type Certificate struct {
	// The Certificate Manager location. If not specified, "global" is used.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Configuration and state of a Managed Certificate.
	   Certificate Manager provisions and renews Managed Certificates
	   automatically, for as long as it's authorized to do so.
	   Structure is documented below.
	*/
	Managed types.Certificatemanager_CertificateManaged `json:"managed,omitempty" yaml:"managed,omitempty"`

	/*
	   A user-defined name of the certificate. Certificate names must be unique
	   The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]- which means the first character must be a letter,
	   and all following characters must be a dash, underscore, letter or digit.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The scope of the certificate.
	   DEFAULT: Certificates with default scope are served from core Google data centers.
	   If unsure, choose this option.
	   EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served from Edge Points of Presence.
	   See https://cloud.google.com/vpc/docs/edge-locations.
	   ALL_REGIONS: Certificates with ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs).
	   See https://cloud.google.com/compute/docs/regions-zones
	*/
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	/*
	   Certificate data for a SelfManaged Certificate.
	   SelfManaged Certificates are uploaded by the user. Updating such
	   certificates before they expire remains the user's responsibility.
	   Structure is documented below.
	*/
	SelfManaged types.Certificatemanager_CertificateSelfManaged `json:"selfManaged,omitempty" yaml:"selfManaged,omitempty"`

	// A human-readable description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Set of label tags associated with the Certificate resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
