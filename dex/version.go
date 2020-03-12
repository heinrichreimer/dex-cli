package dex

import (
	"fmt"
	"github.com/dexidp/dex/api"
	"golang.org/x/net/context"
)

func (dex Dex) GetVersion() error {
	req := &api.VersionReq{}
	res, err := dex.client.GetVersion(context.TODO(), req)
	if err != nil {
		return err
	}
	var serverVersion string
	if len(res.Server) != 0 {
		serverVersion = res.Server
	} else {
		serverVersion = "Unknown"
	}
	fmt.Printf("%s (%d)\n", serverVersion, res.Api)
	return nil
}
