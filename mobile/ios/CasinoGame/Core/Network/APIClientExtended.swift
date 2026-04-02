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
