package mongo

import (
	"gylogger"
	"gopkg.in/mgo.v2"
	"github.com/stvp/go-toml-config"
)

var (
	dialContext *DialContext
	dbConfig *config.ConfigSet

	dbUrl string
	DBName string
	dbConnPoolSize int
)

func InitMongo(path string) {
	loadConfig(path)
	mongodbContext, err := Dial(dbUrl, dbConnPoolSize)
	if err != nil {
		logger.Debug("Could not establish connection with db:", err)
	} else {
		logger.Debugf("Connected to db %s.", dbUrl)
		dialContext = mongodbContext
	}
}

func loadConfig(path string) {
	dbConfig = config.NewConfigSet("dbConfig", config.ExitOnError)
	dbConfig.StringVar(&dbUrl, "db_url", "mongodb://127.0.0.1:27017")
	dbConfig.StringVar(&DBName, "db_name", "default_schema")
	dbConfig.IntVar(&dbConnPoolSize, "db_conn_pool_size", 5)
	err := dbConfig.Parse(path)
	if err != nil {
		logger.Warnf("load dbconfig error, %v", err)
	} else {
		logger.Info("loaded dbconfig")
	}
}

func getSession() *Session {
	context := dialContext.Ref()
	return context
}

func returnSession(session *Session) {
	dialContext.UnRef(session)
}

func GetCollection(name string) *mgo.Collection {
	session := getSession()
	defer returnSession(session)
	return session.Session.DB(DBName).C(name)
}

func GetGridFS(category string) *mgo.GridFS {
	session := getSession()
	defer returnSession(session)
	return session.DB(DBName).GridFS(category)
}