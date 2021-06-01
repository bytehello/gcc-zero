package internal

import "testing"

func TestKVer_Put(t *testing.T) {
	kver := NewKVer([]string{"127.0.0.1:2379"})
	prevKeyValue, err := kver.Put("test", "hello world")
	if err != nil {
		t.Fatal("put error:", err)
	}
	if prevKeyValue != nil {
		t.Logf("prevKeyValue's Key is:%s, value is: %s", prevKeyValue.Key, prevKeyValue.Value)
	}

	keyValue, err := kver.Get("test")
	if err != nil {
		t.Fatal("Get error:", err)
	}

	t.Log("value is ", string(keyValue.Value))
	t.Log("value version is ", keyValue.Version)
}
