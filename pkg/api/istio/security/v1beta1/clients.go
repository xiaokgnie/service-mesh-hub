package v1beta1

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	. "istio.io/client-go/pkg/apis/security/v1beta1"
)

// clienset for the security/v1beta1 APIs
type Clientset interface {
	// clienset for the security/v1beta1/v1beta1 APIs
	AuthorizationPolicies() AuthorizationPolicyClient
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

// clienset for the security/v1beta1/v1beta1 APIs
func (c *clientSet) AuthorizationPolicies() AuthorizationPolicyClient {
	return NewAuthorizationPolicyClient(c.client)
}

// Reader knows how to read and list AuthorizationPolicys.
type AuthorizationPolicyReader interface {
	// Get retrieves a AuthorizationPolicy for the given object key
	GetAuthorizationPolicy(ctx context.Context, key client.ObjectKey) (*AuthorizationPolicy, error)

	// List retrieves list of AuthorizationPolicys for a given namespace and list options.
	ListAuthorizationPolicy(ctx context.Context, opts ...client.ListOption) (*AuthorizationPolicyList, error)
}

// Writer knows how to create, delete, and update AuthorizationPolicys.
type AuthorizationPolicyWriter interface {
	// Create saves the AuthorizationPolicy object.
	CreateAuthorizationPolicy(ctx context.Context, obj *AuthorizationPolicy, opts ...client.CreateOption) error

	// Delete deletes the AuthorizationPolicy object.
	DeleteAuthorizationPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error

	// Update updates the given AuthorizationPolicy object.
	UpdateAuthorizationPolicy(ctx context.Context, obj *AuthorizationPolicy, opts ...client.UpdateOption) error

	// If the AuthorizationPolicy object exists, update its spec. Otherwise, create the AuthorizationPolicy object.
	UpsertAuthorizationPolicySpec(ctx context.Context, obj *AuthorizationPolicy, opts ...client.UpdateOption) error

	// Patch patches the given AuthorizationPolicy object.
	PatchAuthorizationPolicy(ctx context.Context, obj *AuthorizationPolicy, patch client.Patch, opts ...client.PatchOption) error

	// DeleteAllOf deletes all AuthorizationPolicy objects matching the given options.
	DeleteAllOfAuthorizationPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error
}

// StatusWriter knows how to update status subresource of a AuthorizationPolicy object.
type AuthorizationPolicyStatusWriter interface {
	// Update updates the fields corresponding to the status subresource for the
	// given AuthorizationPolicy object.
	UpdateAuthorizationPolicyStatus(ctx context.Context, obj *AuthorizationPolicy, opts ...client.UpdateOption) error

	// Patch patches the given AuthorizationPolicy object's subresource.
	PatchAuthorizationPolicyStatus(ctx context.Context, obj *AuthorizationPolicy, patch client.Patch, opts ...client.PatchOption) error
}

// Client knows how to perform CRUD operations on AuthorizationPolicys.
type AuthorizationPolicyClient interface {
	AuthorizationPolicyReader
	AuthorizationPolicyWriter
	AuthorizationPolicyStatusWriter
}

type authorizationPolicyClient struct {
	client client.Client
}

func NewAuthorizationPolicyClient(client client.Client) *authorizationPolicyClient {
	return &authorizationPolicyClient{client: client}
}

func (c *authorizationPolicyClient) GetAuthorizationPolicy(ctx context.Context, key client.ObjectKey) (*AuthorizationPolicy, error) {
	obj := &AuthorizationPolicy{}
	if err := c.client.Get(ctx, key, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (c *authorizationPolicyClient) ListAuthorizationPolicy(ctx context.Context, opts ...client.ListOption) (*AuthorizationPolicyList, error) {
	list := &AuthorizationPolicyList{}
	if err := c.client.List(ctx, list, opts...); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *authorizationPolicyClient) CreateAuthorizationPolicy(ctx context.Context, obj *AuthorizationPolicy, opts ...client.CreateOption) error {
	return c.client.Create(ctx, obj, opts...)
}

func (c *authorizationPolicyClient) DeleteAuthorizationPolicy(ctx context.Context, key client.ObjectKey, opts ...client.DeleteOption) error {
	obj := &AuthorizationPolicy{}
	obj.SetName(key.Name)
	obj.SetNamespace(key.Namespace)
	return c.client.Delete(ctx, obj, opts...)
}

func (c *authorizationPolicyClient) UpdateAuthorizationPolicy(ctx context.Context, obj *AuthorizationPolicy, opts ...client.UpdateOption) error {
	return c.client.Update(ctx, obj, opts...)
}

func (c *authorizationPolicyClient) UpsertAuthorizationPolicySpec(ctx context.Context, obj *AuthorizationPolicy, opts ...client.UpdateOption) error {
	existing, err := c.GetAuthorizationPolicy(ctx, client.ObjectKey{Name: obj.GetName(), Namespace: obj.GetNamespace()})
	if err != nil {
		if errors.IsNotFound(err) {
			return c.CreateAuthorizationPolicy(ctx, obj)
		}
		return err
	}
	existing.Spec = obj.Spec
	return c.client.Update(ctx, existing, opts...)
}

func (c *authorizationPolicyClient) PatchAuthorizationPolicy(ctx context.Context, obj *AuthorizationPolicy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Patch(ctx, obj, patch, opts...)
}

func (c *authorizationPolicyClient) DeleteAllOfAuthorizationPolicy(ctx context.Context, opts ...client.DeleteAllOfOption) error {
	obj := &AuthorizationPolicy{}
	return c.client.DeleteAllOf(ctx, obj, opts...)
}

func (c *authorizationPolicyClient) UpdateAuthorizationPolicyStatus(ctx context.Context, obj *AuthorizationPolicy, opts ...client.UpdateOption) error {
	return c.client.Status().Update(ctx, obj, opts...)
}

func (c *authorizationPolicyClient) PatchAuthorizationPolicyStatus(ctx context.Context, obj *AuthorizationPolicy, patch client.Patch, opts ...client.PatchOption) error {
	return c.client.Status().Patch(ctx, obj, patch, opts...)
}
