package redis

import (
	"context"
	"time"
)

func (c *Client) Set(ctx context.Context, key string, value interface{}, expireTime time.Time) error {

	pipe := c.Client.TxPipeline()

	pipe.HSet(ctx, key, value)
	pipe.ExpireAt(ctx, key, expireTime)

	_, err := pipe.Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
