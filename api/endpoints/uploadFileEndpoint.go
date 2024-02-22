package endpoints

import (
	"context"
	"database/sql"

	"github.com/Jose-Gomez-c/challenge/api/adapter"
	"github.com/Jose-Gomez-c/challenge/api/contoller"
	"github.com/Jose-Gomez-c/challenge/api/repositories"
	"github.com/Jose-Gomez-c/challenge/api/services"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type FileEndpoints interface {
	MapEndpoints()
}

type FileEndpointsLayer struct {
	engine *gin.Engine
	db     *sql.DB
	rdb    *redis.Client
	ctx    context.Context
}

func NewFileEndpoints(engine *gin.Engine, db *sql.DB, rdb *redis.Client, ctx context.Context) FileEndpoints {
	return &FileEndpointsLayer{engine: engine, db: db, rdb: rdb, ctx: ctx}
}

func (config *FileEndpointsLayer) MapEndpoints() {
	config.setGroup()
	config.uploadEndpoit()

}

func (config *FileEndpointsLayer) setGroup() {
	config.engine.Group("/file")
}

func (config *FileEndpointsLayer) uploadEndpoit() {
	redisAdapter := adapter.NewRedisAdapter(config.rdb)
	itemRepository := repositories.NewItemRepository(config.db)
	httpAdapter := adapter.NewHttpAdapter(config.engine)
	apiService := services.NewApiservice(httpAdapter, redisAdapter)
	uploadServices := services.NewUploadServices(apiService, itemRepository)
	fileContoller := contoller.NewfileController(uploadServices)
	config.engine.POST("/upload", fileContoller.FillDataBase())
}
