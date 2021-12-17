package server

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	runtime "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

func (c *criService) PodSandboxStats(
	ctx context.Context,
	r *runtime.PodSandboxStatsRequest,
) (*runtime.PodSandboxStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PodSandboxStats not implemented")
}
