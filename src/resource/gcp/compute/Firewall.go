package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type Firewall struct {
	/*
	   If destination ranges are specified, the firewall will apply only to
	   traffic that has destination IP address in these ranges. These ranges
	   must be expressed in CIDR format. IPv4 or IPv6 ranges are supported.
	*/
	DestinationRanges []string `json:"destinationRanges,omitempty" yaml:"destinationRanges,omitempty"`

	/*
	   This field denotes the logging options for a particular firewall rule.
	   If defined, logging is enabled, and logs will be exported to Cloud Logging.
	   Structure is documented below.
	*/
	LogConfig types.Compute_FirewallLogConfig `json:"logConfig,omitempty" yaml:"logConfig,omitempty"`

	/*
	   The name or self_link of the network to attach this firewall to.


	   - - -
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   Priority for this rule. This is an integer between 0 and 65535, both
	   inclusive. When not specified, the value assumed is 1000. Relative
	   priorities determine precedence of conflicting rules. Lower value of
	   priority implies higher precedence (eg, a rule with priority 0 has
	   higher precedence than a rule with priority 1). DENY rules take
	   precedence over ALLOW rules having equal priority.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A list of service accounts indicating sets of instances located in the
	   network that may make network connections as specified in allowed[].
	   targetServiceAccounts cannot be used at the same time as targetTags or
	   sourceTags. If neither targetServiceAccounts nor targetTags are
	   specified, the firewall rule applies to all instances on the specified
	   network.
	*/
	TargetServiceAccounts []string `json:"targetServiceAccounts,omitempty" yaml:"targetServiceAccounts,omitempty"`

	/*
	   An optional description of this resource. Provide this property when
	   you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   This field denotes whether to enable logging for a particular firewall rule.
	   If logging is enabled, logs will be exported to Stackdriver. Deprecated in favor of `log_config`
	*/
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   If source tags are specified, the firewall will apply only to traffic
	   with source IP that belongs to a tag listed in source tags. Source
	   tags cannot be used to control traffic to an instance's external IP
	   address. Because tags are associated with an instance, not an IP
	   address. One or both of sourceRanges and sourceTags may be set. If
	   both properties are set, the firewall will apply to traffic that has
	   source IP address within sourceRanges OR the source IP that belongs to
	   a tag listed in the sourceTags property. The connection does not need
	   to match both properties for the firewall to apply. For INGRESS traffic,
	   one of `source_ranges`, `source_tags` or `source_service_accounts` is required.
	*/
	SourceTags []string `json:"sourceTags,omitempty" yaml:"sourceTags,omitempty"`

	/*
	   A list of instance tags indicating sets of instances located in the
	   network that may make network connections as specified in allowed[].
	   If no targetTags are specified, the firewall rule applies to all
	   instances on the specified network.
	*/
	TargetTags []string `json:"targetTags,omitempty" yaml:"targetTags,omitempty"`

	/*
	   Denotes whether the firewall rule is disabled, i.e not applied to the
	   network it is associated with. When set to true, the firewall rule is
	   not enforced and the network behaves as if it did not exist. If this
	   is unspecified, the firewall rule will be enabled.
	*/
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	/*
	   If source ranges are specified, the firewall will apply only to
	   traffic that has source IP address in these ranges. These ranges must
	   be expressed in CIDR format. One or both of sourceRanges and
	   sourceTags may be set. If both properties are set, the firewall will
	   apply to traffic that has source IP address within sourceRanges OR the
	   source IP that belongs to a tag listed in the sourceTags property. The
	   connection does not need to match both properties for the firewall to
	   apply. IPv4 or IPv6 ranges are supported. For INGRESS traffic, one of
	   `source_ranges`, `source_tags` or `source_service_accounts` is required.
	*/
	SourceRanges []string `json:"sourceRanges,omitempty" yaml:"sourceRanges,omitempty"`

	/*
	   Direction of traffic to which this firewall applies; default is
	   INGRESS. Note: For INGRESS traffic, one of `source_ranges`,
	   `source_tags` or `source_service_accounts` is required.
	   Possible values are: `INGRESS`, `EGRESS`.
	*/
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`

	/*
	   The list of DENY rules specified by this firewall. Each rule specifies
	   a protocol and port-range tuple that describes a denied connection.
	   Structure is documented below.
	*/
	Denies []types.Compute_FirewallDeny `json:"denies,omitempty" yaml:"denies,omitempty"`

	/*
	   If source service accounts are specified, the firewall will apply only
	   to traffic originating from an instance with a service account in this
	   list. Source service accounts cannot be used to control traffic to an
	   instance's external IP address because service accounts are associated
	   with an instance, not an IP address. sourceRanges can be set at the
	   same time as sourceServiceAccounts. If both are set, the firewall will
	   apply to traffic that has source IP address within sourceRanges OR the
	   source IP belongs to an instance with service account listed in
	   sourceServiceAccount. The connection does not need to match both
	   properties for the firewall to apply. sourceServiceAccounts cannot be
	   used at the same time as sourceTags or targetTags. For INGRESS traffic,
	   one of `source_ranges`, `source_tags` or `source_service_accounts` is required.
	*/
	SourceServiceAccounts []string `json:"sourceServiceAccounts,omitempty" yaml:"sourceServiceAccounts,omitempty"`

	/*
	   The list of ALLOW rules specified by this firewall. Each rule
	   specifies a protocol and port-range tuple that describes a permitted
	   connection.
	   Structure is documented below.
	*/
	Allows []types.Compute_FirewallAllow `json:"allows,omitempty" yaml:"allows,omitempty"`
}
