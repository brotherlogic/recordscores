package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gdpb "github.com/brotherlogic/godiscogs"
	pbg "github.com/brotherlogic/goserver/proto"
	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/recordscores/proto"
)

//Server main server type
type Server struct {
	*goserver.GoServer
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
	return []*pbg.State{
		&pbg.State{Key: "magic", Value: int64(13)},
	}
}

func (s *Server) getRecord(ctx context.Context, id int32) (*rcpb.Record, error) {
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

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.PrepServer()
	server.Register = server

	err := server.RegisterServerV2("recordscores", false, true)
	if err != nil {
		return
	}

	fmt.Printf("%v", server.Serve())
}
