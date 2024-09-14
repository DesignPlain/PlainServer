package types

type Emr_ClusterKerberosAttributes struct {
	// Active Directory password for `ad_domain_join_user`. This provider cannot perform drift detection of this configuration.
	AdDomainJoinPassword string `json:"adDomainJoinPassword,omitempty" yaml:"adDomainJoinPassword,omitempty"`

	// Required only when establishing a cross-realm trust with an Active Directory domain. A user with sufficient privileges to join resources to the domain. This provider cannot perform drift detection of this configuration.
	AdDomainJoinUser string `json:"adDomainJoinUser,omitempty" yaml:"adDomainJoinUser,omitempty"`

	// Required only when establishing a cross-realm trust with a KDC in a different realm. The cross-realm principal password, which must be identical across realms. This provider cannot perform drift detection of this configuration.
	CrossRealmTrustPrincipalPassword string `json:"crossRealmTrustPrincipalPassword,omitempty" yaml:"crossRealmTrustPrincipalPassword,omitempty"`

	// Password used within the cluster for the kadmin service on the cluster-dedicated KDC, which maintains Kerberos principals, password policies, and keytabs for the cluster. This provider cannot perform drift detection of this configuration.
	KdcAdminPassword string `json:"kdcAdminPassword,omitempty" yaml:"kdcAdminPassword,omitempty"`

	// Name of the Kerberos realm to which all nodes in a cluster belong. For example, `EC2.INTERNAL`
	Realm string `json:"realm,omitempty" yaml:"realm,omitempty"`
}
