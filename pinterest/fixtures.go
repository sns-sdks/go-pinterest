package pinterest

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

// BCSuite For the tests with app context
type BCSuite struct {
	suite.Suite
	Pin *Client
}

func (bc *BCSuite) SetupSuite() {
	bc.Pin = NewBearerClient("")
}

func (bc *BCSuite) SetupTest() {
	httpmock.ActivateNonDefault(bc.Pin.Cli.GetClient())
}

func (bc *BCSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}
