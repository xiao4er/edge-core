package client

import (
	"github.com/thingio/edge-core/common/proto/resource"
	"github.com/thingio/edge-core/common/service"
	"github.com/thingio/edge-core/datahub/conf"
	"testing"
)

func init() {
	conf.Load("../etc/edge.yaml")
}

func TestNode(t *testing.T) {
	cli, err := NewDatahubClient(conf.Config.Mqtt, service.ApiServer, conf.Config.NodeId)
	if err != nil {
		t.Errorf("fail to create datahub client, err: %s", err.Error())
		return
	}

	res, err := cli.GetResource(resource.KindNode, conf.Config.NodeId)
	if err != nil {
		t.Errorf("fail to get node resource, err: %s", err.Error())
		return
	}

	t.Logf("get node resource: %+v", res)
}