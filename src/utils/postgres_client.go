
package utils

import (
	"fmt"

	"cmd/config"
	
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)
var (
	PostgresClient *sql.DB
	pgHost = config.PostgresHost
	pgPort = config.PostgresPort
	pgUser = config.PostgresUser
	pgPassword = config.PostgresPassword
	pgDbName = config.PostgresDbName
)


func ConnectToPostgres(){
	if PostgresClient == nil || PostgresClient.Ping() != nil {
		// fmt.Println("Connecting to DB...")
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			pgHost, pgPort, pgUser, pgPassword, pgDbName)
		
		var err error
		PostgresClient, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Errorf("Error while creating pg connection : %v", err.Error())
			return
		}
	}
	log.Info("Connected to Postgers!")
}
