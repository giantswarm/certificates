package certs

import "fmt"

// These constants are used when filtering the secrets, to only retrieve the
// ones we are interested in.
const (
	// certificateLabel is the label used in the secret to identify a secret
	// containing the certificate.
	certificateLabel = "giantswarm.io/certificate"
	// clusterIDLabel is the label used in the secret to identify a secret
	// containing the certificate.
	clusterIDLabel = "giantswarm.io/cluster-id"

	// legacyCertificateLabel is the label used in the secret to identify a secret
	// containing the certificate.
	//
	// TODO use certificateLabel instead when all cert secrets have it.
	legacyCertificateLabel = "clusterComponent"
	// legacyClusterIDLabel is the label used in the secret to identify a secret
	// containing the certificate.
	//
	// TODO use clusterIDLabel instead when all cert secrets have it.
	legacyClusterIDLabel = "clusterID"

	SecretNamespace = "default"
)

// Cert is a certificate name.
type Cert string

// These constants used as Cert parsing a secret received from the API.
const (
	APICert                Cert = "api"
	CalicoCert             Cert = "calico"
	CalicoEtcdClientCert   Cert = "calico-etcd-client"
	ClusterOperatorAPICert Cert = "cluster-operator-api"
	EtcdCert               Cert = "etcd"
	FlanneldEtcdClientCert Cert = "flanneld-etcd-client"
	NodeOperatorCert       Cert = "node-operator"
	PrometheusCert         Cert = "prometheus"
	ServiceAccountCert     Cert = "service-account"
	WorkerCert             Cert = "worker"
)

// AllCerts lists all certificates that can be created by cert-operator.
var AllCerts = []Cert{
	APICert,
	CalicoCert,
	CalicoEtcdClientCert,
	ClusterOperatorAPICert,
	EtcdCert,
	FlanneldEtcdClientCert,
	NodeOperatorCert,
	PrometheusCert,
	ServiceAccountCert,
	WorkerCert,
}

// K8sSecretName returns Kubernetes Secret object name for the certificate name
// and the guest cluster ID.
func K8sSecretName(clusterID string, certificate Cert) string {
	return fmt.Sprintf("%s-%s", clusterID, certificate)
}

func K8sSecretLabels(clusterID string, certificate Cert) map[string]string {
	return map[string]string{
		certificateLabel:       string(certificate),
		clusterIDLabel:         clusterID,
		legacyCertificateLabel: string(certificate),
		legacyClusterIDLabel:   clusterID,
	}
}
