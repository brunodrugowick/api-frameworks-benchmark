package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/brunodrugowick/go-http-server-things/pkg/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TopEntity struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	MiddleEntities []MiddleEntity `gorm:"foreignKey:TopEntityID" json:"middleEntities"`
}

type MiddleEntity struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	TopEntityID   uint          `json:"-"`
	InnerEntities []InnerEntity `gorm:"foreignKey:MiddleEntityID" json:"innerEntities"`
}

type InnerEntity struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	MiddleEntityID uint   `json:"-"`
	Text           string `gorm:"type:varchar(100)" json:"text"`
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
			middleEntity := MiddleEntity{TopEntityID: topEntity.ID}
			db.Create(&middleEntity)

			for k := 0; k < 10; k++ {
				innerEntity := InnerEntity{
					MiddleEntityID: middleEntity.ID,
					Text:           fmt.Sprintf("%d-%d-%d", i, j, k),
				}
				db.Create(&innerEntity)
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

	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	log.Printf("Starting server on port %d", port)

	apiPathHandler := server.NewDefaultPathHandlerBuilder("/api").
		WithHandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, Golang!"))
		}).
		WithHandlerFunc("/top-entities", func(w http.ResponseWriter, r *http.Request) {
			var topEntities []TopEntity
			db.Preload("MiddleEntities.InnerEntities").Find(&topEntities)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(topEntities)
		}).
		Build()

	srv := server.NewDefaultServerBuilder().
		SetPort(port).
		WithPathHandler(apiPathHandler).Build()

	log.Fatal(srv.ListenAndServe())
}
