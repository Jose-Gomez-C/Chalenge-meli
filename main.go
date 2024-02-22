package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Jose-Gomez-c/challenge/api/endpoints"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	eng := gin.Default()
	dataSourceName := "root:mysql@tcp(localhost:33306)/sys"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("no base de datos", err)
		return
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	endpoints := endpoints.NewFileEndpoints(eng, db, rdb, ctx)
	endpoints.MapEndpoints()
	eng.Run(":8080")
}
