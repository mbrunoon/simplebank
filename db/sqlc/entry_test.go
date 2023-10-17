package db

import (
	"context"
	"database/sql"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func CreateRandomEntry(t *testing.T) Entry {

	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandonMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, entry)

	assert.Equal(t, entry.Amount, arg.Amount)
	assert.Equal(t, entry.AccountID, arg.AccountID)

	return entry
}

func TestGetEntry(t *testing.T) {

	entry1 := CreateRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, entry2)

	assert.Equal(t, entry1.Amount, entry2.Amount)
	assert.Equal(t, entry1.AccountID, entry2.AccountID)
}

func TestUpdateEntry(t *testing.T) {

	entry1 := CreateRandomEntry(t)

	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandonMoney(),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, entry2)

	assert.Equal(t, entry1.AccountID, entry2.AccountID)
	assert.Equal(t, entry1.ID, entry2.ID)
	assert.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)
}

func TestDeleteEntry(t *testing.T) {

	entry1 := CreateRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	assert.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, entry2)
}

func TestListEntries(t *testing.T) {

	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	assert.NoError(t, err)
	assert.Len(t, entries, 5)

	for _, entry := range entries {
		assert.NotEmpty(t, entry)
	}
}
