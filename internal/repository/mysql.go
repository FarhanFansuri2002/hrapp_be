package repository

import (
	"context"
	"database/sql"

	"server/internal/model"
)

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore(db *sql.DB) *MySQLStore {
	return &MySQLStore{db: db}
}

func (store *MySQLStore) ListEmployees(ctx context.Context) ([]model.Employee, error) {
	rows, err := store.db.QueryContext(ctx, `
		SELECT id, name, role, department, status, joined_at
		FROM employees
		ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := make([]model.Employee, 0)
	for rows.Next() {
		var employee model.Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Role, &employee.Department, &employee.Status, &employee.JoinedAt); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}
