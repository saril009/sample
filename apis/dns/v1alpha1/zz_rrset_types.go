/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ItemsObservation struct {

	// A Boolean flag indicating whether or not parts of the record are unable to be explicitly managed.
	IsProtected *bool `json:"isProtected,omitempty" tf:"is_protected,omitempty"`

	// A unique identifier for the record within its zone.
	RecordHash *string `json:"recordHash,omitempty" tf:"record_hash,omitempty"`

	// The latest version of the record's zone in which its RRSet differs from the preceding version.
	RrsetVersion *string `json:"rrsetVersion,omitempty" tf:"rrset_version,omitempty"`
}

type ItemsParameters struct {

	// The target fully-qualified domain name (FQDN) within the target zone.
	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// (Updatable) The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see Supported DNS Resource Record Types
	// +kubebuilder:validation:Required
	Rdata *string `json:"rdata" tf:"rdata,omitempty"`

	// The canonical name for the record's type, such as A or CNAME. For more information, see Resource Record (RR) TYPEs.
	// +kubebuilder:validation:Required
	Rtype *string `json:"rtype" tf:"rtype,omitempty"`

	// (Updatable) The Time To Live for the record, in seconds.
	// +kubebuilder:validation:Required
	TTL *float64 `json:"ttl" tf:"ttl,omitempty"`
}

type RrsetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Updatable)
	// NOTE Omitting items at time of create, will delete any existing records in the RRSet
	// +kubebuilder:validation:Optional
	Items []ItemsObservation `json:"items,omitempty" tf:"items,omitempty"`
}

type RrsetParameters struct {

	// (Updatable) The OCID of the compartment the resource belongs to.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// The target fully-qualified domain name (FQDN) within the target zone.
	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// (Updatable)
	// NOTE Omitting items at time of create, will delete any existing records in the RRSet
	// +kubebuilder:validation:Optional
	Items []ItemsParameters `json:"items,omitempty" tf:"items,omitempty"`

	// The canonical name for the record's type, such as A or CNAME. For more information, see Resource Record (RR) TYPEs.
	// +kubebuilder:validation:Required
	Rtype *string `json:"rtype" tf:"rtype,omitempty"`

	// Specifies to operate only on resources that have a matching DNS scope.
	// This value will be null for zones in the global DNS and PRIVATE when creating private Rrsets.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The OCID of the view the resource is associated with.
	// +kubebuilder:validation:Optional
	ViewID *string `json:"viewId,omitempty" tf:"view_id,omitempty"`

	// The name or OCID of the target zone.
	// +kubebuilder:validation:Required
	ZoneNameOrID *string `json:"zoneNameOrId" tf:"zone_name_or_id,omitempty"`
}

// RrsetSpec defines the desired state of Rrset
type RrsetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RrsetParameters `json:"forProvider"`
}

// RrsetStatus defines the observed state of Rrset.
type RrsetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RrsetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Rrset is the Schema for the Rrsets API. Provides the Rrset resource in Oracle Cloud Infrastructure DNS service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type Rrset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RrsetSpec   `json:"spec"`
	Status            RrsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RrsetList contains a list of Rrsets
type RrsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rrset `json:"items"`
}

// Repository type metadata.
var (
	Rrset_Kind             = "Rrset"
	Rrset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rrset_Kind}.String()
	Rrset_KindAPIVersion   = Rrset_Kind + "." + CRDGroupVersion.String()
	Rrset_GroupVersionKind = CRDGroupVersion.WithKind(Rrset_Kind)
)

func init() {
	SchemeBuilder.Register(&Rrset{}, &RrsetList{})
}
