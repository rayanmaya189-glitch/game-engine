package model

import (
	"time"
)

type BannerType string

const (
	BannerTypeHero         BannerType = "HERO"
	BannerTypeSidebar      BannerType = "SIDEBAR"
	BannerTypePopup        BannerType = "POPUP"
	BannerTypeInGame       BannerType = "IN_GAME"
	BannerTypeInterstitial BannerType = "INTERSTITIAL"
	BannerTypeNotification BannerType = "NOTIFICATION_BAR"
)

type BannerStatus string

const (
	BannerStatusActive   BannerStatus = "ACTIVE"
	BannerStatusInactive BannerStatus = "INACTIVE"
	BannerStatusScheduled BannerStatus = "SCHEDULED"
	BannerStatusExpired  BannerStatus = "EXPIRED"
)

type Banner struct {
	ID              string       `json:"id"`
	Title           string       `json:"title"`
	Description     string       `json:"description"`
	ImageURL        string       `json:"image_url"`
	ClickURL        string       `json:"click_url"`
	BannerType      BannerType   `json:"banner_type"`
	Status          BannerStatus `json:"status"`
	Priority        int          `json:"priority"`
	Width           int          `json:"width"`
	Height          int          `json:"height"`
	StartDate       *time.Time   `json:"start_date"`
	EndDate         *time.Time   `json:"end_date"`
	TargetCountries []string     `json:"target_countries"`
	TargetVIPLevels []string     `json:"target_vip_levels"`
	TargetGameTypes []string     `json:"target_game_types"`
	ClickCount      int64        `json:"click_count"`
	ImpressionCount int64        `json:"impression_count"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
}

type Announcement struct {
	ID        string       `json:"id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	Type      string       `json:"type"`
	Status    BannerStatus `json:"status"`
	Priority  int          `json:"priority"`
	StartDate *time.Time   `json:"start_date"`
	EndDate   *time.Time   `json:"end_date"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type BannerFilter struct {
	BannerType string
	Status     string
	Country    string
	VIPLevel   string
	GameType   string
	Page       int
	PageSize   int
}

type BannerList struct {
	Banners    []*Banner `json:"banners"`
	Total      int64     `json:"total"`
	Page       int       `json:"page"`
	PageSize   int       `json:"page_size"`
}
