package scrap_paper

import "encore.dev/storage/sqldb"

var pgsql = sqldb.NewDatabase("scrap_paper", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

//encore:service
type ScrapPaperService struct {}

func NewScrapPaperService() *ScrapPaperService {
	return &ScrapPaperService{}
}