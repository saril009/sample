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

type ConnectHarnessObservation struct {

	// The OCID of the connect harness.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Any additional details about the current state of the connect harness.
	LifecycleStateDetails *string `json:"lifecycleStateDetails,omitempty" tf:"lifecycle_state_details,omitempty"`

	// The current state of the connect harness.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the connect harness was created, expressed in in RFC 3339 timestamp format.  Example: 2018-04-20T00:00:07.405Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type ConnectHarnessParameters struct {

	// (Updatable) The OCID of the compartment that contains the connect harness.
	// +crossplane:generate:reference:type=github.com/oracle/provider-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// Reference to a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// Selector for a Compartment in identity to populate compartmentId.
	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.  Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair that is applied with no predefined name, type, or namespace. Exists for cross-compatibility only. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// The name of the connect harness. Avoid entering confidential information.  Example: JDBCConnector
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// ConnectHarnessSpec defines the desired state of ConnectHarness
type ConnectHarnessSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectHarnessParameters `json:"forProvider"`
}

// ConnectHarnessStatus defines the observed state of ConnectHarness.
type ConnectHarnessStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectHarnessObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectHarness is the Schema for the ConnectHarnesss API. Provides the Connect Harness resource in Oracle Cloud Infrastructure Streaming service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type ConnectHarness struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectHarnessSpec   `json:"spec"`
	Status            ConnectHarnessStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectHarnessList contains a list of ConnectHarnesss
type ConnectHarnessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectHarness `json:"items"`
}

// Repository type metadata.
var (
	ConnectHarness_Kind             = "ConnectHarness"
	ConnectHarness_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectHarness_Kind}.String()
	ConnectHarness_KindAPIVersion   = ConnectHarness_Kind + "." + CRDGroupVersion.String()
	ConnectHarness_GroupVersionKind = CRDGroupVersion.WithKind(ConnectHarness_Kind)
)

func init() {
	SchemeBuilder.Register(&ConnectHarness{}, &ConnectHarnessList{})
}
