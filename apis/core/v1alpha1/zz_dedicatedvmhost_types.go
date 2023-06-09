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

type DedicatedVMHostObservation struct {

	// The OCID of the dedicated VM host.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The current available memory of the dedicated VM host, in GBs.
	RemainingMemoryInGbs *float64 `json:"remainingMemoryInGbs,omitempty" tf:"remaining_memory_in_gbs,omitempty"`

	// The current available OCPUs of the dedicated VM host.
	RemainingOcpus *float64 `json:"remainingOcpus,omitempty" tf:"remaining_ocpus,omitempty"`

	// The current state of the dedicated VM host.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the dedicated VM host was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	// The current total memory of the dedicated VM host, in GBs.
	TotalMemoryInGbs *float64 `json:"totalMemoryInGbs,omitempty" tf:"total_memory_in_gbs,omitempty"`

	// The current total OCPUs of the dedicated VM host.
	TotalOcpus *float64 `json:"totalOcpus,omitempty" tf:"total_ocpus,omitempty"`
}

type DedicatedVMHostParameters struct {

	// The availability domain of the dedicated virtual machine host.  Example: Uocm:PHX-AD-1
	// +kubebuilder:validation:Required
	AvailabilityDomain *string `json:"availabilityDomain" tf:"availability_domain,omitempty"`

	// (Updatable) The OCID of the compartment.
	// +kubebuilder:validation:Required
	CompartmentID *string `json:"compartmentId" tf:"compartment_id,omitempty"`

	// The dedicated virtual machine host shape. The shape determines the number of CPUs and other resources available for VM instances launched on the dedicated virtual machine host.
	// +kubebuilder:validation:Required
	DedicatedVMHostShape *string `json:"dedicatedVmHostShape" tf:"dedicated_vm_host_shape,omitempty"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The fault domain for the dedicated virtual machine host's assigned instances. For more information, see Fault Domains. If you do not specify the fault domain, the system selects one for you. To change the fault domain for a dedicated virtual machine host, delete it and create a new dedicated virtual machine host in the preferred fault domain.
	// +kubebuilder:validation:Optional
	FaultDomain *string `json:"faultDomain,omitempty" tf:"fault_domain,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`
}

// DedicatedVMHostSpec defines the desired state of DedicatedVMHost
type DedicatedVMHostSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedVMHostParameters `json:"forProvider"`
}

// DedicatedVMHostStatus defines the observed state of DedicatedVMHost.
type DedicatedVMHostStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedVMHostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedVMHost is the Schema for the DedicatedVMHosts API. Provides the Dedicated Vm Host resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type DedicatedVMHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedVMHostSpec   `json:"spec"`
	Status            DedicatedVMHostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedVMHostList contains a list of DedicatedVMHosts
type DedicatedVMHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedVMHost `json:"items"`
}

// Repository type metadata.
var (
	DedicatedVMHost_Kind             = "DedicatedVMHost"
	DedicatedVMHost_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedVMHost_Kind}.String()
	DedicatedVMHost_KindAPIVersion   = DedicatedVMHost_Kind + "." + CRDGroupVersion.String()
	DedicatedVMHost_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedVMHost_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedVMHost{}, &DedicatedVMHostList{})
}
