package v1alpha1

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// clienset for the networking.zephyr.solo.io/v1alpha1 APIs
type Clientset interface {
	// clienset for the networking.zephyr.solo.io/v1alpha1/v1alpha1 APIs
	TrafficPolicies() TrafficPolicyClient
	// clienset for the networking.zephyr.solo.io/v1alpha1/v1alpha1 APIs
	AccessControlPolicies() AccessControlPolicyClient
	// clienset for the networking.zephyr.solo.io/v1alpha1/v1alpha1 APIs
	VirtualMeshes() VirtualMeshClient
}

type clientSet struct {
	client client.Client
}

func NewClientsetFromConfig(cfg *rest.Config) (*clientSet, error) {
	scheme := scheme.Scheme
	if err := AddToScheme(scheme); err != nil {
		return nil, err
	}
	client, err := client.New(cfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return NewClientset(client), nil
}

func NewClientset(client client.Client) *clientSet {
	return &clientSet{client: client}
}

// clienset for the networking.zephyr.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) TrafficPolicies() TrafficPolicyClient {
	return NewTrafficPolicyClient(c.client)
}

// clienset for the networking.zephyr.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) AccessControlPolicies() AccessControlPolicyClient {
	return NewAccessControlPolicyClient(c.client)
}

// clienset for the networking.zephyr.solo.io/v1alpha1/v1alpha1 APIs
func (c *clientSet) VirtualMeshes() VirtualMeshClient {
	return NewVirtualMeshClient(c.client)
}

// Reader knows how to read and list TrafficPolicys.
type TrafficPolicyReader interface {
	// Get retrieves a TrafficPolicy for the given object key
	GetTrafficPolicy(ctx context.Context, key client.ObjectKey) (*TrafficPolicy, error)

	// List retrieves list of TrafficPolicys for a given namespace and list options.
	ListTrafficPolicy(ctx context.Context, opts ...client.ListOption) (*TrafficPolicyList, error)
}

// Writer knows how to create, delete, and update TrafficPolicys.
type TrafficPolicyWriter interface {
	// Create saves the TrafficPolicy object.
	CreateTrafficPolicy(ctx context.Context, obj *TrafficPolicy, opts ...client.CreateOption) error

	// Delete deletes the TrafficPolicy object.
	DeleteTrafficPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given TrafficPolicy object.
	UpdateTrafficPolicy(ctx context.Context, obj *TrafficPolicy, opts ...client.UpdateOption) error

	// If the TrafficPolicy object exists, update its spec. Otherwise, create the TrafficPolicy object.
	UpsertTrafficPolicySpec(ctx context.Context, obj *TrafficPolicy, opts ...client.UpdateOption) error

	// Patch patches the given TrafficPolicy object.
	PatchTrafficPolicy(ctx context.Context, obj *TrafficPolicy, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all TrafficPolicy objects matching the given options.
	DeleteAllOfTrafficPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a TrafficPolicy object.
type TrafficPolicyStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given TrafficPolicy object.
	UpdateTrafficPolicyStatus(ctx context.Context, obj *TrafficPolicy, opts ...client.UpdateOption) error

	// Patch patches the given TrafficPolicy object's subresource.
	PatchTrafficPolicyStatus(ctx context.Context, obj *TrafficPolicy, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on TrafficPolicys.
type TrafficPolicyClient interface {
	TrafficPolicyReader
	TrafficPolicyWriter
	TrafficPolicyStatusWriter
}

type trafficPolicyClient struct {
	client client.Client
}

func NewTrafficPolicyClient(client client.Client) *trafficPolicyClient {
	return &trafficPolicyClient{client: client}
}

func (c *trafficPolicyClient) GetTrafficPolicy(ctx context.Context, key client.ObjectKey) (*TrafficPolicy, error) {
	obj := &TrafficPolicy{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *trafficPolicyClient) ListTrafficPolicy(ctx context.Context, opts ...client.ListOption) (*TrafficPolicyList, error) {
	list := &TrafficPolicyList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *trafficPolicyClient) CreateTrafficPolicy(ctx context.Context, obj *TrafficPolicy, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *trafficPolicyClient) DeleteTrafficPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &TrafficPolicy{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *trafficPolicyClient) UpdateTrafficPolicy(ctx context.Context, obj *TrafficPolicy, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *trafficPolicyClient) UpsertTrafficPolicySpec(ctx context.Context, obj *TrafficPolicy, opts ...client.UpdateOption) error {
	existing, err := c.GetTrafficPolicy(ctx, client.ObjectKey{Name: obj.GetName(), Namespace: obj.GetNamespace()})
	if err != nil {
		if errors.IsNotFound(err) {
			return c.CreateTrafficPolicy(ctx, obj)
		}
		return err
	}
	existing.Spec = obj.Spec
	return c.client.Update(ctx, existing, opts...)
}

func (c *trafficPolicyClient) PatchTrafficPolicy(ctx context.Context, obj *TrafficPolicy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *trafficPolicyClient) DeleteAllOfTrafficPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &TrafficPolicy{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *trafficPolicyClient) UpdateTrafficPolicyStatus(ctx context.Context, obj *TrafficPolicy, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *trafficPolicyClient) PatchTrafficPolicyStatus(ctx context.Context, obj *TrafficPolicy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list AccessControlPolicys.
type AccessControlPolicyReader interface {
	// Get retrieves a AccessControlPolicy for the given object key
	GetAccessControlPolicy(ctx context.Context, key client.ObjectKey) (*AccessControlPolicy, error)

	// List retrieves list of AccessControlPolicys for a given namespace and list options.
	ListAccessControlPolicy(ctx context.Context, opts ...client.ListOption) (*AccessControlPolicyList, error)
}

// Writer knows how to create, delete, and update AccessControlPolicys.
type AccessControlPolicyWriter interface {
	// Create saves the AccessControlPolicy object.
	CreateAccessControlPolicy(ctx context.Context, obj *AccessControlPolicy, opts ...client.CreateOption) error

	// Delete deletes the AccessControlPolicy object.
	DeleteAccessControlPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given AccessControlPolicy object.
	UpdateAccessControlPolicy(ctx context.Context, obj *AccessControlPolicy, opts ...client.UpdateOption) error

	// If the AccessControlPolicy object exists, update its spec. Otherwise, create the AccessControlPolicy object.
	UpsertAccessControlPolicySpec(ctx context.Context, obj *AccessControlPolicy, opts ...client.UpdateOption) error

	// Patch patches the given AccessControlPolicy object.
	PatchAccessControlPolicy(ctx context.Context, obj *AccessControlPolicy, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all AccessControlPolicy objects matching the given options.
	DeleteAllOfAccessControlPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a AccessControlPolicy object.
type AccessControlPolicyStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given AccessControlPolicy object.
	UpdateAccessControlPolicyStatus(ctx context.Context, obj *AccessControlPolicy, opts ...client.UpdateOption) error

	// Patch patches the given AccessControlPolicy object's subresource.
	PatchAccessControlPolicyStatus(ctx context.Context, obj *AccessControlPolicy, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on AccessControlPolicys.
type AccessControlPolicyClient interface {
	AccessControlPolicyReader
	AccessControlPolicyWriter
	AccessControlPolicyStatusWriter
}

type accessControlPolicyClient struct {
	client client.Client
}

func NewAccessControlPolicyClient(client client.Client) *accessControlPolicyClient {
	return &accessControlPolicyClient{client: client}
}

func (c *accessControlPolicyClient) GetAccessControlPolicy(ctx context.Context, key client.ObjectKey) (*AccessControlPolicy, error) {
	obj := &AccessControlPolicy{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *accessControlPolicyClient) ListAccessControlPolicy(ctx context.Context, opts ...client.ListOption) (*AccessControlPolicyList, error) {
	list := &AccessControlPolicyList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *accessControlPolicyClient) CreateAccessControlPolicy(ctx context.Context, obj *AccessControlPolicy, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *accessControlPolicyClient) DeleteAccessControlPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &AccessControlPolicy{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *accessControlPolicyClient) UpdateAccessControlPolicy(ctx context.Context, obj *AccessControlPolicy, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *accessControlPolicyClient) UpsertAccessControlPolicySpec(ctx context.Context, obj *AccessControlPolicy, opts ...client.UpdateOption) error {
	existing, err := c.GetAccessControlPolicy(ctx, client.ObjectKey{Name: obj.GetName(), Namespace: obj.GetNamespace()})
	if err != nil {
		if errors.IsNotFound(err) {
			return c.CreateAccessControlPolicy(ctx, obj)
		}
		return err
	}
	existing.Spec = obj.Spec
	return c.client.Update(ctx, existing, opts...)
}

func (c *accessControlPolicyClient) PatchAccessControlPolicy(ctx context.Context, obj *AccessControlPolicy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *accessControlPolicyClient) DeleteAllOfAccessControlPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &AccessControlPolicy{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *accessControlPolicyClient) UpdateAccessControlPolicyStatus(ctx context.Context, obj *AccessControlPolicy, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *accessControlPolicyClient) PatchAccessControlPolicyStatus(ctx context.Context, obj *AccessControlPolicy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}

// Reader knows how to read and list VirtualMeshs.
type VirtualMeshReader interface {
	// Get retrieves a VirtualMesh for the given object key
	GetVirtualMesh(ctx context.Context, key client.ObjectKey) (*VirtualMesh, error)

	// List retrieves list of VirtualMeshs for a given namespace and list options.
	ListVirtualMesh(ctx context.Context, opts ...client.ListOption) (*VirtualMeshList, error)
}

// Writer knows how to create, delete, and update VirtualMeshs.
type VirtualMeshWriter interface {
	// Create saves the VirtualMesh object.
	CreateVirtualMesh(ctx context.Context, obj *VirtualMesh, opts ...client.CreateOption) error

	// Delete deletes the VirtualMesh object.
	DeleteVirtualMesh(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given VirtualMesh object.
	UpdateVirtualMesh(ctx context.Context, obj *VirtualMesh, opts ...client.UpdateOption) error

	// If the VirtualMesh object exists, update its spec. Otherwise, create the VirtualMesh object.
	UpsertVirtualMeshSpec(ctx context.Context, obj *VirtualMesh, opts ...client.UpdateOption) error

	// Patch patches the given VirtualMesh object.
	PatchVirtualMesh(ctx context.Context, obj *VirtualMesh, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all VirtualMesh objects matching the given options.
	DeleteAllOfVirtualMesh(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a VirtualMesh object.
type VirtualMeshStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given VirtualMesh object.
	UpdateVirtualMeshStatus(ctx context.Context, obj *VirtualMesh, opts ...client.UpdateOption) error

	// Patch patches the given VirtualMesh object's subresource.
	PatchVirtualMeshStatus(ctx context.Context, obj *VirtualMesh, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on VirtualMeshs.
type VirtualMeshClient interface {
	VirtualMeshReader
	VirtualMeshWriter
	VirtualMeshStatusWriter
}

type virtualMeshClient struct {
	client client.Client
}

func NewVirtualMeshClient(client client.Client) *virtualMeshClient {
	return &virtualMeshClient{client: client}
}

func (c *virtualMeshClient) GetVirtualMesh(ctx context.Context, key client.ObjectKey) (*VirtualMesh, error) {
	obj := &VirtualMesh{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *virtualMeshClient) ListVirtualMesh(ctx context.Context, opts ...client.ListOption) (*VirtualMeshList, error) {
	list := &VirtualMeshList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *virtualMeshClient) CreateVirtualMesh(ctx context.Context, obj *VirtualMesh, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *virtualMeshClient) DeleteVirtualMesh(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &VirtualMesh{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *virtualMeshClient) UpdateVirtualMesh(ctx context.Context, obj *VirtualMesh, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *virtualMeshClient) UpsertVirtualMeshSpec(ctx context.Context, obj *VirtualMesh, opts ...client.UpdateOption) error {
	existing, err := c.GetVirtualMesh(ctx, client.ObjectKey{Name: obj.GetName(), Namespace: obj.GetNamespace()})
	if err != nil {
		if errors.IsNotFound(err) {
			return c.CreateVirtualMesh(ctx, obj)
		}
		return err
	}
	existing.Spec = obj.Spec
	return c.client.Update(ctx, existing, opts...)
}

func (c *virtualMeshClient) PatchVirtualMesh(ctx context.Context, obj *VirtualMesh, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *virtualMeshClient) DeleteAllOfVirtualMesh(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &VirtualMesh{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *virtualMeshClient) UpdateVirtualMeshStatus(ctx context.Context, obj *VirtualMesh, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *virtualMeshClient) PatchVirtualMeshStatus(ctx context.Context, obj *VirtualMesh, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}
