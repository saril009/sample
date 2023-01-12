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

type VnicAttachmentCreateVnicDetailsObservation struct {
}

type VnicAttachmentCreateVnicDetailsParameters struct {

	// +kubebuilder:validation:Optional
	AssignPrivateDNSRecord *bool `json:"assignPrivateDnsRecord,omitempty" tf:"assign_private_dns_record,omitempty"`

	// +kubebuilder:validation:Optional
	AssignPublicIP *string `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Optional
	HostnameLabel *string `json:"hostnameLabel,omitempty" tf:"hostname_label,omitempty"`

	// +kubebuilder:validation:Optional
	NsgIds []*string `json:"nsgIds,omitempty" tf:"nsg_ids,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// +kubebuilder:validation:Optional
	SkipSourceDestCheck *bool `json:"skipSourceDestCheck,omitempty" tf:"skip_source_dest_check,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`
}

type VnicAttachmentObservation struct {
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain,omitempty"`

	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	VlanID *string `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	VlanTag *float64 `json:"vlanTag,omitempty" tf:"vlan_tag,omitempty"`

	VnicID *string `json:"vnicId,omitempty" tf:"vnic_id,omitempty"`
}

type VnicAttachmentParameters struct {

	// +kubebuilder:validation:Required
	CreateVnicDetails []VnicAttachmentCreateVnicDetailsParameters `json:"createVnicDetails" tf:"create_vnic_details,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NicIndex *float64 `json:"nicIndex,omitempty" tf:"nic_index,omitempty"`
}

// VnicAttachmentSpec defines the desired state of VnicAttachment
type VnicAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VnicAttachmentParameters `json:"forProvider"`
}

// VnicAttachmentStatus defines the observed state of VnicAttachment.
type VnicAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VnicAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VnicAttachment is the Schema for the VnicAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type VnicAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VnicAttachmentSpec   `json:"spec"`
	Status            VnicAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VnicAttachmentList contains a list of VnicAttachments
type VnicAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VnicAttachment `json:"items"`
}

// Repository type metadata.
var (
	VnicAttachment_Kind             = "VnicAttachment"
	VnicAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VnicAttachment_Kind}.String()
	VnicAttachment_KindAPIVersion   = VnicAttachment_Kind + "." + CRDGroupVersion.String()
	VnicAttachment_GroupVersionKind = CRDGroupVersion.WithKind(VnicAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&VnicAttachment{}, &VnicAttachmentList{})
}