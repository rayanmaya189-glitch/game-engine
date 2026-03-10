import Foundation
import Combine

// MARK: - Auth Manager
class AuthManager: ObservableObject {
    @Published var isLoggedIn: Bool = false
    @Published var currentUser: User?
    @Published var isLoading: Bool = false
    @Published var error: String?
    
    private let api = APIClient.shared
    
    init() {
        if UserDefaults.standard.string(forKey: "auth_token") != nil {
            isLoggedIn = true
        }
    }
    
    func login(email: String, password: String) async {
        await MainActor.run { isLoading = true }
        
        do {
            let response = try await api.login(email: email, password: password)
            UserDefaults.standard.set(response.token, forKey: "auth_token")
            UserDefaults.standard.set(response.refreshToken, forKey: "refresh_token")
            
            await MainActor.run {
                self.currentUser = response.user
                self.isLoggedIn = true
                self.isLoading = false
                self.error = nil
            }
        } catch {
            await MainActor.run {
                self.error = error.localizedDescription
                self.isLoading = false
            }
        }
    }
    
    func register(email: String, password: String, username: String, phone: String?) async {
        await MainActor.run { isLoading = true }
        
        do {
            let response = try await api.register(email: email, password: password, username: username, phone: phone)
            UserDefaults.standard.set(response.token, forKey: "auth_token")
            UserDefaults.standard.set(response.refreshToken, forKey: "refresh_token")
            
            await MainActor.run {
                self.currentUser = response.user
                self.isLoggedIn = true
                self.isLoading = false
                self.error = nil
            }
        } catch {
            await MainActor.run {
                self.error = error.localizedDescription
                self.isLoading = false
            }
        }
    }
    
    func logout() async {
        do {
            try await api.logout()
        } catch { }
        
        UserDefaults.standard.removeObject(forKey: "auth_token")
        UserDefaults.standard.removeObject(forKey: "refresh_token")
        
        await MainActor.run {
            self.currentUser = nil
            self.isLoggedIn = false
        }
    }
}

// MARK: - App State
class AppState: ObservableObject {
    @Published var selectedTab: Int = 0
}

// MARK: - WebSocket Service
class WebSocketService: ObservableObject {
    private var webSocketTask: URLSessionWebSocketTask?
    private let session = URLSession.shared
    @Published var isConnected: Bool = false
    @Published var messages: [String] = []
    
    func connect() {
        guard let token = UserDefaults.standard.string(forKey: "auth_token"),
              let url = URL(string: "ws://localhost:8080/ws?token=\(token)") else { return }
        
        webSocketTask = session.webSocketTask(with: url)
        webSocketTask?.resume()
        isConnected = true
        
        receiveMessage()
    }
    
    func disconnect() {
        webSocketTask?.cancel(with: .goingAway, reason: nil)
        webSocketTask = nil
        isConnected = false
    }
    
    private func receiveMessage() {
        webSocketTask?.receive { [weak self] result in
            switch result {
            case .success(let message):
                switch message {
                case .string(let text):
                    DispatchQueue.main.async {
                        self?.messages.append(text)
                    }
                default: break
                }
                self?.receiveMessage()
            case .failure:
                DispatchQueue.main.async {
                    self?.isConnected = false
                }
            }
        }
    }
    
    func send(type: String, data: [String: Any]?) {
        let message: [String: Any] = ["type": type, "data": data ?? [:]]
        if let jsonData = try? JSONSerialization.data(withJSONObject: message),
           let jsonString = String(data: jsonData, encoding: .utf8) {
            webSocketTask?.send(.string(jsonString)) { _ in }
        }
    }
}
