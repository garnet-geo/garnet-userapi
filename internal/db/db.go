package db

import (
	"context"
	"fmt"

	"github.com/garnet-geo/garnet-userapi/internal/env"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

var conn *pgxpool.Pool

type logrusTracer struct {
	logger *log.Logger
}

func (l *logrusTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	l.logger.Debugln("Executing command", "SQL", data.SQL, "Args", data.Args)
	return ctx
}

func (l *logrusTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	// Skip...
}

func InitConnection() {
	config, err := pgxpool.ParseConfig(env.GetDatabaseUrl())
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	config.ConnConfig.Tracer = &logrusTracer{
		logger: log.StandardLogger(),
	}

	conn, err = pgxpool.NewWithConfig(Context(), config)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	log.Infoln("Connected to the database")
}

func CloseConnection() {
	conn.Close()

	log.Infoln("Closed connection to the database")
}

func ExecuteQuery(query string, args ...any) (pgx.Rows, error) {
	return conn.Query(Context(), query, args...)
}

func ExecuteQueryRow(query string, args ...any) pgx.Row {
	return conn.QueryRow(Context(), query, args...)
}

func ExecuteTransaction() (pgx.Tx, error) {
	log.Debugln("Executing transaction")
	return conn.Begin(Context())
}

func Context() context.Context {
	return context.Background()
}
