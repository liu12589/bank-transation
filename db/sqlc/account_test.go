package db

import (
	"bank-transation/util"
	"context"
	"database/sql"
	_ "database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateTestAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.GenerateOwner(),
		Balance:  util.GenerateMoney(),
		Currency: util.GenerateCurrency(),
	}

	// 这里 CreateAccount 实际上调用了 QueryRowContext 方法。
	// 而 QueryRowContext 我没找到文档说明，实际测试中发现该方法，可以先执行语句再返回查询。
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

func TestGetAccount(t *testing.T) {
	account1 := CreateTestAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	account1 := CreateTestAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.GenerateMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account1 := CreateTestAccount(t)
	err1 := testQueries.DeleteAccount(context.Background(), account1.ID)
	account2, err2 := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err1)
	require.Error(t, err2)

	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateTestAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
