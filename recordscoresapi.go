package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/recordscores/proto"
)

var (
	scoresGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordscores_scores",
		Help: "The size of the score list",
	})
	overallScores = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordscores_overall_scores",
		Help: "The size of the score list",
	})
)

const (
	// SCORES - All the record scores
	SCORES = "/github.com/brotherlogic/recordscores/scores"
)

func (s *Server) save(ctx context.Context, scores *pb.Scores) error {
	s.metrics(ctx, scores)
	scoresGauge.Set(float64(len(scores.GetScores())))
	return s.KSclient.Save(ctx, SCORES, scores)
}

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
	scoresGauge.Set(float64(len(scores.GetScores())))
	overallScores.Set(float64(len(scores.GetLastScore())))

	s.metrics(ctx, scores)
	return scores, nil
}

//GetScore gets a score
func (s *Server) GetScore(ctx context.Context, req *pb.GetScoreRequest) (*pb.GetScoreResponse, error) {
	scores, err := s.load(ctx)
	if err != nil {
		return nil, err
	}

	subscores := []*pb.Score{}
	for _, score := range scores.GetScores() {
		if score.GetInstanceId() == req.GetInstanceId() {
			subscores = append(subscores, score)
		}
		if score.GetCategory() != rcpb.ReleaseMetadata_UNKNOWN && score.GetCategory() == req.GetCategory() {
			subscores = append(subscores, score)
		}
	}

	cscore, err := s.computeScore(ctx, req.GetInstanceId(), subscores)
	if err != nil {
		return nil, err
	}
	return &pb.GetScoreResponse{Scores: subscores, ComputedScore: cscore}, nil
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

	loaded := false
	if len(subscores) == 0 {
		//Seed the scores
		subscores, err = s.readScores(ctx, req.GetInstanceId())
		if err != nil {
			return nil, err
		}
		loaded = true

		for _, subsc := range subscores {
			scores.Scores = append(scores.Scores, subsc)
		}
	}

	record, err := s.getRecord(ctx, req.GetInstanceId())
	if err != nil {
		if status.Convert(err).Code() == codes.OutOfRange {
			return &rcpb.ClientUpdateResponse{}, nil
		}
		return nil, err
	}

	score, err := s.computeScore(ctx, req.GetInstanceId(), subscores)
	if err == nil && score.GetOverall() != record.GetMetadata().GetOverallScore() {
		s.updateOverallScore(ctx, req.GetInstanceId(), score.GetOverall())
	}

	if record.GetRelease().GetRating() > 0 && !strings.HasPrefix(record.GetMetadata().GetCategory().String(), "PRE") && record.GetMetadata().GetCategory() != rcpb.ReleaseMetadata_UNLISTENED {
		latest := ""
		latestTime := int64(0)
		for _, score := range subscores {
			if score.GetScoreTime() > latestTime {
				latest = score.GetCategory().String()
				latestTime = score.GetScoreTime()
			}
		}

		if record.GetMetadata().GetCategory().String() != latest {
			newScore := &pb.Score{
				ScoreTime:  time.Now().Unix(),
				Rating:     record.GetRelease().GetRating(),
				Category:   record.GetMetadata().GetCategory(),
				InstanceId: record.GetRelease().GetInstanceId(),
			}
			scores.Scores = append(scores.Scores, newScore)
			s.Log(fmt.Sprintf("Adding score to db: %v -> %v", newScore, latest))

			return &rcpb.ClientUpdateResponse{}, s.save(ctx, scores)
		}
	}

	if loaded {
		return &rcpb.ClientUpdateResponse{}, s.save(ctx, scores)
	}

	return &rcpb.ClientUpdateResponse{}, nil
}
