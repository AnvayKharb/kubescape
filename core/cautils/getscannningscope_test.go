package cautils

import (
	"testing"

	"github.com/kubescape/opa-utils/reporthandling"
	"github.com/kubescape/opa-utils/reporthandling/apis"
	reporthandlingv2 "github.com/kubescape/opa-utils/reporthandling/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetScanningScope(t *testing.T) {
	tests := []struct {
		name     string
		metadata reporthandlingv2.ContextMetadata
		want     reporthandling.ScanningScopeType
	}{
		{
			name:     "no cluster context — file scope",
			metadata: reporthandlingv2.ContextMetadata{},
			want:     reporthandling.ScopeFile,
		},
		{
			name: "cluster context with no cloud metadata — cluster scope",
			metadata: reporthandlingv2.ContextMetadata{
				ClusterContextMetadata: &reporthandlingv2.ClusterMetadata{},
			},
			want: reporthandling.ScopeCluster,
		},
		{
			name: "cluster context with nil cloud metadata — cluster scope",
			metadata: reporthandlingv2.ContextMetadata{
				ClusterContextMetadata: &reporthandlingv2.ClusterMetadata{
					CloudMetadata: nil,
				},
			},
			want: reporthandling.ScopeCluster,
		},
		{
			name: "cluster context with empty cloud provider — cluster scope",
			metadata: reporthandlingv2.ContextMetadata{
				ClusterContextMetadata: &reporthandlingv2.ClusterMetadata{
					CloudMetadata: &reporthandlingv2.CloudMetadata{
						CloudProvider: "",
					},
				},
			},
			want: reporthandling.ScopeCluster,
		},
		{
			name: "cluster context with EKS cloud provider — EKS scope",
			metadata: reporthandlingv2.ContextMetadata{
				ClusterContextMetadata: &reporthandlingv2.ClusterMetadata{
					CloudMetadata: &reporthandlingv2.CloudMetadata{
						CloudProvider: apis.EKS,
					},
				},
			},
			want: reporthandling.ScopeCloudEKS,
		},
		{
			name: "cluster context with GKE cloud provider — GKE scope",
			metadata: reporthandlingv2.ContextMetadata{
				ClusterContextMetadata: &reporthandlingv2.ClusterMetadata{
					CloudMetadata: &reporthandlingv2.CloudMetadata{
						CloudProvider: apis.GKE,
					},
				},
			},
			want: reporthandling.ScopeCloudGKE,
		},
		{
			name: "cluster context with AKS cloud provider — AKS scope",
			metadata: reporthandlingv2.ContextMetadata{
				ClusterContextMetadata: &reporthandlingv2.ClusterMetadata{
					CloudMetadata: &reporthandlingv2.CloudMetadata{
						CloudProvider: apis.AKS,
					},
				},
			},
			want: reporthandling.ScopeCloudAKS,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetScanningScope(tt.metadata))
		})
	}
}
