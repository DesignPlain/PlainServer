package ec2

type TrafficMirrorSession struct {
	// ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
	PacketLength int `json:"packetLength,omitempty" yaml:"packetLength,omitempty"`

	// The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets.
	SessionNumber int `json:"sessionNumber,omitempty" yaml:"sessionNumber,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// ID of the traffic mirror filter to be used
	TrafficMirrorFilterId string `json:"trafficMirrorFilterId,omitempty" yaml:"trafficMirrorFilterId,omitempty"`

	// ID of the traffic mirror target to be used
	TrafficMirrorTargetId string `json:"trafficMirrorTargetId,omitempty" yaml:"trafficMirrorTargetId,omitempty"`

	// The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
	VirtualNetworkId int `json:"virtualNetworkId,omitempty" yaml:"virtualNetworkId,omitempty"`

	// A description of the traffic mirror session.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
