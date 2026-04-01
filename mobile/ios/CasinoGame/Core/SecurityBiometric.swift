import UIKit
import LocalAuthentication

extension SecurityService {

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

    func authenticateEnhanced(
        reason: String,
        completion: @escaping (Result<AuthenticationResult, Error>) -> Void
    ) {
        let context = LAContext()
        context.localizedCancelTitle = "Cancel"
        context.localizedFallbackTitle = "Use Passcode"

        context.evaluatePolicy(
            .deviceOwnerAuthenticationWithBiometrics,
            localizedReason: reason
        ) { [weak self] success, error in
            if success {
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
}
