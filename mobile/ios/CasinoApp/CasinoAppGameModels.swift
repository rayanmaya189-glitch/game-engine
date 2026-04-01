import Foundation

struct Game: Codable, Identifiable {
    let id: String
    let name: String
    let provider: String
    let category: String
    let thumbnail: String?
    let status: String
    let minBet: Double
    let maxBet: Double
    let rtp: Double
    let volatility: String?

    enum CodingKeys: String, CodingKey {
        case id, name, provider, category, thumbnail, status, rtp, volatility
        case minBet = "min_bet"
        case maxBet = "max_bet"
    }
}

struct GameDetails: Codable {
    let id: String
    let name: String
    let provider: String
    let category: String
    let thumbnail: String?
    let banner: String?
    let description: String
    let status: String
    let minBet: Double
    let maxBet: Double
    let rtp: Double
    let volatility: String?
}

struct GamesResponse: Codable {
    let games: [Game]
    let total: Int
    let page: Int
    let pages: Int
}

struct Category: Codable, Identifiable {
    let id: String
    let name: String
    let icon: String?
    let gameCount: Int

    enum CodingKeys: String, CodingKey {
        case id, name, icon
        case gameCount = "game_count"
    }
}

struct CategoriesResponse: Codable {
    let categories: [Category]
}

struct Tournament: Codable, Identifiable {
    let id: String
    let name: String
    let description: String?
    let game: String
    let prizePool: Double
    let minBet: Double
    let startDate: String
    let endDate: String
    let status: String
    let playerCount: Int

    enum CodingKeys: String, CodingKey {
        case id, name, description, game, status
        case prizePool = "prize_pool"
        case minBet = "min_bet"
        case startDate = "start_date"
        case endDate = "end_date"
        case playerCount = "player_count"
    }
}

struct TournamentsResponse: Codable {
    let tournaments: [Tournament]
    let total: Int
    let page: Int
    let pages: Int
}

struct Jackpot: Codable, Identifiable {
    let id: String
    let name: String
    let game: String
    let currentAmount: Double
    let minBet: Double
    let maxWin: Double
    let status: String
    let hitCount: Int

    enum CodingKeys: String, CodingKey {
        case id, name, game, status
        case currentAmount = "current_amount"
        case minBet = "min_bet"
        case maxWin = "max_win"
        case hitCount = "hit_count"
    }
}

struct JackpotsResponse: Codable {
    let jackpots: [Jackpot]
}
