package types

type Storagegateway_GatewaySmbActiveDirectorySettings struct {
	/*
	   The organizational unit (OU) is a container in an Active Directory that can hold users, groups,
	   computers, and other OUs and this parameter specifies the OU that the gateway will join within the AD domain.
	*/
	OrganizationalUnit string `json:"organizationalUnit,omitempty" yaml:"organizationalUnit,omitempty"`

	// The password of the user who has permission to add the gateway to the Active Directory domain.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Specifies the time in seconds, in which the JoinDomain operation must complete. The default is `20` seconds.
	TimeoutInSeconds int `json:"timeoutInSeconds,omitempty" yaml:"timeoutInSeconds,omitempty"`

	// The user name of user who has permission to add the gateway to the Active Directory domain.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	//
	ActiveDirectoryStatus string `json:"activeDirectoryStatus,omitempty" yaml:"activeDirectoryStatus,omitempty"`

	/*
	   List of IPv4 addresses, NetBIOS names, or host names of your domain server.
	   If you need to specify the port number include it after the colon (“:”). For example, `mydc.mydomain.com:389`.
	*/
	DomainControllers []string `json:"domainControllers,omitempty" yaml:"domainControllers,omitempty"`

	// The name of the domain that you want the gateway to join.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
}
