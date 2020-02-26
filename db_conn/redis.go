// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-26
// Time: 10:01

package db_conn

import (
	"fmt"

	"ginchat/common"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedisClient() (err error) {
	RedisClient, err = NewClient()
	return err
}

func NewClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", common.RedisIp, common.RedisPort),
	})
	_, err := client.Ping().Result()
	return client, err
}

