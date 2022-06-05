package MongoConnInit

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MDB *mongo.Database
)

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig(host string, port int, username string, db_name string, password string) *DBConfig {
	dbConfig := DBConfig{
		Host:     host,
		Port:     port,
		User:     username,
		DBName:   db_name,
		Password: password,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

//InitializeMongoConn Mongo connection initializer
func InitializeMongoConn(host string, port int, username string, DBName string, password string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//options := options.Client().ApplyURI(DbURL(BuildDBConfig(host, port, username, DBName, password)))
	options := options.Client().ApplyURI("mongodb://localhost:27017/?retryWrites=true&w=majority")

	options.SetMinPoolSize(10)
	//options.SetMaxPoolSize(50)

	options.SetMaxConnIdleTime(time.Second * 15)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		panic(err)
	}

	MDB = client.Database(DBName)
}
