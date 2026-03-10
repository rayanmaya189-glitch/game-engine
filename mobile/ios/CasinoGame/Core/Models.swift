import Foundation

// MARK: - User

struct User: Codable {
    let id: String
    let email: String
    let username: String
    let phone: String?
    let firstName: String?
    let lastName: String?
    let avatarUrl: String?
    let kycLevel: Int
    let isVerified: Bool
    let currency: String
    let language: String
    let createdAt: String
    
    enum CodingKeys: String, CodingKey {
        case id, email, username, phone
        case firstName = "first_name"
        case lastName = "last_name"
        case avatarUrl = "avatar_url"
        case kycLevel = "kyc_level"
        case isVerified = "is_verified"
        case currency, language
        case createdAt = "created_at"
    }
}

// MARK: - Auth Response

struct LoginResponse: Codable {
    let user: User
    let accessToken: String
    let refreshToken: String
    let expiresIn: Int
    
    enum CodingKeys: String, CodingKey {
        case user
        case accessToken = "access_token"
        case refreshToken = "refresh_token"
        case expiresIn = "expires_in"
    }
}

struct RegisterResponse: Codable {
    let user: User
    let accessToken: String
    let refreshToken: String
    let expiresIn: Int
    let message: String
    
    enum CodingKeys: String, CodingKey {
        case user
        case accessToken = "access_token"
        case refreshToken = "refresh_token"
        case expiresIn = "expires_in"
        case message
    }
}

// MARK: - Game

struct Game: Codable, Identifiable {
    let id: String
    let name: String
    let slug: String
    let description: String?
    let provider: GameProvider?
    let category: GameCategory?
    let thumbnailUrl: String?
    let backgroundUrl: String?
    let gameUrl: String?
    let rtp: Double?
    let minBet: Double?
    let maxBet: Double?
    let volatility: String?
    let isFeatured: Bool
    let isNew: Bool
    let isHot: Bool
    let isFavorite: Bool
    let isAvailable: Bool
    let jackpotAmount: Double?
    let playCount: Int
    let rating: Double?
    
    enum CodingKeys: String, CodingKey {
        case id, name, slug, description, provider, category
        case thumbnailUrl = "thumbnail_url"
        case backgroundUrl = "background_url"
        case gameUrl = "game_url"
        case rtp
        case minBet = "min_bet"
        case maxBet = "max_bet"
        case volatility
        case isFeatured = "is_featured"
        case isNew = "is_new"
        case isHot = "is_hot"
        case isFavorite = "is_favorite"
        case isAvailable = "is_available"
        case jackpotAmount = "jackpot_amount"
        case playCount = "play_count"
        case rating
    }
}

struct GameProvider: Codable {
    let id: String
    let name: String
    let logoUrl: String?
    
    enum CodingKeys: String, CodingKey {
        case id, name
        case logoUrl = "logo_url"
    }
}

struct GameCategory: Codable, Identifiable {
    let id: String
    let name: String
    let slug: String
    let icon: String?
    let gameCount: Int
    
    enum CodingKeys: String, CodingKey {
        case id, name, slug, icon
        case gameCount = "game_count"
    }
}

struct GameListResponse: Codable {
    let games: [Game]
    let total: Int
    let page: Int
    let pageSize: Int
    let totalPages: Int
    
    enum CodingKeys: String, CodingKey {
        case games, total, page
        case pageSize = "page_size"
        case totalPages = "total_pages"
    }
}

struct FeaturedGamesResponse: Codable {
    let featured: [Game]
    let popular: [Game]
    let new: [Game]
    let jackpot: [Game]
}

struct CategoriesResponse: Codable {
    let categories: [GameCategory]
}

struct GameSession: Codable {
    let sessionId: String
    let gameId: String
    let playUrl: String
    let funPlayUrl: String?
    let expiresAt: String
    
    enum CodingKeys: String, CodingKey {
        case sessionId = "session_id"
        case gameId = "game_id"
        case playUrl = "play_url"
        case funPlayUrl = "fun_play_url"
        case expiresAt = "expires_at"
    }
}

// MARK: - Wallet

struct Wallet: Codable {
    let id: String
    let userId: String
    let currency: String
    let balance: Double
    let bonusBalance: Double
    let pendingBalance: Double
    let totalWon: Double
    let totalWagered: Double
    let lastUpdated: String
    
    enum CodingKeys: String, CodingKey {
        case id
        case userId = "user_id"
        case currency, balance
        case bonusBalance = "bonus_balance"
        case pendingBalance = "pending_balance"
        case totalWon = "total_won"
        case totalWagered = "total_wagered"
        case lastUpdated = "last_updated"
    }
}

struct WalletBalance: Codable {
    let balance: Double
    let bonusBalance: Double
    let currency: String
    
    enum CodingKeys: String, CodingKey {
        case balance
        case bonusBalance = "bonus_balance"
        case currency
    }
}

struct DepositResponse: Codable {
    let transactionId: String
    let status: String
    let amount: Double
    let currency: String
    let paymentUrl: String?
    let expiresAt: String?
    
    enum CodingKeys: String, CodingKey {
        case transactionId = "transaction_id"
        case status, amount, currency
        case paymentUrl = "payment_url"
        case expiresAt = "expires_at"
    }
}

struct WithdrawResponse: Codable {
    let transactionId: String
    let status: String
    let amount: Double
    let currency: String
    let fee: Double
    let estimatedArrival: String?
    
    enum CodingKeys: String, CodingKey {
        case transactionId = "transaction_id"
        case status, amount, currency, fee
        case estimatedArrival = "estimated_arrival"
    }
}

struct Transaction: Codable, Identifiable {
    let id: String
    let userId: String
    let type: String
    let amount: Double
    let currency: String
    let status: String
    let paymentMethod: String?
    let referenceId: String?
    let description: String?
    let createdAt: String
    let completedAt: String?
    
    enum CodingKeys: String, CodingKey {
        case id
        case userId = "user_id"
        case type, amount, currency, status
        case paymentMethod = "payment_method"
        case referenceId = "reference_id"
        case description
        case createdAt = "created_at"
        case completedAt = "completed_at"
    }
}

struct TransactionListResponse: Codable {
    let transactions: [Transaction]
    let total: Int
    let page: Int
    let pageSize: Int
    let totalPages: Int
    
    enum CodingKeys: String, CodingKey {
        case transactions, total, page
        case pageSize = "page_size"
        case totalPages = "total_pages"
    }
}

struct PaymentMethod: Codable {
    let id: String
    let name: String
    let type: String
    let logoUrl: String?
    let minAmount: Double
    let maxAmount: Double
    let feePercentage: Double
    let processingTime: String
    let isAvailable: Bool
    
    enum CodingKeys: String, CodingKey {
        case id, name, type
        case logoUrl = "logo_url"
        case minAmount = "min_amount"
        case maxAmount = "max_amount"
        case feePercentage = "fee_percentage"
        case processingTime = "processing_time"
        case isAvailable = "is_available"
    }
}

struct PaymentMethodsResponse: Codable {
    let depositMethods: [PaymentMethod]
    let withdrawMethods: [PaymentMethod]
    
    enum CodingKeys: String, CodingKey {
        case depositMethods = "deposit_methods"
        case withdrawMethods = "withdraw_methods"
    }
}
