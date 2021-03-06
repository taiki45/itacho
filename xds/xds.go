package xds

import (
	"fmt"
	"path/filepath"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/golang/protobuf/jsonpb"
)

const (
	// EndpointType for xDS resource type
	EndpointType = cache.EndpointType
	// ClusterType for xDS resource type
	ClusterType = cache.ClusterType
	// RouteType for xDS resource type
	RouteType = cache.RouteType
)

const (
	// CDS flag
	CDS ResponseType = iota
	// RDS flag
	RDS
	// EDS flag
	EDS
)

// ResponseType to switch xDS type in itacho
type ResponseType = int

// CdsPathForNode returns xDS API path for specific node
func CdsPathForNode(nodeCluster string) string {
	return filepath.Join(baseDirCds, nodeCluster)
}

// RdsPathForNode returns xDS API path for specific node
func RdsPathForNode(nodeCluster string) string {
	return filepath.Join(baseDirRds, nodeCluster)
}

const (
	// baseDirCds for object strage base path
	baseDirCds = "/v2/discovery/clusters"
	// baseDirRds for object strage base path
	baseDirRds = "/v2/discovery/routes"
)

// ExtractNodeCluster returns cluster value from Node proto message
func ExtractNodeCluster(node *core.Node) string {
	if node == nil {
		return "unknown"
	}
	return node.Cluster
}

// UnmarshalDiscoveryRequest build Envoy's DiscoveryRequest proto message from JSON string
func UnmarshalDiscoveryRequest(typeURL string, body *[]byte) (*v2.DiscoveryRequest, error) {
	req := &v2.DiscoveryRequest{}
	if err := jsonpb.UnmarshalString(string(*body), req); err != nil {
		return nil, fmt.Errorf("Failed parse JSON body: %s", err)
	}
	req.TypeUrl = typeURL
	return req, nil
}
