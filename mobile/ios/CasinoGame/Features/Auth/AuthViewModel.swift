import Foundation

struct AuthState {
    var isLoading: Bool = false
    var isLoggedIn: Bool = false
    var user: User?
    var error: String?
}

class AuthViewModel {
    
    var onStateChange: ((AuthState) -> Void)?
    
    private(set) var state = AuthState() {
        didSet {
            onStateChange?(state)
        }
    }
    
    private let apiClient = APIClient.shared
    
    func login(email: String, password: String) {
        state.isLoading = true
        state.error = nil
        
        Task {
            do {
                let response = try await apiClient.login(email: email, password: password)
                UserDefaultsManager.shared.saveUserData(from: response)
                
                state.isLoading = false
                state.isLoggedIn = true
                state.user = response.user
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }
    
    func register(
        email: String,
        username: String,
        password: String,
        phone: String?,
        currency: String,
        referralCode: String?
    ) {
        state.isLoading = true
        state.error = nil
        
        Task {
            do {
                let response = try await apiClient.register(
                    email: email,
                    username: username,
                    password: password,
                    phone: phone,
                    currency: currency,
                    referralCode: referralCode
                )
                UserDefaultsManager.shared.saveUserData(from: response)
                
                state.isLoading = false
                state.isLoggedIn = true
                state.user = response.user
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }
    
    func logout() {
        state.isLoading = true
        
        Task {
            do {
                try await apiClient.logout()
            } catch {
                // Continue even if logout API fails
            }
            
            UserDefaultsManager.shared.clearAll()
            state = AuthState()
        }
    }
    
    func resetPassword(email: String) {
        state.isLoading = true
        state.error = nil
        
        // Implement reset password API call
        state.isLoading = false
    }
}
