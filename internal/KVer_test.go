package internal

import "testing"

func TestKVer_Put(t *testing.T) {
	key := "test"
	kver := NewKVer([]string{"127.0.0.1:2379"})
	prevKeyValue, err := kver.Put(key, "hello world3")
	if err != nil {
		t.Fatal("put error:", err)
	}
	if prevKeyValue != nil {
		t.Log(prevKeyValue.CreateRevision, prevKeyValue.ModRevision, prevKeyValue.Version)
		t.Logf("prevKeyValue's Key is:%s, value is: %s", prevKeyValue.Key, prevKeyValue.Value)
	}

	keyValue, err := kver.Get(key)
	if err != nil {
		t.Fatal("Get error:", err)
	}

	t.Log("value is ", string(keyValue.Value))
	t.Log("value version is ", keyValue.Version)
}
