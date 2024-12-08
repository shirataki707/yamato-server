package main

import (
	"log"
	"yamato/internal/domain"
	"yamato/internal/interface/handler"
	"yamato/internal/interface/router"
	"yamato/internal/repository"
	"yamato/internal/usecase"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Mountain{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

}

func main() {
	dsn := "user:password@tcp(localhost:3306)/hyakumeizan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	autoMigrate(db)

	mountainRepo := repository.NewMountainRepository(db)
	mountainUseCase := usecase.NewMountainUseCase(mountainRepo)
	mountainHandler := handler.NewMountainHandler(mountainUseCase)

	r := router.SetupRouter(mountainHandler)

	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
