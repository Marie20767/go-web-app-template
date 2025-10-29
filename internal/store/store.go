package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Marie20767/go-web-app-template/internal/store/sqlc"
	"github.com/Marie20767/go-web-app-template/internal/utils/config"
)

type Store struct {
	conn    *sql.DB
	Queries *sqlc.Queries
}

func connectDB(ctx context.Context, cfg *config.Config) (*sql.DB, error) {
	dbConn, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}

	timeOut := time.Duration(cfg.DbTimeout) * time.Second
	dbCtx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()
	err = dbConn.PingContext(dbCtx)
	if err != nil {
		cErr := dbConn.Close()
		if cErr != nil {
			return nil, fmt.Errorf("failed to ping DB: %v; also failed to close DB: %v", err, cErr)
		}
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return dbConn, nil
}

func NewStore(ctx context.Context, cfg *config.Config) (*Store, error) {
	dbConn, err := connectDB(ctx, cfg)

	if err != nil {
		return nil, err
	}

	return &Store{
		conn:    dbConn,
		Queries: sqlc.New(dbConn),
	}, nil
}

func (s *Store) Close() error {
	return s.conn.Close()
}
