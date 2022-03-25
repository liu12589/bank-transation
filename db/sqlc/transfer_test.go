package db

import (
	"bank-transation/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateTestTransfer(t *testing.T, fromAccount Account, toAccount Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.GenerateMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, fromAccount.ID, transfer.FromAccountID)
	require.Equal(t, toAccount.ID, transfer.ToAccountID)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestGetTransfer(t *testing.T) {
	fromAccount := CreateTestAccount(t)
	toAccount := CreateTestAccount(t)
	transfer := CreateTestTransfer(t, fromAccount, toAccount)
	transfer1, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer1)

	require.Equal(t, transfer.ID, transfer1.ID)
	require.Equal(t, transfer.Amount, transfer1.Amount)
	require.Equal(t, transfer.CreatedAt, transfer1.CreatedAt)
	require.Equal(t, transfer.ToAccountID, transfer1.ToAccountID)
	require.Equal(t, transfer.FromAccountID, transfer1.FromAccountID)
}

func TestListTransfers(t *testing.T) {
	fromAccount := CreateTestAccount(t)
	toAccount := CreateTestAccount(t)
	for i := 0; i < 10; i++ {
		CreateTestTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
		require.NotZero(t, transfer.Amount)
	}
}
