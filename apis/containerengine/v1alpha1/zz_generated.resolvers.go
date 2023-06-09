/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha11 "github.com/oracle/provider-oci/apis/core/v1alpha1"
	v1alpha1 "github.com/oracle/provider-oci/apis/identity/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EndpointConfig); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.EndpointConfig[i3].NsgIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.EndpointConfig[i3].NsgIdsRefs,
			Selector:      mg.Spec.ForProvider.EndpointConfig[i3].NsgIdsSelector,
			To: reference.To{
				List:    &v1alpha11.NetworkSecurityGroupList{},
				Managed: &v1alpha11.NetworkSecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EndpointConfig[i3].NsgIds")
		}
		mg.Spec.ForProvider.EndpointConfig[i3].NsgIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.EndpointConfig[i3].NsgIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.EndpointConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EndpointConfig[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EndpointConfig[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.EndpointConfig[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EndpointConfig[i3].SubnetID")
		}
		mg.Spec.ForProvider.EndpointConfig[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EndpointConfig[i3].SubnetIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Options); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Options[i3].ServiceLBSubnetIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.Options[i3].ServiceLBSubnetIdsRef,
			Selector:      mg.Spec.ForProvider.Options[i3].ServiceLBSubnetIDSelector,
			To: reference.To{
				List:    &v1alpha11.SubnetList{},
				Managed: &v1alpha11.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Options[i3].ServiceLBSubnetIds")
		}
		mg.Spec.ForProvider.Options[i3].ServiceLBSubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Options[i3].ServiceLBSubnetIdsRef = mrsp.ResolvedReferences

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VcnID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VcnIDRef,
		Selector:     mg.Spec.ForProvider.VcnIDSelector,
		To: reference.To{
			List:    &v1alpha11.VcnList{},
			Managed: &v1alpha11.Vcn{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VcnID")
	}
	mg.Spec.ForProvider.VcnID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VcnIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NodePool.
func (mg *NodePool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CompartmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CompartmentIDRef,
		Selector:     mg.Spec.ForProvider.CompartmentIDSelector,
		To: reference.To{
			List:    &v1alpha1.CompartmentList{},
			Managed: &v1alpha1.Compartment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CompartmentID")
	}
	mg.Spec.ForProvider.CompartmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CompartmentIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NodeConfigDetails); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.NodeConfigDetails[i3].NsgIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.NodeConfigDetails[i3].NsgIdsRefs,
			Selector:      mg.Spec.ForProvider.NodeConfigDetails[i3].NsgIdsSelector,
			To: reference.To{
				List:    &v1alpha11.NetworkSecurityGroupList{},
				Managed: &v1alpha11.NetworkSecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NodeConfigDetails[i3].NsgIds")
		}
		mg.Spec.ForProvider.NodeConfigDetails[i3].NsgIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.NodeConfigDetails[i3].NsgIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NodeConfigDetails); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NodeConfigDetails[i3].PlacementConfigs); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeConfigDetails[i3].PlacementConfigs[i4].SubnetID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.NodeConfigDetails[i3].PlacementConfigs[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.NodeConfigDetails[i3].PlacementConfigs[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha11.SubnetList{},
					Managed: &v1alpha11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NodeConfigDetails[i3].PlacementConfigs[i4].SubnetID")
			}
			mg.Spec.ForProvider.NodeConfigDetails[i3].PlacementConfigs[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NodeConfigDetails[i3].PlacementConfigs[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}

	return nil
}
