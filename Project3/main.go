package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	db := flag.String("db", "", "Database name")
	out := flag.String("out", "./backup", "Output directory")
	uri := flag.String("uri", "", "MongoDB URI")
	gzip := flag.Bool("gzip", true, "Enable gzip compression")

	flag.Parse()

	if *db == "" {
		fmt.Println("Error: db is required")
		os.Exit(1)
	}

	// timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	finalOut := fmt.Sprintf("%s_%s", *out, timestamp)

	args := []string{}

	if *uri != "" {
		args = append(args, "--uri", *uri)
	} else {
		args = append(args, "--db", *db)
	}

	if *gzip {
		args = append(args, "--gzip")
		args = append(args, "--archive="+finalOut+".gz")
	} else {
		args = append(args, "--out", finalOut)
	}

	cmd := exec.Command("mongodump", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running mongodump...")

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Backup saved to:", finalOut)
}
