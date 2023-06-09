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

type ResolverEndpointObservation struct {

	// The OCID of the owning compartment. This will match the resolver that the resolver endpoint is under and will be updated if the resolver's compartment is changed.
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The canonical absolute URL of the resource.
	Self *string `json:"self,omitempty" tf:"self,omitempty"`

	// The current state of the resource.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`
}

type ResolverEndpointParameters struct {

	// (Updatable) The type of resolver endpoint. VNIC is currently the only supported type.
	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part of the subnet and will be assigned by the system if unspecified when isForwarding is true.
	// +kubebuilder:validation:Optional
	ForwardingAddress *string `json:"forwardingAddress,omitempty" tf:"forwarding_address,omitempty"`

	// A Boolean flag indicating whether or not the resolver endpoint is for forwarding.
	// +kubebuilder:validation:Required
	IsForwarding *bool `json:"isForwarding" tf:"is_forwarding,omitempty"`

	// A Boolean flag indicating whether or not the resolver endpoint is for listening.
	// +kubebuilder:validation:Required
	IsListening *bool `json:"isListening" tf:"is_listening,omitempty"`

	// An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the subnet and will be assigned by the system if unspecified when isListening is true.
	// +kubebuilder:validation:Optional
	ListeningAddress *string `json:"listeningAddress,omitempty" tf:"listening_address,omitempty"`

	// The name of the resolver endpoint. Must be unique, case-insensitive, within the resolver.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// An array of network security group OCIDs for the resolver endpoint. These must be part of the VCN that the resolver endpoint is a part of.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.NetworkSecurityGroup
	// +kubebuilder:validation:Optional
	NsgIds []*string `json:"nsgIds,omitempty" tf:"nsg_ids,omitempty"`

	// References to NetworkSecurityGroup in core to populate nsgIds.
	// +kubebuilder:validation:Optional
	NsgIdsRefs []v1.Reference `json:"nsgIdsRefs,omitempty" tf:"-"`

	// Selector for a list of NetworkSecurityGroup in core to populate nsgIds.
	// +kubebuilder:validation:Optional
	NsgIdsSelector *v1.Selector `json:"nsgIdsSelector,omitempty" tf:"-"`

	// The OCID of the target resolver.
	// +crossplane:generate:reference:type=Resolver
	// +kubebuilder:validation:Optional
	ResolverID *string `json:"resolverId,omitempty" tf:"resolver_id,omitempty"`

	// Reference to a Resolver to populate resolverId.
	// +kubebuilder:validation:Optional
	ResolverIDRef *v1.Reference `json:"resolverIdRef,omitempty" tf:"-"`

	// Selector for a Resolver to populate resolverId.
	// +kubebuilder:validation:Optional
	ResolverIDSelector *v1.Selector `json:"resolverIdSelector,omitempty" tf:"-"`

	// Value must be PRIVATE when creating private name resolver endpoints.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The OCID of a subnet. Must be part of the VCN that the resolver is attached to.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/core/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in core to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in core to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// ResolverEndpointSpec defines the desired state of ResolverEndpoint
type ResolverEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResolverEndpointParameters `json:"forProvider"`
}

// ResolverEndpointStatus defines the observed state of ResolverEndpoint.
type ResolverEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResolverEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverEndpoint is the Schema for the ResolverEndpoints API. Provides the Resolver Endpoint resource in Oracle Cloud Infrastructure DNS service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type ResolverEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResolverEndpointSpec   `json:"spec"`
	Status            ResolverEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverEndpointList contains a list of ResolverEndpoints
type ResolverEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResolverEndpoint `json:"items"`
}

// Repository type metadata.
var (
	ResolverEndpoint_Kind             = "ResolverEndpoint"
	ResolverEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResolverEndpoint_Kind}.String()
	ResolverEndpoint_KindAPIVersion   = ResolverEndpoint_Kind + "." + CRDGroupVersion.String()
	ResolverEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(ResolverEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&ResolverEndpoint{}, &ResolverEndpointList{})
}
