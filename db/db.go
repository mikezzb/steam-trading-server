package db

import (
	"log"
	"time"

	"github.com/mikezzb/steam-trading-server/pkg/setting"
	"github.com/mikezzb/steam-trading-shared/database"
	"github.com/mikezzb/steam-trading-shared/database/repository"
)

var dbClient *database.DBClient
var Repos repository.RepoFactory

func Setup() {
	var err error
	dbClient, err = database.NewDBClient(
		setting.DB.DatabaseUri,
		setting.DB.DatabaseName,
		10*time.Second,
	)

	if err != nil {
		log.Fatalf("database.NewDBClient err: %v", err)
	}

	Repos = repository.NewRepoFactory(dbClient, nil)

}
