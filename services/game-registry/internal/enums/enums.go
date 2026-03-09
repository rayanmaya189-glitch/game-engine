package enums

// Status represents the general status
type Status int32

const (
	StatusUnspecified Status = iota
	StatusPending
	StatusActive
	StatusInactive
	StatusSuspended
	StatusDeleted
	StatusBlocked
	StatusArchived
	StatusBanned
	StatusLocked
	StatusExpired
	StatusCancelled
	StatusCompleted
	StatusFailed
)

// GameCategory represents game category types
type GameCategory int32

const (
	GameCategoryUnspecified GameCategory = iota
	GameCategorySlots
	GameCategoryTableGames
	GameCategoryLiveCasino
	GameCategorySportsBetting
	GameCategoryEsports
	GameCategoryLottery
	GameCategoryBingo
	GameCategoryPoker
	GameCategoryVirtualSports
	GameCategoryInstantWin
	GameCategoryScratchCards
	GameCategoryKeno
	GameCategoryBlackjack
	GameCategoryRoulette
	GameCategoryBaccarat
	GameCategoryCraps
	GameCategorySicBo
	GameCategoryDragonTiger
	GameCategoryGameShows
	GameCategoryOther
)

// GameProvider represents game provider types
type GameProvider int32

const (
	GameProviderUnspecified GameProvider = iota
	GameProviderInternal
	GameProviderPragmaticPlay
	GameProviderNetEnt
	GameProviderMicrogaming
	GameProviderEvolution
	GameProviderPlaytech
	GameProviderBetSoft
	GameProviderBetGaming
	GameProviderAmatic
	GameProviderBelatra
	GameProviderEGT
	GameProviderIgrosoft
	GameProviderPlayNGO
	GameProviderYggdrasil
	GameProviderQuickSpin
	GameProviderRelaxGaming
	GameProviderThunderkick
	GameProviderELKStudios
	GameProviderNoLimitCity
	GameProviderRedBeardGaming
)

// DeviceType represents device types
type DeviceType int32

const (
	DeviceTypeUnspecified DeviceType = iota
	DeviceTypeDesktop
	DeviceTypeMobile
	DeviceTypeTablet
	DeviceTypeTV
	DeviceTypeWatch
	DeviceTypeVR
	DeviceTypeConsole
)

// GameLanguage represents supported languages
type GameLanguage int32

const (
	LanguageUnspecified GameLanguage = iota
	LanguageEn
	LanguageTh
	LanguageVi
	LanguageId
	LanguageMs
	LanguageZh
	LanguageZhTw
	LanguageJa
	LanguageKo
	LanguageEs
	LanguagePt
	LanguageRu
	LanguageAr
	LanguageFr
	LanguageDe
	LanguageIt
	LanguageNl
	LanguagePl
	LanguageTr
	LanguageHi
)

func (l GameLanguage) String() string {
	return []string{"UNSPECIFIED", "EN", "TH", "VI", "ID", "MS", "ZH", "ZH_TW", "JA", "KO", "ES", "PT", "RU", "AR", "FR", "DE", "IT", "NL", "PL", "TR", "HI"}[l]
}

// Volatility represents game volatility levels
type Volatility string

const (
	VolatilityLow    Volatility = "LOW"
	VolatilityMedium Volatility = "MEDIUM"
	VolatilityHigh   Volatility = "HIGH"
)

// GameType represents the main game type
type GameType string

const (
	GameTypeSlot  GameType = "SLOT"
	GameTypeCard  GameType = "CARD"
	GameTypeDice  GameType = "DICE"
	GameTypeTable GameType = "TABLE"
	GameTypeLive  GameType = "LIVE"
)

// GameFeature represents game features
type GameFeature string

const (
	GameFeatureJackpot    GameFeature = "JACKPOT"
	GameFeatureBonus      GameFeature = "BONUS"
	GameFeatureFreeSpins  GameFeature = "FREESPINS"
	GameFeatureMegaWays   GameFeature = "MEGAWAYS"
	GameFeatureMultiplier GameFeature = "MULTIPLIER"
	GameFeatureWild       GameFeature = "WILD"
	GameFeatureScatter    GameFeature = "SCATTER"
)
