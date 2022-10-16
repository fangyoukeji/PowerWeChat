package shop

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/shop/basic"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/shop/register"
)

type Client struct {
	*kernel.BaseClient

	Register *register.Client
	Basic    *basic.Client
}
