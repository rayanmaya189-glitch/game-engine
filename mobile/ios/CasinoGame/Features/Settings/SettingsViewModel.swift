import Foundation

struct UserSettings {
    var language: String = "en"
    var notificationsEnabled: Bool = true
    var twoFactorEnabled: Bool = false
}

struct SettingsState {
    var isLoading: Bool = false
    var user: User?
    var settings: UserSettings = UserSettings()
    var isLoggedOut: Bool = false
    var error: String?
}

class SettingsViewModel {

    var onStateChange: ((SettingsState) -> Void)?

    private(set) var state = SettingsState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadSettings() {
        state.isLoading = true
        Task {
            do {
                let user = try await apiClient.getCurrentUser()
                state.isLoading = false
                state.user = user
                state.settings.language = user.language
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func updateLanguage(_ language: String) {
        state.settings.language = language
        Task {
            do {
                let _ = try await apiClient.updateProfile(
                    firstName: nil, lastName: nil, phone: nil,
                    language: language, currency: nil
                )
            } catch { }
        }
    }

    func toggleNotifications(_ enabled: Bool) {
        state.settings.notificationsEnabled = enabled
    }

    func toggleTwoFactor(_ enabled: Bool) {
        state.settings.twoFactorEnabled = enabled
    }

    func logout() {
        Task {
            do {
                try await apiClient.logout()
            } catch { }
            UserDefaultsManager.shared.clearAll()
            state.isLoggedOut = true
        }
    }

    func deleteAccount() {
        Task {
            do {
                try await apiClient.deleteAccount()
            } catch { }
            UserDefaultsManager.shared.clearAll()
            state.isLoggedOut = true
        }
    }
}
