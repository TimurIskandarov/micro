package db

import (
	"context"
	"fmt"

	"user/config"

	"go.uber.org/zap"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewSqlDB(ctx context.Context, dbConf config.DB, logger *zap.Logger) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name,
	)

	pgCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Fatal("failed to get db config: ", zap.Error(err))
	}
	
	dbc, err := pgxpool.ConnectConfig(ctx, pgCfg)
	if err != nil {
		logger.Fatal("failed to get db connection: ", zap.Error(err))
	}

	err = dbc.Ping(ctx)
	if err != nil {
		logger.Fatal("ping error: ", zap.Error(err))
	}
	
	return dbc
}
