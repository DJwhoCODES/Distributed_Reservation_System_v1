package postgres

import (
	"context"
	"fmt"

	"github.com/djwhocodes/ticket-reservation/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *configs.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&prefer_simple_protocol=true",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	fmt.Println(dsn)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	return pgxpool.NewWithConfig(context.Background(), config)
}

// dsn = data source name
// fmt.Sprintf = Sprintf formats and returns a string.
// pgxpool = It's the PostgreSQL connection pool from the pgx library. It maintains reusable connections. Connections are borrowed and returned to the pool instead of recreated every time.
// context.Background() = Every database operation in Go takes a context.Context. context.Background() itself never times out—it's simply the root context.
