import Foundation

// MARK: - Extended API Methods
extension APIClient {

    func sendChatMessage(text: String) async throws -> ChatMessage {
        let parameters: [String: Any] = ["text": text]
        return try await request(endpoint: "chat/messages", method: .post, parameters: parameters)
    }

    func getOnlineUsers() async throws -> OnlineUsersResponse {
        return try await request(endpoint: "chat/online-users", method: .get)
    }

    // MARK: - Notifications

    func getNotifications() async throws -> NotificationsResponse {
        return try await request(endpoint: "notifications", method: .get)
    }

    func markNotificationRead(id: String) async throws {
        let _: EmptyResponse = try await request(endpoint: "notifications/\(id)/read", method: .put)
    }

    func deleteNotification(id: String) async throws {
        let _: EmptyResponse = try await request(endpoint: "notifications/\(id)", method: .delete)
    }

    // MARK: - Referral

    func getReferralCode() async throws -> ReferralCodeResponse {
        return try await request(endpoint: "referral/code", method: .get)
    }

    func getReferralStats() async throws -> ReferralStats {
        return try await request(endpoint: "referral/stats", method: .get)
    }

    func getReferralHistory() async throws -> ReferralHistoryResponse {
        return try await request(endpoint: "referral/history", method: .get)
    }

    // MARK: - Live Dealer

    func getLiveDealerTables() async throws -> LiveDealerTablesResponse {
        return try await request(endpoint: "live-dealer/tables", method: .get)
    }

    func joinLiveDealerTable(id: String) async throws -> GameSession {
        return try await request(endpoint: "live-dealer/tables/\(id)/join", method: .post)
    }

    // MARK: - Account

    func deleteAccount() async throws {
        let _: EmptyResponse = try await request(endpoint: "auth/account", method: .delete)
    }

    // MARK: - Payment History

    func getPaymentHistory(type: String? = nil, page: Int = 1, pageSize: Int = 20) async throws -> TransactionListResponse {
        var parameters: [String: Any] = [
            "page": page,
            "page_size": pageSize
        ]
        if let type = type { parameters["type"] = type }
        return try await request(endpoint: "payments/history", method: .get, parameters: parameters)
    }

    // MARK: - KYC

    func getKycStatus() async throws -> KycStatus {
        return try await request(endpoint: "kyc/status", method: .get)
    }

    func uploadKycDocument(type: String, fileName: String, data: Data) async throws -> KycUploadResponse {
        let parameters: [String: Any] = [
            "type": type,
            "file_name": fileName,
            "file_data": data.base64EncodedString()
        ]
        return try await request(endpoint: "kyc/documents", method: .post, parameters: parameters)
    }

    func submitKycVerification() async throws -> KycVerificationResponse {
        return try await request(endpoint: "kyc/submit", method: .post)
    }

    // MARK: - Game Detail

    func getGameDetail(id: String) async throws -> GameDetailResponse {
        return try await request(endpoint: "games/\(id)/detail", method: .get)
    }

    // MARK: - Bet History

    func getBetHistory(result: String? = nil, page: Int = 1, pageSize: Int = 20) async throws -> BetListResponse {
        var parameters: [String: Any] = [
            "page": page,
            "page_size": pageSize
        ]
        if let result = result { parameters["result"] = result }
        return try await request(endpoint: "bets/history", method: .get, parameters: parameters)
    }

    // MARK: - Support

    func getFaq() async throws -> FaqResponse {
        return try await request(endpoint: "support/faq", method: .get)
    }

    func getTickets() async throws -> TicketsResponse {
        return try await request(endpoint: "support/tickets", method: .get)
    }

    func createTicket(subject: String, message: String) async throws -> SupportTicket {
        let parameters: [String: Any] = ["subject": subject, "message": message]
        return try await request(endpoint: "support/tickets", method: .post, parameters: parameters)
    }

    // MARK: - Leaderboard

    func getLeaderboard(period: String) async throws -> LeaderboardResponse {
        let parameters: [String: Any] = ["period": period]
        return try await request(endpoint: "leaderboard", method: .get, parameters: parameters)
    }

    // MARK: - Responsible Gaming

    func getGamingLimits() async throws -> GamingLimitsResponse {
        return try await request(endpoint: "responsible-gaming/limits", method: .get)
    }

    func updateGamingLimits(limits: GamingLimits) async throws -> GamingLimits {
        let parameters: [String: Any] = [
            "deposit_limit": limits.depositLimit,
            "loss_limit": limits.lossLimit,
            "session_limit": limits.sessionLimit,
            "cool_off_enabled": limits.coolOffEnabled,
            "self_exclusion_days": limits.selfExclusionDays
        ]
        return try await request(endpoint: "responsible-gaming/limits", method: .put, parameters: parameters)
    }

    // MARK: - Private Methods
    
    private func request<T: Decodable>(
        endpoint: String,
        method: HTTPMethod,
        parameters: [String: Any]? = nil
    ) async throws -> T {
        let url = baseURL + endpoint
        
        return try await withCheckedThrowingContinuation { continuation in
            session.request(
                url,
                method: method,
                parameters: parameters,
                encoding: JSONEncoding.default
            )
            .validate()
            .responseDecodable(of: T.self) { response in
                switch response.result {
                case .success(let value):
                    continuation.resume(returning: value)
                case .failure(let error):
                    if let data = response.data,
                       let errorResponse = try? JSONDecoder().decode(ErrorResponse.self, from: data) {
                        continuation.resume(throwing: APIError(message: errorResponse.message))
                    } else {
                        continuation.resume(throwing: error)
                    }
                }
            }
        }
    }
}

// MARK: - Auth Interceptor

class AuthInterceptor: RequestInterceptor {
    
    func adapt(_ urlRequest: URLRequest, for session: Session, completion: @escaping (Result<URLRequest, Error>) -> Void) {
        var request = urlRequest
        
        if let token = UserDefaultsManager.shared.accessToken {
            request.setValue("Bearer \(token)", forHTTPHeaderField: "Authorization")
        }
        
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        request.setValue("application/json", forHTTPHeaderField: "Accept")
        
        completion(.success(request))
    }
}

// MARK: - API Error

struct ErrorResponse: Decodable {
    let message: String
    let code: String?
}

enum APIError: LocalizedError {
    case message(String)
    
    var errorDescription: String? {
        switch self {
        case .message(let message):
            return message
        }
    }
}

// MARK: - Empty Response

struct EmptyResponse: Decodable {}
