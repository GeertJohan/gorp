package main

import (
	"errors"
	"fmt"
)

// MysqlFlags can be used on the commands `test all` and `test mysql`.
type MysqlFlags struct {
	DSN *string `long:"mysql-dsn" env:"GORP_MYSQL_DSN" description:"MySQL Data Source Name"`

	Address  *string `long:"mysql-address" env:"GORP_MYSQL_ADDRESS" description:"MySQL server tcp address" default-mask:"127.0.0.1:3306"`
	Username *string `long:"mysql-username" env:"GORP_MYSQL_USERNAME" description:"MySQL username" default-mask:"gorptest"`
	Password *string `long:"mysql-password" env:"GORP_MYSQL_PASSWORD" description:"MySQL password" default-mask:"gorptest"`
	Database *string `long:"mysql-database" env:"GORP_MYSQL_DATABASE" description:"MySQL database" default-mask:"gorptest"`
}

// manual checks on flags
func (m *MysqlFlags) check() error {
	if m.DSN != nil {
		if m.Address != nil {
			return errors.New("cannot use --mysql-address with --mysql-dsn")
		}
		if m.Username != nil {
			return errors.New("cannot use --mysql-username with --mysql-dsn")
		}
		if m.Password != nil {
			return errors.New("cannot use --mysql-password with --mysql-dsn")
		}
		if m.Database != nil {
			return errors.New("cannot use --mysql-database with --mysql-dsn")
		}
	} else {
		localhost := "127.0.0.1:3306"
		gorptest := "gorptest"
		if m.Address == nil {
			m.Address = &localhost
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

type cmdTestMysql struct {
	MysqlFlags
}

func (c *cmdTestMysql) Execute(args []string) error {
	if err := c.MysqlFlags.check(); err != nil {
		return err
	}
	var dsn string
	if c.MysqlFlags.DSN != nil {
		dsn = *c.MysqlFlags.DSN
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", *c.MysqlFlags.Username, *c.MysqlFlags.Password, *c.MysqlFlags.Address, *c.MysqlFlags.Database)
	}
	fmt.Printf("connecting to mysql: %s\n", dsn)
	return nil
}

type cmdTestMysqlDocker struct{}

func (c *cmdTestMysqlDocker) Execute(args []string) error {
	// ++ setup docker
	// ++ create custom flags and cmd
	// ++ manually run cmd
	return nil
}
