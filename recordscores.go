package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gdpb "github.com/brotherlogic/godiscogs/proto"
	pbg "github.com/brotherlogic/goserver/proto"
	rcpb "github.com/brotherlogic/recordcollection/proto"
	rppb "github.com/brotherlogic/recordprocess/proto"
	pb "github.com/brotherlogic/recordscores/proto"
)

// Server main server type
type Server struct {
	*goserver.GoServer
	returnScore     int32
	returnNilScore  bool
	returnRecord    *rcpb.Record
	returnNilRecord bool
	returnUpdate    int32
}

// Init builds the server
func Init() *Server {
	s := &Server{
		GoServer: &goserver.GoServer{},
	}
	return s
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	rcpb.RegisterClientUpdateServiceServer(server, s)
	pb.RegisterRecordScoreServiceServer(server, s)
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{}
}

func (s *Server) getRecord(ctx context.Context, id int32) (*rcpb.Record, error) {
	if s.returnNilRecord {
		return nil, fmt.Errorf("Built to fail")
	}
	if s.returnRecord != nil {
		return s.returnRecord, nil
	}

	conn, err := s.FDialServer(ctx, "recordcollection")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := rcpb.NewRecordCollectionServiceClient(conn)

	resp, err := client.GetRecord(ctx, &rcpb.GetRecordRequest{InstanceId: id})
	if err != nil {
		return nil, err
	}
	return resp.GetRecord(), nil
}

func (s *Server) updateOverallScore(ctx context.Context, id int32, score float32) error {
	if s.returnUpdate > 0 {
		return nil
	}
	conn, err := s.FDialServer(ctx, "recordcollection")
	if err != nil {
		return err
	}
	defer conn.Close()
	client := rcpb.NewRecordCollectionServiceClient(conn)

	_, err = client.UpdateRecord(ctx, &rcpb.UpdateRecordRequest{
		Reason: "scores-push",
		Update: &rcpb.Record{
			Release: &gdpb.Release{
				InstanceId: id,
			},
			Metadata: &rcpb.ReleaseMetadata{
				OverallScore: score,
			},
		},
	})

	return err
}

func (s *Server) readScores(ctx context.Context, iid int32) ([]*pb.Score, error) {
	if s.returnNilScore {
		return nil, fmt.Errorf("Built to fail")
	}
	if s.returnScore > 0 {
		return []*pb.Score{&pb.Score{Rating: s.returnScore}}, nil
	}

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

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.PrepServer("recordscores")
	server.Register = server

	err := server.RegisterServerV2(false)
	if err != nil {
		return
	}

	server.Serve()
}
