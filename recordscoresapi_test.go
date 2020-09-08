package main

import (
	"context"
	"testing"

	"github.com/brotherlogic/keystore/client"

	gdpb "github.com/brotherlogic/godiscogs"
	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/recordscores/proto"
)

func InitTest() *Server {
	s := Init()
	s.SkipLog = true
	s.SkipIssue = true

	s.GoServer.KSclient = *keystoreclient.GetTestClient(".testclient")
	s.GoServer.KSclient.Save(context.Background(), SCORES, &pb.Scores{Scores: []*pb.Score{&pb.Score{InstanceId: 15}}})

	return s
}

func TestBasicInteraction(t *testing.T) {
	s := InitTest()
	s.returnScore = 4
	s.returnRecord = &rcpb.Record{Release: &gdpb.Release{InstanceId: int32(12), Rating: 4}, Metadata: &rcpb.ReleaseMetadata{Category: rcpb.ReleaseMetadata_SOPHMORE}}
	s.returnUpdate = 2

	_, err := s.ClientUpdate(context.Background(), &rcpb.ClientUpdateRequest{InstanceId: int32(12)})
	if err != nil {
		t.Fatalf("Update did not work: %v", err)
	}

	scores, err := s.GetScore(context.Background(), &pb.GetScoreRequest{InstanceId: int32(12)})
	if err != nil {
		t.Fatalf("Get Score did not work: %v", err)
	}

	if len(scores.GetScores()) == 0 {
		t.Errorf("Did not get any scores: %v", scores)
	}
}
