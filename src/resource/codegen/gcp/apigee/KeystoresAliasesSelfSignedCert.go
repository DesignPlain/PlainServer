package apigee

import types "libds/gcp/types"

type KeystoresAliasesSelfSignedCert struct {
	// The Apigee keystore name associated in an Apigee environment
	Keystore string `json:"keystore,omitempty" yaml:"keystore,omitempty"`

	/*
	   Alias for the key/certificate pair. Values must match the regular expression [\w\s-.]{1,255}.
	   This must be provided for all formats except selfsignedcert; self-signed certs may specify the alias in either
	   this parameter or the JSON body.
	*/
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`

	// Validity duration of certificate, in days. Accepts positive non-zero value. Defaults to 365.
	CertValidityInDays int `json:"certValidityInDays,omitempty" yaml:"certValidityInDays,omitempty"`

	// The Apigee environment name
	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`

	/*
	   Subject details.
	   Structure is documented below.
	*/
	Subject types.Apigee_KeystoresAliasesSelfSignedCertSubject `json:"subject,omitempty" yaml:"subject,omitempty"`

	/*
	   List of alternative host names. Maximum length is 255 characters for each value.
	   Structure is documented below.
	*/
	SubjectAlternativeDnsNames types.Apigee_KeystoresAliasesSelfSignedCertSubjectAlternativeDnsNames `json:"subjectAlternativeDnsNames,omitempty" yaml:"subjectAlternativeDnsNames,omitempty"`

	// Key size. Default and maximum value is 2048 bits.
	KeySize string `json:"keySize,omitempty" yaml:"keySize,omitempty"`

	// The Apigee Organization name associated with the Apigee environment
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	// Signature algorithm to generate private key. Valid values are SHA512withRSA, SHA384withRSA, and SHA256withRSA
	SigAlg string `json:"sigAlg,omitempty" yaml:"sigAlg,omitempty"`
}
