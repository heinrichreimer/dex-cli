package dex

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/dexidp/dex/api"
	"github.com/google/uuid"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/net/context"
	"os"
)

func readPassword(
	text string,
	file string,
	stdin bool,
	message string,
) ([]byte, error) {
	var password []byte
	var err error
	if len(text) != 0 {
		return []byte(text), nil
	} else if len(file) != 0 {
		inFile, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer inFile.Close()

		sc := bufio.NewScanner(inFile)
		if !sc.Scan() {
			err := sc.Err()
			if err != nil {
				return nil, err
			}
			return nil, errors.New(fmt.Sprintf("File '%s' doesn't contain anything.\n", file))
		}
		return []byte(sc.Text()), nil
	} else if stdin {
		sc := bufio.NewScanner(os.Stdin)
		if !sc.Scan() {
			err := sc.Err()
			if err != nil {
				return nil, err
			}
			return nil, errors.New(fmt.Sprintf("Standard input stream '%s' doesn't contain anything.\n", file))
		}
		return []byte(sc.Text()), nil
	} else {
		fmt.Print(message)
		password, err = terminal.ReadPassword(int(os.Stdin.Fd()))
		fmt.Print("\n")
		return password, err
	}
}

func (dex Dex) CreatePassword(
	email string,
	username string,
	passwordText string,
	passwordFile string,
	passwordStdin bool,
) error {
	userId, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	password, err := readPassword(passwordText, passwordFile, passwordStdin, fmt.Sprintf("Enter password for user '%s': ", email))
	if err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		return err
	}

	req := &api.CreatePasswordReq{
		Password: &api.Password{
			Email:    email,
			Hash:     hash,
			Username: username,
			UserId:   userId.String(),
		},
	}
	res, err := dex.client.CreatePassword(context.TODO(), req)
	if err != nil {
		return err
	}
	if res.AlreadyExists {
		return errors.New(fmt.Sprintf("User '%s' already exists.\n", req.Password.Email))
	}
	fmt.Printf("Successfully created user %s.\n", req.Password.Email)
	return nil
}

func (dex Dex) UpdatePassword(
	email string,
	username string,
	passwordText string,
	passwordFile string,
	passwordStdin bool,
) error {
	password, err := readPassword(passwordText, passwordFile, passwordStdin, fmt.Sprintf("Enter new password for user '%s': ", email))
	if err != nil {
		return err
	}
	var hash []byte
	if len(password) == 0 {
		hash = nil
	} else {
		hash, err = bcrypt.GenerateFromPassword(password, 10)
		if err != nil {
			return err
		}
	}

	req := &api.UpdatePasswordReq{
		Email:       email,
		NewHash:     hash,
		NewUsername: username,
	}
	res, err := dex.client.UpdatePassword(context.TODO(), req)
	if err != nil {
		return err
	}
	if res.NotFound {
		return errors.New(fmt.Sprintf("User '%s' does not exist.\n", req.Email))
	}
	fmt.Printf("Successfully updated user '%s'.\n", req.Email)
	return nil
}

func (dex Dex) DeletePassword(
	email string,
) error {
	req := &api.DeletePasswordReq{
		Email: email,
	}
	res, err := dex.client.DeletePassword(context.TODO(), req)
	if err != nil {
		return err
	}
	if res.NotFound {
		return errors.New(fmt.Sprintf("User '%s' does not exist.\n", req.Email))
	}
	fmt.Printf("Successfully deleted user '%s'.\n", req.Email)
	return nil
}

func (dex Dex) ListPasswords() error {
	req := &api.ListPasswordReq{}
	res, err := dex.client.ListPasswords(context.TODO(), req)
	if err != nil {
		return err
	}
	if len(res.Passwords) == 0 {
		fmt.Print("No users found.\n")
		return nil
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Email", "Username"})
	table.SetBorder(false)
	for i := range res.Passwords {
		password := res.Passwords[i]
		table.Append([]string{password.UserId, password.Email, password.Username})
	}
	table.Render()
	return nil
}
