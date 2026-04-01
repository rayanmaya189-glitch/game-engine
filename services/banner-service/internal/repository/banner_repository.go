package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"github.com/game_engine/banner-service/internal/config"
	"github.com/game_engine/banner-service/internal/model"
)

type BannerRepository struct {
	db *sql.DB
}

func NewPostgresDB(cfg config.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxConns)
	db.SetMaxIdleConns(cfg.MaxConns / 2)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewBannerRepository(db *sql.DB) *BannerRepository {
	return &BannerRepository{db: db}
}

func (r *BannerRepository) Create(banner *model.Banner) error {
	query := `INSERT INTO banners (id, title, description, image_url, click_url, banner_type, status, priority,
		width, height, start_date, end_date, target_countries, target_vip_levels, target_game_types, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)`

	countries := strings.Join(banner.TargetCountries, ",")
	vipLevels := strings.Join(banner.TargetVIPLevels, ",")
	gameTypes := strings.Join(banner.TargetGameTypes, ",")

	_, err := r.db.Exec(query,
		banner.ID, banner.Title, banner.Description, banner.ImageURL, banner.ClickURL,
		banner.BannerType, banner.Status, banner.Priority, banner.Width, banner.Height,
		banner.StartDate, banner.EndDate, countries, vipLevels, gameTypes,
		banner.CreatedAt, banner.UpdatedAt,
	)
	return err
}

func (r *BannerRepository) GetByID(id string) (*model.Banner, error) {
	query := `SELECT id, title, description, image_url, click_url, banner_type, status, priority,
		width, height, start_date, end_date, target_countries, target_vip_levels, target_game_types,
		click_count, impression_count, created_at, updated_at FROM banners WHERE id = $1`

	var banner model.Banner
	var countries, vipLevels, gameTypes sql.NullString
	var startDate, endDate sql.NullTime

	err := r.db.QueryRow(query, id).Scan(
		&banner.ID, &banner.Title, &banner.Description, &banner.ImageURL, &banner.ClickURL,
		&banner.BannerType, &banner.Status, &banner.Priority, &banner.Width, &banner.Height,
		&startDate, &endDate, &countries, &vipLevels, &gameTypes,
		&banner.ClickCount, &banner.ImpressionCount, &banner.CreatedAt, &banner.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	if startDate.Valid {
		banner.StartDate = &startDate.Time
	}
	if endDate.Valid {
		banner.EndDate = &endDate.Time
	}
	banner.TargetCountries = splitString(countries.String)
	banner.TargetVIPLevels = splitString(vipLevels.String)
	banner.TargetGameTypes = splitString(gameTypes.String)

	return &banner, nil
}

func (r *BannerRepository) List(filter *model.BannerFilter) (*model.BannerList, error) {
	where := []string{"1=1"}
	args := []interface{}{}
	idx := 1

	if filter.BannerType != "" {
		where = append(where, fmt.Sprintf("banner_type = $%d", idx))
		args = append(args, filter.BannerType)
		idx++
	}
	if filter.Status != "" {
		where = append(where, fmt.Sprintf("status = $%d", idx))
		args = append(args, filter.Status)
		idx++
	}

	whereClause := strings.Join(where, " AND ")

	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM banners WHERE %s", whereClause)
	var total int64
	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, err
	}

	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PageSize < 1 {
		filter.PageSize = 20
	}
	offset := (filter.Page - 1) * filter.PageSize

	query := fmt.Sprintf(`SELECT id, title, description, image_url, click_url, banner_type, status, priority,
		width, height, start_date, end_date, click_count, impression_count, created_at, updated_at
		FROM banners WHERE %s ORDER BY priority DESC, created_at DESC LIMIT $%d OFFSET $%d`,
		whereClause, idx, idx+1)

	args = append(args, filter.PageSize, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banners []*model.Banner
	for rows.Next() {
		var b model.Banner
		var startDate, endDate sql.NullTime

		if err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ClickURL,
			&b.BannerType, &b.Status, &b.Priority, &b.Width, &b.Height,
			&startDate, &endDate, &b.ClickCount, &b.ImpressionCount, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		if startDate.Valid {
			b.StartDate = &startDate.Time
		}
		if endDate.Valid {
			b.EndDate = &endDate.Time
		}
		banners = append(banners, &b)
	}

	return &model.BannerList{
		Banners:  banners,
		Total:    total,
		Page:     filter.Page,
		PageSize: filter.PageSize,
	}, nil
}

func (r *BannerRepository) Update(banner *model.Banner) error {
	query := `UPDATE banners SET title=$1, description=$2, image_url=$3, click_url=$4, banner_type=$5,
		status=$6, priority=$7, width=$8, height=$9, start_date=$10, end_date=$11,
		target_countries=$12, target_vip_levels=$13, target_game_types=$14, updated_at=$15 WHERE id=$16`

	countries := strings.Join(banner.TargetCountries, ",")
	vipLevels := strings.Join(banner.TargetVIPLevels, ",")
	gameTypes := strings.Join(banner.TargetGameTypes, ",")

	_, err := r.db.Exec(query,
		banner.Title, banner.Description, banner.ImageURL, banner.ClickURL, banner.BannerType,
		banner.Status, banner.Priority, banner.Width, banner.Height, banner.StartDate, banner.EndDate,
		countries, vipLevels, gameTypes, time.Now(), banner.ID,
	)
	return err
}

func (r *BannerRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM banners WHERE id = $1", id)
	return err
}

func (r *BannerRepository) IncrementImpression(id string) error {
	_, err := r.db.Exec("UPDATE banners SET impression_count = impression_count + 1 WHERE id = $1", id)
	return err
}

func (r *BannerRepository) IncrementClick(id string) error {
	_, err := r.db.Exec("UPDATE banners SET click_count = click_count + 1 WHERE id = $1", id)
	return err
}

func (r *BannerRepository) CreateAnnouncement(a *model.Announcement) error {
	query := `INSERT INTO announcements (id, title, content, type, status, priority, start_date, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.Exec(query, a.ID, a.Title, a.Content, a.Type, a.Status, a.Priority,
		a.StartDate, a.EndDate, a.CreatedAt, a.UpdatedAt)
	return err
}

func (r *BannerRepository) ListActiveAnnouncements() ([]*model.Announcement, error) {
	query := `SELECT id, title, content, type, status, priority, start_date, end_date, created_at, updated_at
		FROM announcements WHERE status = 'ACTIVE' AND (start_date IS NULL OR start_date <= NOW())
		AND (end_date IS NULL OR end_date >= NOW()) ORDER BY priority DESC, created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var announcements []*model.Announcement
	for rows.Next() {
		var a model.Announcement
		var startDate, endDate sql.NullTime
		if err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.Type, &a.Status, &a.Priority,
			&startDate, &endDate, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		if startDate.Valid {
			a.StartDate = &startDate.Time
		}
		if endDate.Valid {
			a.EndDate = &endDate.Time
		}
		announcements = append(announcements, &a)
	}
	return announcements, nil
}

func splitString(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, ",")
}
