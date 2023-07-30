// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package inject

import (
	broker2 "github.com/Goboolean/fetch-server/internal/adapter/broker"
	"github.com/Goboolean/fetch-server/internal/adapter/cache"
	grpc2 "github.com/Goboolean/fetch-server/internal/adapter/grpc"
	"github.com/Goboolean/fetch-server/internal/adapter/meta"
	"github.com/Goboolean/fetch-server/internal/adapter/persistence"
	"github.com/Goboolean/fetch-server/internal/adapter/transaction"
	"github.com/Goboolean/fetch-server/internal/adapter/websocket"
	"github.com/Goboolean/fetch-server/internal/domain/port"
	"github.com/Goboolean/fetch-server/internal/domain/port/in"
	"github.com/Goboolean/fetch-server/internal/domain/port/out"
	"github.com/Goboolean/fetch-server/internal/domain/service/config"
	persistence2 "github.com/Goboolean/fetch-server/internal/domain/service/persistence"
	"github.com/Goboolean/fetch-server/internal/domain/service/relayer"
	"github.com/Goboolean/fetch-server/internal/domain/service/transmission"
	"github.com/Goboolean/fetch-server/internal/infrastructure/cache/redis"
	"github.com/Goboolean/fetch-server/internal/infrastructure/grpc"
	"github.com/Goboolean/fetch-server/internal/infrastructure/prometheus"
	"github.com/Goboolean/fetch-server/internal/infrastructure/ws"
	"github.com/Goboolean/fetch-server/internal/infrastructure/ws/buycycle"
	"github.com/Goboolean/fetch-server/internal/infrastructure/ws/kis"
	"github.com/Goboolean/fetch-server/internal/infrastructure/ws/polygon"
	"github.com/Goboolean/shared/pkg/broker"
	"github.com/Goboolean/shared/pkg/mongo"
	"github.com/Goboolean/shared/pkg/rdbms"
	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/google/wire"
	"os"
	"strconv"
)

// Injectors from infrastructure.go:

func InitMongo() *mongo.DB {
	configMap := provideMongoArgs()
	db := mongo.NewDB(configMap)
	return db
}

func InitMongoQueries() *mongo.Queries {
	configMap := provideMongoArgs()
	db := mongo.NewDB(configMap)
	queries := mongo.New(db)
	return queries
}

func InitPsql() *rdbms.PSQL {
	configMap := providePsqlArgs()
	psql := rdbms.NewDB(configMap)
	return psql
}

func InitPsqlQueries() *rdbms.Queries {
	configMap := providePsqlArgs()
	psql := rdbms.NewDB(configMap)
	queries := rdbms.NewQueries(psql)
	return queries
}

func InitRedis() *redis.Redis {
	configMap := provideRedisArgs()
	redisRedis := redis.NewInstance(configMap)
	return redisRedis
}

func InitRedisQueries() *redis.Queries {
	configMap := provideRedisArgs()
	redisRedis := redis.NewInstance(configMap)
	queries := redis.New(redisRedis)
	return queries
}

func InitKafkaConfigurator() *broker.Configurator {
	configMap := provideKafkaArgs()
	configurator := broker.NewConfigurator(configMap)
	return configurator
}

func InitKafkaPublisher() *broker.Publisher {
	configMap := provideKafkaArgs()
	publisher := broker.NewPublisher(configMap)
	return publisher
}

func InitGrpc(configuratorPort in.ConfiguratorPort) *grpc.Host {
	configMap := provideGrpcArgs()
	stockConfiguratorServer := grpc2.NewAdapter(configuratorPort)
	host := grpc.New(configMap, stockConfiguratorServer)
	return host
}

func InitBuycycle(receiver ws.Receiver) *buycycle.Subscriber {
	configMap := provideBuycycleArgs()
	subscriber := buycycle.New(configMap, receiver)
	return subscriber
}

func InitKIS(receiver ws.Receiver) *kis.Subscriber {
	configMap := provideKISArgs()
	subscriber := kis.New(configMap, receiver)
	return subscriber
}

func InitPolygon(receiver ws.Receiver) *polygon.Subscriber {
	configMap := providePolygonArgs()
	subscriber := polygon.New(configMap, receiver)
	return subscriber
}

func InitPrometheus() *prometheus.Server {
	configMap := providePrometheusArgs()
	server := prometheus.New(configMap)
	return server
}

// Injectors from service.go:

func InitMockRelayer(relayerPort out.RelayerPort) *relayer.RelayerManager {
	stockPersistencePort := persistence.NewMockAdapter()
	tx := transaction.NewMock()
	stockMetadataPort := meta.NewMockAdapter()
	relayerManager := relayer.New(stockPersistencePort, tx, stockMetadataPort, relayerPort)
	return relayerManager
}

func InitMockPersistenceManager(relayerManager *relayer.RelayerManager, option persistence2.Option) *persistence2.PersistenceManager {
	tx := transaction.NewMock()
	stockPersistencePort := persistence.NewMockAdapter()
	stockPersistenceCachePort := cache.NewMockAdapter()
	persistenceManager := persistence2.New(tx, stockPersistencePort, stockPersistenceCachePort, relayerManager, option)
	return persistenceManager
}

func InitMockTransmissionManager(relayerManager *relayer.RelayerManager, option transmission.Option) *transmission.Transmitter {
	transmissionPort := broker2.NewMockAdapter()
	transmitter := transmission.New(transmissionPort, relayerManager, option)
	return transmitter
}

func InitMockConfigurationManager(relayerManager *relayer.RelayerManager, persistenceManager *persistence2.PersistenceManager, transmitter *transmission.Transmitter) *config.ConfigurationManager {
	stockMetadataPort := meta.NewMockAdapter()
	tx := transaction.NewMock()
	configurationManager := config.New(stockMetadataPort, tx, relayerManager, persistenceManager, transmitter)
	return configurationManager
}

func InitTransactor(mongo2 *mongo.DB, psql *rdbms.PSQL) port.TX {
	tx := transaction.New(mongo2, psql)
	return tx
}

func InitRelayer(tx port.TX, queries *mongo.Queries, rdbmsQueries *rdbms.Queries, relayerPort out.RelayerPort, server *prometheus.Server) *relayer.RelayerManager {
	stockPersistencePort := persistence.NewAdapter(rdbmsQueries, queries, server)
	stockMetadataPort := meta.NewAdapter(rdbmsQueries)
	relayerManager := relayer.New(stockPersistencePort, tx, stockMetadataPort, relayerPort)
	return relayerManager
}

func InitTransmission(tx port.TX, option transmission.Option, configurator *broker.Configurator, publisher *broker.Publisher, relayerManager *relayer.RelayerManager, server *prometheus.Server) *transmission.Transmitter {
	transmissionPort := broker2.NewAdapter(configurator, publisher, server)
	transmitter := transmission.New(transmissionPort, relayerManager, option)
	return transmitter
}

func InitPersistenceManager(tx port.TX, option persistence2.Option, queries *redis.Queries, rdbmsQueries *rdbms.Queries, mongoQueries *mongo.Queries, relayerManager *relayer.RelayerManager, server *prometheus.Server) *persistence2.PersistenceManager {
	stockPersistencePort := persistence.NewAdapter(rdbmsQueries, mongoQueries, server)
	stockPersistenceCachePort := cache.NewAdapter(queries)
	persistenceManager := persistence2.New(tx, stockPersistencePort, stockPersistenceCachePort, relayerManager, option)
	return persistenceManager
}

func InitConfigurationManager(tx port.TX, queries *rdbms.Queries, persistenceManager *persistence2.PersistenceManager, transmitter *transmission.Transmitter, relayerManager *relayer.RelayerManager, server *prometheus.Server) *config.ConfigurationManager {
	stockMetadataPort := meta.NewAdapter(queries)
	configurationManager := config.New(stockMetadataPort, tx, relayerManager, persistenceManager, transmitter)
	return configurationManager
}

func InitGrpcWithAdapter(configuratorPort in.ConfiguratorPort) *grpc.Host {
	configMap := provideGrpcArgs()
	stockConfiguratorServer := grpc2.NewAdapter(configuratorPort)
	host := grpc.New(configMap, stockConfiguratorServer)
	return host
}

func InitWs(server *prometheus.Server) *websocket.Adapter {
	v := provideFetcher()
	adapter := websocket.NewAdapter(server, v...)
	return adapter
}

// infrastructure.go:

func provideMongoArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"HOST":     os.Getenv("MONGO_HOST"),
		"USER":     os.Getenv("MONGO_USER"),
		"PORT":     os.Getenv("MONGO_PORT"),
		"PASSWORD": os.Getenv("MONGO_PASS"),
		"DATABASE": os.Getenv("MONGO_DATABASE"),
	}
}

func providePsqlArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"HOST":     os.Getenv("PSQL_HOST"),
		"USER":     os.Getenv("PSQL_USER"),
		"PORT":     os.Getenv("PSQL_PORT"),
		"PASSWORD": os.Getenv("PSQL_PASS"),
		"DATABASE": os.Getenv("PSQL_DATABASE"),
	}
}

func provideRedisArgs() *resolver.ConfigMap {
	database, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		panic(err)
	}

	return &resolver.ConfigMap{
		"HOST":     os.Getenv("REDIS_HOST"),
		"PORT":     os.Getenv("REDIS_PORT"),
		"USER":     os.Getenv("REDIS_USER"),
		"PASSWORD": os.Getenv("REDIS_PASS"),
		"DATABASE": database,
	}
}

func provideKafkaArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"HOST": os.Getenv("KAFKA_HOST"),
		"PORT": os.Getenv("KAFKA_PORT"),
	}
}

func provideGrpcArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"HOST": os.Getenv("GRPC_HOST"),
		"PORT": os.Getenv("GRPC_PORT"),
	}
}

func provideBuycycleArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"HOST": os.Getenv("BUYCYCLE_HOST"),
		"PORT": os.Getenv("BUYCYCLE_PORT"),
	}
}

func provideKISArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"KIS_APPKEY": os.Getenv("KIS_APPKEY"),
		"KIS_SECRET": os.Getenv("KIS_SECRET"),
	}
}

func providePolygonArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"KEY": os.Getenv("POLYGON_API_KEY"),
	}
}

func providePrometheusArgs() *resolver.ConfigMap {
	return &resolver.ConfigMap{
		"PORT": os.Getenv("METRIC_PORT"),
	}
}

var MongoSet = wire.NewSet(
	provideMongoArgs, mongo.NewDB, mongo.New,
)

var PsqlSet = wire.NewSet(
	providePsqlArgs, rdbms.NewDB, rdbms.NewQueries,
)

var KafkaSet = wire.NewSet(
	provideKafkaArgs, broker.NewConfigurator, broker.NewPublisher, broker.NewSubscriber,
)

var RedisSet = wire.NewSet(
	provideRedisArgs, redis.NewInstance, redis.New,
)

var GrpcSet = wire.NewSet(
	provideGrpcArgs, grpc.New,
)

var BuycycleSet = wire.NewSet(
	provideBuycycleArgs, buycycle.New,
)

var KISSet = wire.NewSet(
	provideKISArgs, kis.New,
)

var PolygonSet = wire.NewSet(
	providePolygonArgs, polygon.New,
)

var PrometheusSet = wire.NewSet(
	providePrometheusArgs, prometheus.New,
)

var InfrastructureSet = wire.NewSet(
	InitMongo,
	InitMongoQueries,
	InitPsql,
	InitPsqlQueries,
	InitKafkaConfigurator,
	InitKafkaPublisher,
	InitRedis,
	InitGrpc,
	InitBuycycle,
	InitKIS,
	InitPolygon,
	InitPrometheus,
)

// service.go:

var ServiceSet = wire.NewSet(relayer.New, persistence2.New, transmission.New, config.New)

func provideFetcher() []ws.Fetcher {
	return []ws.Fetcher{}
}

var AdapterArgumentSet = wire.NewSet(
	provideFetcher,
)

var MockAdapterSet = wire.NewSet(
	AdapterArgumentSet, grpc2.NewAdapter, broker2.NewMockAdapter, cache.NewMockAdapter, meta.NewMockAdapter, persistence.NewMockAdapter, transaction.NewMock, websocket.NewAdapter,
)

var AdapterSet = wire.NewSet(
	AdapterArgumentSet, broker2.NewAdapter, grpc2.NewAdapter, cache.NewAdapter, meta.NewAdapter, persistence.NewAdapter, websocket.NewAdapter,
)

func provideTransmissionArgs() transmission.Option {
	return transmission.Option{}
}

func providePersistenceArgs() persistence2.Option {
	return persistence2.Option{}
}
