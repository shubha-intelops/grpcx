package daos

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/daos/clients/sqls"
	"github.com/shubha-intelops/grpcx/profile/pkg/grpc/server/models"
	log "github.com/sirupsen/logrus"
)

type ProfileDataDao struct {
	sqlClient *sqls.MySQLClient
}

func migrateProfileData(r *sqls.MySQLClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS profileData(
		ID int NOT NULL AUTO_INCREMENT,
        
		Address VARCHAR(100) NOT NULL,
		Name VARCHAR(100) NOT NULL,
	    PRIMARY KEY (ID)
	);
	`
	_, err := r.DB.Exec(query)
	return err
}

func NewProfileDataDao() (*ProfileDataDao, error) {
	sqlClient, err := sqls.InitMySQLDB()
	if err != nil {
		return nil, err
	}
	err = migrateProfileData(sqlClient)
	if err != nil {
		return nil, err
	}
	return &ProfileDataDao{
		sqlClient,
	}, nil
}

func (profileDataDao *ProfileDataDao) CreateProfileData(m *models.ProfileData) (*models.ProfileData, error) {
	insertQuery := "INSERT INTO profileData(Address, Name) values(?, ?)"
	res, err := profileDataDao.sqlClient.DB.Exec(insertQuery, m.Address, m.Name)
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
	log.Debugf("profileData created")
	return m, nil
}

func (profileDataDao *ProfileDataDao) UpdateProfileData(id int64, m *models.ProfileData) (*models.ProfileData, error) {
	if id == 0 {
		return nil, errors.New("invalid profileData ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	profileData, err := profileDataDao.GetProfileData(id)
	if err != nil {
		return nil, err
	}
	if profileData == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE profileData SET Address = ?, Name = ? WHERE Id = ?"
	res, err := profileDataDao.sqlClient.DB.Exec(updateQuery, m.Address, m.Name, id)
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

	log.Debugf("profileData updated")
	return m, nil
}

func (profileDataDao *ProfileDataDao) DeleteProfileData(id int64) error {
	deleteQuery := "DELETE FROM profileData WHERE Id = ?"
	res, err := profileDataDao.sqlClient.DB.Exec(deleteQuery, id)
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

	log.Debugf("profileData deleted")
	return nil
}

func (profileDataDao *ProfileDataDao) ListProfileData() ([]*models.ProfileData, error) {
	selectQuery := "SELECT * FROM profileData"
	rows, err := profileDataDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var profileData []*models.ProfileData
	for rows.Next() {
		m := models.ProfileData{}
		if err = rows.Scan(&m.Id, &m.Address, &m.Name); err != nil {
			return nil, err
		}
		profileData = append(profileData, &m)
	}
	if profileData == nil {
		profileData = []*models.ProfileData{}
	}
	log.Debugf("profileData listed")
	return profileData, nil
}

func (profileDataDao *ProfileDataDao) GetProfileData(id int64) (*models.ProfileData, error) {
	selectQuery := "SELECT * FROM profileData WHERE Id = ?"
	row := profileDataDao.sqlClient.DB.QueryRow(selectQuery, id)

	m := models.ProfileData{}
	if err := row.Scan(&m.Id, &m.Address, &m.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}
	log.Debugf("profileData retrieved")
	return &m, nil
}
