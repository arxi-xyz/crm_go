package redis

import (
	"context"
)

func (c *Client) Exist(ctx context.Context, key string) (int64, error) {

	cmd := c.Client.Exists(ctx, key)

	if err := cmd.Err(); err != nil {
		return 0, err
	}

	return cmd.Val(), nil
}
