package testsuite

import (
	"time"

	"github.com/stretchr/testify/suite"
	db "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

const TimeZone = "Canada/Eastern"

var TimeLocation, _ = time.LoadLocation(TimeZone)

type Helper interface {
	SQLBuilder() sqlbuilder.Database
	Session() db.Database

	Adapter() string

	TearUp() error
	TearDown() error
}

type Suite struct {
	suite.Suite

	Helper
}

func (s *Suite) AfterTest(suiteName, testName string) {
	err := s.TearDown()
	s.NoError(err)
}

func (s *Suite) BeforeTest(suiteName, testName string) {
	err := s.TearUp()
	s.NoError(err)
}
