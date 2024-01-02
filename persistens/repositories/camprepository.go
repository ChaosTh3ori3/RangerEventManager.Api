package repositories

import (
	"context"
	"fmt"
	"log"

	entities_camp "github.com/ChaosTh3ori3/RangerEventManager.Api/models/entities/camp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CampRepository struct {
	collection *mongo.Collection
}

func NewCampRepository(collection *mongo.Collection) CampRepository {
	return CampRepository{
		collection: collection,
	}
}

func (cr *CampRepository) GetCampByCampNumber(campnumber int) entities_camp.Camp {

	filter := bson.D{
		{Key: "Number", Value: campnumber},
	}

	var result entities_camp.Camp
	err := cr.collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Dokument nicht gefunden")
		} else {
			log.Fatal(err)
		}
	}

	return result
}

func (cr *CampRepository) CreateNewCamp(camp entities_camp.Camp) {
	result, err := cr.collection.InsertOne(context.Background(), camp)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
