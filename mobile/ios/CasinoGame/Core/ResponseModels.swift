import Foundation

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

struct PaymentMethodsResponse: Codable {
    let depositMethods: [PaymentMethod]
    let withdrawMethods: [PaymentMethod]

    enum CodingKeys: String, CodingKey {
        case depositMethods = "deposit_methods"
        case withdrawMethods = "withdraw_methods"
    }
}
