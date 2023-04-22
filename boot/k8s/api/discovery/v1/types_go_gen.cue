// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go k8s.io/api/discovery/v1

package v1

import (
	metav1 "github.com/defn/m/boot/k8s/apimachinery/pkg/apis/meta/v1"
	"github.com/defn/m/boot/k8s/api/core/v1"
)

// EndpointSlice represents a subset of the endpoints that implement a service.
// For a given service there may be multiple EndpointSlice objects, selected by
// labels, which must be joined to produce the full set of endpoints.
#EndpointSlice: {
	metav1.#TypeMeta

	// Standard object's metadata.
	// +optional
	metadata?: metav1.#ObjectMeta @go(ObjectMeta) @protobuf(1,bytes,opt)

	// addressType specifies the type of address carried by this EndpointSlice.
	// All addresses in this slice must be the same type. This field is
	// immutable after creation. The following address types are currently
	// supported:
	// * IPv4: Represents an IPv4 Address.
	// * IPv6: Represents an IPv6 Address.
	// * FQDN: Represents a Fully Qualified Domain Name.
	addressType: #AddressType @go(AddressType) @protobuf(4,bytes,rep)

	// endpoints is a list of unique endpoints in this slice. Each slice may
	// include a maximum of 1000 endpoints.
	// +listType=atomic
	endpoints: [...#Endpoint] @go(Endpoints,[]Endpoint) @protobuf(2,bytes,rep)

	// ports specifies the list of network ports exposed by each endpoint in
	// this slice. Each port must have a unique name. When ports is empty, it
	// indicates that there are no defined ports. When a port is defined with a
	// nil port value, it indicates "all ports". Each slice may include a
	// maximum of 100 ports.
	// +optional
	// +listType=atomic
	ports: [...#EndpointPort] @go(Ports,[]EndpointPort) @protobuf(3,bytes,rep)
}

// AddressType represents the type of address referred to by an endpoint.
// +enum
#AddressType: string // #enumAddressType

#enumAddressType:
	#AddressTypeIPv4 |
	#AddressTypeIPv6 |
	#AddressTypeFQDN

// AddressTypeIPv4 represents an IPv4 Address.
#AddressTypeIPv4: #AddressType & "IPv4"

// AddressTypeIPv6 represents an IPv6 Address.
#AddressTypeIPv6: #AddressType & "IPv6"

// AddressTypeFQDN represents a FQDN.
#AddressTypeFQDN: #AddressType & "FQDN"

// Endpoint represents a single logical "backend" implementing a service.
#Endpoint: {
	// addresses of this endpoint. The contents of this field are interpreted
	// according to the corresponding EndpointSlice addressType field. Consumers
	// must handle different types of addresses in the context of their own
	// capabilities. This must contain at least one address but no more than
	// 100. These are all assumed to be fungible and clients may choose to only
	// use the first element. Refer to: https://issue.k8s.io/106267
	// +listType=set
	addresses: [...string] @go(Addresses,[]string) @protobuf(1,bytes,rep)

	// conditions contains information about the current status of the endpoint.
	conditions?: #EndpointConditions @go(Conditions) @protobuf(2,bytes,opt)

	// hostname of this endpoint. This field may be used by consumers of
	// endpoints to distinguish endpoints from each other (e.g. in DNS names).
	// Multiple endpoints which use the same hostname should be considered
	// fungible (e.g. multiple A values in DNS). Must be lowercase and pass DNS
	// Label (RFC 1123) validation.
	// +optional
	hostname?: null | string @go(Hostname,*string) @protobuf(3,bytes,opt)

	// targetRef is a reference to a Kubernetes object that represents this
	// endpoint.
	// +optional
	targetRef?: null | v1.#ObjectReference @go(TargetRef,*v1.ObjectReference) @protobuf(4,bytes,opt)

	// deprecatedTopology contains topology information part of the v1beta1
	// API. This field is deprecated, and will be removed when the v1beta1
	// API is removed (no sooner than kubernetes v1.24).  While this field can
	// hold values, it is not writable through the v1 API, and any attempts to
	// write to it will be silently ignored. Topology information can be found
	// in the zone and nodeName fields instead.
	// +optional
	deprecatedTopology?: {[string]: string} @go(DeprecatedTopology,map[string]string) @protobuf(5,bytes,opt)

	// nodeName represents the name of the Node hosting this endpoint. This can
	// be used to determine endpoints local to a Node.
	// +optional
	nodeName?: null | string @go(NodeName,*string) @protobuf(6,bytes,opt)

	// zone is the name of the Zone this endpoint exists in.
	// +optional
	zone?: null | string @go(Zone,*string) @protobuf(7,bytes,opt)

	// hints contains information associated with how an endpoint should be
	// consumed.
	// +optional
	hints?: null | #EndpointHints @go(Hints,*EndpointHints) @protobuf(8,bytes,opt)
}

// EndpointConditions represents the current condition of an endpoint.
#EndpointConditions: {
	// ready indicates that this endpoint is prepared to receive traffic,
	// according to whatever system is managing the endpoint. A nil value
	// indicates an unknown state. In most cases consumers should interpret this
	// unknown state as ready. For compatibility reasons, ready should never be
	// "true" for terminating endpoints, except when the normal readiness
	// behavior is being explicitly overridden, for example when the associated
	// Service has set the publishNotReadyAddresses flag.
	// +optional
	ready?: null | bool @go(Ready,*bool) @protobuf(1,bytes)

	// serving is identical to ready except that it is set regardless of the
	// terminating state of endpoints. This condition should be set to true for
	// a ready endpoint that is terminating. If nil, consumers should defer to
	// the ready condition.
	// +optional
	serving?: null | bool @go(Serving,*bool) @protobuf(2,bytes)

	// terminating indicates that this endpoint is terminating. A nil value
	// indicates an unknown state. Consumers should interpret this unknown state
	// to mean that the endpoint is not terminating.
	// +optional
	terminating?: null | bool @go(Terminating,*bool) @protobuf(3,bytes)
}

// EndpointHints provides hints describing how an endpoint should be consumed.
#EndpointHints: {
	// forZones indicates the zone(s) this endpoint should be consumed by to
	// enable topology aware routing.
	// +listType=atomic
	forZones?: [...#ForZone] @go(ForZones,[]ForZone) @protobuf(1,bytes)
}

// ForZone provides information about which zones should consume this endpoint.
#ForZone: {
	// name represents the name of the zone.
	name: string @go(Name) @protobuf(1,bytes)
}

// EndpointPort represents a Port used by an EndpointSlice
// +structType=atomic
#EndpointPort: {
	// name represents the name of this port. All ports in an EndpointSlice must have a unique name.
	// If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name.
	// Name must either be an empty string or pass DNS_LABEL validation:
	// * must be no more than 63 characters long.
	// * must consist of lower case alphanumeric characters or '-'.
	// * must start and end with an alphanumeric character.
	// Default is empty string.
	name?: null | string @go(Name,*string) @protobuf(1,bytes)

	// protocol represents the IP protocol for this port.
	// Must be UDP, TCP, or SCTP.
	// Default is TCP.
	protocol?: null | v1.#Protocol @go(Protocol,*v1.Protocol) @protobuf(2,bytes)

	// port represents the port number of the endpoint.
	// If this is not specified, ports are not restricted and must be
	// interpreted in the context of the specific consumer.
	port?: null | int32 @go(Port,*int32) @protobuf(3,bytes,opt)

	// The application protocol for this port.
	// This is used as a hint for implementations to offer richer behavior for protocols that they understand.
	// This field follows standard Kubernetes label syntax.
	// Valid values are either:
	//
	// * Un-prefixed protocol names - reserved for IANA standard service names (as per
	// RFC-6335 and https://www.iana.org/assignments/service-names).
	//
	// * Kubernetes-defined prefixed names:
	//   * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540
	//
	// * Other protocols should use implementation-defined prefixed names such as
	// mycompany.com/my-custom-protocol.
	// +optional
	appProtocol?: null | string @go(AppProtocol,*string) @protobuf(4,bytes)
}

// EndpointSliceList represents a list of endpoint slices
#EndpointSliceList: {
	metav1.#TypeMeta

	// Standard list metadata.
	// +optional
	metadata?: metav1.#ListMeta @go(ListMeta) @protobuf(1,bytes,opt)

	// items is the list of endpoint slices
	items: [...#EndpointSlice] @go(Items,[]EndpointSlice) @protobuf(2,bytes,rep)
}
