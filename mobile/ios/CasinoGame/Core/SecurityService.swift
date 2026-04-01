import UIKit
import LocalAuthentication
import KeychainAccess
import CryptoKit

class SecurityService {

    static let shared = SecurityService()

    private let keychain: Keychain
    private let keyTag = "com.casino.app.securekey"
    private var attestedKeyId: String?

    private init() {
        keychain = Keychain(service: "com.casino.app")
            .accessibility(.whenUnlockedThisDeviceOnly)
            .authenticationPolicy(.biometryCurrentSet)
    }

    // MARK: - Secure Enclave Key Management

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

    // MARK: - Enhanced Keychain Operations

    func saveSecureValue(key: String, value: String) throws {
        let secureKeychain = Keychain(service: "com.casino.app")
            .accessibility(.whenUnlockedThisDeviceOnly, authenticationPolicy: .biometryCurrentSet)
            .authenticationPrompt("Authenticate to access secure data")
        try secureKeychain.set(value, key: key)
    }

    func getSecureValue(key: String) throws -> String? {
        let secureKeychain = Keychain(service: "com.casino.app")
            .accessibility(.whenUnlockedThisDeviceOnly, authenticationPolicy: .biometryCurrentSet)
            .authenticationPrompt("Authenticate to access secure data")
        return try secureKeychain.get(key)
    }

    func saveValue(key: String, value: String) throws {
        try keychain
            .accessibility(.whenUnlockedThisDeviceOnly)
            .set(value, key: key)
    }

    func getValue(key: String) -> String? {
        return try? keychain.get(key)
    }

    // MARK: - Anti-Tampering

    func performIntegrityCheck() -> IntegrityResult {
        var checks: [IntegrityCheck] = []

        if isDeviceJailbroken() {
            checks.append(.jailbreakDetected)
        }

        if isDebuggerAttached() {
            checks.append(.debuggerAttached)
        }

        if !verifyAppIntegrity() {
            checks.append(.appIntegrityFailed)
        }

        if isRuntimeManipulated() {
            checks.append(.runtimeManipulation)
        }

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

    // MARK: - Device Fingerprint

    func getDeviceFingerprint() -> DeviceFingerprint {
        var components: [String] = []

        if let vendorId = UIDevice.current.identifierForVendor?.uuidString {
            components.append("vendor:\(vendorId)")
        }

        components.append("model:\(UIDevice.current.model)")
        components.append("system:\(UIDevice.current.systemVersion)")
        components.append("name:\(UIDevice.current.name)")

        let screen = UIScreen.main.bounds
        components.append("screen:\(Int(screen.width))x\(Int(screen.height))")

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

extension String {
    func sha256Hash() -> String {
        guard let data = self.data(using: .utf8) else { return "" }
        let hash = SHA256.hash(data: data)
        return hash.compactMap { String(format: "%02x", $0) }.joined()
    }
}
