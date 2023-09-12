package daos

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
    "github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/daos/clients/sqls"
	"github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/models"
    log "github.com/sirupsen/logrus"
)

type DataDao struct {
	sqlClient *sqls.MySQLClient
}

func migrateData(r *sqls.MySQLClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS data(
		ID int NOT NULL AUTO_INCREMENT,
        
		Ggg VARCHAR(100) NOT NULL,
	    PRIMARY KEY (ID)
	);
	`
	_, err := r.DB.Exec(query)
	return err
}

func NewDataDao() (*DataDao, error) {
	sqlClient, err := sqls.InitMySQLDB()
	if err != nil {
		return nil, err
	}
	err = migrateData(sqlClient)
    if err != nil {
        return nil, err
    }
	return &DataDao{
		sqlClient,
	}, nil
}

func (dataDao *DataDao) CreateData(m *models.Data) (*models.Data, error) {
    insertQuery := "INSERT INTO data(Ggg) values(?)"
    res, err := dataDao.sqlClient.DB.Exec(insertQuery, m.Ggg)
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
	log.Debugf("data created")
    return m, nil
}

func (dataDao *DataDao) UpdateData(id int64, m *models.Data) (*models.Data, error)  {
	if id == 0 {
        return nil, errors.New("invalid data ID")
    }
    if id != m.Id {
        return nil, errors.New("id and payload don't match")
    }

    data, err := dataDao.GetData(id)
    if err != nil {
        return nil, err
    }
    if data == nil {
        return nil, sql.ErrNoRows
    }

    updateQuery := "UPDATE data SET Ggg = ? WHERE Id = ?"
    res, err := dataDao.sqlClient.DB.Exec(updateQuery, m.Ggg, id)
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

	log.Debugf("data updated")
    return m, nil
}

func (dataDao *DataDao) DeleteData(id int64) error {
	deleteQuery := "DELETE FROM data WHERE Id = ?"
    res, err := dataDao.sqlClient.DB.Exec(deleteQuery, id)
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

	log.Debugf("data deleted")
	return nil
}

func (dataDao *DataDao) ListData() ([]*models.Data, error) {
	selectQuery := "SELECT * FROM data"
    rows, err := dataDao.sqlClient.DB.Query(selectQuery)
    if err != nil {
        return nil, err
    }
    defer func(rows *sql.Rows) {
        _ = rows.Close()
    }(rows)

    var data []*models.Data
    for rows.Next() {
        m := models.Data {
           
               
           
        }
        if err = rows.Scan(&m.Id, &m.Ggg); err != nil {
            return nil, err
        }
        data = append(data, &m)
    }
    if data == nil {
        data = []*models.Data{}
    }
	log.Debugf("data listed")
	return data, nil
}

func (dataDao *DataDao) GetData(id int64) (*models.Data, error) {
	selectQuery := "SELECT * FROM data WHERE Id = ?"
    row := dataDao.sqlClient.DB.QueryRow(selectQuery, id)

    m := models.Data {
       
           
       
    }
    if err := row.Scan(&m.Id, &m.Ggg); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, sqls.ErrNotExists
        }
        return nil, err
    }
	log.Debugf("data retrieved")
    return &m, nil
}
