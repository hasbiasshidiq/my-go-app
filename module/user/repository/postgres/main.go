package repository

type postgresUserRepo struct {
	// db *sql.DB or *bun.DB or whatever
}

func NewPostgresUserRepo() *postgresUserRepo {
	return &postgresUserRepo{}
}
