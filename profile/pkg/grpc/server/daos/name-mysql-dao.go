package daos

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/daos/clients/sqls"
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/models"
	log "github.com/sirupsen/logrus"
)

type NameDao struct {
	sqlClient *sqls.MySQLClient
}

func migrateNames(r *sqls.MySQLClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS names(
		ID int NOT NULL AUTO_INCREMENT,
        
		Address VARCHAR(100) NOT NULL,
		Name VARCHAR(100) NOT NULL,
		Age INT NOT NULL,
		Email_address VARCHAR(100) NOT NULL,
	    PRIMARY KEY (ID)
	);
	`
	_, err := r.DB.Exec(query)
	return err
}

func NewNameDao() (*NameDao, error) {
	sqlClient, err := sqls.InitMySQLDB()
	if err != nil {
		return nil, err
	}
	err = migrateNames(sqlClient)
	if err != nil {
		return nil, err
	}
	return &NameDao{
		sqlClient,
	}, nil
}

func (nameDao *NameDao) CreateName(m *models.Name) (*models.Name, error) {
	insertQuery := "INSERT INTO names(Address, Name, Age, Email_address) values(?, ?, ?, ?)"
	res, err := nameDao.sqlClient.DB.Exec(insertQuery, m.Address, m.Name, m.Age, m.Email_address)
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
	log.Debugf("name created")
	return m, nil
}

func (nameDao *NameDao) UpdateName(id int64, m *models.Name) (*models.Name, error) {
	if id == 0 {
		return nil, errors.New("invalid name ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	name, err := nameDao.GetName(id)
	if err != nil {
		return nil, err
	}
	if name == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE names SET Address = ?, Name = ?, Age = ?, Email_address = ? WHERE Id = ?"
	res, err := nameDao.sqlClient.DB.Exec(updateQuery, m.Address, m.Name, m.Age, m.Email_address, id)
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

	log.Debugf("name updated")
	return m, nil
}

func (nameDao *NameDao) DeleteName(id int64) error {
	deleteQuery := "DELETE FROM names WHERE Id = ?"
	res, err := nameDao.sqlClient.DB.Exec(deleteQuery, id)
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

	log.Debugf("name deleted")
	return nil
}

func (nameDao *NameDao) ListNames() ([]*models.Name, error) {
	selectQuery := "SELECT * FROM names"
	rows, err := nameDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var names []*models.Name
	for rows.Next() {
		m := models.Name{}
		if err = rows.Scan(&m.Id, &m.Address, &m.Name, &m.Age, &m.Email_address); err != nil {
			return nil, err
		}
		names = append(names, &m)
	}
	if names == nil {
		names = []*models.Name{}
	}
	log.Debugf("name listed")
	return names, nil
}

func (nameDao *NameDao) GetName(id int64) (*models.Name, error) {
	selectQuery := "SELECT * FROM names WHERE Id = ?"
	row := nameDao.sqlClient.DB.QueryRow(selectQuery, id)

	m := models.Name{}
	if err := row.Scan(&m.Id, &m.Address, &m.Name, &m.Age, &m.Email_address); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}
	log.Debugf("name retrieved")
	return &m, nil
}
