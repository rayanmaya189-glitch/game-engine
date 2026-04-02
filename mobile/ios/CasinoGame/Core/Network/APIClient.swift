import Foundation
import Alamofire
import Combine

class APIClient {
    
    static let shared = APIClient()
    
    private let baseURL: String
    private let session: Session
    
    private init() {
        // Use environment configuration for API base URL
        #if DEBUG
        baseURL = ProcessInfo.processInfo.environment["API_BASE_URL"] ?? "https://sandbox.api.casino-game.engine/v1/"
        #else
        baseURL = ProcessInfo.processInfo.environment["API_BASE_URL"] ?? "https://api.casino-game.engine/v1/"
        #endif
        
        let interceptor = AuthInterceptor()
        session = Session(interceptor: interceptor)
    }
    
    // MARK: - Auth
    
    func login(email: String, password: String) async throws -> LoginResponse {
        let parameters: [String: Any] = [
            "email": email,
            "password": password
        ]
        
        return try await request(endpoint: "auth/login", method: .post, parameters: parameters)
    }
    
    func register(
        email: String,
        username: String,
        password: String,
        phone: String?,
        currency: String,
        referralCode: String?
    ) async throws -> RegisterResponse {
        var parameters: [String: Any] = [
            "email": email,
            "username": username,
            "password": password,
            "currency": currency
        ]
        
        if let phone = phone { parameters["phone"] = phone }
        if let referralCode = referralCode { parameters["referral_code"] = referralCode }
        
        return try await request(endpoint: "auth/register", method: .post, parameters: parameters)
    }
    
    func logout() async throws {
        let _: EmptyResponse = try await request(endpoint: "auth/logout", method: .post)
    }
    
    func getCurrentUser() async throws -> User {
        return try await request(endpoint: "auth/me", method: .get)
    }
    
    func updateProfile(
        firstName: String?,
        lastName: String?,
        phone: String?,
        language: String?,
        currency: String?
    ) async throws -> User {
        var parameters: [String: Any] = [:]
        
        if let firstName = firstName { parameters["first_name"] = firstName }
        if let lastName = lastName { parameters["last_name"] = lastName }
        if let phone = phone { parameters["phone"] = phone }
        if let language = language { parameters["language"] = language }
        if let currency = currency { parameters["currency"] = currency }
        
        return try await request(endpoint: "auth/profile", method: .put, parameters: parameters)
    }
    
    // MARK: - Wallet
    
    func getWallet() async throws -> Wallet {
        return try await request(endpoint: "wallet", method: .get)
    }
    
    func getBalance() async throws -> WalletBalance {
        return try await request(endpoint: "wallet/balance", method: .get)
    }
    
    func deposit(amount: Double, paymentMethod: String, paymentId: String?, currency: String) async throws -> DepositResponse {
        var parameters: [String: Any] = [
            "amount": amount,
            "payment_method": paymentMethod,
            "currency": currency
        ]
        
        if let paymentId = paymentId { parameters["payment_id"] = paymentId }
        
        return try await request(endpoint: "wallet/deposit", method: .post, parameters: parameters)
    }
    
    func withdraw(amount: Double, paymentMethod: String, paymentDetails: String, currency: String) async throws -> WithdrawResponse {
        let parameters: [String: Any] = [
            "amount": amount,
            "payment_method": paymentMethod,
            "payment_details": paymentDetails,
            "currency": currency
        ]
        
        return try await request(endpoint: "wallet/withdraw", method: .post, parameters: parameters)
    }
    
    func getTransactions(page: Int = 1, pageSize: Int = 20, type: String? = nil, status: String? = nil) async throws -> TransactionListResponse {
        var parameters: [String: Any] = [
            "page": page,
            "page_size": pageSize
        ]
        
        if let type = type { parameters["type"] = type }
        if let status = status { parameters["status"] = status }
        
        return try await request(endpoint: "wallet/transactions", method: .get, parameters: parameters)
    }
    
    func getPaymentMethods() async throws -> PaymentMethodsResponse {
        return try await request(endpoint: "wallet/payment-methods", method: .get)
    }
    
    // MARK: - Games
    
    func getGames(page: Int = 1, pageSize: Int = 20, category: String? = nil, provider: String? = nil, search: String? = nil) async throws -> GameListResponse {
        var parameters: [String: Any] = [
            "page": page,
            "page_size": pageSize
        ]
        
        if let category = category { parameters["category"] = category }
        if let provider = provider { parameters["provider"] = provider }
        if let search = search { parameters["search"] = search }
        
        return try await request(endpoint: "games", method: .get, parameters: parameters)
    }
    
    func getFeaturedGames() async throws -> FeaturedGamesResponse {
        return try await request(endpoint: "games/featured", method: .get)
    }
    
    func getCategories() async throws -> CategoriesResponse {
        return try await request(endpoint: "games/categories", method: .get)
    }
    
    func getGame(id: String) async throws -> Game {
        return try await request(endpoint: "games/\(id)", method: .get)
    }
    
    func startGame(id: String, isFunPlay: Bool = false) async throws -> GameSession {
        let parameters: [String: Any] = [
            "game_id": id,
            "is_fun_play": isFunPlay
        ]
        
        return try await request(endpoint: "games/\(id)/start", method: .post, parameters: parameters)
    }
    
    func addToFavorites(gameId: String) async throws {
        let _: EmptyResponse = try await request(endpoint: "games/\(gameId)/favorite", method: .post)
    }
    
    func removeFromFavorites(gameId: String) async throws {
        let _: EmptyResponse = try await request(endpoint: "games/\(gameId)/favorite", method: .delete)
    }

    func spinGame(betAmount: Double) async throws -> SpinResult {
        let parameters: [String: Any] = ["bet_amount": betAmount]
        return try await request(endpoint: "games/slots/spin", method: .post, parameters: parameters)
    }

    // MARK: - Jackpot

    func getJackpotWinners() async throws -> [JackpotWinner] {
        return try await request(endpoint: "jackpot/winners", method: .get)
    }

    // MARK: - Bonus

    func getBonuses() async throws -> [Bonus] {
        return try await request(endpoint: "bonuses", method: .get)
    }

    func claimBonus(id: String) async throws {
        let _: EmptyResponse = try await request(endpoint: "bonuses/\(id)/claim", method: .post)
    }

    // MARK: - Chat

    func getChatMessages() async throws -> ChatMessagesResponse {
        return try await request(endpoint: "chat/messages", method: .get)
    }