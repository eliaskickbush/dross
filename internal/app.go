package internal

import (
	"context"
	"database/sql"
	"dross/internal/db"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func RunApp() {

	sqliteDb, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()

	res, err := sqliteDb.ExecContext(ctx, "SELECT 1;")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(res.RowsAffected())

	// params := db.CreateChunkParams{
	// 	Content: "Racing es el mas grande",
	// 	Tags:    nil,
	// }

	// created, err := db.New(sqliteDb).CreateChunk(ctx, params)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Println("Created chunk: ", created)

	chunks, err := db.New(sqliteDb).GetChunks(ctx, db.GetChunksParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(chunks)

}
