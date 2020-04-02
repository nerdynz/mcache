package mcache

import (
	"context"
	"io"
	"time"

	// "github.com/bradfitz/gomemcache/memcache"

	"google.golang.org/appengine/memcache"
)

type logger interface {
	SetOutput(out io.Writer)
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
}

type Cache struct {
	// Client *memmemcache
}

func New(logger logger) *Cache {
	// mc := memcache.New()
	// opts := &xredis.Options{
	// 	Host:     "localhost",
	// 	Port:     6379,
	// 	Password: "", // no password set
	// 	// DB:       0,  // use default DB
	// }

	// if redisURL != "" {
	// 	opts = &xredis.Options{}
	// 	u, err := url.Parse(redisURL)
	// 	if err != nil {
	// 		logger.Error(err)
	// 		return nil
	// 	}
	// 	opts.Host = u.Host
	// 	if strings.Contains(opts.Host, ":") {
	// 		opts.Host = strings.Split(opts.Host, ":")[0]
	// 	}
	// 	p, _ := u.User.Password()
	// 	opts.Password = p
	// 	// opts.User = u.User.Username()
	// 	port, err := strconv.Atoi(u.Port())
	// 	if err != nil {
	// 		logger.Error("cache couldn't parse port")
	// 		return nil
	// 	}
	// 	opts.Port = port
	// }

	// client := xredis.SetupClient(opts)
	// pong, err := client.Ping()
	// if err != nil {
	// 	logger.Error(err)
	// 	return nil
	// }

	// logger.Info("cache running", pong)
	return &Cache{
		// Client: mc,
	}
}

func (cache *Cache) Get(key string) (string, error) {
	value, err := cache.GetBytes(key)
	if err != nil {
		return "", err
	}
	return string(value[:]), nil
}

func (cache *Cache) Del(key string) error {
	ctx := context.Background()
	err := memcache.Delete(ctx, key)
	return err
}

func (cache *Cache) Expire(key string) error {
	err := cache.Del(key)
	return err
}

func (cache *Cache) GetBytes(key string) ([]byte, error) { // Encourages BLOATED interfaces
	ctx := context.Background()
	item, err := memcache.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	return item.Value, err
}

func (cache *Cache) Set(key string, value string, duration time.Duration) error {
	return cache.SetBytes(key, []byte(value), duration)
}

func (cache *Cache) SetBytes(key string, value []byte, duration time.Duration) error { // Encourages BLOATED interfaces
	// result := string(value[:])
	ctx := context.Background()

	item := &memcache.Item{}
	// secs := int32(duration / time.Second)
	item.Expiration = duration
	item.Key = key
	item.Value = value
	return memcache.Set(ctx, item)
}

func (cache *Cache) FlushDB() error {
	ctx := context.Background()
	return memcache.Flush(ctx)
}
