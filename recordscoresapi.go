package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	rppb "github.com/brotherlogic/recordprocess/proto"
	pb "github.com/brotherlogic/recordscores/proto"
)

const (
	// SCORES - All the record scores
	SCORES = "/github.com/brotherlogic/recordscores/scores"
)

func (s *Server) load(ctx context.Context) (*pb.Scores, error) {
	data, _, err := s.KSclient.Read(ctx, SCORES, &pb.Scores{})

	code := status.Convert(err).Code()

	// InvalidArgument is an empty read
	if code == codes.InvalidArgument {
		data = &pb.Scores{}
	} else if code != codes.OK {
		return nil, err
	}
	scores := data.(*pb.Scores)

	return scores, nil
}

func (s *Server) readScores(ctx context.Context, iid int32) ([]*pb.Score, error) {
	conn, err := s.FDialServer(ctx, "recordprocess")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := rppb.NewScoreServiceClient(conn)
	res, err := client.GetScore(ctx, &rppb.GetScoreRequest{InstanceId: iid})
	if err != nil {
		return nil, err
	}

	scores := []*pb.Score{}
	for _, rs := range res.GetScores() {
		scores = append(scores, &pb.Score{InstanceId: rs.GetInstanceId(), Rating: rs.GetRating(), Category: rs.GetCategory(), ScoreTime: rs.GetScoreTime()})
	}
	return scores, nil
}

//ClientUpdate on an updated record
func (s *Server) ClientUpdate(ctx context.Context, req *rcpb.ClientUpdateRequest) (*rcpb.ClientUpdateResponse, error) {
	scores, err := s.load(ctx)
	if err != nil {
		return nil, err
	}

	subscores := []*pb.Score{}
	for _, score := range scores.GetScores() {
		if score.GetInstanceId() == req.GetInstanceId() {
			subscores = append(subscores, score)
		}
	}

	if len(subscores) == 0 {
		//Seed the scores
		subscores, err = s.readScores(ctx, req.GetInstanceId())
		if err != nil {
			return nil, err
		}
	}

	s.Log(fmt.Sprintf("Updating score for %v (%v scores in the db)", req.GetInstanceId(), len(subscores)))
	return &rcpb.ClientUpdateResponse{}, nil
}
