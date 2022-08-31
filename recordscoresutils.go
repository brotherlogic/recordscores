package main

import (
	"fmt"
	"sort"

	"golang.org/x/net/context"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/recordscores/proto"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	scoreCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "recordscores_score_counts",
		Help: "The size of the score list",
	}, []string{"folder", "score"})
	avgScore = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "recordscores_avg_score",
		Help: "The score for a given purchase location",
	}, []string{"location"})
)

func (s *Server) metrics(ctx context.Context, scores *pb.Scores) {
	scoreCount.Reset()

	for _, score := range scores.GetLastScore() {
		scoreCount.With(prometheus.Labels{"folder": fmt.Sprintf("%v", score.GetCurrFolder()), "score": fmt.Sprintf("%v", score.GetOverall())}).Inc()
	}

	count := make(map[string]float64)
	scs := make(map[string]float64)
	for _, score := range scores.GetLastScore() {
		count[score.GetLocation().String()]++
		scs[score.GetLocation().String()] += float64(score.GetOverall())
	}

	for location, sc := range scs {
		avgScore.With(prometheus.Labels{"location": location}).Set(sc / count[location])
	}
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
	return s.computeScoreInternal(ctx, rec, scores)
}

func (s *Server) computeScoreInternal(ctx context.Context, rec *rcpb.Record, scores []*pb.Score) (*pb.ComputedScore, error) {
	sort.SliceStable(scores, func(i, j int) bool {
		return scores[i].GetScoreTime() > scores[j].GetScoreTime()
	})

	s.CtxLog(ctx, fmt.Sprintf("SCORESARE %v", scores))

	sum := float32(0)
	for i := 0; i < min(3, len(scores)); i++ {
		sum += float32(scores[i].GetRating())
	}

	cs := &pb.ComputedScore{
		BaseRating: int32(2 * (sum / float32(min(3, len(scores))))),
		CurrFolder: rec.GetRelease().GetFolderId(),
		Location:   rec.GetMetadata().GetPurchaseLocation(),
	}
	s.CtxLog(ctx, fmt.Sprintf("Base: %v", cs))

	if rec.GetMetadata().GetKeep() == rcpb.ReleaseMetadata_NOT_KEEPER {
		cs.Adjustments = append(cs.Adjustments, &pb.ScoreAdjustment{
			Type:        pb.ScoreAdjustment_KEEP_ADJUSTMENT,
			ValueChange: -3,
		})
	}

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

		if rec.GetMetadata().GetMatch() == rcpb.ReleaseMetadata_FULL_MATCH ||
			rec.GetMetadata().GetMatch() == rcpb.ReleaseMetadata_PARTIAL_MATCH {
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
