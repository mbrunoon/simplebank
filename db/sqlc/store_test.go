package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testTrasferTx(t *testing.T) {

	store := NewStore(testDB)

	account1 := createRandonAccount(t)
	account2 := createRandonAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()

		for i := 0; i < n; i++ {
			err := <-errs
			assert.NoError(t, err)

			result := <-results
			assert.NotEmpty(t, result)

			transfer := result.Transfer
			assert.NotEmpty(t, transfer)
			assert.Equal(t, account1.ID, transfer.FromAccountID)
			assert.Equal(t, account2.ID, transfer.ToAccountID)
			assert.Equal(t, amount, transfer.Amount)
			assert.NotZero(t, transfer.ID)
			assert.NotZero(t, transfer.CreatedAt)

			_, err = store.GetTransfer(context.Background(), transfer.ID)
			assert.NoError(t, err)

			fromEntry := result.FromEntry
			assert.NotEmpty(t, fromEntry.ID)
			assert.Equal(t, account1.ID, fromEntry.AccountID)
			assert.Equal(t, -amount, fromEntry.Amount)
			assert.NotZero(t, fromEntry.ID)
			assert.NotZero(t, fromEntry.CreatedAt)

			_, err = store.GetEntry(context.Background(), fromEntry.ID)
			assert.NoError(t, err)

			toEntry := result.ToEntry
			assert.NotEmpty(t, toEntry.ID)
			assert.Equal(t, account1.ID, toEntry.AccountID)
			assert.Equal(t, amount, toEntry.Amount)
			assert.NotZero(t, toEntry.ID)
			assert.NotZero(t, toEntry.CreatedAt)

			_, err = store.GetEntry(context.Background(), toEntry.ID)
			assert.NoError(t, err)

			fromAccount := result.FromAccount
			assert.NotEmpty(t, fromAccount)
			assert.Equal(t, fromAccount.ID, account1.ID)

			toAccount := result.ToAccount
			assert.NotEmpty(t, toAccount)
			assert.Equal(t, toAccount.ID, account2.ID)

			diff1 := fromAccount.Balance - account1.Balance
			diff2 := toAccount.Balance - account2.Balance
			assert.Equal(t, diff1, diff2)
			assert.True(t, diff1 > 0)
			assert.True(t, diff1%amount == 0)

			k := int(diff1 / amount)
			assert.True(t, k >= 1 && k <= n)
			assert.NotContains(t, existed, k)
			existed[k] = true
		}

	}
}
