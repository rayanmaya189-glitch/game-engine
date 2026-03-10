import Foundation

class UserDefaultsManager {
    
    static let shared = UserDefaultsManager()
    
    private let defaults = UserDefaults.standard
    
    private enum Keys {
        static let accessToken = "access_token"
        static let refreshToken = "refresh_token"
        static let userId = "user_id"
        static let userEmail = "user_email"
        static let username = "username"
        static let currency = "currency"
        static let language = "language"
        static let biometricEnabled = "biometric_enabled"
    }
    
    private init() {}
    
    // MARK: - Access Token
    
    var accessToken: String? {
        get { defaults.string(forKey: Keys.accessToken) }
        set { defaults.set(newValue, forKey: Keys.accessToken) }
    }
    
    // MARK: - Refresh Token
    
    var refreshToken: String? {
        get { defaults.string(forKey: Keys.refreshToken) }
        set { defaults.set(newValue, forKey: Keys.refreshToken) }
    }
    
    // MARK: - User ID
    
    var userId: String? {
        get { defaults.string(forKey: Keys.userId) }
        set { defaults.set(newValue, forKey: Keys.userId) }
    }
    
    // MARK: - User Email
    
    var userEmail: String? {
        get { defaults.string(forKey: Keys.userEmail) }
        set { defaults.set(newValue, forKey: Keys.userEmail) }
    }
    
    // MARK: - Username
    
    var username: String? {
        get { defaults.string(forKey: Keys.username) }
        set { defaults.set(newValue, forKey: Keys.username) }
    }
    
    // MARK: - Currency
    
    var currency: String? {
        get { defaults.string(forKey: Keys.currency) }
        set { defaults.set(newValue, forKey: Keys.currency) }
    }
    
    // MARK: - Language
    
    var language: String? {
        get { defaults.string(forKey: Keys.language) }
        set { defaults.set(newValue, forKey: Keys.language) }
    }
    
    // MARK: - Biometric Enabled
    
    var biometricEnabled: Bool {
        get { defaults.bool(forKey: Keys.biometricEnabled) }
        set { defaults.set(newValue, forKey: Keys.biometricEnabled) }
    }
    
    // MARK: - Helpers
    
    var isLoggedIn: Bool {
        return accessToken != nil && !accessToken!.isEmpty
    }
    
    func saveUserData(from response: LoginResponse) {
        accessToken = response.accessToken
        refreshToken = response.refreshToken
        userId = response.user.id
        userEmail = response.user.email
        username = response.user.username
        currency = response.user.currency
    }
    
    func clearAll() {
        accessToken = nil
        refreshToken = nil
        userId = nil
        userEmail = nil
        username = nil
        currency = nil
        language = nil
    }
}
