package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/ullas-jain/simplebank/util"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)
	queriedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, queriedAccount)

	require.Equal(t, createdAccount.ID, queriedAccount.ID)
	require.Equal(t, createdAccount.Owner, queriedAccount.Owner)
	require.Equal(t, createdAccount.Balance, queriedAccount.Balance)
	require.Equal(t, createdAccount.Currency, queriedAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, queriedAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	createdAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: util.RandomMoney(),
	}

	queriedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, queriedAccount)

	require.Equal(t, createdAccount.ID, queriedAccount.ID)
	require.Equal(t, createdAccount.Owner, queriedAccount.Owner)
	require.Equal(t, arg.Balance, queriedAccount.Balance)
	require.Equal(t, createdAccount.Currency, queriedAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, queriedAccount.CreatedAt, time.Second)
}