package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db := flag.String("db", "", "Database name")
	out := flag.String("out", "./backup", "Output directory")
	uri := flag.String("uri", "mongodb://localhost:27017", "MongoDB URI")

	flag.Parse()

	if *db == "" {
		fmt.Println("Error: db is required")
		os.Exit(1)
	}

	// timestamp folder
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	finalOut := filepath.Join(*out, *db+"_"+timestamp)

	os.MkdirAll(finalOut, os.ModePerm)

	// connect to mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(*uri))
	if err != nil {
		panic(err)
	}

	collections, err := client.Database(*db).ListCollectionNames(context.TODO(), map[string]interface{}{})
	if err != nil {
		panic(err)
	}

	// export each collection
	for _, coll := range collections {
		outFile := filepath.Join(finalOut, coll+".json")

		cmd := exec.Command(
			"mongoexport",
			"--uri="+*uri,
			"--db="+*db,
			"--collection="+coll,
			"--out="+outFile,
			"--jsonArray", // مهم علشان Compass
		)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Println("Exporting:", coll)

		if err := cmd.Run(); err != nil {
			fmt.Println("Error exporting", coll, err)
		}
	}

	fmt.Println("JSON backup ready at:", finalOut)
}
