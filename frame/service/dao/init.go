package dao

import (
	"FRAME/conf"
	"FRAME/service/dao/query"
	"fmt"
	"github.com/panjiawan/go-lib/pkg/plog"
	"github.com/panjiawan/go-lib/pkg/pmysql"
	"github.com/panjiawan/go-lib/pkg/predis"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	DB_SHIP_USEKEY = "default"
	MYSQL_KEY      = "default"
	REDIS_KEY      = "default"
)

func Run() {
	InitMysql()
	InitRedis()
}

var (
	mysqlHandles   map[string]*pmysql.Service
	redisHandles   map[string]*predis.Service
	redisKeyPrefix map[string]string
)

func InitMysql() {
	mysqlConf := conf.GetHandle().GetMysqlConf()

	mysqlHandles = make(map[string]*pmysql.Service)

	if len(mysqlConf.Hosts) == 1 {
		MYSQL_KEY = mysqlConf.Hosts[0].Name
	}

	for _, cfg := range mysqlConf.Hosts {
		mysqlHandles[cfg.Name] = pmysql.New(
			pmysql.WithConnection(cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DB),
			pmysql.WithLimit(cfg.MaxIdle, cfg.MaxOpen),
			pmysql.WithPrefix(cfg.Prefix),
			pmysql.WithDebug(cfg.Debug),
			pmysql.WithCharset("utf8mb4"),
		)
	}

	for k, f := range mysqlHandles {
		if err := f.Run(); err != nil {
			plog.Error("mysql start error", zap.String("key", k), zap.Error(err))
			panic(err)
		} else {
			plog.Info("mysql started", zap.String("key", k))
		}
	}
}

func InitRedis() {
	redisConf := conf.GetHandle().GetRedisConf()

	redisHandles = make(map[string]*predis.Service)
	redisKeyPrefix = make(map[string]string)
	if len(redisConf.Hosts) == 1 {
		REDIS_KEY = redisConf.Hosts[0].Name
	}

	for _, cfg := range redisConf.Hosts {
		timeout := time.Duration(cfg.Timeout) * time.Second
		redisHandles[cfg.Name] = predis.New(
			predis.WithConnection(cfg.Host, cfg.Port),
			predis.WithAuth(cfg.Auth),
			predis.WithDB(cfg.DB),
			predis.WithLimit(cfg.MinIdle, cfg.MaxIdle),
			predis.WithReadTimeout(timeout),
			predis.WithWriteTimeout(timeout),
		)

		redisKeyPrefix[cfg.Name] = cfg.Prefix
	}

	for k, f := range redisHandles {
		if err := f.Run(); err != nil {
			plog.Error("redis start error", zap.String("key", k), zap.Error(err))
			panic(err)
		} else {
			plog.Info("redis started", zap.String("key", k))
		}
	}
}

func Mysql(key ...string) *gorm.DB {
	if len(key) == 1 {
		return mysqlHandles[key[0]].Handle()
	}
	return mysqlHandles[MYSQL_KEY].Handle()
}

func Query(key ...string) *query.Query {
	db := Mysql(key...)
	return query.Use(db)
}

func Redis(key ...string) *redis.Client {
	if len(key) == 1 {
		return redisHandles[key[0]].GetConn()
	}
	return redisHandles[REDIS_KEY].GetConn()
}

func FormatRedisKey(key string) string {
	return fmt.Sprintf("%s:%s", redisKeyPrefix[REDIS_KEY], key)
}
