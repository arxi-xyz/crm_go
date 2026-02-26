package redis

import (
	"context"
	"crm_go/pkg/appError"
	"errors"
)

func (c *Client) IsMember(
	ctx context.Context,
	key string,
	member string,
) (bool, *appError.AppError) {

	if key == "" || member == "" {
		return false, appError.Internal(errors.New("key is empty"))
	}

	res, err := c.Client.SIsMember(ctx, key, member).Result()

	if err != nil {
		return false, appError.Internal(err)
	}

	return res, nil
}
