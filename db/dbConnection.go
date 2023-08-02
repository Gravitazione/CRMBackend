package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	articleModels "backendCRM/models/article"
	authModels "backendCRM/models/authentication"
	masterModels "backendCRM/models/master"
)

var Context *gorm.DB

// code to connect your DB
// User env constant from .env file here
func Connect() {
	dsn := "postgresql://masteruser:O-jBhqtpZ0vQDsAyR56smg@crm-database-5118.6xw.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	Context = connection
	// authentication
	connection.AutoMigrate(&authModels.Customer{})
	connection.AutoMigrate(&authModels.Auth{})
	connection.AutoMigrate(&authModels.Role{})
	//article
	connection.AutoMigrate(&articleModels.Article{})
	connection.AutoMigrate(&articleModels.ArticleCategory{})
	connection.AutoMigrate(&articleModels.ArticleCategoryMapping{})
	connection.AutoMigrate(&articleModels.ArticleTierMapping{})
	connection.AutoMigrate(&articleModels.ArticleReward{})
	connection.AutoMigrate(&articleModels.ArticleRewardCriteria{})
	//master
	connection.AutoMigrate(&masterModels.Tier{})
}
