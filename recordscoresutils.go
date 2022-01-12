package main

import (
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/recordscores/proto"
)

func (s *Server) computeScore(ctx context.Context, iid int32, scores []*pb.Score) (*pb.ComputedScore, error) {
	rec, err := s.getRecord(ctx, iid)
	if err != nil {
		return nil, err
	}

	var cs *pb.ComputedScore
	if rec.GetRelease().GetRating() > 0 {
		cs = &pb.ComputedScore{BaseRating: rec.GetRelease().GetRating()}
	} else {
		br := int32(0)
		bt := int64(0)
		for _, score := range scores {
			if score.GetScoreTime() > bt {
				br = (score.GetRating())
				bt = (score.GetScoreTime())
			}
		}

		cs = &pb.ComputedScore{BaseRating: br}
	}

	if len(rec.GetRelease().GetOtherVersions()) > 0 {
		cs.Adjustments = append(cs.Adjustments, &pb.ScoreAdjustment{
			Type:        pb.ScoreAdjustment_OTHER_VERSIONS_ADJUSTMENT,
			ValueChange: -0.5,
		})
	}

	overall := float32(cs.BaseRating)
	for _, adjustment := range cs.Adjustments {
		overall += adjustment.GetValueChange()
	}
	cs.Overall = overall

	return cs, nil
}
