package routes

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestWrapHttpHandlers_GetProducts(t *testing.T) {
	//	db, mock := NewMock()
	//	defer db.Close()
	//	wh := WrapHttpHandlers{&models.Storage{Db: db}}
	//
	//	r := httptest.NewRequest("GET", "http://127.0.0.1/products", nil)
	//	w := httptest.NewRecorder()
	//
	//	rowsProds := sqlmock.NewRows([]string{"p.id", "p.name", "p.item_number", "p.manufacturer_id", "m.name As manufacturer_name"})
	//	mock.ExpectQuery("^SELECT (.+) FROM products p*").
	//		WillReturnRows(rowsProds)
	//
	//	rowsCount := sqlmock.NewRows([]string{"count"})
	//	rowsCount.AddRow(0)
	//	mock.ExpectQuery("^SELECT COUNT(.+) as count FROM*").
	//		WillReturnRows(rowsCount)
	//
	//	wh.GetProducts(w, r)
	//	// TODO: check
}

func TestWrapHttpHandlers_CreateProduct(t *testing.T) {
	//db, _ := NewMock()
	//defer db.Close()
	//wh := WrapHttpHandlers{&models.Storage{Db: db}}
	//
	//bodyReq := bytes.NewReader([]byte(""))
	//r := httptest.NewRequest("POST", "http://127.0.0.1/products", bodyReq)
	//w := httptest.NewRecorder()
	//
	////rowsProds := sqlmock.NewRows([]string{"p.id", "p.name", "p.item_number", "p.manufacturer_id", "m.name As manufacturer_name"})
	////mock.ExpectQuery("^SELECT (.+) FROM products p*").
	////	WillReturnRows(rowsProds)
	////
	////rowsCount := sqlmock.NewRows([]string{"count"})
	////rowsCount.AddRow(0)
	////mock.ExpectQuery("^SELECT COUNT(.+) as count FROM*").
	////	WillReturnRows(rowsCount)
	//
	//wh.GetProducts(w, r)
	//// TODO: check
}
