package accounts

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProfileUpdaterTestSuite struct {
	suite.Suite
}

func TestProfileUpdater(t *testing.T) {
	suite.Run(t, new(ProfileUpdaterTestSuite))
}
