package certificateauthority

import types "DesignSphere_Server/src/resource/gcp/types"

type Authority struct {
	// The user provided Resource ID for this Certificate Authority.
	CertificateAuthorityId string `json:"certificateAuthorityId,omitempty" yaml:"certificateAuthorityId,omitempty"`

	/*
	   This field allows the CA to be deleted even if the CA has active certs. Active certs include both unrevoked and unexpired certs.
	   Use with care. Defaults to `false`.
	*/
	IgnoreActiveCertificatesOnDeletion bool `json:"ignoreActiveCertificatesOnDeletion,omitempty" yaml:"ignoreActiveCertificatesOnDeletion,omitempty"`

	/*
	   If this flag is set, the Certificate Authority will be deleted as soon as
	   possible without a 30-day grace period where undeletion would have been
	   allowed. If you proceed, there will be no way to recover this CA.
	   Use with care. Defaults to `false`.
	*/
	SkipGracePeriod bool `json:"skipGracePeriod,omitempty" yaml:"skipGracePeriod,omitempty"`

	/*
	   The Type of this CertificateAuthority.
	   > --Note:-- For `SUBORDINATE` Certificate Authorities, they need to
	   be activated before they can issue certificates.
	   Default value is `SELF_SIGNED`.
	   Possible values are: `SELF_SIGNED`, `SUBORDINATE`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   The config used to create a self-signed X.509 certificate or CSR.
	   Structure is documented below.
	*/
	Config types.Certificateauthority_AuthorityConfig `json:"config,omitempty" yaml:"config,omitempty"`

	/*
	   Whether or not to allow Terraform to destroy the CertificateAuthority. Unless this field is set to false in Terraform
	   state, a 'terraform destroy' or 'terraform apply' that would delete the instance will fail.
	*/
	DeletionProtection bool `json:"deletionProtection,omitempty" yaml:"deletionProtection,omitempty"`

	/*
	   The name of a Cloud Storage bucket where this CertificateAuthority will publish content,
	   such as the CA certificate and CRLs. This must be a bucket name, without any prefixes
	   (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named
	   my-bucket, you would simply specify `my-bucket`. If not specified, a managed bucket will be
	   created.
	*/
	GcsBucket string `json:"gcsBucket,omitempty" yaml:"gcsBucket,omitempty"`

	/*
	   Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority
	   is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA
	   certificate. Otherwise, it is used to sign a CSR.
	   Structure is documented below.
	*/
	KeySpec types.Certificateauthority_AuthorityKeySpec `json:"keySpec,omitempty" yaml:"keySpec,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The desired lifetime of the CA certificate. Used to create the "notBeforeTime" and
	   "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	   fractional digits, terminated by 's'. Example: "3.5s".
	*/
	Lifetime string `json:"lifetime,omitempty" yaml:"lifetime,omitempty"`

	// The signed CA certificate issued from the subordinated CA's CSR. This is needed when activating the subordiante CA with a third party issuer.
	PemCaCertificate string `json:"pemCaCertificate,omitempty" yaml:"pemCaCertificate,omitempty"`

	/*
	   If this is a subordinate CertificateAuthority, this field will be set
	   with the subordinate configuration, which describes its issuers.
	   Structure is documented below.
	*/
	SubordinateConfig types.Certificateauthority_AuthoritySubordinateConfig `json:"subordinateConfig,omitempty" yaml:"subordinateConfig,omitempty"`

	// Desired state of the CertificateAuthority. Set this field to `STAGED` to create a `STAGED` root CA.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`

	/*
	   Labels with user-defined metadata.
	   An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	   "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Location of the CertificateAuthority. A full list of valid locations can be found by
	   running `gcloud privateca locations list`.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The name of the CaPool this Certificate Authority belongs to.
	Pool string `json:"pool,omitempty" yaml:"pool,omitempty"`
}
