package rds

type ProxyTarget struct {
	/*
	   DB cluster identifier.

	   --NOTE:-- Either `db_instance_identifier` or `db_cluster_identifier` should be specified and both should not be specified together
	*/
	DbClusterIdentifier string `json:"dbClusterIdentifier,omitempty" yaml:"dbClusterIdentifier,omitempty"`

	// DB instance identifier.
	DbInstanceIdentifier string `json:"dbInstanceIdentifier,omitempty" yaml:"dbInstanceIdentifier,omitempty"`

	// The name of the DB proxy.
	DbProxyName string `json:"dbProxyName,omitempty" yaml:"dbProxyName,omitempty"`

	// The name of the target group.
	TargetGroupName string `json:"targetGroupName,omitempty" yaml:"targetGroupName,omitempty"`
}
