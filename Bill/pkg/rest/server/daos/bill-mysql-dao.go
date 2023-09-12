package daos

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/shubha-intelops/grpcx/bill/pkg/rest/server/daos/clients/sqls"
	"github.com/shubha-intelops/grpcx/bill/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type BillDao struct {
	sqlClient *sqls.MySQLClient
}

func migrateBills(r *sqls.MySQLClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS bills(
		ID int NOT NULL AUTO_INCREMENT,
        
		Ammount INT NOT NULL,
		Name INT NOT NULL,
	    PRIMARY KEY (ID)
	);
	`
	_, err := r.DB.Exec(query)
	return err
}

func NewBillDao() (*BillDao, error) {
	sqlClient, err := sqls.InitMySQLDB()
	if err != nil {
		return nil, err
	}
	err = migrateBills(sqlClient)
	if err != nil {
		return nil, err
	}
	return &BillDao{
		sqlClient,
	}, nil
}

func (billDao *BillDao) CreateBill(m *models.Bill) (*models.Bill, error) {
	insertQuery := "INSERT INTO bills(Ammount, Name) values(?, ?)"
	res, err := billDao.sqlClient.DB.Exec(insertQuery, m.Ammount, m.Name)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1062 {
				return nil, sqls.ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id
	log.Debugf("bill created")
	return m, nil
}

func (billDao *BillDao) UpdateBill(id int64, m *models.Bill) (*models.Bill, error) {
	if id == 0 {
		return nil, errors.New("invalid bill ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	bill, err := billDao.GetBill(id)
	if err != nil {
		return nil, err
	}
	if bill == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE bills SET Ammount = ?, Name = ? WHERE Id = ?"
	res, err := billDao.sqlClient.DB.Exec(updateQuery, m.Ammount, m.Name, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("bill updated")
	return m, nil
}

func (billDao *BillDao) DeleteBill(id int64) error {
	deleteQuery := "DELETE FROM bills WHERE Id = ?"
	res, err := billDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("bill deleted")
	return nil
}

func (billDao *BillDao) ListBills() ([]*models.Bill, error) {
	selectQuery := "SELECT * FROM bills"
	rows, err := billDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var bills []*models.Bill
	for rows.Next() {
		m := models.Bill{}
		if err = rows.Scan(&m.Id, &m.Ammount, &m.Name); err != nil {
			return nil, err
		}
		bills = append(bills, &m)
	}
	if bills == nil {
		bills = []*models.Bill{}
	}
	log.Debugf("bill listed")
	return bills, nil
}

func (billDao *BillDao) GetBill(id int64) (*models.Bill, error) {
	selectQuery := "SELECT * FROM bills WHERE Id = ?"
	row := billDao.sqlClient.DB.QueryRow(selectQuery, id)

	m := models.Bill{}
	if err := row.Scan(&m.Id, &m.Ammount, &m.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}
	log.Debugf("bill retrieved")
	return &m, nil
}
