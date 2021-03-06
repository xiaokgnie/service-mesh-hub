package cert_manager

import (
	"fmt"

	"github.com/rotisserie/eris"
	zephyr_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	zephyr_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	zephyr_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	zephyr_security_types "github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1/types"
	"istio.io/istio/pkg/spiffe"
)

const (
	DefaultIstioOrg              = "Istio"
	DefaultCitadelServiceAccount = "istio-citadel" // The default SPIFFE URL value for trust domain
	DefaultTrustDomain           = "cluster.local"
)

var (
	IncorrectMeshTypeError = func(mesh *zephyr_discovery.Mesh) error {
		return eris.Errorf("invalid mesh type (%T) passed into istio certificate config producer",
			mesh.Spec.GetMeshType())
	}
)

func NewIstioCertConfigProducer() IstioCertConfigProducer {
	return &istioCertConfigProducer{}
}

type istioCertConfigProducer struct{}

type IstioCertConfigProducer CertConfigProducer

func BuildSpiffeURI(trustDomain, namespace, sa string) string {
	return fmt.Sprintf("%s%s/ns/%s/sa/%s", spiffe.URIPrefix, trustDomain, namespace, sa)
}

func (i *istioCertConfigProducer) ConfigureCertificateInfo(
	vm *zephyr_networking.VirtualMesh,
	mesh *zephyr_discovery.Mesh,
) (*zephyr_security_types.VirtualMeshCertificateSigningRequestSpec_CertConfig, error) {
	istioMesh := mesh.Spec.GetIstio()
	if istioMesh == nil {
		return nil, IncorrectMeshTypeError(mesh)
	}

	trustDomain := DefaultTrustDomain
	citadelServiceAccount := DefaultCitadelServiceAccount
	citadelNamespace := istioMesh.GetInstallation().GetInstallationNamespace()

	if istioMesh.GetCitadelInfo().GetTrustDomain() != "" {
		trustDomain = istioMesh.GetCitadelInfo().GetTrustDomain()
	}
	if istioMesh.GetCitadelInfo().GetCitadelNamespace() != "" {
		citadelNamespace = istioMesh.GetCitadelInfo().GetCitadelNamespace()
	}
	if istioMesh.GetCitadelInfo().GetCitadelServiceAccount() != "" {
		citadelServiceAccount = istioMesh.GetCitadelInfo().GetCitadelServiceAccount()
	}
	return &zephyr_security_types.VirtualMeshCertificateSigningRequestSpec_CertConfig{
		// TODO: Make citadel namespace discoverable
		Hosts:    []string{BuildSpiffeURI(trustDomain, citadelNamespace, citadelServiceAccount)},
		Org:      DefaultIstioOrg,
		MeshType: zephyr_core_types.MeshType_ISTIO,
	}, nil
}
