package shop

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/shop/basic"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/shop/register"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseClient: baseClient,
		Register:   &register.Client{},
		Basic:      &basic.Client{},
	}, nil
}
