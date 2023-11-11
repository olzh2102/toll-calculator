package client

import (
	"context"

	"github.com/olzh2102/toll-calculator/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}
