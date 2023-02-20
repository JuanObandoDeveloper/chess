package db

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func New(lc fx.Lifecycle, l *zap.Logger) (*PrismaClient, error) {
	prisma := NewClient()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			l.Debug("Connecting to database")
			return prisma.Connect()
		},
		OnStop: func(context.Context) error {
			l.Debug("Disconnecting from database")
			return prisma.Disconnect()
		},
	})

	return prisma, nil
}

var Module = fx.Options(
	fx.Provide(New),
)
