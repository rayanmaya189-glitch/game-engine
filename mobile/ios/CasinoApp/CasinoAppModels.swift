import Foundation

// MARK: - API Client
class APIClient {
    static let shared = APIClient()

    private let baseURL = "http://localhost:8080/api/v1"
    private let session: URLSession

    private init() {
        let config = URLSessionConfiguration.default
        config.timeoutIntervalForRequest = 30
        session = URLSession(configuration: config)
    }

    private var authToken: String? {
        UserDefaults.standard.string(forKey: "auth_token")
    }

    private func createRequest(endpoint: String, method: String = "GET", body: Data? = nil) -> URLRequest? {
        guard let url = URL(string: "\(baseURL)\(endpoint)") else { return nil }
        var request = URLRequest(url: url)
        request.httpMethod = method
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        if let token = authToken {
            request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        }
        request.httpBody = body
        return request
    }

    // MARK: - Auth
    func login(email: String, password: String) async throws -> LoginResponse {
        let body = try? JSONEncoder().encode(["email": email, "password": password])
        guard let request = createRequest(endpoint: "/player/auth/login", method: "POST", body: body) else {
            throw APIError.invalidRequest
        }
        let (data, response) = try await session.data(for: request)
        return try JSONDecoder().decode(LoginResponse.self, from: data)
    }

    func register(email: String, password: String, username: String, phone: String?) async throws -> LoginResponse {
        var dict: [String: Any] = ["email": email, "password": password, "username": username]
        if let phone = phone { dict["phone"] = phone }
        let body = try? JSONSerialization.data(withJSONObject: dict)
        guard let request = createRequest(endpoint: "/player/auth/register", method: "POST", body: body) else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(LoginResponse.self, from: data)
    }

    func logout() async throws {
        guard let request = createRequest(endpoint: "/player/auth/logout", method: "POST") else {
            throw APIError.invalidRequest
        }
        _ = try await session.data(for: request)
    }

    // MARK: - Profile
    func getProfile() async throws -> UserProfile {
        guard let request = createRequest(endpoint: "/player/profile") else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(UserProfile.self, from: data)
    }

    // MARK: - Wallet
    func getBalance() async throws -> BalanceResponse {
        guard let request = createRequest(endpoint: "/player/wallet/balance") else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(BalanceResponse.self, from: data)
    }

    func getTransactions(page: Int = 1, limit: Int = 20) async throws -> TransactionsResponse {
        guard let request = createRequest(endpoint: "/player/wallet/transactions?page=\(page)&limit=\(limit)") else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(TransactionsResponse.self, from: data)
    }

    // MARK: - Games
    func getGames(category: String? = nil, search: String? = nil, page: Int = 1) async throws -> GamesResponse {
        var endpoint = "/player/games?page=\(page)"
        if let category = category { endpoint += "&category=\(category)" }
        if let search = search { endpoint += "&search=\(search)" }
        guard let request = createRequest(endpoint: endpoint) else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(GamesResponse.self, from: data)
    }

    func getGameDetails(id: String) async throws -> GameDetails {
        guard let request = createRequest(endpoint: "/player/games/\(id)") else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(GameDetails.self, from: data)
    }

    func getCategories() async throws -> CategoriesResponse {
        guard let request = createRequest(endpoint: "/player/games/categories") else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(CategoriesResponse.self, from: data)
    }

    func getFeaturedGames() async throws -> GamesResponse {
        guard let request = createRequest(endpoint: "/player/games/featured") else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(GamesResponse.self, from: data)
    }

    // MARK: - Tournaments
    func getTournaments(status: String? = nil) async throws -> TournamentsResponse {
        var endpoint = "/player/tournaments"
        if let status = status { endpoint += "?status=\(status)" }
        guard let request = createRequest(endpoint: endpoint) else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(TournamentsResponse.self, from: data)
    }

    // MARK: - Jackpots
    func getJackpots() async throws -> JackpotsResponse {
        guard let request = createRequest(endpoint: "/player/jackpots") else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(JackpotsResponse.self, from: data)
    }

    // MARK: - Bonuses
    func getBonuses() async throws -> BonusesResponse {
        guard let request = createRequest(endpoint: "/player/bonuses") else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(BonusesResponse.self, from: data)
    }

    func claimBonus(id: String) async throws -> ClaimBonusResponse {
        guard let request = createRequest(endpoint: "/player/bonuses/\(id)/claim", method: "POST") else {
            throw APIError.invalidRequest
        }
        let (data, _) = try await session.data(for: request)
        return try JSONDecoder().decode(ClaimBonusResponse.self, from: data)
    }
}

enum APIError: Error {
    case invalidRequest
    case invalidResponse
    case unauthorized
}

// MARK: - Models
struct LoginResponse: Codable {
    let user: User
    let token: String
    let refreshToken: String

    enum CodingKeys: String, CodingKey {
        case user, token
        case refreshToken = "refresh_token"
    }
}

struct User: Codable, Identifiable {
    let id: String
    let username: String
    let email: String
    let status: String
    let kycLevel: String
    let createdAt: String

    enum CodingKeys: String, CodingKey {
        case id, username, email, status
        case kycLevel = "kyc_level"
        case createdAt = "created_at"
    }
}

struct UserProfile: Codable {
    let id: String
    let username: String
    let email: String
    let phone: String?
    let status: String
    let kycLevel: String
    let createdAt: String
    let avatarUrl: String?

    enum CodingKeys: String, CodingKey {
        case id, username, email, phone, status
        case kycLevel = "kyc_level"
        case createdAt = "created_at"
        case avatarUrl = "avatar_url"
    }
}
