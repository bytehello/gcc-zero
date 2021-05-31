package internal

import (
	"context"
	"testing"
)

func TestRegistry_GetConn(t *testing.T) {
	cli, err := GetRegistry().GetConn([]string{"127.0.0.1:2379"})
	if err != nil {
		t.Fatal("GetConn err:", err)
	}
	resp, err := cli.Grant(context.Background(), 100)
	if err != nil {
		t.Fatal("get resp err:", err)
	}
	t.Log("id:", resp.ID)

	putResp, err := cli.Put(context.Background(), "test", "value")
	if err != nil {
		t.Fatal("Put resp err:", err)
	}
	t.Log("putResp:", putResp.PrevKv)

	getResp, err := cli.Get(context.Background(), "test")
	if err != nil {
		t.Fatal("getResp err:", err)
	}
	t.Log("getResp:", getResp.Kvs)
}
