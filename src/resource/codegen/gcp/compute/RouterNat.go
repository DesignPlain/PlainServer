package compute

import types "libds/gcp/types"

type RouterNat struct {
	/*
	   A list of URLs of the IP resources to be drained. These IPs must be
	   valid static external IPs that have been assigned to the NAT.
	*/
	DrainNatIps []string `json:"drainNatIps,omitempty" yaml:"drainNatIps,omitempty"`

	/*
	   Enable Dynamic Port Allocation.
	   If minPortsPerVm is set, minPortsPerVm must be set to a power of two greater than or equal to 32.
	   If minPortsPerVm is not set, a minimum of 32 ports will be allocated to a VM from this NAT config.
	   If maxPortsPerVm is set, maxPortsPerVm must be set to a power of two greater than minPortsPerVm.
	   If maxPortsPerVm is not set, a maximum of 65536 ports will be allocated to a VM from this NAT config.
	   Mutually exclusive with enableEndpointIndependentMapping.
	*/
	EnableDynamicPortAllocation bool `json:"enableDynamicPortAllocation,omitempty" yaml:"enableDynamicPortAllocation,omitempty"`

	// Region where the router and NAT reside.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   Timeout (in seconds) for TCP transitory connections.
	   Defaults to 30s if not set.
	*/
	TcpTransitoryIdleTimeoutSec int `json:"tcpTransitoryIdleTimeoutSec,omitempty" yaml:"tcpTransitoryIdleTimeoutSec,omitempty"`

	// Timeout (in seconds) for UDP connections. Defaults to 30s if not set.
	UdpIdleTimeoutSec int `json:"udpIdleTimeoutSec,omitempty" yaml:"udpIdleTimeoutSec,omitempty"`

	// Minimum number of ports allocated to a VM from this NAT. Defaults to 64 for static port allocation and 32 dynamic port allocation if not set.
	MinPortsPerVm int `json:"minPortsPerVm,omitempty" yaml:"minPortsPerVm,omitempty"`

	/*
	   Name of the NAT service. The name must be 1-63 characters long and
	   comply with RFC1035.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The name of the Cloud Router in which this NAT will be configured.


	   - - -
	*/
	Router string `json:"router,omitempty" yaml:"router,omitempty"`

	/*
	   How NAT should be configured per Subnetwork.
	   If `ALL_SUBNETWORKS_ALL_IP_RANGES`, all of the
	   IP ranges in every Subnetwork are allowed to Nat.
	   If `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, all of the primary IP
	   ranges in every Subnetwork are allowed to Nat.
	   `LIST_OF_SUBNETWORKS`: A list of Subnetworks are allowed to Nat
	   (specified in the field subnetwork below). Note that if this field
	   contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	   ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	   other RouterNat section in any Router for this network in this region.
	   Possible values are: `ALL_SUBNETWORKS_ALL_IP_RANGES`, `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES`, `LIST_OF_SUBNETWORKS`.
	*/
	SourceSubnetworkIpRangesToNat string `json:"sourceSubnetworkIpRangesToNat,omitempty" yaml:"sourceSubnetworkIpRangesToNat,omitempty"`

	/*
	   One or more subnetwork NAT configurations. Only used if
	   `source_subnetwork_ip_ranges_to_nat` is set to `LIST_OF_SUBNETWORKS`
	   Structure is documented below.
	*/
	Subnetworks []types.Compute_RouterNatSubnetwork `json:"subnetworks,omitempty" yaml:"subnetworks,omitempty"`

	/*
	   Timeout (in seconds) for TCP established connections.
	   Defaults to 1200s if not set.
	*/
	TcpEstablishedIdleTimeoutSec int `json:"tcpEstablishedIdleTimeoutSec,omitempty" yaml:"tcpEstablishedIdleTimeoutSec,omitempty"`

	// Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.
	IcmpIdleTimeoutSec int `json:"icmpIdleTimeoutSec,omitempty" yaml:"icmpIdleTimeoutSec,omitempty"`

	/*
	   Configuration for logging on NAT
	   Structure is documented below.
	*/
	LogConfig types.Compute_RouterNatLogConfig `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`

	/*
	   Timeout (in seconds) for TCP connections that are in TIME_WAIT state.
	   Defaults to 120s if not set.
	*/
	TcpTimeWaitTimeoutSec int `json:"tcpTimeWaitTimeoutSec,omitempty" yaml:"tcpTimeWaitTimeoutSec,omitempty"`

	/*
	   Enable endpoint independent mapping.
	   For more information see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs).
	*/
	EnableEndpointIndependentMapping bool `json:"enableEndpointIndependentMapping,omitempty" yaml:"enableEndpointIndependentMapping,omitempty"`

	/*
	   Maximum number of ports allocated to a VM from this NAT.
	   This field can only be set when enableDynamicPortAllocation is enabled.
	*/
	MaxPortsPerVm int `json:"maxPortsPerVm,omitempty" yaml:"maxPortsPerVm,omitempty"`

	/*
	   How external IPs should be allocated for this NAT. Valid values are
	   `AUTO_ONLY` for only allowing NAT IPs allocated by Google Cloud
	   Platform, or `MANUAL_ONLY` for only user-allocated NAT IP addresses.
	   Possible values are: `MANUAL_ONLY`, `AUTO_ONLY`.
	*/
	NatIpAllocateOption string `json:"natIpAllocateOption,omitempty" yaml:"natIpAllocateOption,omitempty"`

	/*
	   Self-links of NAT IPs. Only valid if natIpAllocateOption
	   is set to MANUAL_ONLY.
	*/
	NatIps []string `json:"natIps,omitempty" yaml:"natIps,omitempty"`

	/*
	   A list of rules associated with this NAT.
	   Structure is documented below.
	*/
	Rules []types.Compute_RouterNatRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	/*
	   Indicates whether this NAT is used for public or private IP translation.
	   If unspecified, it defaults to PUBLIC.
	   If `PUBLIC` NAT used for public IP translation.
	   If `PRIVATE` NAT used for private IP translation.
	   Default value is `PUBLIC`.
	   Possible values are: `PUBLIC`, `PRIVATE`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
