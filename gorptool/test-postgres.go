package main

import (
	"errors"
	"fmt"
	"strconv"
)

// PostgresFlags can be used on the commands `test all` and `test posgres`.
type PostgresFlags struct {
	DSN *string `long:"postgres-dsn" env:"GORP_MYSQL_DSN" description:"Postgres Data Source Name"`

	Host     *string `long:"postgres-host" env:"GORP_POSTGRES_HOST" description:"Postgres server host" default-mask:"127.0.0.1"`
	Port     *uint16 `long:"postgres-port" env:"GORP_POSTGRES_PORT" description:"Postgres server port" default-mask:"3306"`
	Username *string `long:"postgres-username" env:"GORP_POSTGRES_USERNAME" description:"Postgres username" default-mask:"gorptest"`
	Password *string `long:"postgres-password" env:"GORP_POSTGRES_PASSWORD" description:"Postgres password" default-mask:"gorptest"`
	Database *string `long:"postgres-database" env:"GORP_POSTGRES_DATABASE" description:"Postgres database" default-mask:"gorptest"`
}

// extra manual flags checks for postgres specific
func (m *PostgresFlags) check() error {
	if m.DSN != nil {
		if m.Host != nil {
			return errors.New("cannot use --postgres-host with --postgres-dsn")
		}
		if m.Port != nil {
			return errors.New("cannot use --postgres-port with --postgres-dsn")
		}
		if m.Username != nil {
			return errors.New("cannot use --postgres-username with --postgres-dsn")
		}
		if m.Password != nil {
			return errors.New("cannot use --postgres-password with --postgres-dsn")
		}
		if m.Database != nil {
			return errors.New("cannot use --postgres-database with --postgres-dsn")
		}
	} else {
		defaulthost := "127.0.0.1"
		defaultport := uint16(3306)
		gorptest := "gorptest"
		if m.Host == nil {
			m.Host = &defaulthost
		}
		if m.Port == nil {
			m.Port = &defaultport
		}
		if m.Username == nil {
			m.Username = &gorptest
		}
		if m.Password == nil {
			m.Password = &gorptest
		}
		if m.Database == nil {
			m.Database = &gorptest
		}
	}
	return nil
}

type cmdTestPostgres struct {
	PostgresFlags
}

func (c *cmdTestPostgres) Execute(args []string) error {
	if err := c.PostgresFlags.check(); err != nil {
		return err
	}
	var dsn string
	if c.PostgresFlags.DSN != nil {
		dsn = *c.PostgresFlags.DSN
	} else {
		// * dbname - The name of the database to connect to
		// * user - The user to sign in as
		// * password - The user's password
		// * host - The host to connect to. Values that start with / are for unix domain sockets. (default is localhost)
		// * port - The port to bind to. (default is 5432)
		// * sslmode - Whether or not to use SSL (default is require, this is not the default for libpq)
		// * fallback_application_name - An application_name to fall back to if one isn't provided.
		// * connect_timeout - Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
		// * sslcert - Cert file location. The file must contain PEM encoded data.
		// * sslrootcert - The location of the root certificate file. The file must contain PEM encoded data.
		// * sslkey - Key file location. The file must contain PEM encoded data.
		dsn = "user=" + *c.PostgresFlags.Username
		dsn += " password=" + *c.PostgresFlags.Password
		dsn += " host=" + *c.PostgresFlags.Host
		dsn += " port=" + strconv.Itoa(int(*c.PostgresFlags.Port))
		dsn += " dbname=" + *c.PostgresFlags.Database
	}
	fmt.Printf("connecting to postgres: %s\n", dsn)
	return nil
}

type cmdTestPostgresDocker struct{}

func (c *cmdTestPostgresDocker) Execute(args []string) error {
	// ++ setup docker
	// ++ create custom flags and cmd
	// ++ manually run cmd
	return nil
}
