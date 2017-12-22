
package main

import (
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/configuration"
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"strings"
	"github.com/SantoDE/datahamster/storage"
	"github.com/SantoDE/datahamster/server"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "server.address",
			Usage:  "The Database Name to Backup",
			EnvVar: "DATABASE_NAME",
		},
		cli.StringFlag{
			Name:   "dump.identifier",
			Usage:  "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "SCHEDULE_DURATION",
		},
		cli.StringFlag{
			Name:   "dump.storage.type",
			Usage:  "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "SCHEDULE_DURATION",
		},
		cli.StringFlag{
			Name:   "dump.storage.file.dir",
			Usage:  "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "SCHEDULE_DURATION",
		},
		cli.StringFlag{
			Name:   "log.level",
			Usage:  "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "SCHEDULE_DURATION",
		},
	}

	app.Action = func(c *cli.Context) error {
		globalConfiguration := initConfig(c)
		level, err := logrus.ParseLevel(strings.ToLower(globalConfiguration.LogLevel))
		if err != nil {
			log.Error("Error getting level", err)
		}
		log.SetLevel(level)

		cfg := globalConfiguration.Server

		fmt.Printf("Server Adress %s", cfg.Address)
		server.StartRpc()

		return nil
	}

	app.Name = "Datahamster - Worker"
	app.Usage = "Worker to automatically get databse dumps and forward them to the server"

	app.Run(os.Args)
}

func initConfig(c *cli.Context) configuration.GlobalConfiguration {

	var serverAddress = c.String("server.address")
	var logLevel = c.String("log.level")
	var dumpIdentifier = c.String("dump.identifier")


	dumpConfig := new(configuration.DumpConfiguration)
	dumpConfig.Identifier = dumpIdentifier
	storageConfig := new(storage.StorageConfiguration)

	switch storageType := c.String("dump.storage.type"); storageType {

	case "file":
		var storageDir = c.String("dump.storage.file.dir")
		storageConfig.File = storage.FileStorageConfiguration{
			Dir: storageDir,
		}

	default:
		storageConfig.Type = storageType
	}

	dumps := []configuration.DumpConfiguration{}

	dumps = append(dumps, *dumpConfig)

	config := configuration.GlobalConfiguration{
		Server: configuration.ServerConfiguration{
			Address: serverAddress,
		},
		Dumps:  dumps,
		LogLevel: logLevel,
	}

	return config
}