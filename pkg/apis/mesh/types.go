// Package mesh provides the types and functions for interacting with
// the NGINX Service Mesh API and configuration.
package mesh

// injector annotations and labels.
const (
	// InjectedAnnotation tells us if a pod has been injected.
	InjectedAnnotation = "injector.nsm.nginx.com/status"
	// AutoInjectLabel tells whether a pod should be injected with the sidecar.
	AutoInjectLabel = "injector.nsm.nginx.com/auto-inject"
)

// Injected is used as the value in the InjectedAnnotation.
const Injected = "injected"

// IgnoredNamespaces is a map of the namespaces that the service mesh will ignore.
var IgnoredNamespaces = map[string]bool{
	"kube-system": true,
}

// DeployLabel is the label key for deployment type of the resource.
const DeployLabel = "nsm.nginx.com/"

// NATS channel names.
const (
	// NatsAgentConfigChannel sends the mesh config from mesh-api to agent.
	NatsAgentConfigChannel = "nginx.nsm.agent.config"
	// NatsAgentSubChannel sends a subscription and version notice from agent to mesh-api.
	NatsAgentSubChannel = "nginx.nsm.agent.subscription"
	// NatsAPIPingChannel sends a ping from mesh-api to agent on restart.
	NatsAPIPingChannel = "nginx.nsm.api.ping"
)

// k8s static resource names.
const (
	// MeshConfigMap is the name of the config map that holds the mesh config.
	MeshConfigMap = "mesh-config"
	// MeshConfigFileName is the name of the file where the mesh config is stored.
	MeshConfigFileName = "mesh-config.json"
	// NatsServer is the name of the nats-server service.
	NatsServer = "nats-server"
	// MeshAPI is the name of the mesh api.
	MeshAPI = "nginx-mesh-api"
	// MeshCertReloader is the name of the mesh cert reloader image.
	MeshCertReloader = "nginx-mesh-cert-reloader"
	// MeshSidecar is the name of the mesh sidecar.
	MeshSidecar = "nginx-mesh-sidecar"
	// MeshSidecarInit is the name of the mesh init container.
	MeshSidecarInit = "nginx-mesh-init"
	// MetricsService is the name of the traffic metrics service.
	MetricsService = "nginx-mesh-metrics-svc"
	// MetricsServiceAccount is the name of the service account of traffic metrics.
	MetricsServiceAccount = "nginx-mesh-metrics"
	// MetricsDeployment is the name of the traffic metrics deployment.
	MetricsDeployment = MetricsServiceAccount
	// HTTPRouteGroupKind is the kind for HTTPRouteGroups.
	HTTPRouteGroupKind = "HTTPRouteGroup"
	// TCPRouteKind is the kind of TcpRoutes.
	TCPRouteKind = "TCPRoute"
)

// field names for inject payload.
const (
	// IgnoreIncomingPortsField is the field that lists ports to ignore for incoming traffic.
	IgnoreIncomingPortsField = "ignoreIncomingPorts"
	// IgnoreOutgoingPortsField is the field that lists ports to ignore for outgoing traffic.
	IgnoreOutgoingPortsField = "ignoreOutgoingPorts"
	// FileField is the field that contains the binary data of the file to inject.
	FileField = "file"
)

// MetricsConfig holds the data that may be dynamically updated at runtime for the nginx-mesh-metrics component.
type MetricsConfig struct {
	PromAddr *string `json:"PrometheusAddress,omitempty"`
}

// MtlsModes are the supported mtls modes.
var MtlsModes = map[string]struct{}{
	string(Off):        {},
	string(Permissive): {},
	string(Strict):     {},
}

// LoadBalancingMethods are the available NGINX load balancing methods.
var LoadBalancingMethods = map[string]struct{}{
	string(MeshConfigLoadBalancingMethodRoundRobin):                 {},
	string(MeshConfigLoadBalancingMethodLeastConn):                  {},
	string(MeshConfigLoadBalancingMethodLeastTime):                  {},
	string(MeshConfigLoadBalancingMethodLeastTimeLastByte):          {},
	string(MeshConfigLoadBalancingMethodLeastTimeLastByteInflight):  {},
	string(MeshConfigLoadBalancingMethodRandom):                     {},
	string(MeshConfigLoadBalancingMethodRandomTwo):                  {},
	string(MeshConfigLoadBalancingMethodRandomTwoLeastConn):         {},
	string(MeshConfigLoadBalancingMethodRandomTwoLeastTime):         {},
	string(MeshConfigLoadBalancingMethodRandomTwoLeastTimeLastByte): {},
}

// NGINXLogFormats are the supported NGINX log formats.
var NGINXLogFormats = map[string]struct{}{
	string(MeshConfigNginxLogFormatDefault): {},
	string(MeshConfigNginxLogFormatJson):    {},
}

// TracingBackends are the supported tracing backends.
var TracingBackends = map[string]struct{}{
	string(Zipkin):  {},
	string(Jaeger):  {},
	string(Datadog): {},
}

// Environments are the supported kubernetes environments.
var Environments = map[string]struct{}{
	string(Kubernetes): {},
	string(Openshift):  {},
}
