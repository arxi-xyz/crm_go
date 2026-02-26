package redis

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func (c *Client) SetList(ctx context.Context, key string, value interface{}, expireTime time.Time) error {
	rv := reflect.ValueOf(value)
	if !rv.IsValid() {
		return fmt.Errorf("value is invalid")
	}

	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return fmt.Errorf("value is nil")
		}
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return fmt.Errorf("SetList expects slice/array, got %s", rv.Kind())
	}

	items := make([]interface{}, 0, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		items = append(items, rv.Index(i).Interface())
	}

	pipe := c.Client.TxPipeline()

	pipe.Del(ctx, key)

	if len(items) > 0 {
		pipe.RPush(ctx, key, items...)
		pipe.ExpireAt(ctx, key, expireTime)
	}

	_, err := pipe.Exec(ctx)
	return err
}
