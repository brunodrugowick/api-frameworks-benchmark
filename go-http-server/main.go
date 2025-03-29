package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

type TopEntity struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	MiddleEntities []MiddleEntity `gorm:"many2many:top_entity_middle_entity;" json:"middleEntities"`
}

type MiddleEntity struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	TopEntities   []TopEntity   `gorm:"many2many:top_entity_middle_entity;" json:"-"`
	InnerEntities []InnerEntity `gorm:"many2many:middle_entity_inner_entity" json:"innerEntities"`
}

type InnerEntity struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	MiddleEntities []MiddleEntity `gorm:"many2many:middle_entity_inner_entity"`
	Text           string         `gorm:"type:varchar(100)" json:"text"`
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Retrieve the underlying *sql.DB object
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get *sql.DB: %v", err)
	}
	// Configure the connection pool
	sqlDB.SetMaxIdleConns(10)           // Set the maximum number of idle connections
	sqlDB.SetMaxOpenConns(100)          // Set the maximum number of open connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Set the maximum lifetime of a connection

	// Clear DB
	db.Delete(&InnerEntity{}, "1=1")
	db.Delete(&MiddleEntity{}, "1=1")
	db.Delete(&TopEntity{}, "1=1")
	// Migrate the schema
	db.AutoMigrate(&TopEntity{}, &MiddleEntity{}, &InnerEntity{})

	for i := 0; i < 10; i++ {
		topEntity := TopEntity{}
		db.Create(&topEntity)

		for j := 0; j < 10; j++ {
			middleEntity := MiddleEntity{TopEntities: []TopEntity{topEntity}}
			middleEntity.TopEntities = append(middleEntity.TopEntities, topEntity)
			db.Create(&middleEntity)
			topEntity.MiddleEntities = append(topEntity.MiddleEntities, middleEntity)

			for k := 0; k < 10; k++ {
				innerEntity := InnerEntity{
					Text: fmt.Sprintf("%d-%d-%d", i, j, k),
				}
				innerEntity.MiddleEntities = append(innerEntity.MiddleEntities, middleEntity)
				db.Create(&innerEntity)
				middleEntity.InnerEntities = append(middleEntity.InnerEntities, innerEntity)
			}
		}
	}

	var count int64
	db.Model(&InnerEntity{}).Count(&count)
	log.Printf("Total inner entities: %d", count)

	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "9096"
	}

	log.Printf("Starting server on port :%s", portString)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.Writer.Write([]byte("Hello, Golang"))
		c.Writer.WriteHeader(http.StatusOK)
		return
	})
	r.GET("/api/top-entities", func(c *gin.Context) {
		var topEntities []TopEntity
		if err := db.Preload("MiddleEntities.InnerEntities").Find(&topEntities).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, topEntities)
	})

	err = r.Run(fmt.Sprintf(":%s", portString))
	if err != nil {
		fmt.Print(err)
	}
}
