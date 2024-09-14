package rds

import types "libds/aws/types"

type ProxyDefaultTargetGroup struct {
	// The settings that determine the size and behavior of the connection pool for the target group.
	ConnectionPoolConfig types.Rds_ProxyDefaultTargetGroupConnectionPoolConfig `json:"connectionPoolConfig,omitempty" yaml:"connectionPoolConfig,omitempty"`

	// Name of the RDS DB Proxy.
	DbProxyName string `json:"dbProxyName,omitempty" yaml:"dbProxyName,omitempty"`
}
