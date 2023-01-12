/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LoadBalancerListenerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LoadBalancerListenerParameters struct {

	// +kubebuilder:validation:Required
	DefaultBackendSetName *string `json:"defaultBackendSetName" tf:"default_backend_set_name,omitempty"`

	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=LoadBalancerNetworkLoadBalancer
	// +kubebuilder:validation:Optional
	NetworkLoadBalancerID *string `json:"networkLoadBalancerId,omitempty" tf:"network_load_balancer_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkLoadBalancerIDRef *v1.Reference `json:"networkLoadBalancerIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkLoadBalancerIDSelector *v1.Selector `json:"networkLoadBalancerIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

// LoadBalancerListenerSpec defines the desired state of LoadBalancerListener
type LoadBalancerListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerListenerParameters `json:"forProvider"`
}

// LoadBalancerListenerStatus defines the observed state of LoadBalancerListener.
type LoadBalancerListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerListener is the Schema for the LoadBalancerListeners API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type LoadBalancerListener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerListenerSpec   `json:"spec"`
	Status            LoadBalancerListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerListenerList contains a list of LoadBalancerListeners
type LoadBalancerListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerListener `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerListener_Kind             = "LoadBalancerListener"
	LoadBalancerListener_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerListener_Kind}.String()
	LoadBalancerListener_KindAPIVersion   = LoadBalancerListener_Kind + "." + CRDGroupVersion.String()
	LoadBalancerListener_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerListener_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerListener{}, &LoadBalancerListenerList{})
}