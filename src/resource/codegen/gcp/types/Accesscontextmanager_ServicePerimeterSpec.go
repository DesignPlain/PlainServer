package types

type Accesscontextmanager_ServicePerimeterSpec struct {
	/*
	   List of EgressPolicies to apply to the perimeter. A perimeter may
	   have multiple EgressPolicies, each of which is evaluated separately.
	   Access is granted if any EgressPolicy grants it. Must be empty for
	   a perimeter bridge.
	   Structure is documented below.
	*/
	EgressPolicies []Accesscontextmanager_ServicePerimeterSpecEgressPolicy `json:"egressPolicies,omitempty" yaml:"egressPolicies,omitempty"`

	/*
	   List of `IngressPolicies` to apply to the perimeter. A perimeter may
	   have multiple `IngressPolicies`, each of which is evaluated
	   separately. Access is granted if any `Ingress Policy` grants it.
	   Must be empty for a perimeter bridge.
	   Structure is documented below.
	*/
	IngressPolicies []Accesscontextmanager_ServicePerimeterSpecIngressPolicy `json:"ingressPolicies,omitempty" yaml:"ingressPolicies,omitempty"`

	/*
	   A list of GCP resources that are inside of the service perimeter.
	   Currently only projects are allowed.
	   Format: projects/{project_number}
	*/
	Resources []string `json:"resources,omitempty" yaml:"resources,omitempty"`

	/*
	   GCP services that are subject to the Service Perimeter
	   restrictions. Must contain a list of services. For example, if
	   `storage.googleapis.com` is specified, access to the storage
	   buckets inside the perimeter must meet the perimeter's access
	   restrictions.
	*/
	RestrictedServices []string `json:"restrictedServices,omitempty" yaml:"restrictedServices,omitempty"`

	/*
	   Specifies how APIs are allowed to communicate within the Service
	   Perimeter.
	   Structure is documented below.
	*/
	VpcAccessibleServices Accesscontextmanager_ServicePerimeterSpecVpcAccessibleServices `json:"vpcAccessibleServices,omitempty" yaml:"vpcAccessibleServices,omitempty"`

	/*
	   A list of AccessLevel resource names that allow resources within
	   the ServicePerimeter to be accessed from the internet.
	   AccessLevels listed must be in the same policy as this
	   ServicePerimeter. Referencing a nonexistent AccessLevel is a
	   syntax error. If no AccessLevel names are listed, resources within
	   the perimeter can only be accessed via GCP calls with request
	   origins within the perimeter. For Service Perimeter Bridge, must
	   be empty.
	   Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
	*/
	AccessLevels []string `json:"accessLevels,omitempty" yaml:"accessLevels,omitempty"`
}
