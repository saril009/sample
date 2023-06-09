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

type BackendsObservation struct {

	// The IP address of the backend server. Example: 10.0.0.3
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Whether the network load balancer should treat this server as a backup unit. If true, then the network load balancer forwards no ingress traffic to this backend server unless all other backend servers not marked as "isBackup" fail the health check policy.  Example: false
	IsBackup *bool `json:"isBackup,omitempty" tf:"is_backup,omitempty"`

	// Whether the network load balancer should drain this server. Servers marked "isDrain" receive no incoming traffic.  Example: false
	IsDrain *bool `json:"isDrain,omitempty" tf:"is_drain,omitempty"`

	// Whether the network load balancer should treat this server as offline. Offline servers receive no incoming traffic.  Example: false
	IsOffline *bool `json:"isOffline,omitempty" tf:"is_offline,omitempty"`

	// A user-friendly name for the backend set that must be unique and cannot be changed.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Updatable) The backend server port against which to run the health check. If the port is not specified, then the network load balancer uses the port information from the Backend object. The port must be specified if the backend port is 0.  Example: 8080
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The IP OCID/Instance OCID associated with the backend server. Example: ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`

	// The network load balancing policy weight assigned to the server. Backend servers with a higher weight receive a larger proportion of incoming traffic. For example, a server weighted '3' receives three times the number of new connections as a server weighted '1'. For more information about load balancing policies, see How Network Load Balancing Policies Work.  Example: 3
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type BackendsParameters struct {
}

type HealthCheckerObservation struct {
}

type HealthCheckerParameters struct {

	// (Updatable) The interval between health checks, in milliseconds. The default value is 10000 (10 seconds).  Example: 10000
	// +kubebuilder:validation:Optional
	IntervalInMillis *float64 `json:"intervalInMillis,omitempty" tf:"interval_in_millis,omitempty"`

	// (Updatable) The backend server port against which to run the health check. If the port is not specified, then the network load balancer uses the port information from the Backend object. The port must be specified if the backend port is 0.  Example: 8080
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// (Updatable) The protocol the health check must use; either HTTP or HTTPS, or UDP or TCP.  Example: HTTP
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// (Updatable) Base64 encoded pattern to be sent as UDP or TCP health check probe.
	// +kubebuilder:validation:Optional
	RequestData *string `json:"requestData,omitempty" tf:"request_data,omitempty"`

	// (Updatable) A regular expression for parsing the response body from the backend server.  Example: ^((?!false).|\s)*$
	// +kubebuilder:validation:Optional
	ResponseBodyRegex *string `json:"responseBodyRegex,omitempty" tf:"response_body_regex,omitempty"`

	// (Updatable) Base64 encoded pattern to be validated as UDP or TCP health check probe response.
	// +kubebuilder:validation:Optional
	ResponseData *string `json:"responseData,omitempty" tf:"response_data,omitempty"`

	// (Updatable) The number of retries to attempt before a backend server is considered "unhealthy". This number also applies when recovering a server to the "healthy" state. The default value is 3.  Example: 3
	// +kubebuilder:validation:Optional
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`

	// (Updatable) The status code a healthy backend server should return. If you configure the health check policy to use the HTTP protocol, then you can use common HTTP status codes such as "200".  Example: 200
	// +kubebuilder:validation:Optional
	ReturnCode *float64 `json:"returnCode,omitempty" tf:"return_code,omitempty"`

	// (Updatable) The maximum time, in milliseconds, to wait for a reply to a health check. A health check is successful only if a reply returns within this timeout period. The default value is 3000 (3 seconds).  Example: 3000
	// +kubebuilder:validation:Optional
	TimeoutInMillis *float64 `json:"timeoutInMillis,omitempty" tf:"timeout_in_millis,omitempty"`

	// (Updatable) The path against which to run the health check.  Example: /healthcheck
	// +kubebuilder:validation:Optional
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`
}

type LoadBalancerBackendSetObservation struct {

	// Array of backends.
	Backends []BackendsObservation `json:"backends,omitempty" tf:"backends,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LoadBalancerBackendSetParameters struct {

	// (Updatable) The health check policy configuration. For more information, see Editing Health Check Policies.
	// +kubebuilder:validation:Required
	HealthChecker []HealthCheckerParameters `json:"healthChecker" tf:"health_checker,omitempty"`

	// (Updatable) IP version associated with the backend set.
	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// (Updatable) If this parameter is enabled, then the network load balancer preserves the source IP of the packet when it is forwarded to backends. Backends see the original source IP. If the isPreserveSourceDestination parameter is enabled for the network load balancer resource, then this parameter cannot be disabled. The value is true by default.
	// +kubebuilder:validation:Optional
	IsPreserveSource *bool `json:"isPreserveSource,omitempty" tf:"is_preserve_source,omitempty"`

	// A user-friendly name for the backend set that must be unique and cannot be changed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The OCID of the network load balancer to update.
	// +crossplane:generate:reference:type=LoadBalancerNetworkLoadBalancer
	// +kubebuilder:validation:Optional
	NetworkLoadBalancerID *string `json:"networkLoadBalancerId,omitempty" tf:"network_load_balancer_id,omitempty"`

	// Reference to a LoadBalancerNetworkLoadBalancer to populate networkLoadBalancerId.
	// +kubebuilder:validation:Optional
	NetworkLoadBalancerIDRef *v1.Reference `json:"networkLoadBalancerIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancerNetworkLoadBalancer to populate networkLoadBalancerId.
	// +kubebuilder:validation:Optional
	NetworkLoadBalancerIDSelector *v1.Selector `json:"networkLoadBalancerIdSelector,omitempty" tf:"-"`

	// (Updatable) The network load balancer policy for the backend set.  Example: `FIVE_TUPLE“
	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`
}

// LoadBalancerBackendSetSpec defines the desired state of LoadBalancerBackendSet
type LoadBalancerBackendSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerBackendSetParameters `json:"forProvider"`
}

// LoadBalancerBackendSetStatus defines the observed state of LoadBalancerBackendSet.
type LoadBalancerBackendSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerBackendSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerBackendSet is the Schema for the LoadBalancerBackendSets API. Provides the Backend Set resource in Oracle Cloud Infrastructure Network Load Balancer service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type LoadBalancerBackendSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerBackendSetSpec   `json:"spec"`
	Status            LoadBalancerBackendSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerBackendSetList contains a list of LoadBalancerBackendSets
type LoadBalancerBackendSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerBackendSet `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancerBackendSet_Kind             = "LoadBalancerBackendSet"
	LoadBalancerBackendSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancerBackendSet_Kind}.String()
	LoadBalancerBackendSet_KindAPIVersion   = LoadBalancerBackendSet_Kind + "." + CRDGroupVersion.String()
	LoadBalancerBackendSet_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancerBackendSet_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancerBackendSet{}, &LoadBalancerBackendSetList{})
}
