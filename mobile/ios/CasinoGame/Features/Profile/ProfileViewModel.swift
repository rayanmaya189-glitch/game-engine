import Foundation

struct ProfileState {
    var isLoading: Bool = false
    var user: User?
    var isLoggedOut: Bool = false
    var error: String?
}

class ProfileViewModel {
    
    var onStateChange: ((ProfileState) -> Void)?
    
    private(set) var state = ProfileState() {
        didSet {
            onStateChange?(state)
        }
    }
    
    private let apiClient = APIClient.shared
    
    func loadUser() {
        state.isLoading = true
        
        Task {
            do {
                let user = try await apiClient.getCurrentUser()
                state.isLoading = false
                state.user = user
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
            }
        }
    }
    
    func logout() {
        Task {
            do {
                try await apiClient.logout()
            } catch {
                // Continue even if logout API fails
            }
            
            UserDefaultsManager.shared.clearAll()
            state.isLoggedOut = true
        }
    }
}
