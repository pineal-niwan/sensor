package main

import (
	"context"
	"github.com/pineal-niwan/sensor/cache/key_value_service/client"
	"testing"
	"time"
)

func TestStringKV1(t *testing.T) {
	clientKV := client.NewStringAsKeyClient(`localhost:10001`, 6000)
	err := clientKV.Dial(time.Second * 3)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}

	ctx := context.Background()

	err = clientKV.Set(ctx, `1`, []byte{1, 2, 3})
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}

	l1, l2, err := clientKV.GetLen(ctx)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}
	t.Logf(`l1:%+v -l2:%+v`, l1, l2)

	buf, err := clientKV.Get(ctx, `1`)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}
	t.Logf(`buf:%+v`, buf)
}

func TestStringKV2(t *testing.T) {
	clientKV := client.NewStringAsKeyClient(`localhost:10001`, 6000)
	err := clientKV.Dial(time.Second * 3)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}

	ctx := context.Background()

	err = clientKV.SetWithTimeout(ctx, `1`, []byte{1, 2, 3}, 2000)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}

	l1, l2, err := clientKV.GetLen(ctx)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}
	t.Logf(`l1:%+v -l2:%+v`, l1, l2)

	time.Sleep(time.Second * 6)

	buf, err := clientKV.Get(ctx, `1`)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}
	t.Logf(`buf:%+v`, buf)
}

func TestStringKV3(t *testing.T) {
	clientKV := client.NewStringAsKeyClient(`localhost:10001`, 6000)
	err := clientKV.Dial(time.Second * 3)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}

	ctx := context.Background()

	err = clientKV.Set(ctx, `1`, nil)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}

	l1, l2, err := clientKV.GetLen(ctx)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}
	t.Logf(`l1:%+v -l2:%+v`, l1, l2)

	buf, err := clientKV.Get(ctx, `1`)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}
	t.Logf(`buf:%+v`, buf)
}

func TestStringKV4(t *testing.T) {
	clientKV := client.NewStringAsKeyClient(`localhost:10001`, 6000)
	err := clientKV.Dial(time.Second * 3)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}

	ctx := context.Background()

	err = clientKV.Set(ctx, `1`, make([]byte, 0))
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}

	l1, l2, err := clientKV.GetLen(ctx)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}
	t.Logf(`l1:%+v -l2:%+v`, l1, l2)

	buf, err := clientKV.Get(ctx, `1`)
	if err != nil {
		t.Errorf(`err:%+v`, err)
		t.Fail()
	}
	t.Logf(`buf:%+v`, buf)
}
