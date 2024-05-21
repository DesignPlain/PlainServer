package types

type Certificatemanager_TrustConfigTrustStore struct {
	/*
	   Set of intermediate CA certificates used for the path building phase of chain validation.
	   The field is currently not supported if trust config is used for the workload certificate feature.
	   Structure is documented below.
	*/
	IntermediateCas []Certificatemanager_TrustConfigTrustStoreIntermediateCa `json:"intermediateCas,omitempty" yaml:"intermediateCas,omitempty"`

	/*
	   List of Trust Anchors to be used while performing validation against a given TrustStore.
	   Structure is documented below.
	*/
	TrustAnchors []Certificatemanager_TrustConfigTrustStoreTrustAnchor `json:"trustAnchors,omitempty" yaml:"trustAnchors,omitempty"`
}
