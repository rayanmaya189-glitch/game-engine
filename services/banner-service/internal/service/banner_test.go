package service

import (
	"errors"
	"testing"

	"github.com/game_engine/banner-service/internal/model"
)

type mockBannerRepo struct {
	banners           map[string]*model.Banner
	announcements     []*model.Announcement
	createErr         error
	getErr            error
	updateErr         error
	deleteErr         error
	listResult        *model.BannerList
	listErr           error
	impressionErr     error
	clickErr          error
	activeAnnouncements []*model.Announcement
}

func newMockBannerRepo() *mockBannerRepo {
	return &mockBannerRepo{
		banners: make(map[string]*model.Banner),
	}
}

func (m *mockBannerRepo) Create(banner *model.Banner) error {
	if m.createErr != nil {
		return m.createErr
	}
	m.banners[banner.ID] = banner
	return nil
}

func (m *mockBannerRepo) GetByID(id string) (*model.Banner, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	b, ok := m.banners[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return b, nil
}

func (m *mockBannerRepo) List(filter *model.BannerFilter) (*model.BannerList, error) {
	if m.listErr != nil {
		return nil, m.listErr
	}
	if m.listResult != nil {
		return m.listResult, nil
	}
	var result []*model.Banner
	for _, b := range m.banners {
		result = append(result, b)
	}
	return &model.BannerList{Banners: result, Total: int64(len(result)), Page: 1, PageSize: 20}, nil
}

func (m *mockBannerRepo) Update(banner *model.Banner) error {
	if m.updateErr != nil {
		return m.updateErr
	}
	m.banners[banner.ID] = banner
	return nil
}

func (m *mockBannerRepo) Delete(id string) error {
	if m.deleteErr != nil {
		return m.deleteErr
	}
	delete(m.banners, id)
	return nil
}

func (m *mockBannerRepo) IncrementImpression(id string) error {
	return m.impressionErr
}

func (m *mockBannerRepo) IncrementClick(id string) error {
	return m.clickErr
}

func (m *mockBannerRepo) CreateAnnouncement(a *model.Announcement) error {
	m.announcements = append(m.announcements, a)
	return nil
}

func (m *mockBannerRepo) ListActiveAnnouncements() ([]*model.Announcement, error) {
	return m.activeAnnouncements, nil
}

func TestCreateBanner(t *testing.T) {
	tests := []struct {
		name    string
		banner  *model.Banner
		wantErr bool
		errMsg  string
	}{
		{
			"valid banner",
			&model.Banner{Title: "Test", ImageURL: "https://img.png"},
			false, "",
		},
		{
			"missing title",
			&model.Banner{ImageURL: "https://img.png"},
			true, "title is required",
		},
		{
			"missing image",
			&model.Banner{Title: "Test"},
			true, "image_url is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockBannerRepo()
			svc := NewBannerService(repo)
			result, err := svc.CreateBanner(tt.banner)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if err.Error() != tt.errMsg {
					t.Errorf("expected %q, got %q", tt.errMsg, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.ID == "" {
				t.Error("expected ID to be set")
			}
			if result.Status != model.BannerStatusActive {
				t.Errorf("expected ACTIVE status, got %s", result.Status)
			}
			if result.Width != 1920 {
				t.Errorf("expected default width 1920, got %d", result.Width)
			}
			if result.Height != 600 {
				t.Errorf("expected default height 600, got %d", result.Height)
			}
		})
	}
}

func TestGetBanner(t *testing.T) {
	repo := newMockBannerRepo()
	svc := NewBannerService(repo)

	_, err := svc.GetBanner("")
	if err == nil || err.Error() != "id is required" {
		t.Error("expected id required error")
	}

	banner := &model.Banner{Title: "T", ImageURL: "img", ID: "b1"}
	repo.banners["b1"] = banner

	got, err := svc.GetBanner("b1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != "b1" {
		t.Errorf("expected b1, got %s", got.ID)
	}
}

func TestListBanners(t *testing.T) {
	repo := newMockBannerRepo()
	svc := NewBannerService(repo)

	repo.banners["1"] = &model.Banner{ID: "1", Title: "A", ImageURL: "x"}
	repo.banners["2"] = &model.Banner{ID: "2", Title: "B", ImageURL: "y"}

	filter := &model.BannerFilter{Page: 0, PageSize: 0}
	result, err := svc.ListBanners(filter)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if filter.Page != 1 {
		t.Error("expected page default to 1")
	}
	if filter.PageSize != 20 {
		t.Error("expected page size default to 20")
	}
	if result.Total != 2 {
		t.Errorf("expected 2 banners, got %d", result.Total)
	}
}

func TestListBannersPageSizeCap(t *testing.T) {
	repo := newMockBannerRepo()
	svc := NewBannerService(repo)

	filter := &model.BannerFilter{PageSize: 200}
	svc.ListBanners(filter)

	if filter.PageSize != 100 {
		t.Errorf("expected page size capped at 100, got %d", filter.PageSize)
	}
}

func TestUpdateBanner(t *testing.T) {
	repo := newMockBannerRepo()
	svc := NewBannerService(repo)

	existing := &model.Banner{ID: "b1", Title: "Old", ImageURL: "old.png", Description: "desc"}
	repo.banners["b1"] = existing

	update := &model.Banner{Title: "New", ClickURL: "https://click.com"}
	result, err := svc.UpdateBanner("b1", update)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Title != "New" {
		t.Errorf("expected title 'New', got %q", result.Title)
	}
	if result.ClickURL != "https://click.com" {
		t.Errorf("expected click URL updated")
	}
	if result.Description != "desc" {
		t.Error("description should not change when empty in update")
	}
}

func TestUpdateBannerNotFound(t *testing.T) {
	repo := newMockBannerRepo()
	svc := NewBannerService(repo)

	_, err := svc.UpdateBanner("missing", &model.Banner{Title: "X"})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestDeleteBanner(t *testing.T) {
	repo := newMockBannerRepo()
	svc := NewBannerService(repo)

	err := svc.DeleteBanner("")
	if err == nil || err.Error() != "id is required" {
		t.Error("expected id required error")
	}

	repo.banners["b1"] = &model.Banner{ID: "b1"}
	err = svc.DeleteBanner("b1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, exists := repo.banners["b1"]; exists {
		t.Error("banner should be deleted")
	}
}
