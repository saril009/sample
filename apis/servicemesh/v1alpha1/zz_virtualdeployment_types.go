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

type ServiceDiscoveryObservation struct {
}

type ServiceDiscoveryParameters struct {

	// +kubebuilder:validation:Required
	Hostname *string `json:"hostname" tf:"hostname,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type VirtualDeploymentAccessLoggingObservation struct {
}

type VirtualDeploymentAccessLoggingParameters struct {

	// +kubebuilder:validation:Optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled,omitempty"`
}

type VirtualDeploymentListenersObservation struct {
}

type VirtualDeploymentListenersParameters struct {

	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type VirtualDeploymentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	SystemTags map[string]*string `json:"systemTags,omitempty" tf:"system_tags,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated,omitempty"`
}

type VirtualDeploymentParameters struct {

	// +kubebuilder:validation:Optional
	AccessLogging []VirtualDeploymentAccessLoggingParameters `json:"accessLogging,omitempty" tf:"access_logging,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Required
	Listeners []VirtualDeploymentListenersParameters `json:"listeners" tf:"listeners,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ServiceDiscovery []ServiceDiscoveryParameters `json:"serviceDiscovery" tf:"service_discovery,omitempty"`

	// +crossplane:generate:reference:type=VirtualService
	// +kubebuilder:validation:Optional
	VirtualServiceID *string `json:"virtualServiceId,omitempty" tf:"virtual_service_id,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualServiceIDRef *v1.Reference `json:"virtualServiceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VirtualServiceIDSelector *v1.Selector `json:"virtualServiceIdSelector,omitempty" tf:"-"`
}

// VirtualDeploymentSpec defines the desired state of VirtualDeployment
type VirtualDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualDeploymentParameters `json:"forProvider"`
}

// VirtualDeploymentStatus defines the observed state of VirtualDeployment.
type VirtualDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualDeployment is the Schema for the VirtualDeployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type VirtualDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualDeploymentSpec   `json:"spec"`
	Status            VirtualDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualDeploymentList contains a list of VirtualDeployments
type VirtualDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualDeployment `json:"items"`
}

// Repository type metadata.
var (
	VirtualDeployment_Kind             = "VirtualDeployment"
	VirtualDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualDeployment_Kind}.String()
	VirtualDeployment_KindAPIVersion   = VirtualDeployment_Kind + "." + CRDGroupVersion.String()
	VirtualDeployment_GroupVersionKind = CRDGroupVersion.WithKind(VirtualDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualDeployment{}, &VirtualDeploymentList{})
}