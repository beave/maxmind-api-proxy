package main

import (

	"fmt"
	"context"
	"strconv"
	"time"
	"log"

	"github.com/redis/go-redis/v9"
	)

var rdb *redis.Client
var ctx = context.Background()

func Redis_Init() {

	connect_string := fmt.Sprintf("%s:%d", Config.Redis_Host, Config.Redis_Port)

	database, _ := strconv.Atoi(Config.Redis_Database)

        rdb = redis.NewClient(&redis.Options{
                Addr:     connect_string,
                Password: Config.Redis_Password,
                DB:       database,
        })

}

func Redis_Check_Cache(ip_address string) []byte {

        key := fmt.Sprintf("geoip:%s", ip_address)

        val, err := rdb.Get(ctx, key).Result()

        if err != nil {
                return nil
        }

        return []byte(val)
}

func Redis_Store_Cache(json_data string, ip_address string) {

        key := fmt.Sprintf("geoip:%s", ip_address)

        err := rdb.Set(ctx, key, json_data, time.Duration(Config.Redis_Cache_Time)*time.Hour).Err()

        if err != nil {
                log.Printf("ERROR: Can't SET into Redis!")
                return
        }

}

