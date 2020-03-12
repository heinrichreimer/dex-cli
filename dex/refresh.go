package dex

import (
	"errors"
	"fmt"
	"github.com/dexidp/dex/api"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/net/context"
	"os"
)

func (dex Dex) ListRefresh(
	userId string,
) error {
	req := &api.ListRefreshReq{
		UserId: userId,
	}
	res, err := dex.client.ListRefresh(context.TODO(), req)
	if err != nil {
		return err
	}
	if len(res.RefreshTokens) == 0 {
		fmt.Printf("User '%s' has no refresh tokens yet.\n", req.UserId)
		return nil
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Client ID", "Last used", "Created at"})
	table.SetBorder(false)
	for i := range res.RefreshTokens {
		token := res.RefreshTokens[i]
		table.Append([]string{token.Id, token.ClientId, string(token.LastUsed), string(token.CreatedAt)})
	}
	table.Render()
	return nil
}

func (dex Dex) RevokeRefresh(
	userId string,
	clientId string,
) error {
	req := &api.RevokeRefreshReq{
		UserId:   userId,
		ClientId: clientId,
	}
	res, err := dex.client.RevokeRefresh(context.TODO(), req)
	if err != nil {
		return err
	}
	if res.NotFound {
		return errors.New(fmt.Sprintf(
			"Refresh token of user '%s' for client '%s' was not found.\n",
			req.UserId, req.ClientId,
		))
	}
	fmt.Printf("Successfully revoked refresh token of user '%s' for client '%s'.\n", req.UserId, req.ClientId)
	return nil
}
