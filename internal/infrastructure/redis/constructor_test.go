package redis_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/Goboolean/fetch-server.v1/internal/infrastructure/redis"
	_ "github.com/Goboolean/fetch-server.v1/internal/util/env"
	"github.com/Goboolean/shared/pkg/resolver"
)

var (
	instance *redis.Redis
	queries  *redis.Queries
)

func Setup() {

	database, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		panic(err)
	}

	instance, err = redis.NewInstance(&resolver.ConfigMap{
		"HOST":     os.Getenv("REDIS_HOST"),
		"PORT":     os.Getenv("REDIS_PORT"),
		"USER":     os.Getenv("REDIS_USER"),
		"PASSWORD": os.Getenv("REDIS_PASS"),
		"DATABASE": database,
	})
	if err != nil {
		panic(err)
	}

	queries = redis.New(instance)
}

func Teardown() {
	instance.Close()
}

func TestMain(m *testing.M) {
	Setup()
	code := m.Run()
	Teardown()
	os.Exit(code)
}

func Test_Constructor(t *testing.T) {
	if err := instance.Ping(); err != nil {
		t.Errorf("Ping() failed: %v", err)
	}
}
