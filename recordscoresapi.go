package main

import (
	"fmt"

	"golang.org/x/net/context"

	rcpb "github.com/brotherlogic/recordcollection/proto"
)

//ClientUpdate on an updated record
func (s *Server) ClientUpdate(ctx context.Context, req *rcpb.ClientUpdateRequest) (*rcpb.ClientUpdateResponse, error) {
	s.Log(fmt.Sprintf("Updating score for %v", req.GetInstanceId()))
	return &rcpb.ClientUpdateResponse{}, nil
}
