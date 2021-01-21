package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetCustomer(t *testing.T) {
	//r := mux.NewRouter()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Error occured, while opening db stub connection")
	}

	defer db.Close()

	// Mock columns: []string{"ID", "Name"}

	// Expect Transaction begin.
	mock.ExpectBegin()
	// Query expected, no arguments in expected query.
	myMockRows := sqlmock.NewRows([]string{"ID", "Name"}).AddRow("1", "Ishan Pandey")
	myMockRows2 := sqlmock.NewRows([]string{"ID", "Name"}).AddRow("2", "Varun Singh")
	mock.ExpectQuery("SELECT * FROM Customers").WithArgs().WillReturnRows(myMockRows, myMockRows2)
	// Expect a transaction commit/over in database.
	mock.ExpectCommit()

	// Run the function in dev file.
	err = GetCustomer(db)
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
	}

	// Make sure all expectations are met.
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations, FAILED")
	}

}
