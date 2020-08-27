package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/recordscores/proto"
)

const (
	// SCORES - All the record scores
	SCORES = "/github.com/brotherlogic/recordscores/scores"
)

//ClientUpdate on an updated record
func (s *Server) ClientUpdate(ctx context.Context, req *rcpb.ClientUpdateRequest) (*rcpb.ClientUpdateResponse, error) {
	data, _, err := s.KSclient.Read(ctx, SCORES, &pb.Scores{})

	code := status.Convert(err).Code()
	if code != codes.OK {
		return nil, err
	}
	scores := data.(*pb.Scores)

	s.Log(fmt.Sprintf("Updating score for %v (%v scores in the db)", req.GetInstanceId(), len(scores.GetScores())))
	return &rcpb.ClientUpdateResponse{}, nil
}
