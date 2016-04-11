package etcd

import (
	"github.com/stvp/go-toml-config"
	"time"
	"gylogger"
	"strings"
)

var (
	etcdHosts []string
	dialTimeout time.Duration
	retryDelay time.Duration
	clientFiles []string
	ServiceName string
	ServiceIp string
	ServicePort string
	ActionStart int
	ActionEnd int
)

func InitEtcd(path string) {
	// 加载ETCD配置
	err := loadConfig(path)
	if err != nil {
		logger.Error("load discover config error, %v", err)
		return
	}

	// 初始化ETCD客户端连接
	logger.Debug("init default pool")
	_default_pool.initPool()
	// 连接client_files
	logger.Debug("connect to all service nodes")
	_default_pool.connect()
	// 开启监控
	if len(clientFiles) > 0 {
		logger.Infof("target etcd client count %d, start watcher.", len(clientFiles))
		go _default_pool.watcher()
	} else {
		logger.Infof("target etcd client count %d, no watcher.", len(clientFiles))
	}
}

func loadConfig(path string) (err error) {
	var dialTimeoutSec int64
	var retryDelaySec int64
	var etcdHostStr string
	var clientFileStr string
	discoverConfig := config.NewConfigSet("discoverConfig", config.ExitOnError)
	discoverConfig.StringVar(&etcdHostStr, "etcd_hosts", "http://127.0.0.1:4001")
	discoverConfig.Int64Var(&dialTimeoutSec, "dial_timeout", 10)
	discoverConfig.Int64Var(&retryDelaySec, "retry_delay", 10)
	discoverConfig.StringVar(&clientFileStr, "client_files", "")
	discoverConfig.StringVar(&ServiceName, "service_name", "/gy/test/1")
	discoverConfig.StringVar(&ServiceIp, "service_ip", "127.0.0.1")
	discoverConfig.StringVar(&ServicePort, "service_port", "8010")
	discoverConfig.IntVar(&ActionStart, "action_start", -1)
	discoverConfig.IntVar(&ActionEnd, "action_end", -1)
	err = discoverConfig.Parse(path)
	if err != nil {
		return
	}
	dialTimeout = time.Second * time.Duration(dialTimeoutSec)
	retryDelay = time.Second * time.Duration(retryDelaySec)
	if clientFileStr != "" {
		clientFiles = strings.Split(clientFileStr, ",")
	}
	etcdHosts = strings.Split(etcdHostStr, ",")
	return
}