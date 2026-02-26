package redis

import (
	"context"
	"crm_go/pkg/appError"
	"errors"
	"time"
)

func (c *Client) SetSet(
	ctx context.Context,
	key string,
	members []interface{},
	expireAt time.Time,
) *appError.AppError {

	if key == "" {
		return appError.Internal(errors.New("key is empty"))
	}

	pipe := c.Client.TxPipeline()

	pipe.Del(ctx, key)

	if len(members) > 0 {
		values := make([]interface{}, 0, len(members))
		for _, m := range members {
			values = append(values, m)
		}

		pipe.SAdd(ctx, key, values...)
		pipe.ExpireAt(ctx, key, expireAt)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return appError.Internal(err)
	}

	return nil
}
