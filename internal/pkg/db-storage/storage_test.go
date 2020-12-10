package db_storage

import (
	"os"
	"testing"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"github.com/stretchr/testify/suite"
)

type TestStorageSuite struct {
	suite.Suite

	db *sqlx.DB

	storage *Storage
}

func (s *TestStorageSuite) SetupTest() {
	s.db = s.initDB()

	s.storage = &Storage{db: s.db}
}

func (s *TestStorageSuite) TearDownTest() {
	s.truncateTables()
	s.db.Close()
}

func (s *TestStorageSuite) initDB() *sqlx.DB {
	dsn := os.Getenv("DSN_TEST")
	if dsn == "" {
		s.T().Fatal("Нужен DSN_TEST")
	}

	gooseDB := sqlx.MustConnect("pgx", dsn)
	if err := goose.Up(gooseDB.DB, "../../../migrations"); err != nil {
		s.T().Fatal("не получилось накатить миграции")
	}

	if err := gooseDB.Close(); err != nil {
		s.T().Fatal("не удалось закрыть конекшен к БД")
	}

	driverConfig := &stdlib.DriverConfig{
		ConnConfig: pgx.ConnConfig{
			PreferSimpleProtocol: true,
		}}
	stdlib.RegisterDriverConfig(driverConfig)

	db := sqlx.MustConnect("pgx", driverConfig.ConnectionString(dsn))

	return db
}

func (s *TestStorageSuite) truncateTables() {
	s.db.MustExec(`TRUNCATE users, user_location CASCADE;`)
}

func TestStorageSuiteSuite(t *testing.T) {
	suite.Run(t, new(TestStorageSuite))
}
