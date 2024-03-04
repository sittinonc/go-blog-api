package main

import (
	"fmt"
	"os"
	"web-api/Controllers"
	"web-api/Repositories/GormPsqlRepo"
	FiberServer "web-api/Server"
	"web-api/Services"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initEnvironment() FiberServer.ServerConfig {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return FiberServer.ServerConfig{}
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	return FiberServer.ServerConfig{
		Port:            os.Getenv("SERVER_PORT"),
		DatabaseAddress: fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode),
	}
}
func InitDB(config FiberServer.ServerConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DatabaseAddress), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func InitRepositories(db *gorm.DB) (GormPsqlRepo.IUserRepository, GormPsqlRepo.IBlogRepository, GormPsqlRepo.ICommentRepository, GormPsqlRepo.ITagRepository) {
	return GormPsqlRepo.NewUserRepository(db), GormPsqlRepo.NewBlogRepository(db), GormPsqlRepo.NewCommentRepository(db), GormPsqlRepo.NewTagRepository(db)
}

func InitServices(userRepo GormPsqlRepo.IUserRepository, blogRepo GormPsqlRepo.IBlogRepository, commentRepo GormPsqlRepo.ICommentRepository, tagRepo GormPsqlRepo.ITagRepository) (Services.IUserService, Services.IBlogService, Services.ICommentService, Services.ITagService) {
	return Services.NewUserService(userRepo), Services.NewBlogService(blogRepo), Services.NewCommentService(commentRepo), Services.NewTagService(tagRepo)
}

func main() {
	config := initEnvironment()
	db := InitDB(config)
	userRepo, blogRepo, commentRepo, tagRepo := InitRepositories(db)
	userService, blogService, commentService, tagService := InitServices(userRepo, blogRepo, commentRepo, tagRepo)

	controllers := Controllers.New(userService, blogService, commentService, tagService)

	server := FiberServer.NewFiberServer(&config, controllers)

	server.Fiber.Listen(":" + config.Port)
}
