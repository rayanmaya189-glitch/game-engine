import UIKit
import CryptoKit
import KeychainAccess

extension SecurityService {

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

    @available(iOS 14.0, *)
    func attestDevice(completion: @escaping (Result<String, Error>) -> Void) {
        let nonce = generateNonce()
        completion(.success(nonce))
    }

    @available(iOS 14.0, *)
    func generateAttestationClaim(nonce: String, data: Data) -> Data? {
        return try? signWithBiometric(data: data)
    }

    // MARK: - Certificate Pinning

    private var pinnedPublicKeyHashes: Set<String> {
        return [
            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
            "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB="
        ]
    }

    func validateCertificatePinning(serverTrust: SecTrust, host: String) -> Bool {
        guard let certificate = SecTrustGetCertificateAtIndex(serverTrust, 0) else {
            return false
        }

        guard let publicKey = SecCertificateCopyKey(certificate) else {
            return false
        }

        var error: Unmanaged<CFError>?
        guard let keyData = SecKeyCopyExternalRepresentation(publicKey, &error) as Data? else {
            return false
        }

        let hash = SHA256.hash(data: keyData)
        let keyHash = Data(hash).base64EncodedString()

        return pinnedPublicKeyHashes.contains(keyHash)
    }
}
