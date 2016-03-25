package main

import
(
  "gopkg.in/redis.v3"
  "time"
   "log"
   "fmt"
   "os"
)


type RedisCache struct {

	RedisClient *redis.Client
}


func NewRedisCache() Cache {
    
     r := new(RedisCache)
     r.RedisClient = redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        PoolSize: 20,
        ReadTimeout: 30 * time.Second,	
    })

      f, err := os.OpenFile("/home/joemir/logs/luchettacache.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    if err != nil {
        fmt.Printf("error opening file: %v", err)
    }

    defer f.Close()

    
    //log.SetOutput(f)

     return r
}

func (redisC *RedisCache) Get(key string)  []byte {  


   valueResult, err := redisC.RedisClient.Get(key).Result()
    if err != nil {	
        
              log.Printf(fmt.Sprintf("Error from redis cache key:%s value:%s"), key, valueResult)
    }
       
       //log.Printf(fmt.Sprintf("Get from redis cache key:%s value:%s"), key, valueResult)
    fmt.Println( "valueResult  "+valueResult)
    if valueResult == ""{
       log.Printf(fmt.Sprintf("Get status MISSING from redis cache key:%s  return nil value"), key)
      return nil      
    }

//log.Printf(fmt.Sprintf("Get status hit from redis cache key:%s  return value %s"), key, valueResult)
    return  []byte(valueResult);
}


func (redisC *RedisCache) Put(key string, content []byte)  {

   fmt.Println(" key ="+key)
  log.Printf(fmt.Sprintf("Start Put SUCCESS on redis cache key:%s"), key)
   err := redisC.RedisClient.Set(key, content, 0).Err()
    if err != nil {
         log.Printf(fmt.Sprintf("Error on Put redis cache key:%s "), key)
    }

    log.Printf(fmt.Sprintf("Put SUCCESS on redis cache key:%s"), key)
}	