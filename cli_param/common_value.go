package cli_param

import (
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/urfave/cli.v1/altsrc"
)

var (

	/*--- common flag -----*/

	AddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `address`,
		Usage: `server listen address`,
	})

	LogLevelFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `logLevel`,
		Usage: `log level - debug/warn/info/error`,
	})

	LoadConfigFileFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `loadConfigFile`,
		Usage: `load configuration file for command param`,
	})

	PProfAddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `pProfAddress`,
		Usage: `pProf address`,
	})

	/*----database flag -----*/

	DbAddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `dbAddress`,
		Usage: `database address`,
	})

	DbUserFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `dbUser`,
		Usage: `database user`,
	})

	DbPassFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `dbPass`,
		Usage: `encrypt password`,
	})

	DbMaxLifeTimeFlag = altsrc.NewInt64Flag(cli.Int64Flag{
		Name: `dbMaxLifeTime`,
		Usage: `database connection max unused life time(second),
				if timeout, the connection will be dropped form pool.`,
	})

	DbMaxPoolSizeFlag = altsrc.NewIntFlag(cli.IntFlag{
		Name:  `dbMaxPoolSize`,
		Usage: `max connection pool size of database`,
	})

	DbIdlePoolSizeFlag = altsrc.NewIntFlag(cli.IntFlag{
		Name:  `dbIdlePoolSize`,
		Usage: `max idle connection number of db connection pool`,
	})

	/*--- key value server flag ----*/

	KeyValueAddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `keyServerAddress`,
		Usage: `key - value server address`,
	})

	/*--- http server flag ---*/

	HttpAddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `httpAddress`,
		Usage: `http server address`,
	})

	HttpUrlPrefixFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `httpUrlPrefix`,
		Usage: `http url prefix`,
	})

	JsonSchemaFileFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `jsonSchemaFile`,
		Usage: `json schema file`,
	})

	InternalHttpAddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `internalHttpAddress`,
		Usage: `internal http server address`,
	})

	InternalHttpUrlPrefixFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  `internalHttpUrlPrefix`,
		Usage: `internal http url prefix`,
	})

	/*---- gRpc flag -----*/

	GRPCMaxBackOffFlag = altsrc.NewInt64Flag(cli.Int64Flag{
		Name:  `gRpcMaxBackOff`,
		Usage: `max back off time(ms) for grpc reconnect`,
	})
)
