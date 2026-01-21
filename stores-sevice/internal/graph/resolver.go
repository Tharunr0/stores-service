package graph

import (
	"context"
	"database/sql"

	"stores/internal/store"
)

type Resolver struct {
	DB *sql.DB
}

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) Stores(ctx context.Context) ([]*store.Store, error) {
	rows, err := r.DB.Query(
		`SELECT id, name, revenue FROM stores ORDER BY revenue DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*store.Store
	for rows.Next() {
		s := &store.Store{}
		if err := rows.Scan(&s.ID, &s.Name, &s.Revenue); err != nil {
			return nil, err
		}
		results = append(results, s)
	}
	return results, nil
}
