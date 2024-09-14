package types

type Msk_ClusterBrokerNodeGroupInfo struct {
	// A block that contains information about storage volumes attached to MSK broker nodes. See below.
	StorageInfo Msk_ClusterBrokerNodeGroupInfoStorageInfo `json:"storageInfo,omitempty" yaml:"storageInfo,omitempty"`

	// The distribution of broker nodes across availability zones ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-model-brokerazdistribution)). Currently the only valid value is `DEFAULT`.
	AzDistribution string `json:"azDistribution,omitempty" yaml:"azDistribution,omitempty"`

	// A list of subnets to connect to in client VPC ([documentation](https://docs.aws.amazon.com/msk/1.0/apireference/clusters.html#clusters-prop-brokernodegroupinfo-clientsubnets)).
	ClientSubnets []string `json:"clientSubnets,omitempty" yaml:"clientSubnets,omitempty"`

	// Information about the cluster access configuration. See below. For security reasons, you can't turn on public access while creating an MSK cluster. However, you can update an existing cluster to make it publicly accessible. You can also create a new cluster and then update it to make it publicly accessible ([documentation](https://docs.aws.amazon.com/msk/latest/developerguide/public-access.html)).
	ConnectivityInfo Msk_ClusterBrokerNodeGroupInfoConnectivityInfo `json:"connectivityInfo,omitempty" yaml:"connectivityInfo,omitempty"`

	// Specify the instance type to use for the kafka brokersE.g., kafka.m5.large. ([Pricing info](https://aws.amazon.com/msk/pricing/))
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// A list of the security groups to associate with the elastic network interfaces to control who can communicate with the cluster.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`
}
