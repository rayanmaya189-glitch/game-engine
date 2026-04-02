import Foundation

struct ChatMessage: Codable, Identifiable {
    let id: String
    let senderId: String
    let senderName: String
    let text: String
    let createdAt: String

    enum CodingKeys: String, CodingKey {
        case id
        case senderId = "sender_id"
        case senderName = "sender_name"
        case text
        case createdAt = "created_at"
    }
}

struct ChatMessagesResponse: Codable {
    let messages: [ChatMessage]
    let total: Int
}

struct OnlineUser: Codable {
    let id: String
    let username: String
    let avatarUrl: String?

    enum CodingKeys: String, CodingKey {
        case id, username
        case avatarUrl = "avatar_url"
    }
}

struct OnlineUsersResponse: Codable {
    let users: [OnlineUser]
    let count: Int
}

struct ChatState {
    var isLoading: Bool = false
    var messages: [ChatMessage] = []
    var onlineUsers: [OnlineUser] = []
    var error: String?
}

class ChatViewModel {

    var onStateChange: ((ChatState) -> Void)?

    private(set) var state = ChatState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadMessages() {
        state.isLoading = true
        Task {
            do {
                let response = try await apiClient.getChatMessages()
                state.isLoading = false
                state.messages = response.messages
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func loadOnlineUsers() {
        Task {
            do {
                let response = try await apiClient.getOnlineUsers()
                state.onlineUsers = response.users
            } catch { }
        }
    }

    func sendMessage(_ text: String) {
        guard !text.isEmpty else { return }
        Task {
            do {
                let message = try await apiClient.sendChatMessage(text: text)
                state.messages.append(message)
            } catch let error as APIError {
                state.error = error.errorDescription
            } catch {
                state.error = error.localizedDescription
            }
        }
    }
}
