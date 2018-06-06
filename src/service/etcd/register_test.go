package etcd

import (
	"testing"
	"service/action"
)

func TestRegisterService(t *testing.T) {
	InitEtcd("./etcd_config.conf", action.Profile_service_range)
	err := _default_pool.registerService()
	t.Log(err)
}