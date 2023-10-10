package rdb_test

import (
	"context"
	"go-layered-architecture-sample/internal/config"
	"go-layered-architecture-sample/internal/gateway/rdb"
	"go-layered-architecture-sample/pkg/logger"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/suite"
)

type RDBTestSuite struct {
	repository *rdb.Repository
	suite.Suite
}

func (s *RDBTestSuite) SetupSuite() {
	logger := logger.New()
	conf := getTestDBConfig()
	primaryDB, _, err := rdb.NewRepository(conf, logger)
	s.Require().NoError(err)

	s.repository = primaryDB
}

func getTestDBConfig() *config.Database {
	return &config.Database{
		PrimaryHost:  "localhost",
		PrimaryPort:  3306,
		ReplicaHost:  "localhost",
		ReplicaPort:  3306,
		DatabaseName: "go-layered-architecture-sample",
		User:         "go-layered-architecture-sample",
		Password:     "P@ssw0rd",
	}
}

func newContext() context.Context {
	return context.Background()
}

func ignoreFieldsFor(obj interface{}, fields ...string) cmp.Option {
	return cmpopts.IgnoreFields(obj, fields...)
}
