package main

import (
	"context"
	"testing"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/recordscores/proto"
	"google.golang.org/protobuf/proto"
)

func TestNoKeepScoredLower(t *testing.T) {
	r1 := &rcpb.Record{Metadata: &rcpb.ReleaseMetadata{Keep: rcpb.ReleaseMetadata_KEEPER}}
	r2 := proto.Clone(r1).(*rcpb.Record)

	r2.GetMetadata().Keep = rcpb.ReleaseMetadata_NOT_KEEPER

	s := InitTest()
	scores := []*pb.Score{{Rating: 5}}
	s1, err1 := s.computeScoreInternal(context.Background(), r1, scores)
	s2, err2 := s.computeScoreInternal(context.Background(), r2, scores)

	if err1 != nil || err2 != nil {
		t.Fatalf("Bad compute: %v or %v", err1, err2)
	}

	if s1.GetOverall() <= s2.GetOverall() {
		t.Errorf("Score 1 was higher than score 2: %v vs %v", s1, s2)
	}
}
