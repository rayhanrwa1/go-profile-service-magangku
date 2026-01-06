package repository

import (
	"context"

	"go-profile-service-magangku/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepository struct {
	DB *pgxpool.Pool
}

func NewProfileRepository(db *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (r *ProfileRepository) GetByUserID(
	ctx context.Context,
	userID string,
) (*domain.Profile, error) {

	query := `
		SELECT user_id, full_name, phone_number, photo, city, country
		FROM user_profiles
		WHERE user_id = $1
	`

	row := r.DB.QueryRow(ctx, query, userID)

	var p domain.Profile
	err := row.Scan(
		&p.UserID,
		&p.FullName,
		&p.PhoneNumber,
		&p.Photo,
		&p.City,
		&p.Country,
	)

	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *ProfileRepository) Create(
	ctx context.Context,
	p *domain.Profile,
) error {

	query := `
		INSERT INTO user_profiles (
			user_id, full_name, phone_number, photo, city, country
		)
		VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		p.UserID,
		p.FullName,
		p.PhoneNumber,
		p.Photo,
		p.City,
		p.Country,
	)

	return err
}

func (r *ProfileRepository) Update(
	ctx context.Context,
	p *domain.Profile,
) error {

	query := `
		UPDATE user_profiles
		SET
			full_name=$1,
			phone_number=$2,
			photo=$3,
			city=$4,
			country=$5,
			updated_at=now()
		WHERE user_id=$6
	`

	cmd, err := r.DB.Exec(
		ctx,
		query,
		p.FullName,
		p.PhoneNumber,
		p.Photo,
		p.City,
		p.Country,
		p.UserID,
	)

	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
