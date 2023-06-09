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

type AccessPolicyObservation struct {

	// Unique identifier that is immutable on creation.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details,omitempty"`

	// The current state of the Resource.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. Example: {"orcl-cloud.free-tier-retained": "true"}
	SystemTags map[string]*string `json:"systemTags,omitempty" tf:"system_tags,omitempty"`

	// The time when this resource was created in an RFC3339 formatted datetime string.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The time when this resource was updated in an RFC3339 formatted datetime string.
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`
}

type AccessPolicyParameters struct {

	// (Updatable) The OCID of the compartment.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: {"foo-namespace.bar-key": "value"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Description of the resource. It can be changed after creation. Avoid entering confidential information.  Example: This is my new resource
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: {"bar-key": "value"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The OCID of the service mesh in which this access policy is created.
	// +crossplane:generate:reference:type=Mesh
	// +kubebuilder:validation:Optional
	MeshID *string `json:"meshId,omitempty" tf:"mesh_id,omitempty"`

	// Reference to a Mesh to populate meshId.
	// +kubebuilder:validation:Optional
	MeshIDRef *v1.Reference `json:"meshIdRef,omitempty" tf:"-"`

	// Selector for a Mesh to populate meshId.
	// +kubebuilder:validation:Optional
	MeshIDSelector *v1.Selector `json:"meshIdSelector,omitempty" tf:"-"`

	// A user-friendly name. The name has to be unique within the same service mesh and cannot be changed after creation. Avoid entering confidential information.  Example: My unique resource name
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// (Updatable) List of applicable rules
	// +kubebuilder:validation:Required
	Rules []RulesParameters `json:"rules" tf:"rules,omitempty"`
}

type DestinationObservation struct {
}

type DestinationParameters struct {

	// (Applicable when type=EXTERNAL_SERVICE) (Updatable) The hostnames of the external service. Only applicable for HTTP and HTTPS protocols. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", ".example.com", ".com", "". Hostname "" can be used to allow all hosts.
	// +kubebuilder:validation:Optional
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`

	// (Applicable when type=EXTERNAL_SERVICE) (Updatable) The ipAddresses of the external service in CIDR notation. Only applicable for TCP protocol. All requests matching the given CIDR notation will pass through. In case a wildcard CIDR "0.0.0.0/0" is provided, the same port cannot be used for a virtual service communication.
	// +kubebuilder:validation:Optional
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// (Updatable) The OCID of the ingress gateway resource.
	// +crossplane:generate:reference:type=IngressGateway
	// +kubebuilder:validation:Optional
	IngressGatewayID *string `json:"ingressGatewayId,omitempty" tf:"ingress_gateway_id,omitempty"`

	// Reference to a IngressGateway to populate ingressGatewayId.
	// +kubebuilder:validation:Optional
	IngressGatewayIDRef *v1.Reference `json:"ingressGatewayIdRef,omitempty" tf:"-"`

	// Selector for a IngressGateway to populate ingressGatewayId.
	// +kubebuilder:validation:Optional
	IngressGatewayIDSelector *v1.Selector `json:"ingressGatewayIdSelector,omitempty" tf:"-"`

	// (Applicable when type=EXTERNAL_SERVICE) (Updatable) Ports exposed by an external service. If left empty all ports will be allowed.
	// +kubebuilder:validation:Optional
	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`

	// (Applicable when type=EXTERNAL_SERVICE) (Updatable) Protocol of the external service
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Updatable) Traffic type of the target.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// (Updatable) The OCID of the virtual service resource.
	// +kubebuilder:validation:Optional
	VirtualServiceID *string `json:"virtualServiceId,omitempty" tf:"virtual_service_id,omitempty"`
}

type RulesObservation struct {
}

type RulesParameters struct {

	// (Updatable) Action for the traffic between the source and the destination.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// (Updatable) Target of the access policy. This can either be the source or the destination of the traffic.
	// +kubebuilder:validation:Required
	Destination []DestinationParameters `json:"destination" tf:"destination,omitempty"`

	// (Updatable) Target of the access policy. This can either be the source or the destination of the traffic.
	// +kubebuilder:validation:Required
	Source []SourceParameters `json:"source" tf:"source,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// (Applicable when type=EXTERNAL_SERVICE) (Updatable) The hostnames of the external service. Only applicable for HTTP and HTTPS protocols. Wildcard hostnames are supported in the prefix form. Examples of valid hostnames are "www.example.com", ".example.com", ".com", "". Hostname "" can be used to allow all hosts.
	// +kubebuilder:validation:Optional
	Hostnames []*string `json:"hostnames,omitempty" tf:"hostnames,omitempty"`

	// (Applicable when type=EXTERNAL_SERVICE) (Updatable) The ipAddresses of the external service in CIDR notation. Only applicable for TCP protocol. All requests matching the given CIDR notation will pass through. In case a wildcard CIDR "0.0.0.0/0" is provided, the same port cannot be used for a virtual service communication.
	// +kubebuilder:validation:Optional
	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	// (Updatable) The OCID of the ingress gateway resource.
	// +kubebuilder:validation:Optional
	IngressGatewayID *string `json:"ingressGatewayId,omitempty" tf:"ingress_gateway_id,omitempty"`

	// (Applicable when type=EXTERNAL_SERVICE) (Updatable) Ports exposed by an external service. If left empty all ports will be allowed.
	// +kubebuilder:validation:Optional
	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`

	// (Applicable when type=EXTERNAL_SERVICE) (Updatable) Protocol of the external service
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Updatable) Traffic type of the target.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// (Updatable) The OCID of the virtual service resource.
	// +kubebuilder:validation:Optional
	VirtualServiceID *string `json:"virtualServiceId,omitempty" tf:"virtual_service_id,omitempty"`
}

// AccessPolicySpec defines the desired state of AccessPolicy
type AccessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessPolicyParameters `json:"forProvider"`
}

// AccessPolicyStatus defines the observed state of AccessPolicy.
type AccessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPolicy is the Schema for the AccessPolicys API. Provides the Access Policy resource in Oracle Cloud Infrastructure Service Mesh service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type AccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessPolicySpec   `json:"spec"`
	Status            AccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPolicyList contains a list of AccessPolicys
type AccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	AccessPolicy_Kind             = "AccessPolicy"
	AccessPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessPolicy_Kind}.String()
	AccessPolicy_KindAPIVersion   = AccessPolicy_Kind + "." + CRDGroupVersion.String()
	AccessPolicy_GroupVersionKind = CRDGroupVersion.WithKind(AccessPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessPolicy{}, &AccessPolicyList{})
}
