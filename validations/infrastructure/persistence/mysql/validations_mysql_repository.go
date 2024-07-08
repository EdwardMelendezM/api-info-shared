package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	errDomain "github.com/EdwardMelendezM/api-info-shared/error-log"

	"github.com/EdwardMelendezM/api-info-shared/validations/domain"
)

type validationsMySQLRepo struct {
	client  *sql.DB
	timeout time.Duration
}

func NewValidationsRepository(db *sql.DB, mongoTimeout int) domain.ValidationRepository {
	rep := &validationsMySQLRepo{
		client:  db,
		timeout: time.Duration(mongoTimeout) * time.Second,
	}
	return rep
}

func (r validationsMySQLRepo) RecordExists(
	ctx context.Context,
	params domain.ExistsParams,
) (bool, error) {
	var exists int
	var query string
	var args []interface{}

	query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", params.Table, params.IdColumnName)
	args = append(args, params.IdValue)

	if params.StatusColumnName != nil && params.StatusValue != nil {
		query += fmt.Sprintf(" AND %s = ?", *params.StatusColumnName)
		args = append(args, *params.StatusValue)
	}

	if params.StatusColumnName != nil && params.StatusValue == nil {
		query += fmt.Sprintf(" AND %s is null", *params.StatusColumnName)
	}

	err := r.client.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil {
		return false, errDomain.NewErr().
			SetRaw(err).
			SetLayer(errDomain.Infra).
			SetFunction("RecordExists")
	}
	if exists == 0 {
		return false, nil
	}

	return true, nil
}

func (r validationsMySQLRepo) ValidateExistence(
	ctx context.Context,
	params domain.ExistsParams,
) (bool, error) {
	var exists int
	var query string
	var args []interface{}

	query = fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = ?", params.Table, params.IdColumnName)
	args = append(args, params.IdValue)

	if params.StatusColumnName != nil && params.StatusValue != nil {
		query += fmt.Sprintf(" AND %s = ?", *params.StatusColumnName)
		args = append(args, *params.StatusValue)
	}

	if params.StatusColumnName != nil && params.StatusValue == nil {
		query += fmt.Sprintf(" AND %s is null", *params.StatusColumnName)
	}

	err := r.client.QueryRowContext(ctx, query, args...).Scan(&exists)
	if err != nil {
		return false, errDomain.NewErr().
			SetRaw(err).
			SetLayer(errDomain.Infra).
			SetFunction("ValidateExistence")
	}
	if exists != 0 {
		return true, nil
	}
	return false, nil
}
