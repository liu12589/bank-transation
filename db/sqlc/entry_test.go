package db

import (
	"bank-transation/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateTestEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.GenerateMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)
	return entry
}

func TestGetEntry(t *testing.T) {
	var account Account
	account = CreateTestAccount(t)
	entry := CreateTestEntry(t, account)

	entry1, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)

	require.Equal(t, entry.Amount, entry1.Amount)
	require.Equal(t, entry.AccountID, entry1.AccountID)
	require.Equal(t, entry.CreatedAt, entry1.CreatedAt)
	require.Equal(t, entry.ID, entry1.ID)
}

func TestListEntries(t *testing.T) {
	var account Account
	account = CreateTestAccount(t)
	for i := 0; i < 10; i++ {
		CreateTestEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID)
		require.NotZero(t, entry.ID)
		require.NotZero(t, entry.Amount)
		require.NotZero(t, entry.CreatedAt)
	}
}
