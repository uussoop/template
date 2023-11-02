package generics

import (
	"context"
	"fmt"
	"test-qdrant-golang/vector"

	pb "github.com/qdrant/go-client/qdrant"

	log "github.com/sirupsen/logrus"
)

type SearchResponse struct {
	Score   float32
	Payload map[string]*string
}

func DeleteCollection(collectionName string) {
	if !CollectionExists(collectionName) {
		log.Info("Collection ", collectionName, " does not exist")
		return
	}
	_, err := vector.GetCollectionsClient().Delete(context.TODO(), &pb.DeleteCollection{
		CollectionName: collectionName,
	})
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Collection ", collectionName, " deleted")
	}
}

func UpsertPoints(collectionName string, vectors *[][]float32, payload []map[string]*string) error {
	if !CollectionExists(collectionName) {
		return fmt.Errorf("collection %s does not exist", collectionName)
	}
	collection, _ := vector.GetCollectionsClient().
		Get(context.TODO(), &pb.GetCollectionInfoRequest{CollectionName: collectionName})
	CollectionLen := collection.Result.VectorsCount
	fmt.Println(CollectionLen)
	waitUpsert := true
	insertLimit := uint64(100)
	var upsertPoints []*pb.PointStruct
	for i, v := range *vectors {
		pl := map[string]*pb.Value{}
		for k, v := range payload[i] {
			pl[k] = &pb.Value{Kind: &pb.Value_StringValue{StringValue: *v}}
		}
		upsertPoints = append(upsertPoints, &pb.PointStruct{

			Id: &pb.PointId{
				PointIdOptions: &pb.PointId_Num{Num: uint64(i) + CollectionLen},
			},
			Vectors: &pb.Vectors{
				VectorsOptions: &pb.Vectors_Vector{
					Vector: &pb.Vector{Data: v},
				},
			},

			Payload: pl,
		})
		if len(upsertPoints) == int(insertLimit) {
			_, err := vector.GetPointsClient().Upsert(context.TODO(), &pb.UpsertPoints{
				CollectionName: collectionName,
				Wait:           &waitUpsert,
				Points:         upsertPoints,
			})
			if err != nil {
				log.Error(err)
				return err

			} else {
				log.Info("Upsert", len(upsertPoints), "points")
			}
			upsertPoints = nil
		}

	}
	_, err := vector.GetPointsClient().Upsert(context.TODO(), &pb.UpsertPoints{
		CollectionName: collectionName,
		Wait:           &waitUpsert,
		Points:         upsertPoints,
	})
	if err != nil {
		log.Error(err)
		return err
	} else {
		log.Info("Upsert", len(upsertPoints), "points")
	}
	return nil
}
func CollectionExists(collectionName string) bool {
	allcoll, err := vector.GetCollectionsClient().List(context.TODO(), &pb.ListCollectionsRequest{})
	if err != nil {
		log.Error(err)
		return false
	}
	for _, v := range allcoll.Collections {
		if v.Name == collectionName {
			return true
		}
	}
	return false
}
func SearchCollection(
	collectionName string,
	Inputvector []float32,
	limit uint64,
) ([]*SearchResponse, error) {
	if !CollectionExists(collectionName) {
		return nil, fmt.Errorf("collection %s does not exist", collectionName)
	}
	unfilteredSearchResult, err := vector.GetPointsClient().Search(context.TODO(), &pb.SearchPoints{
		CollectionName: collectionName,
		Vector:         Inputvector,
		Limit:          limit,
		// Include all payload and vectors in the search result
		WithPayload: &pb.WithPayloadSelector{
			SelectorOptions: &pb.WithPayloadSelector_Enable{Enable: true},
		},
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	fmt.Println("unfilteredSearchResult: ", unfilteredSearchResult)
	var searchResult []*SearchResponse
	for _, v := range unfilteredSearchResult.Result {
		// pl := v.Payload["data"].GetStringValue()
		pl := map[string]*string{}

		for k, v := range v.Payload {
			val := v.GetStringValue()
			pl[k] = &val
		}
		fmt.Println("payload: ", v.Payload)
		searchResult = append(searchResult, &SearchResponse{
			Score:   v.Score,
			Payload: pl,
		})
	}
	return searchResult, nil
}
