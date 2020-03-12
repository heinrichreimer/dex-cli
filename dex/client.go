package dex

import (
	"errors"
	"fmt"
	"github.com/dexidp/dex/api"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/net/context"
)

func (dex Dex) CreateClient(
	id string,
	redirectUris []string,
	trustedPeers []string,
	public bool,
	name string,
	logoUrl string,
) error {
	secret, err := password.Generate(64, 10, 0, false, false)
	if err != nil {
		return err
	}

	req := &api.CreateClientReq{
		Client: &api.Client{
			Id:           id,
			Secret:       secret,
			RedirectUris: redirectUris,
			TrustedPeers: trustedPeers,
			Public:       public,
			Name:         name,
			LogoUrl:      logoUrl,
		},
	}
	res, err := dex.client.CreateClient(context.TODO(), req)
	if err != nil {
		return err
	}
	if res.AlreadyExists {
		return errors.New(fmt.Sprintf("Client '%s' already exists.\n", req.Client.Id))
	}
	fmt.Printf("Successfully created client '%s'.", res.Client.Id)
	fmt.Print("Secret:")
	fmt.Print(res.Client.Secret)
	return nil
}

func (dex Dex) DeleteClient(id string) error {
	req := &api.DeleteClientReq{
		Id: id,
	}
	res, err := dex.client.DeleteClient(context.TODO(), req)
	if err != nil {
		return err
	}
	if res.NotFound {
		return errors.New(fmt.Sprintf("Client '%s' does not exist.\n", req.Id))
	}
	fmt.Printf("Successfully deleted client '%s'.\n", req.Id)
	return nil
}
