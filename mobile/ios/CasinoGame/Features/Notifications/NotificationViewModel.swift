import Foundation

enum NotificationType: String, Codable {
    case bonus
    case tournament
    case jackpot
    case system
}

struct AppNotification: Codable, Identifiable {
    let id: String
    let type: NotificationType
    let title: String
    let body: String
    let isRead: Bool
    let createdAt: String

    enum CodingKeys: String, CodingKey {
        case id, type, title, body
        case isRead = "is_read"
        case createdAt = "created_at"
    }
}

struct NotificationsResponse: Codable {
    let notifications: [AppNotification]
    let total: Int
    let unreadCount: Int

    enum CodingKeys: String, CodingKey {
        case notifications, total
        case unreadCount = "unread_count"
    }
}

struct NotificationState {
    var isLoading: Bool = false
    var notifications: [AppNotification] = []
    var unreadCount: Int = 0
    var error: String?
}

class NotificationViewModel {

    var onStateChange: ((NotificationState) -> Void)?

    private(set) var state = NotificationState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadNotifications() {
        state.isLoading = true
        Task {
            do {
                let response = try await apiClient.getNotifications()
                state.isLoading = false
                state.notifications = response.notifications
                state.unreadCount = response.unreadCount
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func markRead(_ notificationId: String) {
        Task {
            do {
                try await apiClient.markNotificationRead(id: notificationId)
                if let index = state.notifications.firstIndex(where: { $0.id == notificationId }) {
                    var updated = state.notifications[index]
                    let newNotification = AppNotification(
                        id: updated.id, type: updated.type, title: updated.title,
                        body: updated.body, isRead: true, createdAt: updated.createdAt
                    )
                    state.notifications[index] = newNotification
                    state.unreadCount = max(0, state.unreadCount - 1)
                }
            } catch { }
        }
    }

    func delete(_ notificationId: String) {
        Task {
            do {
                try await apiClient.deleteNotification(id: notificationId)
                state.notifications.removeAll { $0.id == notificationId }
            } catch { }
        }
    }
}
