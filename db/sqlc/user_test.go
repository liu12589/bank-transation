package db

import (
	"bank-transaction/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createTestUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.GenerateString(10))
	arg := CreateUserParams{
		Username:       util.GenerateOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.GenerateOwner(),
		Email:          util.GenerateEmail(10),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createTestUser(t)
}

func TestGetUser(t *testing.T) {
	user := createTestUser(t)

	user1, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)

	require.NotEmpty(t, user1)
	require.Equal(t, user.Username, user1.Username)
	require.Equal(t, user.CreatedAt, user1.CreatedAt)
	require.Equal(t, user.Email, user1.Email)
	require.Equal(t, user.FullName, user1.FullName)
	require.Equal(t, user.HashedPassword, user1.HashedPassword)
}
