package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "app",
		Usage: "Simple app for a keda related talk",
		Commands: []cli.Command{
			{
				Name:  "app",
				Usage: "Runs simple http server on 3232",
				Action: func(c *cli.Context) error {
					StartWebserver()
					return nil
				},
			},
			{
				Name:  "api",
				Usage: "Runs simple metric API server on 3232",
				Action: func(c *cli.Context) error {
					StartAPI()
					return nil
				},
			},
			{
				Name:  "redis",
				Usage: "redis methods",
				Subcommands: []cli.Command{
					{
						Name:  "publish",
						Usage: "Publishes messages to Redis list",
						Action: func(c *cli.Context) error {
							result, err := Publish()
							if err != nil {
								fmt.Println("Failed to publish messages")
								log.Fatal(err)
							} else {
								fmt.Printf("Published messages, actual list len: %d\n", result)
							}
							return nil
						},
					},
					{
						Name:  "drain",
						Usage: "Drains the Redis list",
						Action: func(c *cli.Context) error {
							_, err := Drain()
							if err != nil {
								fmt.Println("Failed to drain Redis list")
								log.Fatal(err)
							} else {
								fmt.Println("Queue drained.")
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "mysql",
				Usage: "mysql methods",
				Subcommands: []cli.Command{
					{
						Name:  "insert",
						Usage: "Insert some values to database",
						Action: func(c *cli.Context) error {
							err := Insert()
							if err != nil {
								fmt.Println("Failed to insert values")
								log.Fatal(err)
							} else {
								fmt.Println("Records inserted")
							}
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "Delete all records from database",
						Action: func(c *cli.Context) error {
							err := Delete()
							if err != nil {
								fmt.Println("Failed to delete records")
								log.Fatal(err)
							} else {
								fmt.Println("Records deleted")
							}
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
