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

type EgressSecurityRulesIcmpOptionsObservation struct {
}

type EgressSecurityRulesIcmpOptionsParameters struct {

	// (Updatable) The ICMP code .
	// +kubebuilder:validation:Optional
	Code *float64 `json:"code,omitempty" tf:"code,omitempty"`

	// (Updatable) The ICMP type.
	// +kubebuilder:validation:Required
	Type *float64 `json:"type" tf:"type,omitempty"`
}

type EgressSecurityRulesObservation struct {
}

type EgressSecurityRulesParameters struct {

	// (Updatable) An optional description of your choice for the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Conceptually, this is the range of IP addresses that a packet originating from the instance can go to.
	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// (Updatable) Type of destination for the rule. The default is CIDR_BLOCK.
	// +kubebuilder:validation:Optional
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// (Updatable) Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code as defined in:
	// +kubebuilder:validation:Optional
	IcmpOptions []EgressSecurityRulesIcmpOptionsParameters `json:"icmpOptions,omitempty" tf:"icmp_options,omitempty"`

	// (Updatable) The transport protocol. Specify either all or an IPv4 protocol number as defined in Protocol Numbers. Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58").
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// (Updatable) A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if egress traffic allows TCP destination port 80, there should be an ingress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic.
	// +kubebuilder:validation:Optional
	Stateless *bool `json:"stateless,omitempty" tf:"stateless,omitempty"`

	// (Updatable) Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed.
	// +kubebuilder:validation:Optional
	TCPOptions []EgressSecurityRulesTCPOptionsParameters `json:"tcpOptions,omitempty" tf:"tcp_options,omitempty"`

	// (Updatable) Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed.
	// +kubebuilder:validation:Optional
	UDPOptions []EgressSecurityRulesUDPOptionsParameters `json:"udpOptions,omitempty" tf:"udp_options,omitempty"`
}

type EgressSecurityRulesTCPOptionsObservation struct {
}

type EgressSecurityRulesTCPOptionsParameters struct {

	// (Updatable) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// (Updatable) The minimum port number. Must not be greater than the maximum port number.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`

	// (Updatable)
	// +kubebuilder:validation:Optional
	SourcePortRange []EgressSecurityRulesTCPOptionsSourcePortRangeParameters `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`
}

type EgressSecurityRulesTCPOptionsSourcePortRangeObservation struct {
}

type EgressSecurityRulesTCPOptionsSourcePortRangeParameters struct {

	// (Updatable) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// (Updatable) The minimum port number. Must not be greater than the maximum port number.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type EgressSecurityRulesUDPOptionsObservation struct {
}

type EgressSecurityRulesUDPOptionsParameters struct {

	// (Updatable) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// (Updatable) The minimum port number. Must not be greater than the maximum port number.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`

	// (Updatable)
	// +kubebuilder:validation:Optional
	SourcePortRange []EgressSecurityRulesUDPOptionsSourcePortRangeParameters `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`
}

type EgressSecurityRulesUDPOptionsSourcePortRangeObservation struct {
}

type EgressSecurityRulesUDPOptionsSourcePortRangeParameters struct {

	// (Updatable) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// (Updatable) The minimum port number. Must not be greater than the maximum port number.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type IngressSecurityRulesIcmpOptionsObservation struct {
}

type IngressSecurityRulesIcmpOptionsParameters struct {

	// (Updatable) The ICMP code .
	// +kubebuilder:validation:Optional
	Code *float64 `json:"code,omitempty" tf:"code,omitempty"`

	// (Updatable) The ICMP type.
	// +kubebuilder:validation:Required
	Type *float64 `json:"type" tf:"type,omitempty"`
}

type IngressSecurityRulesObservation struct {
}

type IngressSecurityRulesParameters struct {

	// (Updatable) An optional description of your choice for the rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Updatable) Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code as defined in:
	// +kubebuilder:validation:Optional
	IcmpOptions []IngressSecurityRulesIcmpOptionsParameters `json:"icmpOptions,omitempty" tf:"icmp_options,omitempty"`

	// (Updatable) The transport protocol. Specify either all or an IPv4 protocol number as defined in Protocol Numbers. Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58").
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// (Updatable) Conceptually, this is the range of IP addresses that a packet coming into the instance can come from.
	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`

	// (Updatable) Type of source for the rule. The default is CIDR_BLOCK.
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// (Updatable) A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if egress traffic allows TCP destination port 80, there should be an ingress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic.
	// +kubebuilder:validation:Optional
	Stateless *bool `json:"stateless,omitempty" tf:"stateless,omitempty"`

	// (Updatable) Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed.
	// +kubebuilder:validation:Optional
	TCPOptions []IngressSecurityRulesTCPOptionsParameters `json:"tcpOptions,omitempty" tf:"tcp_options,omitempty"`

	// (Updatable) Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed.
	// +kubebuilder:validation:Optional
	UDPOptions []IngressSecurityRulesUDPOptionsParameters `json:"udpOptions,omitempty" tf:"udp_options,omitempty"`
}

type IngressSecurityRulesTCPOptionsObservation struct {
}

type IngressSecurityRulesTCPOptionsParameters struct {

	// (Updatable) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// (Updatable) The minimum port number. Must not be greater than the maximum port number.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`

	// (Updatable)
	// +kubebuilder:validation:Optional
	SourcePortRange []IngressSecurityRulesTCPOptionsSourcePortRangeParameters `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`
}

type IngressSecurityRulesTCPOptionsSourcePortRangeObservation struct {
}

type IngressSecurityRulesTCPOptionsSourcePortRangeParameters struct {

	// (Updatable) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// (Updatable) The minimum port number. Must not be greater than the maximum port number.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type IngressSecurityRulesUDPOptionsObservation struct {
}

type IngressSecurityRulesUDPOptionsParameters struct {

	// (Updatable) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Optional
	Max *float64 `json:"max,omitempty" tf:"max,omitempty"`

	// (Updatable) The minimum port number. Must not be greater than the maximum port number.
	// +kubebuilder:validation:Optional
	Min *float64 `json:"min,omitempty" tf:"min,omitempty"`

	// (Updatable)
	// +kubebuilder:validation:Optional
	SourcePortRange []IngressSecurityRulesUDPOptionsSourcePortRangeParameters `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`
}

type IngressSecurityRulesUDPOptionsSourcePortRangeObservation struct {
}

type IngressSecurityRulesUDPOptionsSourcePortRangeParameters struct {

	// (Updatable) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// (Updatable) The minimum port number. Must not be greater than the maximum port number.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type SecurityListObservation struct {

	// The security list's Oracle Cloud ID (OCID).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The security list's current state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the security list was created, in the format defined by RFC3339.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type SecurityListParameters struct {

	// (Updatable) The OCID of the compartment to contain the security list.
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

	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// (Updatable) Rules for allowing egress IP packets.
	// +kubebuilder:validation:Optional
	EgressSecurityRules []EgressSecurityRulesParameters `json:"egressSecurityRules,omitempty" tf:"egress_security_rules,omitempty"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags.  Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// (Updatable) Rules for allowing ingress IP packets.
	// +kubebuilder:validation:Optional
	IngressSecurityRules []IngressSecurityRulesParameters `json:"ingressSecurityRules,omitempty" tf:"ingress_security_rules,omitempty"`

	// The OCID of the VCN the security list belongs to.
	// +crossplane:generate:reference:type=Vcn
	// +kubebuilder:validation:Optional
	VcnID *string `json:"vcnId,omitempty" tf:"vcn_id,omitempty"`

	// Reference to a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDRef *v1.Reference `json:"vcnIdRef,omitempty" tf:"-"`

	// Selector for a Vcn to populate vcnId.
	// +kubebuilder:validation:Optional
	VcnIDSelector *v1.Selector `json:"vcnIdSelector,omitempty" tf:"-"`
}

// SecurityListSpec defines the desired state of SecurityList
type SecurityListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityListParameters `json:"forProvider"`
}

// SecurityListStatus defines the observed state of SecurityList.
type SecurityListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityList is the Schema for the SecurityLists API. Provides the Security List resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type SecurityList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityListSpec   `json:"spec"`
	Status            SecurityListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityListList contains a list of SecurityLists
type SecurityListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityList `json:"items"`
}

// Repository type metadata.
var (
	SecurityList_Kind             = "SecurityList"
	SecurityList_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityList_Kind}.String()
	SecurityList_KindAPIVersion   = SecurityList_Kind + "." + CRDGroupVersion.String()
	SecurityList_GroupVersionKind = CRDGroupVersion.WithKind(SecurityList_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityList{}, &SecurityListList{})
}
