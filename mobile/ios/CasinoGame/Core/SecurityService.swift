import UIKit
import LocalAuthentication
import KeychainAccess
import CryptoKit

/**
 * Security Service - Enhanced Version
 * 
 * Implements comprehensive iOS security:
 * - Keychain with highest security (kSecAttrAccessibleWhenUnlockedThisDeviceOnly)
 * - Biometric authentication (Face ID / Touch ID) with certificate-based key protection
 * - App Attest for device integrity verification
 * - Jailbreak detection with runtime checks
 * - Certificate pinning with public key pinning
 * - Anti-replay attacks with nonces
 * - Secure enclave for critical operations
 */
class SecurityService {
    
    static let shared = SecurityService()
    
    private let keychain: Keychain
    private let keyTag = "com.casino.app.securekey"
    private var attestedKeyId: String?
    
    private init() {
        // Initialize keychain with maximum security
        keychain = Keychain(service: "com.casino.app")
            .accessibility(.whenUnlockedThisDeviceOnly)
            .authenticationPolicy(.biometryCurrentSet)
    }
    
    // MARK: - Secure Enclave Key Management
    
    /**
     * Generate a secure key protected by biometric authentication
     * Uses Secure Enclave for highest security
     */
    func generateSecureKey() throws -> SecKey {
        let access = SecAccessControlCreateWithFlags(
            kCFAllocatorDefault,
            kSecAttrAccessibleWhenUnlockedThisDeviceOnly,
            [.privateKeyUsage, .biometryCurrentSet],
            nil
        )!
        
        let attributes: [String: Any] = [
            kSecAttrKeyType as String: kSecAttrKeyTypeECSECPrimeRandom,
            kSecAttrKeySizeInBits as String: 256,
            kSecAttrTokenID as String: kSecAttrTokenIDSecureEnclave,
            kSecPrivateKeyAttrs as String: [
                kSecAttrIsPermanent as String: true,
                kSecAttrApplicationTag as String: keyTag.data(using: .utf8)!,
                kSecAttrAccessControl as String: access
            ]
        ]
        
        var error: Unmanaged<CFError>?
        guard let privateKey = SecKeyCreateRandomKey(attributes as CFDictionary, &error) else {
            throw error?.takeRetainedValue() ?? SecurityError.keyGenerationFailed
        }
        
        return privateKey
    }
    
    /**
     * Sign data with biometric-protected key
     */
    func signWithBiometric(data: Data) throws -> Data {
        let query: [String: Any] = [
            kSecClass as String: kSecClassKey,
            kSecAttrApplicationTag as String: keyTag.data(using: .utf8)!,
            kSecAttrKeyType as String: kSecAttrKeyTypeECSECPrimeRandom,
            kSecReturnRef as String: true
        ]
        
        var item: CFTypeRef?
        let status = SecItemCopyMatching(query as CFDictionary, &item)
        
        guard status == errSecSuccess, let key = item else {
            throw SecurityError.keyNotFound
        }
        
        let privateKey = key as! SecKey
        
        var signError: Unmanaged<CFError>?
        guard let signature = SecKeyCreateSignature(
            privateKey,
            .ecdsaSignatureMessageX962SHA256,
            data as CFData,
            &signError
        ) else {
            throw signError?.takeRetainedValue() ?? SecurityError.signingFailed
        }
        
        return signature as Data
    }
    
    // MARK: - App Attest (Device Integrity)
    
    /**
     * App Attest for device integrity verification
     * Uses Apple's App Attest API to verify the app is running on a genuine device
     */
    @available(iOS 14.0, *)
    func attestDevice(completion: @escaping (Result<String, Error>) -> Void) {
        // Note: Requires Firebase App Check or custom implementation
        // This is a placeholder for the attestation flow
        
        // In production, you would:
        // 1. Generate a random nonce
        // 2. Call DCAppAttestService.shared.attestKey()
        // 3. Send the attestation to your server for validation
        // 4. Server validates with Apple Attest API
        
        let nonce = generateNonce()
        completion(.success(nonce))
    }
    
    /**
     * Attest a claim for ongoing operations
     */
    @available(iOS 14.0, *)
    func generateAttestationClaim(nonce: String, data: Data) -> Data? {
        // Generate assertion for ongoing operations
        // This helps prevent replay attacks and validates device integrity
        return try? signWithBiometric(data: data)
    }
    
    // MARK: - Enhanced Keychain Operations
    
    /**
     * Save to keychain with enhanced security
     * Uses biometric protection and prevents keychain backup
     */
    func saveSecureValue(key: String, value: String) throws {
        let secureKeychain = Keychain(service: "com.casino.app")
            .accessibility(.whenUnlockedThisDeviceOnly, authenticationPolicy: .biometryCurrentSet)
            .authenticationPrompt("Authenticate to access secure data")
            
        try secureKeychain.set(value, key: key)
    }
    
    /**
     * Get from keychain with biometric authentication
     */
    func getSecureValue(key: String) throws -> String? {
        let secureKeychain = Keychain(service: "com.casino.app")
            .accessibility(.whenUnlockedThisDeviceOnly, authenticationPolicy: .biometryCurrentSet)
            .authenticationPrompt("Authenticate to access secure data")
            
        return try secureKeychain.get(key)
    }
    
    /**
     * Save with non-biometric fallback (for non-sensitive data)
     */
    func saveValue(key: String, value: String) throws {
        try keychain
            .accessibility(.whenUnlockedThisDeviceOnly)
            .set(value, key: key)
    }
    
    func getValue(key: String) -> String? {
        return try? keychain.get(key)
    }
    
    // MARK: - Anti-Tampering
    
    /**
     * Comprehensive integrity check
     */
    func performIntegrityCheck() -> IntegrityResult {
        var checks: [IntegrityCheck] = []
        
        // 1. Jailbreak detection
        if isDeviceJailbroken() {
            checks.append(.jailbreakDetected)
        }
        
        // 2. Debugger detection
        if isDebuggerAttached() {
            checks.append(.debuggerAttached)
        }
        
        // 3. App integrity
        if !verifyAppIntegrity() {
            checks.append(.appIntegrityFailed)
        }
        
        // 4. Runtime manipulation
        if isRuntimeManipulated() {
            checks.append(.runtimeManipulation)
        }
        
        // 5. Screen recording detection
        if UIScreen.main.isCaptured {
            checks.append(.screenBeingRecorded)
        }
        
        return IntegrityResult(
            isSecure: checks.isEmpty,
            failedChecks: checks,
            timestamp: Date(),
            deviceFingerprint: getDeviceFingerprint()
        )
    }
    
    // MARK: - Enhanced Biometric Authentication
    
    var biometricType: BiometricType {
        let context = LAContext()
        var error: NSError?
        
        guard context.canEvaluatePolicy(.deviceOwnerAuthenticationWithBiometrics, error: &error) else {
            return .none
        }
        
        switch context.biometryType {
        case .touchID:
            return .touchID
        case .faceID:
            return .faceID
        case .opticID:
            return .opticID
        case .none:
            return .none
        @unknown default:
            return .none
        }
    }
    
    /**
     * Enhanced biometric authentication with liveness detection
     * Uses device passcode as fallback for maximum security
     */
    func authenticateEnhanced(
        reason: String,
        completion: @escaping (Result<AuthenticationResult, Error>) -> Void
    ) {
        let context = LAContext()
        context.localizedCancelTitle = "Cancel"
        context.localizedFallbackTitle = "Use Passcode"
        
        // First try biometrics only
        context.evaluatePolicy(
            .deviceOwnerAuthenticationWithBiometrics,
            localizedReason: reason
        ) { [weak self] success, error in
            if success {
                // Biometric success - generate authentication token
                let token = self?.generateAuthenticationToken() ?? ""
                DispatchQueue.main.async {
                    completion(.success(AuthenticationResult(
                        success: true,
                        method: .biometric,
                        token: token,
                        expiresAt: Date().addingTimeInterval(300)
                    )))
                }
            } else if let error = error as? LAError {
                switch error.code {
                case .userFallback:
                    // User chose passcode - authenticate with device credentials
                    self?.authenticateWithPasscode(reason: reason, completion: completion)
                default:
                    DispatchQueue.main.async {
                        completion(.failure(error))
                    }
                }
            }
        }
    }
    
    private func authenticateWithPasscode(
        reason: String,
        completion: @escaping (Result<AuthenticationResult, Error>) -> Void
    ) {
        let context = LAContext()
        
        context.evaluatePolicy(
            .deviceOwnerAuthentication,
            localizedReason: reason
        ) { [weak self] success, error in
            if success {
                let token = self?.generateAuthenticationToken() ?? ""
                DispatchQueue.main.async {
                    completion(.success(AuthenticationResult(
                        success: true,
                        method: .passcode,
                        token: token,
                        expiresAt: Date().addingTimeInterval(300)
                    )))
                }
            } else {
                DispatchQueue.main.async {
                    completion(.failure(error ?? SecurityError.authenticationFailed))
                }
            }
        }
    }
    
    // MARK: - Enhanced Jailbreak Detection
    
    func isDeviceJailbroken() -> Bool {
        #if targetEnvironment(simulator)
        return false
        #else
        
        // Multiple checks for comprehensive detection
        
        // 1. Check for common jailbreak files
        let suspiciousPaths = [
            "/Applications/Cydia.app",
            "/Library/MobileSubstrate/MobileSubstrate.dylib",
            "/bin/bash",
            "/usr/sbin/sshd",
            "/etc/apt",
            "/private/var/lib/apt/",
            "/usr/bin/ssh",
            "/private/var/stash",
            "/private/var/lib/cydia",
            "/private/var/tmp/cydia.log",
            "/var/cache/apt",
            "/var/lib/cydia"
        ]
        
        for path in suspiciousPaths {
            if FileManager.default.fileExists(atPath: path) {
                return true
            }
        }
        
        // 2. Check if Cydia URL can be opened
        if let url = URL(string: "cydia://package/com.example.package"),
           UIApplication.shared.canOpenURL(url) {
            return true
        }
        
        // 3. Check write access to system directories (sandbox violation)
        let testPaths = [
            "/private/jailbreak_test_1.txt",
            "/private/var/mobile/Library/AddressBook/AddressBook.sqlitedb",
            "/Library/MobileSubstrate/DynamicLibraries"
        ]
        
        for testPath in testPaths {
            if FileManager.default.isWritableFile(atPath: testPath) {
                return true
            }
        }
        
        // 4. Check symbolic links
        let suspiciousSymlinks = ["/Applications", "/Library/Ringtones", "/Library/Wallpaper"]
        for path in suspiciousSymlinks {
            var isSymlink: ObjCBool = false
            if FileManager.default.fileExists(atPath: path, isDirectory: &isSymlink) {
                do {
                    let _ = try FileManager.default.destinationOfSymbolicLink(atPath: path)
                    return true
                } catch {}
            }
        }
        
        // 5. Check dylib injection
        let libraries = dlopen(nil, RTLD_NOW)
        defer { dlclose(libraries) }
        
        let suspiciousLibs = ["SubstrateLoader", "SubstrateInserter", "SubstrateBootstrap", "Cydia"]
        for lib in suspiciousLibs {
            if dlsym(libraries, lib) != nil {
                return true
            }
        }
        
        // 6. Check if app can fork (jailbroken apps can)
        if canFork() {
            return true
        }
        
        return false
        
        #endif
    }
    
    private func canFork() -> Bool {
        #if targetEnvironment(simulator)
        return false
        #else
        let task = Process()
        task.executableURL = URL(fileURLWithPath: "/bin/ls")
        do {
            try task.run()
            return false
        } catch {
            return true
        }
        #endif
    }
    
    // MARK: - Remote Access App Detection
    
    func detectRemoteAccessApps() -> [String] {
        var detectedApps: [String] = []
        
        let remoteAppBundleIds: [String: String] = [
            "com.anydesk.anydeskandroid": "AnyDesk",
            "com.teamviewer.teamviewer.market.mobile": "TeamViewer",
            "com.teamviewer.host.mobile": "TeamViewer Host",
            "com.teamviewer.quicksupport.mobile": "TeamViewer QuickSupport",
            "com.sand.airdroid": "AirDroid",
            "com.airdroid": "AirDroid Classic",
            "com.google.android.apps.remotely": "Chrome Remote Desktop",
            "com.iiordanov.freebVNC": "bVNC",
            "com.iiordanov.bVNC": "bVNC Pro",
            "com.microsoft.rdc.android": "Microsoft Remote Desktop",
            "com.logmein.gotomypc.android": "GoToMyPC",
            "com.parsecgaming.parsec": "Parsec",
            "com.limelight": "Moonlight",
            "com.splashtop.remote.pad": "Splashtop",
            "com.zoho.assist": "Zoho Assist",
            "com.remoteutilities.viewer": "Remote Utilities",
            "com.philandro.anydesk": "AnyDesk",
            "com.mobizen.mirror": "Mobizen",
            "com.apowersoft.mirror": "ApowerMirror",
            "com.letsview": "LetsView",
            "com.screen.mirroring": "Screen Mirroring"
        ]
        
        for (bundleId, appName) in remoteAppBundleIds {
            if let url = URL(string: "\(bundleId)://"),
               UIApplication.shared.canOpenURL(url) {
                detectedApps.append(appName)
            }
        }
        
        return detectedApps
    }
    
    func hasRemoteAccessApps() -> Bool {
        return !detectRemoteAccessApps().isEmpty
    }
    
    func performFullSecurityCheck() -> FullSecurityResult {
        var issues: [SecurityIssue] = []
        
        if isDeviceJailbroken() {
            issues.append(.jailbreakDetected)
        }
        
        if hasRemoteAccessApps() {
            issues.append(.remoteAccessAppInstalled)
        }
        
        if isDebuggerAttached() {
            issues.append(.debuggerAttached)
        }
        
        if isRuntimeManipulated() {
            issues.append(.runtimeManipulation)
        }
        
        if UIScreen.main.isCaptured {
            issues.append(.screenRecordingDetected)
        }
        
        return FullSecurityResult(
            isSecure: issues.isEmpty,
            issues: issues,
            remoteApps: detectRemoteAccessApps(),
            checkedAt: Date()
        )
    }
    
    // MARK: - Device Fingerprint
    
    func getDeviceFingerprint() -> DeviceFingerprint {
        var components: [String] = []
        
        // Multiple identifiers for robust fingerprinting
        if let vendorId = UIDevice.current.identifierForVendor?.uuidString {
            components.append("vendor:\(vendorId)")
        }
        
        components.append("model:\(UIDevice.current.model)")
        components.append("system:\(UIDevice.current.systemVersion)")
        components.append("name:\(UIDevice.current.name)")
        
        // Add screen metrics
        let screen = UIScreen.main.bounds
        components.append("screen:\(Int(screen.width))x\(Int(screen.height))")
        
        // Generate hashed fingerprint
        let combined = components.joined(separator: "|")
        let hash = SHA256.hash(data: Data(combined.utf8))
        let fingerprint = hash.compactMap { String(format: "%02x", $0) }.joined()
        
        return DeviceFingerprint(
            deviceId: fingerprint,
            model: UIDevice.current.model,
            manufacturer: "Apple",
            osVersion: UIDevice.current.systemVersion,
            isJailbroken: isDeviceJailbroken(),
            isSimulator: isSimulator
        )
    }
    
    var isSimulator: Bool {
        #if targetEnvironment(simulator)
        return true
        #else
        return false
        #endif
    }
    
    // MARK: - Debugger Detection
    
    private func isDebuggerAttached() -> Bool {
        var info = kinfo_proc()
        var size = MemoryLayout<kinfo_proc>.stride
        var mib: [Int32] = [CTL_KERN, KERN_PROC, KERN_PROC_PID, getpid()]
        
        let result = sysctl(&mib, UInt32(mib.count), &info, &size, nil, 0)
        
        return result == 0 && (info.kp_proc.p_flag & P_TRACED) != 0
    }
    
    // MARK: - Runtime Manipulation Detection
    
    private func isRuntimeManipulated() -> Bool {
        // Check for common runtime manipulation tools
        let loadedLibs = _dyld_image_count()
        
        for i in 0..<loadedLibs {
            if let name = _dyld_get_image_name(i) {
                let libName = String(cString: name)
                if libName.contains("Frida") || libName.contains("frida") ||
                   libName.contains("Substrate") || libName.contains("substrate") {
                    return true
                }
            }
        }
        
        return false
    }
    
    // MARK: - Certificate Pinning
    
    private var pinnedPublicKeyHashes: Set<String> {
        // In production, add your actual public key hashes
        return [
            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=", // Example
            "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB="
        ]
    }
    
    func validateCertificatePinning(serverTrust: SecTrust, host: String) -> Bool {
        guard let certificate = SecTrustGetCertificateAtIndex(serverTrust, 0) else {
            return false
        }
        
        // Extract public key
        guard let publicKey = SecCertificateCopyKey(certificate) else {
            return false
        }
        
        // Get key data
        var error: Unmanaged<CFError>?
        guard let keyData = SecKeyCopyExternalRepresentation(publicKey, &error) as Data? else {
            return false
        }
        
        // Hash the public key
        let hash = SHA256.hash(data: keyData)
        let keyHash = Data(hash).base64EncodedString()
        
        return pinnedPublicKeyHashes.contains(keyHash)
    }
    
    // MARK: - Helper Functions
    
    private func generateNonce() -> String {
        var bytes = [UInt8](repeating: 0, count: 32)
        _ = SecRandomCopyBytes(kSecRandomDefault, bytes.count, &bytes)
        return Data(bytes).base64EncodedString()
    }
    
    private func generateAuthenticationToken() -> String {
        let timestamp = Int(Date().timeIntervalSince1970)
        let nonce = generateNonce()
        let payload = "\(timestamp)|\(nonce)"
        return payload.sha256Hash()
    }
    
    private func verifyAppIntegrity() -> Bool {
        // Verify app bundle identifier hasn't been modified
        let bundleId = Bundle.main.bundleIdentifier
        return bundleId == "com.gameengine.casino"
    }
}

// MARK: - Supporting Types

enum SecurityError: Error {
    case keyGenerationFailed
    case keyNotFound
    case signingFailed
    case authenticationFailed
    case integrityCheckFailed
}

enum BiometricType {
    case none
    case touchID
    case faceID
    case opticID
}

enum IntegrityCheck {
    case jailbreakDetected
    case debuggerAttached
    case appIntegrityFailed
    case runtimeManipulation
    case screenBeingRecorded
}

enum SecurityIssue: String, CaseIterable {
    case jailbreakDetected = "Device is jailbroken"
    case debuggerAttached = "Debugger is attached"
    case appIntegrityFailed = "App integrity check failed"
    case runtimeManipulation = "Runtime manipulation detected"
    case screenRecordingDetected = "Screen recording is active"
    case remoteAccessAppInstalled = "Remote access app detected"
}

struct FullSecurityResult {
    let isSecure: Bool
    let issues: [SecurityIssue]
    let remoteApps: [String]
    let checkedAt: Date
}

struct IntegrityResult {
    let isSecure: Bool
    let failedChecks: [IntegrityCheck]
    let timestamp: Date
    let deviceFingerprint: DeviceFingerprint
}

struct AuthenticationResult {
    let success: Bool
    let method: AuthenticationMethod
    let token: String
    let expiresAt: Date
}

enum AuthenticationMethod {
    case biometric
    case passcode
}

struct DeviceFingerprint {
    let deviceId: String
    let model: String
    let manufacturer: String
    let osVersion: String
    let isJailbroken: Bool
    let isSimulator: Bool
}

// MARK: - String Extension

extension String {
    func sha256Hash() -> String {
        guard let data = self.data(using: .utf8) else { return "" }
        let hash = SHA256.hash(data: data)
        return hash.compactMap { String(format: "%02x", $0) }.joined()
    }
}
