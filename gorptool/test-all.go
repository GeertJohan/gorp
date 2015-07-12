package main

type cmdTestAll struct {
	MysqlFlags
	PostgresFlags
}

func (c *cmdTestAll) Execute(args []string) error {
	var err error

	mysql := &cmdTestMysql{
		MysqlFlags: c.MysqlFlags,
	}
	err = mysql.Execute(args)
	if err != nil {
		return err
	}

	postgres := &cmdTestPostgres{
		PostgresFlags: c.PostgresFlags,
	}
	err = postgres.Execute(args)
	if err != nil {
		return err
	}

	return nil
}

type cmdAllDocker struct{}

func (c *cmdAllDocker) Execute(args []string) error {
	var err error

	err = (&cmdTestMysqlDocker{}).Execute(args)
	if err != nil {
		return err
	}

	err = (&cmdTestPostgresDocker{}).Execute(args)
	if err != nil {
		return err
	}

	return nil
}
