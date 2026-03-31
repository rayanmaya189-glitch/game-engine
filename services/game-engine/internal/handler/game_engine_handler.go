package handler

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/game-engine/game-engine/internal/config"
	"github.com/game-engine/game-engine/internal/registry"
	"github.com/game-engine/game-engine/internal/rng"
)

type GameEngineHandler struct {
	rng      *rng.RNGService
	registry *registry.GameRegistry
	cfg      config.GameEngineConfig
}

func NewGameEngineHandler(rngService *rng.RNGService, reg *registry.GameRegistry, cfg config.GameEngineConfig) *GameEngineHandler {
	return &GameEngineHandler{
		rng:      rngService,
		registry: reg,
		cfg:      cfg,
	}
}

func (h *GameEngineHandler) RegisterServices(server *grpc.Server) {
	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(server, healthSrv)
	reflection.Register(server)
}

func (h *GameEngineHandler) Spin(ctx context.Context, req *SpinRequest) (*SpinResponse, error) {
	if req.GameId == "" {
		return nil, status.Error(codes.InvalidArgument, "game_id is required")
	}
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	gameConfig, err := h.registry.GetGame(req.GameId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "game not found: %v", err)
	}

	reelResult, err := h.rng.GenerateReelResult(gameConfig)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate reel result: %v", err)
	}

	return &SpinResponse{
		GameId:      req.GameId,
		UserId:      req.UserId,
		ReelResult:  reelResult,
		GameConfig:  gameConfig,
	}, nil
}

func (h *GameEngineHandler) GenerateRandom(ctx context.Context, req *RandomRequest) (*RandomResponse, error) {
	if req.Length <= 0 {
		req.Length = 32
	}
	if req.Length > 256 {
		return nil, status.Error(codes.InvalidArgument, "length exceeds maximum of 256")
	}

	randomBytes, err := h.rng.GenerateBytes(int(req.Length))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate random bytes: %v", err)
	}

	return &RandomResponse{
		Bytes: randomBytes,
	}, nil
}

func (h *GameEngineHandler) VerifyProvablyFair(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {
	if req.ServerSeed == "" || req.ClientSeed == "" || req.Nonce == "" {
		return nil, status.Error(codes.InvalidArgument, "server_seed, client_seed, and nonce are required")
	}

	result, err := h.rng.VerifyProvablyFair(req.ServerSeed, req.ClientSeed, req.Nonce)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "verification failed: %v", err)
	}

	return &VerifyResponse{
		Valid:       result.Valid,
		ResultHash:  result.ResultHash,
		ResultValue: fmt.Sprintf("%d", result.ResultValue),
	}, nil
}

type SpinRequest struct {
	GameId string
	UserId string
	Bet    float64
}

type SpinResponse struct {
	GameId     string
	UserId     string
	ReelResult interface{}
	GameConfig interface{}
}

type RandomRequest struct {
	Length int32
}

type RandomResponse struct {
	Bytes []byte
}

type VerifyRequest struct {
	ServerSeed string
	ClientSeed string
	Nonce      string
}

type VerifyResponse struct {
	Valid       bool
	ResultHash  string
	ResultValue string
}
