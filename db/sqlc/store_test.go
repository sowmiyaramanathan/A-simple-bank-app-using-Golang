package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTranferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TranferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TranferTx(context.Background(), TransferParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			// receving errors and results in the channel
			errs <- err
			results <- result
		}()
	}
	// checking results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		tranfer := result.Transfer
		require.NotEmpty(t, tranfer)
		require.NotZero(t, tranfer.ID)
		require.Equal(t, account1.ID, tranfer.FromAccountID)
		require.Equal(t, account2.ID, tranfer.ToAccountID)
		require.Equal(t, amount, tranfer.Amount)
		require.NotZero(t, tranfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), tranfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.NotZero(t, fromEntry.ID)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.NotZero(t, toEntry.ID)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)
	}
}
