package perun

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Red-Sock/Perun/internal/domain"
	"github.com/Red-Sock/Perun/pkg/perun_api"
)

func (s *Impl) ListNodes(ctx context.Context, req *perun_api.ListNodes_Request) (*perun_api.ListNodes_Response, error) {
	listReq := domain.ListVelezNodes{
		SearchPattern: req.GetSearchPattern(),
		Paging:        fromPaging(req.Paging),
	}

	nodes, err := s.nodeService.ListNodes(ctx, listReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := &perun_api.ListNodes_Response{
		Nodes: toVelezNodes(nodes),
	}

	return out, nil
}

func fromPaging(in *perun_api.ListPaging) domain.Paging {
	return domain.Paging{
		Limit:  in.GetLimit(),
		Offset: in.Offset,
	}
}

func toVelezNodes(nodes []domain.VelezConnection) []*perun_api.Node {
	out := make([]*perun_api.Node, 0, len(nodes))

	for _, n := range nodes {
		out = append(out, toVelezNode(n.Node))
	}

	return out
}

func toVelezNode(node domain.Velez) *perun_api.Node {
	out := &perun_api.Node{
		Name: node.Name,
		Addr: node.Addr,
	}

	if node.CustomVelezKeyPath == "" {
		out.CustomVelezKeyPath = &node.CustomVelezKeyPath
	}

	if node.IsInsecure {
		out.SecurityDisabled = &node.IsInsecure
	}

	return out
}
