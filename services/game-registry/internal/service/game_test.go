package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/game_engine/game-registry/internal/config"
	"github.com/game_engine/game-registry/internal/enums"
	"github.com/game_engine/game-registry/internal/model"
)

type mockGameRepo struct {
	games     map[string]*model.Game
	gameList  []model.GameSummary
	categories []model.GameCategory
	providers  []model.GameProvider
	pagination model.PaginationResult
	err       error
}

func newMockGameRepo() *mockGameRepo {
	return &mockGameRepo{
		games: make(map[string]*model.Game),
	}
}

func (m *mockGameRepo) GetGameByID(ctx context.Context, gameID string) (*model.Game, error) {
	if m.err != nil {
		return nil, m.err
	}
	g, ok := m.games[gameID]
	if !ok {
		return nil, nil
	}
	return g, nil
}

func (m *mockGameRepo) ListGames(ctx context.Context, filter model.GameListFilter) ([]model.GameSummary, model.PaginationResult, error) {
	if m.err != nil {
		return nil, model.PaginationResult{}, m.err
	}
	return m.gameList, m.pagination, nil
}

func (m *mockGameRepo) CreateGame(ctx context.Context, game *model.Game) error {
	if m.err != nil {
		return m.err
	}
	m.games[game.ID] = game
	return nil
}

func (m *mockGameRepo) UpdateGame(ctx context.Context, game *model.Game) error {
	m.games[game.ID] = game
	return m.err
}

func (m *mockGameRepo) ToggleGame(ctx context.Context, gameID string, status enums.Status) error {
	if g, ok := m.games[gameID]; ok {
		g.Status = status
	}
	return m.err
}

func (m *mockGameRepo) SetGameOrder(ctx context.Context, gameID string, sortOrder int) error {
	return m.err
}

func (m *mockGameRepo) GetCategories(ctx context.Context, includeGamesCount bool) ([]model.GameCategory, error) {
	return m.categories, m.err
}

func (m *mockGameRepo) GetProviders(ctx context.Context, activeOnly bool) ([]model.GameProvider, error) {
	return m.providers, m.err
}

func (m *mockGameRepo) SearchGames(ctx context.Context, query string, limit int, categoryID string) ([]model.GameSummary, error) {
	return m.gameList, m.err
}

func (m *mockGameRepo) GetFeaturedGames(ctx context.Context, limit int, categoryID string) ([]model.GameSummary, error) {
	return m.gameList, m.err
}

func (m *mockGameRepo) GetPopularGames(ctx context.Context, limit int, categoryID string) ([]model.GameSummary, error) {
	return m.gameList, m.err
}

func (m *mockGameRepo) GetNewGames(ctx context.Context, limit int, categoryID string) ([]model.GameSummary, error) {
	return m.gameList, m.err
}

func (m *mockGameRepo) CreateGameConfig(ctx context.Context, config *model.GameConfig) error {
	return m.err
}

func setupGameServiceTest(t *testing.T) (*GameService, *mockGameRepo) {
	t.Helper()
	cfg := &config.Config{
		Game: config.GameConfig{
			LaunchURLTemplate: "https://play.example.com/{game_id}?token={token}&session={session}",
			SessionTTL:        3600,
			CacheTTL:          600,
		},
	}

	repo := newMockGameRepo()
	svc := &GameService{
		repo:   nil,
		config: cfg,
		nc:     nil,
	}

	_ = svc
	return svc, repo
}

func TestListGames_PaginationDefaults(t *testing.T) {
	svc, repo := setupGameServiceTest(t)

	repo.gameList = []model.GameSummary{
		{GameID: "g1", Name: "Slots Game"},
		{GameID: "g2", Name: "Roulette"},
	}
	repo.pagination = model.PaginationResult{
		Page:       1,
		PageSize:   20,
		TotalCount: 2,
		TotalPages: 1,
	}

	req := &ListGamesRequest{}
	req.Pagination.Page = 0
	req.Pagination.PageSize = 0

	resp, err := svc.ListGames(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response")
	}
}

func TestListGames_RepoError(t *testing.T) {
	svc, repo := setupGameServiceTest(t)
	repo.err = fmt.Errorf("database connection failed")

	_, err := svc.ListGames(context.Background(), &ListGamesRequest{})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestListGames_ValidRequest(t *testing.T) {
	svc, repo := setupGameServiceTest(t)

	repo.gameList = []model.GameSummary{
		{GameID: "g1", Name: "Blackjack"},
	}
	repo.pagination = model.PaginationResult{Page: 1, PageSize: 10, TotalCount: 1, TotalPages: 1}

	req := &ListGamesRequest{}
	req.Pagination.Page = 1
	req.Pagination.PageSize = 10

	resp, err := svc.ListGames(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Games) != 1 {
		t.Errorf("expected 1 game, got %d", len(resp.Games))
	}
	if resp.Games[0].Name != "Blackjack" {
		t.Errorf("expected Blackjack, got %s", resp.Games[0].Name)
	}
}

func TestGetGame_Found(t *testing.T) {
	svc, repo := setupGameServiceTest(t)

	repo.games["game-1"] = &model.Game{
		ID:   "game-1",
		Name: "Test Slots",
		Type: enums.GameCategorySlots,
	}

	game, err := svc.GetGame(context.Background(), "game-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if game.Name != "Test Slots" {
		t.Errorf("expected Test Slots, got %s", game.Name)
	}
}

func TestGetGame_NotFound(t *testing.T) {
	svc, repo := setupGameServiceTest(t)

	game, err := svc.GetGame(context.Background(), "missing")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if game != nil {
		t.Error("expected nil game for missing ID")
	}
}

func TestGetGame_RepoError(t *testing.T) {
	svc, repo := setupGameServiceTest(t)
	repo.err = fmt.Errorf("db error")

	_, err := svc.GetGame(context.Background(), "game-1")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestCreateGame_Success(t *testing.T) {
	svc, repo := setupGameServiceTest(t)

	game := &model.Game{
		ID:         "new-game",
		Name:       "New Poker",
		ProviderID: "provider-1",
		Type:       enums.GameCategoryCards,
		Status:     enums.StatusActive,
	}

	err := svc.CreateGame(context.Background(), game)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, ok := repo.games["new-game"]; !ok {
		t.Error("expected game to be stored in repo")
	}
}

func TestCreateGame_RepoError(t *testing.T) {
	svc, repo := setupGameServiceTest(t)
	repo.err = fmt.Errorf("insert failed")

	game := &model.Game{ID: "fail", Name: "Fail Game"}
	err := svc.CreateGame(context.Background(), game)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestListGames_WithFilters(t *testing.T) {
	svc, repo := setupGameServiceTest(t)

	repo.gameList = []model.GameSummary{
		{GameID: "g1", Name: "Featured Slots", IsFeatured: true},
	}
	repo.pagination = model.PaginationResult{Page: 1, PageSize: 20, TotalCount: 1, TotalPages: 1}

	req := &ListGamesRequest{
		IsFeatured: true,
		Query:      "slots",
	}
	req.Pagination.Page = 1
	req.Pagination.PageSize = 20

	resp, err := svc.ListGames(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Games) != 1 {
		t.Errorf("expected 1 game, got %d", len(resp.Games))
	}
}

func TestGetGameConfig_NotFound(t *testing.T) {
	svc, repo := setupGameServiceTest(t)
	repo.games["missing"] = nil

	_, err := svc.GetGameConfig(context.Background(), "missing", "user-1", enums.DeviceTypeDesktop, enums.GameLanguageEN, "USD", "session-1")
	if err == nil {
		t.Fatal("expected error for missing game")
	}
}
