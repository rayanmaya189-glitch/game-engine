import Foundation

// MARK: - Extended Models
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

// MARK: - KYC

struct KycStatus: Codable {
    let currentLevel: Int
    let maxLevel: Int
    let documents: [KycDocument]
    let isFullyVerified: Bool

    enum CodingKeys: String, CodingKey {
        case currentLevel = "current_level"
        case maxLevel = "max_level"
        case documents
        case isFullyVerified = "is_fully_verified"
    }
}

struct KycDocument: Codable, Identifiable {
    let id: String
    let type: String
    let status: String
    let fileName: String?
    let uploadedAt: String?
    let reviewedAt: String?
    let rejectionReason: String?

    enum CodingKeys: String, CodingKey {
        case id, type, status
        case fileName = "file_name"
        case uploadedAt = "uploaded_at"
        case reviewedAt = "reviewed_at"
        case rejectionReason = "rejection_reason"
    }
}

struct KycUploadResponse: Codable {
    let document: KycDocument
    let message: String
}

struct KycVerificationResponse: Codable {
    let status: String
    let message: String
    let newLevel: Int?

    enum CodingKeys: String, CodingKey {
        case status, message
        case newLevel = "new_level"
    }
}

// MARK: - Game Detail

struct RecentWin: Codable, Identifiable {
    let id: String
    let username: String
    let gameId: String
    let gameName: String
    let amount: Double
    let multiplier: Double
    let currency: String
    let createdAt: String

    enum CodingKeys: String, CodingKey {
        case id, username
        case gameId = "game_id"
        case gameName = "game_name"
        case amount, multiplier, currency
        case createdAt = "created_at"
    }
}

// MARK: - Bet History

struct Bet: Codable, Identifiable {
    let id: String
    let gameId: String
    let gameName: String
    let gameThumbnailUrl: String?
    let stake: Double
    let winAmount: Double?
    let result: String
    let currency: String
    let multiplier: Double?
    let placedAt: String
    let settledAt: String?

    enum CodingKeys: String, CodingKey {
        case id
        case gameId = "game_id"
        case gameName = "game_name"
        case gameThumbnailUrl = "game_thumbnail_url"
        case stake
        case winAmount = "win_amount"
        case result, currency, multiplier
        case placedAt = "placed_at"
        case settledAt = "settled_at"
    }
}

struct BetListResponse: Codable {
    let bets: [Bet]
    let total: Int
    let page: Int
    let pageSize: Int
    let totalPages: Int

    enum CodingKeys: String, CodingKey {
        case bets, total, page
        case pageSize = "page_size"
        case totalPages = "total_pages"
    }
}
