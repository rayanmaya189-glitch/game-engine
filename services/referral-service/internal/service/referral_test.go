package service

import (
	"errors"
	"testing"

	"github.com/game_engine/referral-service/internal/model"
)

type mockReferralRepo struct {
	codes        map[string]*model.ReferralCode
	referrals    map[string]*model.Referral
	referrerByC  map[string]string
	createErr    error
	getCodeErr   error
	getReferrerE error
	createRefErr error
	markQualErr  error
	markRewErr   error
	claimErr     error
	trackErr     error
}

func newMockReferralRepo() *mockReferralRepo {
	return &mockReferralRepo{
		codes:       make(map[string]*model.ReferralCode),
		referrals:   make(map[string]*model.Referral),
		referrerByC: make(map[string]string),
	}
}

func (m *mockReferralRepo) CreateReferralCode(rc *model.ReferralCode) error {
	if m.createErr != nil {
		return m.createErr
	}
	m.codes[rc.PlayerID] = rc
	return nil
}

func (m *mockReferralRepo) GetReferralCodeByPlayer(playerID string) (*model.ReferralCode, error) {
	if m.getCodeErr != nil {
		return nil, m.getCodeErr
	}
	rc, ok := m.codes[playerID]
	if !ok {
		return nil, errors.New("not found")
	}
	return rc, nil
}

func (m *mockReferralRepo) GetReferrerByCode(code string) (string, error) {
	if m.getReferrerE != nil {
		return "", m.getReferrerE
	}
	pid, ok := m.referrerByC[code]
	if !ok {
		return "", errors.New("code not found")
	}
	return pid, nil
}

func (m *mockReferralRepo) CreateReferral(ref *model.Referral) error {
	if m.createRefErr != nil {
		return m.createRefErr
	}
	m.referrals[ref.ID] = ref
	return nil
}

func (m *mockReferralRepo) GetReferralByID(id string) (*model.Referral, error) {
	ref, ok := m.referrals[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return ref, nil
}

func (m *mockReferralRepo) MarkQualified(id string) error   { return m.markQualErr }
func (m *mockReferralRepo) MarkRewarded(id string, amount float64, rt model.RewardType) error {
	return m.markRewErr
}
func (m *mockReferralRepo) ClaimReward(id string) error { return m.claimErr }
func (m *mockReferralRepo) TrackReferralSignup(code string) error { return m.trackErr }
func (m *mockReferralRepo) TrackReferralClick(code string) error  { return nil }
func (m *mockReferralRepo) GetReferralsByReferrer(_ string, _ *model.ReferralFilter) (*model.ReferralList, error) {
	return &model.ReferralList{}, nil
}
func (m *mockReferralRepo) GetReferralStats(_ string) (*model.ReferralStats, error) {
	return &model.ReferralStats{}, nil
}
func (m *mockReferralRepo) GetActiveRewards() ([]*model.ReferralReward, error) {
	return nil, nil
}
func (m *mockReferralRepo) GetChildReferrals(_ string) ([]*model.Referral, error) {
	return nil, nil
}

func TestGenerateReferralCode(t *testing.T) {
	tests := []struct {
		name     string
		playerID string
		wantErr  bool
	}{
		{"valid", "player-1", false},
		{"empty player", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockReferralRepo()
			svc := NewReferralService(repo)
			rc, err := svc.GenerateReferralCode(tt.playerID)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if rc.Code == "" {
				t.Error("expected non-empty code")
			}
			if rc.PlayerID != tt.playerID {
				t.Errorf("expected player %s, got %s", tt.playerID, rc.PlayerID)
			}
			if rc.ReferralURL == "" {
				t.Error("expected non-empty referral URL")
			}
		})
	}
}

func TestGenerateReferralCode_Existing(t *testing.T) {
	repo := newMockReferralRepo()
	repo.codes["p1"] = &model.ReferralCode{PlayerID: "p1", Code: "EXISTING"}
	svc := NewReferralService(repo)

	rc, err := svc.GenerateReferralCode("p1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rc.Code != "EXISTING" {
		t.Errorf("expected existing code, got %s", rc.Code)
	}
}

func TestTrackReferral(t *testing.T) {
	repo := newMockReferralRepo()
	repo.referrerByC["REF123"] = "referrer-1"
	svc := NewReferralService(repo)

	tests := []struct {
		name       string
		code       string
		refereeID  string
		wantErr    bool
		errContain string
	}{
		{"valid", "REF123", "referee-1", false, ""},
		{"empty code", "", "r1", true, "referral code is required"},
		{"empty referee", "REF123", "", true, "referee_id is required"},
		{"self referral", "REF123", "referrer-1", true, "cannot refer yourself"},
		{"invalid code", "BAD", "r1", true, "invalid referral code"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ref, err := svc.TrackReferral(tt.code, tt.refereeID, "web")

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if ref.Status != model.ReferralStatusActive {
				t.Errorf("expected ACTIVE, got %s", ref.Status)
			}
			if ref.ReferrerID != "referrer-1" {
				t.Errorf("expected referrer-1, got %s", ref.ReferrerID)
			}
		})
	}
}

func TestQualifyReferral(t *testing.T) {
	repo := newMockReferralRepo()
	svc := NewReferralService(repo)

	err := svc.QualifyReferral("ref-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRewardReferral(t *testing.T) {
	repo := newMockReferralRepo()
	svc := NewReferralService(repo)

	tests := []struct {
		name    string
		amount  float64
		wantErr bool
	}{
		{"valid", 100.0, false},
		{"zero amount", 0, true},
		{"negative amount", -10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.RewardReferral("ref-1", tt.amount, model.RewardTypeCash)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestClaimReward(t *testing.T) {
	repo := newMockReferralRepo()
	repo.referrals["r1"] = &model.Referral{
		ID: "r1", ReferrerID: "p1", Status: model.ReferralStatusRewarded, RewardClaimed: false,
	}
	svc := NewReferralService(repo)

	err := svc.ClaimReward("r1", "p1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestClaimReward_NotAuthorized(t *testing.T) {
	repo := newMockReferralRepo()
	repo.referrals["r1"] = &model.Referral{
		ID: "r1", ReferrerID: "p1", Status: model.ReferralStatusRewarded, RewardClaimed: false,
	}
	svc := NewReferralService(repo)

	err := svc.ClaimReward("r1", "wrong-user")
	if err == nil {
		t.Fatal("expected authorization error")
	}
}

func TestClaimReward_NotRewarded(t *testing.T) {
	repo := newMockReferralRepo()
	repo.referrals["r1"] = &model.Referral{
		ID: "r1", ReferrerID: "p1", Status: model.ReferralStatusActive, RewardClaimed: false,
	}
	svc := NewReferralService(repo)

	err := svc.ClaimReward("r1", "p1")
	if err == nil {
		t.Fatal("expected error for non-rewarded status")
	}
}

func TestClaimReward_AlreadyClaimed(t *testing.T) {
	repo := newMockReferralRepo()
	repo.referrals["r1"] = &model.Referral{
		ID: "r1", ReferrerID: "p1", Status: model.ReferralStatusRewarded, RewardClaimed: true,
	}
	svc := NewReferralService(repo)

	err := svc.ClaimReward("r1", "p1")
	if err == nil {
		t.Fatal("expected error for already claimed")
	}
}
