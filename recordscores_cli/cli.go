package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

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

				conn2, err := utils.LFDialServer(ctx, "recordcollection")
				if err != nil {
					log.Fatalf("Cannot reach rc: %v", err)
				}
				rcclient := rcpb.NewRecordCollectionServiceClient(conn2)
				rec, err := rcclient.GetRecord(ctx, &rcpb.GetRecordRequest{InstanceId: int32(*id)})
				if err != nil {
					log.Fatalf("Cannot find record: %v", err)
				}

				fmt.Printf("%v - %v\n", rec.GetRecord().GetRelease().GetArtists()[0].GetName(), rec.GetRecord().GetRelease().GetTitle())
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
	case "ping":
		id, err := strconv.Atoi(os.Args[2])
		sclient := rcpb.NewClientUpdateServiceClient(conn)
		_, err = sclient.ClientUpdate(ctx, &rcpb.ClientUpdateRequest{InstanceId: int32(id)})
		if err != nil {
			log.Fatalf("Error on GET: %v", err)
		}

	case "fullping":
		conn2, err2 := utils.LFDialServer(ctx, "recordcollection")
		if err2 != nil {
			log.Fatalf("Can't dial RC: %v", err2)
		}
		rcclient := rcpb.NewRecordCollectionServiceClient(conn2)
		ids, err := rcclient.QueryRecords(ctx, &rcpb.QueryRecordsRequest{Query: &rcpb.QueryRecordsRequest_All{}})
		if err != nil {
			log.Fatalf("Err: %v", err)
		}

		sclient := rcpb.NewClientUpdateServiceClient(conn)

		for _, id := range ids.GetInstanceIds() {
			_, err = sclient.ClientUpdate(ctx, &rcpb.ClientUpdateRequest{InstanceId: int32(id)})
			if err != nil {
				log.Fatalf("Error on GET: %v", err)
			}
			sc, err := client.GetScore(ctx, &pb.GetScoreRequest{InstanceId: id})
			if err != nil {
				log.Fatalf("Error on getScore %v", err)
			}
			fmt.Printf("%v - %v\n", sc.GetComputedScore().GetOverall(), id)
		}
	}
}
