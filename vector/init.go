package vector

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	pb "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	uri     string
	EmbdKey string

	conn              *grpc.ClientConn
	collectionsClient pb.CollectionsClient
	pointsClient      pb.PointsClient

	vectorSize uint64 = 1536
	distance          = pb.Distance_Cosine

	Qa_collection_name = "question_answer"
)

func InitDB() (err error) {
	uri := os.Getenv("QDRANT_URI")
	EmbdKey = os.Getenv("EMBD_KEY")
	if uri == "" || EmbdKey == "" {
		return fmt.Errorf("please check QDRANT_URI or EmbdKey environment variable")
	}

	conn, err = grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error(err)
		return
	}
	createCollection(Qa_collection_name)
	return
}
func createCollection(collectionName string) {
	// Create new collection
	var defaultSegmentNumber uint64 = 2
	_, err := GetCollectionsClient().Create(context.TODO(), &pb.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: &pb.VectorsConfig{Config: &pb.VectorsConfig_Params{
			Params: &pb.VectorParams{
				Size:     vectorSize,
				Distance: distance,
			},
		}},
		OptimizersConfig: &pb.OptimizersConfigDiff{
			DefaultSegmentNumber: &defaultSegmentNumber,
		},
	})
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Collection ", collectionName, " created")
	}

}
func GetCollectionsClient() pb.CollectionsClient {
	if collectionsClient == nil {
		collectionsClient = pb.NewCollectionsClient(conn)
	}
	return collectionsClient
}
func GetPointsClient() pb.PointsClient {
	if pointsClient == nil {
		pointsClient = pb.NewPointsClient(conn)
	}
	return pointsClient
}
