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

type BalancerBackendObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A read-only field showing the IP address and port that uniquely identify this backend server in the backend set.  Example: 10.0.0.3:8080
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type BalancerBackendParameters struct {

	// The name of the backend set to add the backend server to.  Example: example_backend_set
	// +crossplane:generate:reference:type=BalancerBackendSet
	// +kubebuilder:validation:Optional
	BackendsetName *string `json:"backendsetName,omitempty" tf:"backendset_name,omitempty"`

	// Reference to a BalancerBackendSet to populate backendsetName.
	// +kubebuilder:validation:Optional
	BackendsetNameRef *v1.Reference `json:"backendsetNameRef,omitempty" tf:"-"`

	// Selector for a BalancerBackendSet to populate backendsetName.
	// +kubebuilder:validation:Optional
	BackendsetNameSelector *v1.Selector `json:"backendsetNameSelector,omitempty" tf:"-"`

	// (Updatable) Whether the load balancer should treat this server as a backup unit. If true, the load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "backup" fail the health check policy.
	// +kubebuilder:validation:Optional
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// (Updatable) Whether the load balancer should drain this server. Servers marked "drain" receive no new incoming traffic.  Example: false
	// +kubebuilder:validation:Optional
	Drain *bool `json:"drain,omitempty" tf:"drain,omitempty"`

	// The IP address of the backend server.  Example: 10.0.0.3
	// +kubebuilder:validation:Required
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// The OCID of the load balancer associated with the backend set and servers.
	// +crossplane:generate:reference:type=BalancerLoadBalancer
	// +kubebuilder:validation:Optional
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// Reference to a BalancerLoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// Selector for a BalancerLoadBalancer to populate loadBalancerId.
	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// (Updatable) Whether the load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: false
	// +kubebuilder:validation:Optional
	Offline *bool `json:"offline,omitempty" tf:"offline,omitempty"`

	// The communication port for the backend server.  Example: 8080
	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// (Updatable) The load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives 3 times the number of new connections as a server weighted '1'. For more information on load balancing policies, see How Load Balancing Policies Work.  Example: 3
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// BalancerBackendSpec defines the desired state of BalancerBackend
type BalancerBackendSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BalancerBackendParameters `json:"forProvider"`
}

// BalancerBackendStatus defines the observed state of BalancerBackend.
type BalancerBackendStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BalancerBackendObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerBackend is the Schema for the BalancerBackends API. Provides the Backend resource in Oracle Cloud Infrastructure Load Balancer service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type BalancerBackend struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BalancerBackendSpec   `json:"spec"`
	Status            BalancerBackendStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerBackendList contains a list of BalancerBackends
type BalancerBackendList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BalancerBackend `json:"items"`
}

// Repository type metadata.
var (
	BalancerBackend_Kind             = "BalancerBackend"
	BalancerBackend_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BalancerBackend_Kind}.String()
	BalancerBackend_KindAPIVersion   = BalancerBackend_Kind + "." + CRDGroupVersion.String()
	BalancerBackend_GroupVersionKind = CRDGroupVersion.WithKind(BalancerBackend_Kind)
)

func init() {
	SchemeBuilder.Register(&BalancerBackend{}, &BalancerBackendList{})
}
