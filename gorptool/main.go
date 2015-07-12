package main

import (
	"fmt"
	"os"

	goflags "github.com/jessevdk/go-flags"
)

// flags
var flags struct {
	Verbose bool `long:"verbose" short:"v" description:"Show verbose debug information"`

	Test struct {
		All      cmdTestAll      `command:"all" description:"Test all drivers/dialects with local servers"`
		Mysql    cmdTestMysql    `command:"mysql" description:"Test MySQL with local server"`
		Postgres cmdTestPostgres `command:"postgres" description:"Test Postgres with local server"`

		AllDocker      cmdAllDocker          `command:"all-docker" description:"Test all drivers/dialects with databases in docker containers"`
		MysqlDocker    cmdTestMysqlDocker    `command:"mysql-docker" description:"Test MySQL driver+dialect with docker container"`
		PostgresDocker cmdTestPostgresDocker `command:"postgres-docker" description:"Test Postgres driver+dialect with docker container"`
	} `command:"test"`
}

// flags parser
var flagsParser *goflags.Parser

func main() {
	// create flags parser in global var, for flagsParser.Active.Name (operation)
	flagsParser = goflags.NewParser(&flags, goflags.Default)

	// parse flags
	args, err := flagsParser.Parse()
	if err != nil {
		// assert the err to be a flags.Error
		flagError, ok := err.(*goflags.Error)
		if !ok {
			// not a flags error
			os.Exit(1)
		}
		if flagError.Type == goflags.ErrHelp {
			// exitcode 0 when user asked for help
			os.Exit(0)
		}
		if flagError.Type == goflags.ErrUnknownFlag {
			fmt.Println("run with --help to view available options")
		}
		os.Exit(1)
	}

	// error on left-over arguments
	if len(args) > 0 {
		fmt.Printf("unexpected arguments: %s, all arguments after `--` are ignored gracefully\n", args)
		os.Exit(0)
	}
}

func verbosef(format string, args ...interface{}) {
	if flags.Verbose {
		fmt.Printf(format, args...)
	}
}
