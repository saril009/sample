/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this FirewallNetworkFirewall.
func (mg *FirewallNetworkFirewall) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this FirewallNetworkFirewall.
func (mg *FirewallNetworkFirewall) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this FirewallNetworkFirewall.
func (mg *FirewallNetworkFirewall) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this FirewallNetworkFirewall.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *FirewallNetworkFirewall) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this FirewallNetworkFirewall.
func (mg *FirewallNetworkFirewall) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this FirewallNetworkFirewall.
func (mg *FirewallNetworkFirewall) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this FirewallNetworkFirewall.
func (mg *FirewallNetworkFirewall) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this FirewallNetworkFirewall.
func (mg *FirewallNetworkFirewall) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this FirewallNetworkFirewall.
func (mg *FirewallNetworkFirewall) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this FirewallNetworkFirewall.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *FirewallNetworkFirewall) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this FirewallNetworkFirewall.
func (mg *FirewallNetworkFirewall) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this FirewallNetworkFirewall.
func (mg *FirewallNetworkFirewall) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this FirewallNetworkFirewallPolicy.
func (mg *FirewallNetworkFirewallPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this FirewallNetworkFirewallPolicy.
func (mg *FirewallNetworkFirewallPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this FirewallNetworkFirewallPolicy.
func (mg *FirewallNetworkFirewallPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this FirewallNetworkFirewallPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *FirewallNetworkFirewallPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this FirewallNetworkFirewallPolicy.
func (mg *FirewallNetworkFirewallPolicy) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this FirewallNetworkFirewallPolicy.
func (mg *FirewallNetworkFirewallPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this FirewallNetworkFirewallPolicy.
func (mg *FirewallNetworkFirewallPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this FirewallNetworkFirewallPolicy.
func (mg *FirewallNetworkFirewallPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this FirewallNetworkFirewallPolicy.
func (mg *FirewallNetworkFirewallPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this FirewallNetworkFirewallPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *FirewallNetworkFirewallPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this FirewallNetworkFirewallPolicy.
func (mg *FirewallNetworkFirewallPolicy) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this FirewallNetworkFirewallPolicy.
func (mg *FirewallNetworkFirewallPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LoadBalancerBackend.
func (mg *LoadBalancerBackend) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LoadBalancerBackend.
func (mg *LoadBalancerBackend) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LoadBalancerBackend.
func (mg *LoadBalancerBackend) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LoadBalancerBackend.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LoadBalancerBackend) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LoadBalancerBackend.
func (mg *LoadBalancerBackend) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LoadBalancerBackend.
func (mg *LoadBalancerBackend) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LoadBalancerBackend.
func (mg *LoadBalancerBackend) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LoadBalancerBackend.
func (mg *LoadBalancerBackend) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LoadBalancerBackend.
func (mg *LoadBalancerBackend) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LoadBalancerBackend.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LoadBalancerBackend) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LoadBalancerBackend.
func (mg *LoadBalancerBackend) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LoadBalancerBackend.
func (mg *LoadBalancerBackend) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LoadBalancerBackendSet.
func (mg *LoadBalancerBackendSet) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LoadBalancerBackendSet.
func (mg *LoadBalancerBackendSet) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LoadBalancerBackendSet.
func (mg *LoadBalancerBackendSet) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LoadBalancerBackendSet.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LoadBalancerBackendSet) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LoadBalancerBackendSet.
func (mg *LoadBalancerBackendSet) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LoadBalancerBackendSet.
func (mg *LoadBalancerBackendSet) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LoadBalancerBackendSet.
func (mg *LoadBalancerBackendSet) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LoadBalancerBackendSet.
func (mg *LoadBalancerBackendSet) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LoadBalancerBackendSet.
func (mg *LoadBalancerBackendSet) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LoadBalancerBackendSet.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LoadBalancerBackendSet) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LoadBalancerBackendSet.
func (mg *LoadBalancerBackendSet) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LoadBalancerBackendSet.
func (mg *LoadBalancerBackendSet) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LoadBalancerListener.
func (mg *LoadBalancerListener) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LoadBalancerListener.
func (mg *LoadBalancerListener) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LoadBalancerListener.
func (mg *LoadBalancerListener) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LoadBalancerListener.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LoadBalancerListener) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LoadBalancerListener.
func (mg *LoadBalancerListener) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LoadBalancerListener.
func (mg *LoadBalancerListener) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LoadBalancerListener.
func (mg *LoadBalancerListener) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LoadBalancerListener.
func (mg *LoadBalancerListener) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LoadBalancerListener.
func (mg *LoadBalancerListener) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LoadBalancerListener.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LoadBalancerListener) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LoadBalancerListener.
func (mg *LoadBalancerListener) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LoadBalancerListener.
func (mg *LoadBalancerListener) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LoadBalancerNetworkLoadBalancer.
func (mg *LoadBalancerNetworkLoadBalancer) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LoadBalancerNetworkLoadBalancer.
func (mg *LoadBalancerNetworkLoadBalancer) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LoadBalancerNetworkLoadBalancer.
func (mg *LoadBalancerNetworkLoadBalancer) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LoadBalancerNetworkLoadBalancer.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LoadBalancerNetworkLoadBalancer) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LoadBalancerNetworkLoadBalancer.
func (mg *LoadBalancerNetworkLoadBalancer) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LoadBalancerNetworkLoadBalancer.
func (mg *LoadBalancerNetworkLoadBalancer) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LoadBalancerNetworkLoadBalancer.
func (mg *LoadBalancerNetworkLoadBalancer) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LoadBalancerNetworkLoadBalancer.
func (mg *LoadBalancerNetworkLoadBalancer) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LoadBalancerNetworkLoadBalancer.
func (mg *LoadBalancerNetworkLoadBalancer) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LoadBalancerNetworkLoadBalancer.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LoadBalancerNetworkLoadBalancer) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LoadBalancerNetworkLoadBalancer.
func (mg *LoadBalancerNetworkLoadBalancer) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LoadBalancerNetworkLoadBalancer.
func (mg *LoadBalancerNetworkLoadBalancer) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LoadBalancerNetworkLoadBalancersBackendSetsUnified.
func (mg *LoadBalancerNetworkLoadBalancersBackendSetsUnified) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
