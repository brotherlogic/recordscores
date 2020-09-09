package main

import (
	"context"
	"testing"
	"time"

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

func TestReadFail(t *testing.T) {
	s := InitTest()

	s.GoServer.KSclient.Fail = true
	s.returnScore = 4
	s.returnRecord = &rcpb.Record{Release: &gdpb.Release{InstanceId: int32(12), Rating: 4}, Metadata: &rcpb.ReleaseMetadata{Category: rcpb.ReleaseMetadata_SOPHMORE}}
	s.returnUpdate = 2

	_, err := s.ClientUpdate(context.Background(), &rcpb.ClientUpdateRequest{InstanceId: int32(12)})
	if err == nil {
		t.Fatalf("Update did not work: %v", err)
	}

	_, err = s.GetScore(context.Background(), &pb.GetScoreRequest{InstanceId: int32(12)})
	if err == nil {
		t.Fatalf("Update score: %v", err)
	}
}

func TestReadEmpty(t *testing.T) {
	s := InitTest()

	s.GoServer.KSclient = *keystoreclient.GetTestClient(".testclient")
	s.returnScore = 4
	s.returnRecord = &rcpb.Record{Release: &gdpb.Release{InstanceId: int32(12), Rating: 4}, Metadata: &rcpb.ReleaseMetadata{Category: rcpb.ReleaseMetadata_SOPHMORE}}
	s.returnUpdate = 2

	_, err := s.ClientUpdate(context.Background(), &rcpb.ClientUpdateRequest{InstanceId: int32(12)})
	if err != nil {
		t.Fatalf("Update did not work: %v", err)
	}
}

func TestBasicInteractionWithPreload(t *testing.T) {
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

func TestBasicInteraction(t *testing.T) {
	s := InitTest()

	s.returnScore = 4
	s.returnRecord = &rcpb.Record{Release: &gdpb.Release{InstanceId: int32(12), Rating: 4}, Metadata: &rcpb.ReleaseMetadata{Category: rcpb.ReleaseMetadata_SOPHMORE}}
	s.returnUpdate = 2
	s.GoServer.KSclient.Save(context.Background(), SCORES, &pb.Scores{Scores: []*pb.Score{&pb.Score{InstanceId: 12, Category: rcpb.ReleaseMetadata_PROFESSOR, ScoreTime: time.Now().Unix()}}})

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

func TestBasicInteractionWithBadRecord(t *testing.T) {
	s := InitTest()

	s.returnScore = 4
	s.returnNilRecord = true
	s.returnUpdate = 2
	s.GoServer.KSclient.Save(context.Background(), SCORES, &pb.Scores{Scores: []*pb.Score{&pb.Score{InstanceId: 12, Category: rcpb.ReleaseMetadata_PROFESSOR, ScoreTime: time.Now().Unix()}}})

	_, err := s.ClientUpdate(context.Background(), &rcpb.ClientUpdateRequest{InstanceId: int32(12)})
	if err == nil {
		t.Fatalf("Update did not work: %v", err)
	}

}
func TestBasicInteractionWithBadScore(t *testing.T) {
	s := InitTest()

	s.returnNilScore = true
	s.returnRecord = &rcpb.Record{Release: &gdpb.Release{InstanceId: int32(12), Rating: 0}, Metadata: &rcpb.ReleaseMetadata{Category: rcpb.ReleaseMetadata_SOPHMORE}}
	s.returnUpdate = 2

	_, err := s.ClientUpdate(context.Background(), &rcpb.ClientUpdateRequest{InstanceId: int32(12)})
	if err == nil {
		t.Fatalf("Update did not work: %v", err)
	}

}

func TestBasicInteractionWithNoChange(t *testing.T) {
	s := InitTest()

	s.returnScore = 4
	s.returnRecord = &rcpb.Record{Release: &gdpb.Release{InstanceId: int32(12), Rating: 0}, Metadata: &rcpb.ReleaseMetadata{Category: rcpb.ReleaseMetadata_PRE_SOPHMORE}}
	s.returnUpdate = 2
	s.GoServer.KSclient.Save(context.Background(), SCORES, &pb.Scores{Scores: []*pb.Score{&pb.Score{InstanceId: 12, Category: rcpb.ReleaseMetadata_PROFESSOR, ScoreTime: time.Now().Unix()}}})

	_, err := s.ClientUpdate(context.Background(), &rcpb.ClientUpdateRequest{InstanceId: int32(12)})
	if err != nil {
		t.Fatalf("Update did not work: %v", err)
	}

}

func TestBasicInteractionLoadOnly(t *testing.T) {
	s := InitTest()

	s.returnScore = 4
	s.returnRecord = &rcpb.Record{Release: &gdpb.Release{InstanceId: int32(12), Rating: 0}, Metadata: &rcpb.ReleaseMetadata{Category: rcpb.ReleaseMetadata_SOPHMORE}}
	s.returnUpdate = 2

	_, err := s.ClientUpdate(context.Background(), &rcpb.ClientUpdateRequest{InstanceId: int32(12)})
	if err != nil {
		t.Fatalf("Update did not work: %v", err)
	}

	scores, err := s.GetScore(context.Background(), &pb.GetScoreRequest{InstanceId: int32(12)})
	if err != nil {
		t.Fatalf("Get Score did not work: %v", err)
	}

	if len(scores.GetScores()) != 0 {
		t.Errorf("Did not get any scores: %v", scores)
	}
}
