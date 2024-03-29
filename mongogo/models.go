package mongogo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gauravsarma1992/gostructs"
)

type (
	DBConfig struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		DbName   string `json:"db_name"`
	}
	DB struct {
		config   *DBConfig
		Conn     *mongo.Client
		Database *mongo.Database
		decoder  *gostructs.Decoder

		registeredModels map[string]*gostructs.DecodedResult
	}
)

func (db *DB) GetUrl() (url string) {
	if db.config.Username != "" && db.config.Password != "" {
		url = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
			db.config.Username,
			db.config.Password,
			db.config.Host,
			db.config.Port,
			db.config.DbName,
		)
		return
	}
	url = fmt.Sprintf("mongodb://%s:%s/%s/",
		db.config.Host,
		db.config.Port,
		db.config.DbName,
	)
	return
}

func GetDefaultDBConfig() (dbConfig *DBConfig) {
	dbConfig = &DBConfig{
		Host:   "127.0.0.1",
		Port:   "27017",
		DbName: "dev",
	}
	return
}

func NewDB(dbConfig *DBConfig) (db *DB, err error) {
	if dbConfig == nil {
		dbConfig = GetDefaultDBConfig()
	}
	db = &DB{
		registeredModels: make(map[string]*gostructs.DecodedResult),
		config:           dbConfig,
	}
	if db.decoder, err = gostructs.NewDecoder(&gostructs.DecoderConfig{true}); err != nil {
		return
	}
	if db.Conn, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.GetUrl())); err != nil {
		return
	}
	db.Database = db.Conn.Database(db.config.DbName)

	return
}
