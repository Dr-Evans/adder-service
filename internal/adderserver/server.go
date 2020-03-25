package adderserver

import (
	"context"

	"github.com/Dr-Evans/adder-service/rpc/adder"
)

// Server implements the Haberdasher service
type Server struct{}

func (s Server) AddTwo(ctx context.Context, req *adder.AddTwoReq) (*adder.AddTwoResp, error) {
	sum := req.A + req.B

	return &adder.AddTwoResp{
		Sum: sum,
	}, nil
}
