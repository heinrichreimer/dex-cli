package dex

import (
	"github.com/dexidp/dex/api"
	"google.golang.org/grpc"
)

func UseDex(target string, use func(dex Dex) error) error {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := api.NewDexClient(conn)
	return use(Dex{client: client})
}

type Dex struct {
	client api.DexClient
}
