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

	gdpb "github.com/brotherlogic/godiscogs"
	rcpb "github.com/brotherlogic/recordcollection/proto"
	rppb "github.com/brotherlogic/recordprocess/proto"
	pb "github.com/brotherlogic/recordscores/proto"
)

var (
	scoresGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "recordscores_scores",
		Help: "The size of the score list",
	})
)

const (
	// SCORES - All the record scores
	SCORES = "/github.com/brotherlogic/recordscores/scores"
)

func (s *Server) save(ctx context.Context, scores *pb.Scores) error {
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
	}

	return &pb.GetScoreResponse{Scores: subscores}, nil
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

	// Do we need to add the current score
	conn, err := s.FDialServer(ctx, "recordcollection")
	if err != nil {
		return nil, err
	}
	client := rcpb.NewRecordCollectionServiceClient(conn)

	resp, err := client.GetRecord(ctx, &rcpb.GetRecordRequest{InstanceId: req.GetInstanceId()})
	if err != nil {
		return nil, err
	}

	if resp.GetRecord().GetRelease().GetRating() > 0 && !strings.HasPrefix(resp.GetRecord().GetMetadata().GetCategory().String(), "PRE") && resp.GetRecord().GetMetadata().GetCategory() != rcpb.ReleaseMetadata_UNLISTENED {
		latest := ""
		latestTime := int64(0)
		for _, score := range subscores {
			if score.GetScoreTime() > latestTime {
				latest = score.GetCategory().String()
				latestTime = score.GetScoreTime()
			}
		}

		if resp.GetRecord().GetMetadata().GetCategory().String() != latest {
			newScore := &pb.Score{
				ScoreTime:  time.Now().Unix(),
				Rating:     resp.GetRecord().GetRelease().GetRating(),
				Category:   resp.GetRecord().GetMetadata().GetCategory(),
				InstanceId: resp.GetRecord().GetRelease().GetInstanceId(),
			}
			scores.Scores = append(scores.Scores, newScore)
			s.Log(fmt.Sprintf("Adding score to db: %v -> %v", newScore, latest))

			sc := float32(resp.GetRecord().GetRelease().GetRating())
			count := float32(1)
			for _, subscore := range subscores {
				sc += float32(subscore.GetRating())
				count++
			}

			client.UpdateRecord(ctx, &rcpb.UpdateRecordRequest{
				Reason: "scores-push",
				Update: &rcpb.Record{
					Release: &gdpb.Release{
						InstanceId: resp.GetRecord().GetRelease().GetInstanceId(),
					},
					Metadata: &rcpb.ReleaseMetadata{
						OverallScore: sc / count,
					},
				},
			})

			return &rcpb.ClientUpdateResponse{}, s.save(ctx, scores)
		}
	}

	if loaded {
		return &rcpb.ClientUpdateResponse{}, s.save(ctx, scores)
	}

	return &rcpb.ClientUpdateResponse{}, nil
}
