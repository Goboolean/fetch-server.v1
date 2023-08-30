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
	"github.com/Goboolean/fetch-server/internal/domain/service/relay"
	"github.com/Goboolean/fetch-server/internal/domain/service/transmission"
	"github.com/Goboolean/fetch-server/internal/infrastructure/grpc"
	"github.com/Goboolean/fetch-server/internal/infrastructure/rdbms"
	"github.com/Goboolean/fetch-server/internal/infrastructure/redis"
	"github.com/Goboolean/fetch-server/internal/infrastructure/ws"
	"github.com/Goboolean/fetch-server/internal/infrastructure/ws/buycycle"
	"github.com/Goboolean/fetch-server/internal/infrastructure/ws/kis"
	"github.com/Goboolean/fetch-server/internal/infrastructure/ws/mock"
	"github.com/Goboolean/fetch-server/internal/infrastructure/ws/polygon"
	"github.com/Goboolean/fetch-server/internal/util/prometheus"
	"github.com/Goboolean/shared/pkg/broker"
	"github.com/Goboolean/fetch-server/internal/infrastructure/mongo"
	"github.com/Goboolean/shared/pkg/resolver"
	"github.com/google/wire"
	"os"
	"strconv"
	"time"
)

// Injectors from infrastructure.go:

func InitMongo() (*mongo.DB, error) {
	configMap := provideMongoArgs()
	db, err := mongo.NewDB(configMap)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitPsql() (*rdbms.PSQL, error) {
	configMap := providePsqlArgs()
	psql, err := rdbms.NewDB(configMap)
	if err != nil {
		return nil, err
	}
	return psql, nil
}

func InitRedis() (*redis.Redis, error) {
	configMap := provideRedisArgs()
	redisRedis, err := redis.NewInstance(configMap)
	if err != nil {
		return nil, err
	}
	return redisRedis, nil
}

func InitKafkaConfigurator() (*broker.Configurator, error) {
	configMap := provideKafkaArgs()
	configurator, err := broker.NewConfigurator(configMap)
	if err != nil {
		return nil, err
	}
	return configurator, nil
}

func InitKafkaPublisher() (*broker.Publisher, error) {
	configMap := provideKafkaArgs()
	publisher, err := broker.NewPublisher(configMap)
	if err != nil {
		return nil, err
	}
	return publisher, nil
}

func InitGrpc(configuratorPort in.ConfiguratorPort) (*grpc.Host, error) {
	configMap := provideGrpcArgs()
	stockConfiguratorServer := grpc2.NewAdapter(configuratorPort)
	host, err := grpc.New(configMap, stockConfiguratorServer)
	if err != nil {
		return nil, err
	}
	return host, nil
}

func InitGrpcClient() (*grpc.Client, error) {
	configMap := provideGrpcArgs()
	client, err := grpc.NewClient(configMap)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func InitBuycycle(receiver ws.Receiver) (*buycycle.Subscriber, error) {
	configMap := provideBuycycleArgs()
	subscriber, err := buycycle.New(configMap, receiver)
	if err != nil {
		return nil, err
	}
	return subscriber, nil
}

func InitKIS(receiver ws.Receiver) (*kis.Subscriber, error) {
	configMap := provideKISArgs()
	subscriber, err := kis.New(configMap, receiver)
	if err != nil {
		return nil, err
	}
	return subscriber, nil
}

func InitPolygon(receiver ws.Receiver) (*polygon.Subscriber, error) {
	configMap := providePolygonArgs()
	subscriber, err := polygon.New(configMap, receiver)
	if err != nil {
		return nil, err
	}
	return subscriber, nil
}

func InitMockWebsocket(duration time.Duration, receiver ws.Receiver) *mock.MockFetcher {
	mockFetcher := mock.New(duration, receiver)
	return mockFetcher
}

func InitPrometheus() (*prometheus.Server, error) {
	configMap := providePrometheusArgs()
	server, err := prometheus.New(configMap)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// Injectors from service.go:

func InitMockRelayer(relayerPort out.RelayerPort) (*relay.Manager, error) {
	stockPersistencePort := persistence.NewMockAdapter()
	tx := transaction.NewMock()
	stockMetadataPort := meta.NewMockAdapter()
	manager, err := relay.New(stockPersistencePort, tx, stockMetadataPort, relayerPort)
	if err != nil {
		return nil, err
	}
	return manager, nil
}

func InitMockPersister(manager *relay.Manager, option persistence2.Option) (*persistence2.Manager, error) {
	tx := transaction.NewMock()
	stockPersistencePort := persistence.NewMockAdapter()
	stockPersistenceCachePort := cache.NewMockAdapter()
	persistenceManager, err := persistence2.New(tx, stockPersistencePort, stockPersistenceCachePort, manager, option)
	if err != nil {
		return nil, err
	}
	return persistenceManager, nil
}

func InitMockTransmitter(manager *relay.Manager, option transmission.Option) (*transmission.Manager, error) {
	transmissionPort := broker2.NewMockAdapter()
	transmissionManager, err := transmission.New(transmissionPort, manager, option)
	if err != nil {
		return nil, err
	}
	return transmissionManager, nil
}

func InitMockConfigurator(manager *relay.Manager, persistenceManager *persistence2.Manager, transmissionManager *transmission.Manager) (*config.Manager, error) {
	stockMetadataPort := meta.NewMockAdapter()
	tx := transaction.NewMock()
	configManager, err := config.New(stockMetadataPort, tx, manager, persistenceManager, transmissionManager)
	if err != nil {
		return nil, err
	}
	return configManager, nil
}

func InitTransactor(mongo2 *mongo.DB, psql *rdbms.PSQL) port.TX {
	tx := transaction.New(mongo2, psql)
	return tx
}

func InitRelayer(tx port.TX, queries *mongo.Queries, rdbmsQueries *rdbms.Queries, relayerPort out.RelayerPort) (*relay.Manager, error) {
	stockPersistencePort := persistence.NewAdapter(rdbmsQueries, queries)
	stockMetadataPort := meta.NewAdapter(rdbmsQueries)
	manager, err := relay.New(stockPersistencePort, tx, stockMetadataPort, relayerPort)
	if err != nil {
		return nil, err
	}
	return manager, nil
}

func InitTransmitter(tx port.TX, option transmission.Option, configurator *broker.Configurator, publisher *broker.Publisher, manager *relay.Manager) (*transmission.Manager, error) {
	transmissionPort := broker2.NewAdapter(configurator, publisher)
	transmissionManager, err := transmission.New(transmissionPort, manager, option)
	if err != nil {
		return nil, err
	}
	return transmissionManager, nil
}

func InitPersister(tx port.TX, option persistence2.Option, queries *redis.Queries, rdbmsQueries *rdbms.Queries, mongoQueries *mongo.Queries, manager *relay.Manager) (*persistence2.Manager, error) {
	stockPersistencePort := persistence.NewAdapter(rdbmsQueries, mongoQueries)
	stockPersistenceCachePort := cache.NewAdapter(queries)
	persistenceManager, err := persistence2.New(tx, stockPersistencePort, stockPersistenceCachePort, manager, option)
	if err != nil {
		return nil, err
	}
	return persistenceManager, nil
}

func InitConfigurator(tx port.TX, queries *rdbms.Queries, manager *persistence2.Manager, transmissionManager *transmission.Manager, relayManager *relay.Manager) (*config.Manager, error) {
	stockMetadataPort := meta.NewAdapter(queries)
	configManager, err := config.New(stockMetadataPort, tx, relayManager, manager, transmissionManager)
	if err != nil {
		return nil, err
	}
	return configManager, nil
}

func InitGrpcWithAdapter(configuratorPort in.ConfiguratorPort) (*grpc.Host, error) {
	configMap := provideGrpcArgs()
	stockConfiguratorServer := grpc2.NewAdapter(configuratorPort)
	host, err := grpc.New(configMap, stockConfiguratorServer)
	if err != nil {
		return nil, err
	}
	return host, nil
}

func InitWs() *websocket.Adapter {
	v := provideFetcher()
	adapter := websocket.NewAdapter(v...)
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
		"PORT": os.Getenv("SERVER_PORT"),
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
	provideGrpcArgs, grpc.New, grpc.NewClient,
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

// service.go:

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
