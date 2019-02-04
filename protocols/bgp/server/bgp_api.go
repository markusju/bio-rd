package server

import (
	"context"
	"fmt"

	"github.com/bio-routing/bio-rd/protocols/bgp/api"
	"github.com/bio-routing/bio-rd/route"

	bnet "github.com/bio-routing/bio-rd/net"
	routeapi "github.com/bio-routing/bio-rd/route/api"
)

type BGPAPIServer struct {
	srv BGPServer
}

// NewBGPAPIServer creates a new BGP API Server
func NewBGPAPIServer(s BGPServer) *BGPAPIServer {
	return &BGPAPIServer{
		srv: s,
	}
}

func (s *BGPAPIServer) ListSessions(ctx context.Context, in *api.ListSessionsRequest) (*api.ListSessionsResponse, error) {
	return nil, fmt.Errorf("Not implemented yet.")
}

// DumpRIBIn dumps the RIB in of a peer for a given AFI/SAFI
//func (s *BGPAPIServer) DumpRIBIn(ctx context.Context, in *api.DumpRIBRequest) (api.BgpService_DumpRIBInClient, error) {
func (s *BGPAPIServer) DumpRIBIn(in *api.DumpRIBRequest, stream api.BgpService_DumpRIBInServer) error {
	for _, r := range s.srv.DumpRIBIn(bnet.IPFromProtoIP(*in.Peer), uint16(in.Afi), uint8(in.Safi)) {
		x := r.ToProto()
		err := stream.Send(x)
		if err != nil {
			return err
		}
	}

	return nil
}

// DumpRIBOut dumps the RIB out of a peer for a given AFI/SAFI
func (s *BGPAPIServer) DumpRIBOut(in *api.DumpRIBRequest, stream api.BgpService_DumpRIBOutServer) error {
	for _, r := range s.srv.DumpRIBOut(bnet.IPFromProtoIP(*in.Peer), uint16(in.Afi), uint8(in.Safi)) {
		err := stream.Send(r.ToProto())
		if err != nil {
			return err
		}
	}

	return nil
}

func routesToProto(dump []*route.Route) []*routeapi.Route {
	routes := make([]*routeapi.Route, len(dump))
	for i := range dump {
		routes[i] = dump[i].ToProto()
	}

	return routes
}
