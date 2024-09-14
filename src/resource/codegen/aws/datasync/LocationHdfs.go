package datasync

import types "libds/aws/types"

type LocationHdfs struct {
	// The Kerberos key table (keytab) that contains mappings between the defined Kerberos principal and the encrypted keys. Use `kerberos_keytab_base64` instead whenever the value is not a valid UTF-8 string. If `KERBEROS` is specified for `authentication_type`, this parameter (or `kerberos_keytab_base64`) is required.
	KerberosKeytab string `json:"kerberosKeytab,omitempty" yaml:"kerberosKeytab,omitempty"`

	// The URI of the HDFS cluster's Key Management Server (KMS).
	KmsKeyProviderUri string `json:"kmsKeyProviderUri,omitempty" yaml:"kmsKeyProviderUri,omitempty"`

	// The user name used to identify the client on the host operating system. If `SIMPLE` is specified for `authentication_type`, this parameter is required.
	SimpleUser string `json:"simpleUser,omitempty" yaml:"simpleUser,omitempty"`

	// Use instead of `kerberos_keytab` to pass base64-encoded binary data directly. If `KERBEROS` is specified for `authentication_type`, this parameter (or `kerberos_keytab`) is required.
	KerberosKeytabBase64 string `json:"kerberosKeytabBase64,omitempty" yaml:"kerberosKeytabBase64,omitempty"`

	// The krb5.conf file that contains the Kerberos configuration information. Use `kerberos_krb5_conf_base64` instead whenever the value is not a valid UTF-8 string. If `KERBEROS` is specified for `authentication_type`, this parameter (or `kerberos_krb5_conf_base64`) is required.
	KerberosKrb5Conf string `json:"kerberosKrb5Conf,omitempty" yaml:"kerberosKrb5Conf,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The size of data blocks to write into the HDFS cluster. The block size must be a multiple of 512 bytes. The default block size is 128 mebibytes (MiB).
	BlockSize int `json:"blockSize,omitempty" yaml:"blockSize,omitempty"`

	// The Kerberos principal with access to the files and folders on the HDFS cluster. If `KERBEROS` is specified for `authentication_type`, this parameter is required.
	KerberosPrincipal string `json:"kerberosPrincipal,omitempty" yaml:"kerberosPrincipal,omitempty"`

	// The NameNode that manages the HDFS namespace. The NameNode performs operations such as opening, closing, and renaming files and directories. The NameNode contains the information to map blocks of data to the DataNodes. You can use only one NameNode. See configuration below.
	NameNodes []types.Datasync_LocationHdfsNameNode `json:"nameNodes,omitempty" yaml:"nameNodes,omitempty"`

	// The Quality of Protection (QOP) configuration specifies the Remote Procedure Call (RPC) and data transfer protection settings configured on the Hadoop Distributed File System (HDFS) cluster. If `qop_configuration` isn't specified, `rpc_protection` and `data_transfer_protection` default to `PRIVACY`. If you set RpcProtection or DataTransferProtection, the other parameter assumes the same value.  See configuration below.
	QopConfiguration types.Datasync_LocationHdfsQopConfiguration `json:"qopConfiguration,omitempty" yaml:"qopConfiguration,omitempty"`

	// The number of DataNodes to replicate the data to when writing to the HDFS cluster. By default, data is replicated to three DataNodes.
	ReplicationFactor int `json:"replicationFactor,omitempty" yaml:"replicationFactor,omitempty"`

	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns []string `json:"agentArns,omitempty" yaml:"agentArns,omitempty"`

	// The type of authentication used to determine the identity of the user. Valid values are `SIMPLE` and `KERBEROS`.
	AuthenticationType string `json:"authenticationType,omitempty" yaml:"authenticationType,omitempty"`

	// Use instead of `kerberos_krb5_conf` to pass base64-encoded binary data directly. If `KERBEROS` is specified for `authentication_type`, this parameter (or `kerberos_krb5_conf`) is required.
	KerberosKrb5ConfBase64 string `json:"kerberosKrb5ConfBase64,omitempty" yaml:"kerberosKrb5ConfBase64,omitempty"`

	// A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to /.
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`
}
