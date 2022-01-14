package main

import (
	"fmt"
	"sort"

	"golang.org/x/net/context"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/recordscores/proto"
)

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func (s *Server) computeScore(ctx context.Context, iid int32, scores []*pb.Score) (*pb.ComputedScore, error) {
	rec, err := s.getRecord(ctx, iid)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(scores, func(i, j int) bool {
		return scores[i].GetScoreTime() > scores[j].GetScoreTime()
	})

	s.Log(fmt.Sprintf("SCORESARE %v", scores))

	sum := float32(0)
	for i := 0; i < min(3, len(scores)); i++ {
		sum += float32(scores[i].GetRating())
	}

	cs := &pb.ComputedScore{BaseRating: int32(2 * (sum / float32(min(3, len(scores)))))}

	if rec.GetMetadata().GetKeep() != rcpb.ReleaseMetadata_KEEPER {
		if len(rec.GetRelease().GetOtherVersions()) > 0 {
			cs.Adjustments = append(cs.Adjustments, &pb.ScoreAdjustment{
				Type:        pb.ScoreAdjustment_OTHER_VERSIONS_ADJUSTMENT,
				ValueChange: -1,
			})
		}

		if len(rec.GetRelease().GetDigitalVersions()) > 0 {
			cs.Adjustments = append(cs.Adjustments, &pb.ScoreAdjustment{
				Type:        pb.ScoreAdjustment_DIGITAL_VERSIONS_ADJUSTMENT,
				ValueChange: -2,
			})
		}
	}

	overall := float32(cs.BaseRating)
	for _, adjustment := range cs.Adjustments {
		overall += adjustment.GetValueChange()
	}
	cs.Overall = overall

	return cs, nil
}
