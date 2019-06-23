package cli_param

import (
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/urfave/cli.v1/altsrc"
)

const (
	/*--- common flag -----*/
	NameAddress             = `address`
	NameLogLevel            = `logLevel`
	NameLoadConfigFile      = `loadConfigFile`
	NamePProfAddress        = `pProfAddress`
	NameSessionCacheAddress = `sessionCacheAddress`
	NameCypherKey           = `cypherKey`

	/*----database flag -----*/
	NameDbAddress      = `dbAddress`
	NameDbUser         = `dbUser`
	NameDbPass         = `dbPass`
	NameDbSchema       = `dbSchema`
	NameDbMaxLifeTime  = `dbMaxLifeTime`
	NameDbMaxPoolSize  = `dbMaxPoolSize`
	NameDbIdlePoolSize = `dbIdlePoolSize`

	/*--- key value server flag ----*/
	NameKeyServerAddress = `keyServerAddress`

	/*--- http server flag ---*/
	NameHttpAddress           = `httpAddress`
	NameHttpUrlPrefix         = `httpUrlPrefix`
	NameJsonSchemaFile        = `jsonSchemaFile`
	NameInternalHttpAddress   = `internalHttpAddress`
	NameInternalHttpUrlPrefix = `internalHttpUrlPrefix`
	NameSessionKey          = `sessionKey`

	/*---- gRpc flag -----*/
	NameGRpcMaxBackOff = `gRpcMaxBackOff`
)

var (

	/*--- common flag -----*/

	AddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameAddress,
		Usage: `server listen address`,
	})

	LogLevelFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameLogLevel,
		Usage: `log level - debug/warn/info/error`,
	})

	LoadConfigFileFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameLoadConfigFile,
		Usage: `load configuration file for command param`,
	})

	PProfAddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NamePProfAddress,
		Usage: `pProf address`,
	})

	SessionCacheAddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameSessionCacheAddress,
		Usage: `session cache server address`,
	})

	CypherKeyFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameCypherKey,
		Usage: `cypher key`,
	})

	/*----database flag -----*/

	DbAddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameDbAddress,
		Usage: `database address`,
	})

	DbUserFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameDbUser,
		Usage: `database user`,
	})

	DbPassFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameDbPass,
		Usage: `encrypt password`,
	})

	DbSchemaFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameDbSchema,
		Usage: `data base schema`,
	})

	DbMaxLifeTimeFlag = altsrc.NewInt64Flag(cli.Int64Flag{
		Name: NameDbMaxLifeTime,
		Usage: `database connection max unused life time(second),
				if timeout, the connection will be dropped form pool.`,
	})

	DbMaxPoolSizeFlag = altsrc.NewIntFlag(cli.IntFlag{
		Name:  NameDbMaxPoolSize,
		Usage: `max connection pool size of database`,
	})

	DbIdlePoolSizeFlag = altsrc.NewIntFlag(cli.IntFlag{
		Name:  NameDbIdlePoolSize,
		Usage: `max idle connection number of db connection pool`,
	})

	/*--- key value server flag ----*/

	KeyValueAddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameKeyServerAddress,
		Usage: `key - value server address`,
	})

	/*--- http server flag ---*/

	HttpAddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameHttpAddress,
		Usage: `http server address`,
	})

	HttpUrlPrefixFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameHttpUrlPrefix,
		Usage: `http url prefix`,
	})

	JsonSchemaFileFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameJsonSchemaFile,
		Usage: `json schema file`,
	})

	InternalHttpAddressFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameInternalHttpAddress,
		Usage: `internal http server address`,
	})

	InternalHttpUrlPrefixFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameInternalHttpUrlPrefix,
		Usage: `internal http url prefix`,
	})

	SessionKeyFlag = altsrc.NewStringFlag(cli.StringFlag{
		Name:  NameSessionKey,
		Usage: `session key`,
		Value: `tqx-i1-pt-re`,
	})

	/*---- gRpc flag -----*/

	GRPCMaxBackOffFlag = altsrc.NewInt64Flag(cli.Int64Flag{
		Name:  NameGRpcMaxBackOff,
		Usage: `max back off time(ms) for grpc reconnect`,
	})
)

var (
	/*--- common flag -----*/
	GetAddress             = func(c *cli.Context) string { return c.String(NameAddress) }
	GetLogLevel            = func(c *cli.Context) string { return c.String(NameLogLevel) }
	GetLoadConfigFile      = func(c *cli.Context) string { return c.String(NameLoadConfigFile) }
	GetPProfAddress        = func(c *cli.Context) string { return c.String(NamePProfAddress) }
	GetSessionCacheAddress = func(c *cli.Context) string { return c.String(NameSessionCacheAddress) }
	GetCypherKey           = func(c *cli.Context) string { return c.String(NameCypherKey) }

	/*----database flag -----*/
	GetDbAddress      = func(c *cli.Context) string { return c.String(NameDbAddress) }
	GetDbUser         = func(c *cli.Context) string { return c.String(NameDbUser) }
	GetDbPass         = func(c *cli.Context) string { return c.String(NameDbPass) }
	GetDbSchema       = func(c *cli.Context) string { return c.String(NameDbSchema) }
	GetDbMaxLifeTime  = func(c *cli.Context) int64 { return c.Int64(NameDbMaxLifeTime) }
	GetDbMaxPoolSize  = func(c *cli.Context) int { return c.Int(NameDbMaxPoolSize) }
	GetDbIdlePoolSize = func(c *cli.Context) int { return c.Int(NameDbIdlePoolSize) }

	/*--- key value server flag ----*/
	GetKeyServerAddress = func(c *cli.Context) string { return c.String(NameKeyServerAddress) }

	/*--- http server flag ---*/
	GetHttpAddress           = func(c *cli.Context) string { return c.String(NameHttpAddress) }
	GetHttpUrlPrefix         = func(c *cli.Context) string { return c.String(NameHttpUrlPrefix) }
	GetJsonSchemaFile        = func(c *cli.Context) string { return c.String(NameJsonSchemaFile) }
	GetInternalHttpAddress   = func(c *cli.Context) string { return c.String(NameInternalHttpAddress) }
	GetInternalHttpUrlPrefix = func(c *cli.Context) string { return c.String(NameInternalHttpUrlPrefix) }

	/*---- gRpc flag -----*/
	GetGRpcMaxBackOff = func(c *cli.Context) int64 { return c.Int64(NameGRpcMaxBackOff) }
)
