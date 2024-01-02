package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var CampCollection *mongo.Collection

type Database struct {
	connectionstring *string
}

func NewDatabase(connectionstring *string) Database {
	return Database{
		connectionstring: connectionstring,
	}
}

func (db Database) InitDatabase() {
	// MongoDB-Verbindung herstellen
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(*db.connectionstring))
	if err != nil {
		log.Fatal(err)
	}

	// Überprüfen, ob die Verbindung erfolgreich ist
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Erfolgreich mit MongoDB verbunden")

	// Datenbank auswählen
	DB = client.Database("RangerEventManager")
	CampCollection = DB.Collection("Camps")
}

// Verwende diese Funktionen in deinem Code
func (db Database) GetCampsCollection() *mongo.Collection {
	return CampCollection
}
