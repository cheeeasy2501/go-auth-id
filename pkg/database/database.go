package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDatabase interface {
	OpenConnection()
	CloseConnection()
}

type Database struct {
	config IConfig
	conn *sql.DB
}

func NewDB(config IConfig) *Database {
	return &Database{
		config: config,
	}
}

func (db *Database) getDsnString() string {
	return "host=" + db.config.getHost() + " user=" + db.config.getUser() +
	" password=" + db.config.getPassword() + " dbname=" + db.config.getDBName() +
	 " port=" + db.config.getPort() + " sslmode=" + db.config.getSSLMode() + 
	 " timezone=" + db.config.getTimezone() 
}

func (db *Database) OpenConnection() (*gorm.DB, error) {
   // инициализируем коннекшен
   log.Println(db.getDsnString())
   conn, err := gorm.Open(postgres.New(postgres.Config{
	DriverName: "postgres",
	DSN: db.getDsnString(),
  }))
  
  if err != nil {
	return nil, err
  }

  sqlDB, err := conn.DB()

  // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
  sqlDB.SetMaxIdleConns(10)
  
  // SetMaxOpenConns sets the maximum number of open connections to the database.
  sqlDB.SetMaxOpenConns(100)
  
  // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
  sqlDB.SetConnMaxLifetime(time.Hour)

   if err != nil {
	 panic("failed to connect database")
   }

   db.conn = sqlDB

   return conn, nil
}

func (db *Database) CloseConnection() {
	db.conn.Close()
}
