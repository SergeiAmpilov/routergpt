package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"routergpt/internal/domain/models/model"

	_ "github.com/lib/pq"
)

type postgresqlRepository struct {
	db *sql.DB
}

func NewPostgreSQLRepository(connString string) (ModelsRepository, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	repo := &postgresqlRepository{db: db}
	
	// Run migration to create table if not exists
	if err := repo.migrate(); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return repo, nil
}

func (r *postgresqlRepository) migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS ai_models (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL UNIQUE,
		provider VARCHAR(100) NOT NULL,
		version VARCHAR(50) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := r.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create ai_models table: %w", err)
	}

	log.Println("ai_models table created or already exists")
	return nil
}

func (r *postgresqlRepository) Create(ctx context.Context, aiModel *model.AIModel) error {
	query := `
	INSERT INTO ai_models (name, provider, version) 
	VALUES ($1, $2, $3)
	ON CONFLICT (name) DO NOTHING;
	`

	_, err := r.db.ExecContext(ctx, query, aiModel.Name, aiModel.Provider, aiModel.Version)
	if err != nil {
		return fmt.Errorf("failed to insert ai_model: %w", err)
	}

	return nil
}

func (r *postgresqlRepository) GetByName(ctx context.Context, name string) (*model.AIModel, error) {
	query := `SELECT name, provider, version FROM ai_models WHERE name = $1 LIMIT 1;`
	
	var aiModel model.AIModel
	err := r.db.QueryRowContext(ctx, query, name).Scan(&aiModel.Name, &aiModel.Provider, &aiModel.Version)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ai model with name %s not found", name)
		}
		return nil, fmt.Errorf("failed to query ai_model: %w", err)
	}

	return &aiModel, nil
}