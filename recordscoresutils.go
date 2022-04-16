package main

import (
	"fmt"
	"sort"

	"golang.org/x/net/context"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/recordscores/proto"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	scoreCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "recordscores_score_counts",
		Help: "The size of the score list",
	}, []string{"folder", "score"})
)

func (s *Server) metrics(ctx context.Context, scores *pb.Scores) {
	//scoreCount.Reset()

	for _, score := range scores.GetLastScore() {
		scoreCount.With(prometheus.Labels{"folder": fmt.Sprintf("%v", score.GetCurrFolder()), "score": fmt.Sprintf("%v", score.GetOverall())}).Set(float64(1))
	}

	s.CtxLog(ctx, fmt.Sprintf("Read for %v", len(scores.GetLastScore())))
}

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
		if rec.GetMetadata().GetKeep() != rcpb.ReleaseMetadata_DIGITAL_KEEPER {
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

			found := false
			for _, sc := range scores {
				if sc.GetCategory() == rcpb.ReleaseMetadata_LISTED_TO_SELL {
					found = true
				}
			}
			if found {
				cs.Adjustments = append(cs.Adjustments, &pb.ScoreAdjustment{
					Type:        pb.ScoreAdjustment_PREVIOUSY_SOLD_ADJUSTMENT,
					ValueChange: -2,
				})
			}
		}

		if rec.GetMetadata().GetMatch() == rcpb.ReleaseMetadata_FULL_MATCH {
			if rec.GetMetadata().GetKeep() == rcpb.ReleaseMetadata_DIGITAL_KEEPER {
				// Harsher adjustment if we've added this contigency
				cs.Adjustments = append(cs.Adjustments, &pb.ScoreAdjustment{
					Type:        pb.ScoreAdjustment_OWN_OTHER_ADJUSTMENT,
					ValueChange: -10,
				})
			} else {
				cs.Adjustments = append(cs.Adjustments, &pb.ScoreAdjustment{
					Type:        pb.ScoreAdjustment_OWN_OTHER_ADJUSTMENT,
					ValueChange: -5,
				})
			}
		}

	}

	overall := float32(cs.BaseRating)
	for _, adjustment := range cs.Adjustments {
		overall += adjustment.GetValueChange()
	}
	cs.Overall = overall

	// Can't use the default value here - so bump it
	if cs.Overall == 0 {
		cs.Overall = 0.1
	}

	return cs, nil
}
