package config

type Config struct {
	MongoDBConnection MongoDBConnection
	KeyCloakSettings  KeyCloakSettings
}

type MongoDBConnection struct {
	ConnectionString string
}

type KeyCloakSettings struct {
	DiscoveryUrl string
	ClientId     string
	ClientSecret string
}
