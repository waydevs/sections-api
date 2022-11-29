package repository

// Esta interfaz es de ejemplo, seguramente la interfaz que vamos a usar nos la provee un ORM
type dbConn interface {
	Get(key string) string
}

// Repository has the Data Base connection
type Repository struct {
	connExample dbConn
}

// NewRepository return a new instance of Repository
func NewRepository() (*Repository, error) {
	db := NewDb()
	return &Repository{
		connExample: db,
	}, nil
}

// Esto es un mock solo para este ejemplo
type db struct{}

func NewDb() *db {
	return &db{}
}

func (db *db) Get(key string) string {
	return "hello"
}
