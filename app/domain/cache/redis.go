package cache

import (
	"encoding/json"
	"os"
	"stncCms/app/domain/helpers/stnccollection"
	"time"

	"github.com/go-redis/redis"
)

var (
	cacheClient = &redisClient{}
)

type redisClient struct {
	r *redis.Client
}

func RedisDBInit() *redisClient {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDb := os.Getenv("REDIS_DB")
	c := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,                       // no password set
		DB:       stnccollection.StringToint(redisDb), // use default DB
	})

	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	cacheClient.r = c
	return cacheClient
}

//GetKey get key
func (cacheClient *redisClient) GetKey(key string) ([]byte, error) {

	cachedProducts, err := cacheClient.r.Get(key).Bytes()
	if err == redis.Nil || err != nil {
		return nil, err
	}
	// err = json.Unmarshal([]byte(val), &src)
	// err = json.Unmarshal(cachedProducts, &src)
	// if err != nil {
	// 	return err
	// }

	return cachedProducts, nil
}

func (cacheClient *redisClient) FlushAll() {
	cacheClient.r.FlushAll()
}

//SetKey set key
func (cacheClient *redisClient) SetKey(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = cacheClient.r.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// type valueEx struct {
// 	Name  string
// 	Email string
// }

// func main() {
// 	redisClient := RedisDBInit()
// 	key1 := "sampleKey"
// 	value1 := &valueEx{Name: "someName", Email: "someemail@abc.com"}
// 	err := redisClient.SetKey(key1, value1, time.Minute*1)
// 	if err != nil {
// 		log.Fatalf("Error: %v", err.Error())
// 	}

// 	value2 := &valueEx{}
// 	err = redisClient.GetKey(key1, value2)
// 	if err != nil {
// 		log.Fatalf("Error: %v", err.Error())
// 	}

// 	log.Printf("Name: %s", value2.Name)
// 	log.Printf("Email: %s", value2.Email)
// }
