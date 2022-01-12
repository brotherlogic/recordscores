package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/brotherlogic/goserver/utils"

	rcpb "github.com/brotherlogic/recordcollection/proto"
	pb "github.com/brotherlogic/recordscores/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&utils.DiscoveryClientResolverBuilder{})
}

func main() {
	ctx, cancel := utils.BuildContext("recordader-cli", "recordscores")
	defer cancel()

	conn, err := utils.LFDialServer(ctx, "recordscores")
	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewRecordScoreServiceClient(conn)

	switch os.Args[1] {
	case "get":
		addFlags := flag.NewFlagSet("AddRecords", flag.ExitOnError)
		var id = addFlags.Int("id", -1, "Id of the record to add")

		if err := addFlags.Parse(os.Args[2:]); err == nil {
			if *id > 0 {
				res, err := client.GetScore(ctx, &pb.GetScoreRequest{InstanceId: int32(*id)})
				if err != nil {
					log.Fatalf("Error on Add Record: %v", err)
				}
				fmt.Printf("%v Scores\n", len(res.GetScores()))
				for i, score := range res.GetScores() {
					fmt.Printf("%v. %v\n", i, score)
				}
				fmt.Printf("Base: %v\n", res.GetComputedScore().GetBaseRating())
				for _, adjust := range res.GetComputedScore().GetAdjustments() {
					fmt.Printf("Adju: %v [%v]\n", adjust.GetValueChange(), adjust.GetType())
				}
				fmt.Printf("--------\n")
				fmt.Printf("Over: %v\n", res.GetComputedScore().GetOverall())
			}
		}
	case "offline":
		res, err := client.GetScore(ctx, &pb.GetScoreRequest{Category: rcpb.ReleaseMetadata_SOLD_OFFLINE})
		if err != nil {
			log.Fatalf("Error on Add Record: %v", err)
		}
		fmt.Printf("%v Scores\n", len(res.GetScores()))
		for i, score := range res.GetScores() {
			fmt.Printf("%v. %v\n", i, score)
		}
	}
}
