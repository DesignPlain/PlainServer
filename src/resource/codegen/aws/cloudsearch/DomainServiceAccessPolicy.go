package cloudsearch

type DomainServiceAccessPolicy struct {
	// The access rules you want to configure. These rules replace any existing rules. See the [AWS documentation](https://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-access.html) for details.
	AccessPolicy string `json:"accessPolicy,omitempty" yaml:"accessPolicy,omitempty"`

	// The CloudSearch domain name the policy applies to.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
}
