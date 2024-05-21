package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type MangedSslCertificate struct {
	/*
	   Enum field whose value is always `MANAGED` - used to signal to the API
	   which type this is.
	   Default value is `MANAGED`.
	   Possible values are: `MANAGED`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The unique identifier for the resource.
	CertificateId int `json:"certificateId,omitempty" yaml:"certificateId,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Properties relevant to a managed certificate.  These will be used if the
	   certificate is managed (as indicated by a value of `MANAGED` in `type`).
	   Structure is documented below.
	*/
	Managed types.Compute_MangedSslCertificateManaged `json:"managed,omitempty" yaml:"managed,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.

	   These are in the same namespace as the managed SSL certificates.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
