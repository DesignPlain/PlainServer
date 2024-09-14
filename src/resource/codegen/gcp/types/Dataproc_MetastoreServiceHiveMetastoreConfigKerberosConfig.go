package types

type Dataproc_MetastoreServiceHiveMetastoreConfigKerberosConfig struct {
	/*
	   A Kerberos keytab file that can be used to authenticate a service principal with a Kerberos Key Distribution Center (KDC).
	   Structure is documented below.
	*/
	Keytab Dataproc_MetastoreServiceHiveMetastoreConfigKerberosConfigKeytab `json:"keytab,omitempty" yaml:"keytab,omitempty"`

	// A Cloud Storage URI that specifies the path to a krb5.conf file. It is of the form gs://{bucket_name}/path/to/krb5.conf, although the file does not need to be named krb5.conf explicitly.
	Krb5ConfigGcsUri string `json:"krb5ConfigGcsUri,omitempty" yaml:"krb5ConfigGcsUri,omitempty"`

	// A Kerberos principal that exists in the both the keytab the KDC to authenticate as. A typical principal is of the form "primary/instance@REALM", but there is no exact format.
	Principal string `json:"principal,omitempty" yaml:"principal,omitempty"`
}
