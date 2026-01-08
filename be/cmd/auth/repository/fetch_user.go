package authrepository

import (
	"context"
	authdomain "expense-management-system/cmd/auth/domain"
	authquery "expense-management-system/cmd/auth/repository/query"
	"fmt"
	"strings"
)

func (r *AuthrepositoryImpl) FetchUser(c context.Context, q authquery.FetchUser) (authdomain.User, error) {
	var data authdomain.User

	sql := `select id, email, name, role, password, created_at from users`

	sql, values := r.queryFetchUser(sql, q)

	err := r.DB.Conn.QueryRow(c, sql, values...).Scan(
		&data.ID,
		&data.Email,
		&data.Name,
		&data.Role,
		&data.Password,
		&data.CreatedAt,
	)
	return data, err
}

func (r *AuthrepositoryImpl) queryFetchUser(sql string, q authquery.FetchUser) (string, []any) {
	var (
		keys   []string
		values []any
	)

	if len(q.Email) > 0 {
		keys = append(keys, fmt.Sprintf("email = $%d", len(values)+1))
		values = append(values, q.Email)
	}

	if len(keys) > 0 {
		sql += " where " + strings.Join(keys, " and ")
	}

	return sql, values
}
