import Foundation

struct KycDocumentItem: Identifiable {
    let id: String
    let type: String
    let title: String
    let iconName: String
    var status: String
    var fileName: String?
}

struct KycVerificationState {
    var isLoading: Bool = false
    var currentLevel: Int = 0
    var maxLevel: Int = 3
    var documents: [KycDocumentItem] = [
        KycDocumentItem(id: "id", type: "id_card", title: "Government ID", iconName: "person.text.rectangle", status: "not_uploaded"),
        KycDocumentItem(id: "address", type: "proof_of_address", title: "Proof of Address", iconName: "house.fill", status: "not_uploaded"),
        KycDocumentItem(id: "selfie", type: "selfie", title: "Selfie Verification", iconName: "camera.fill", status: "not_uploaded")
    ]
    var isFullyVerified: Bool = false
    var canSubmit: Bool = false
    var submitSuccess: Bool = false
    var error: String?
}

class KycVerificationViewModel {

    var onStateChange: ((KycVerificationState) -> Void)?

    private(set) var state = KycVerificationState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadStatus() {
        state.isLoading = true

        Task {
            do {
                let status = try await apiClient.getKycStatus()
                state.isLoading = false
                state.currentLevel = status.currentLevel
                state.maxLevel = status.maxLevel
                state.isFullyVerified = status.isFullyVerified
                syncDocuments(with: status.documents)
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func uploadDocument(type: String, fileName: String, data: Data) {
        state.isLoading = true

        Task {
            do {
                let response = try await apiClient.uploadKycDocument(type: type, fileName: fileName, data: data)
                state.isLoading = false
                if let index = state.documents.firstIndex(where: { $0.type == type }) {
                    state.documents[index].status = response.document.status
                    state.documents[index].fileName = fileName
                }
                updateCanSubmit()
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func submitVerification() {
        let allUploaded = state.documents.allSatisfy { $0.status == "uploaded" || $0.status == "verified" }
        guard allUploaded else {
            state.error = "Please upload all required documents"
            return
        }

        state.isLoading = true

        Task {
            do {
                let response = try await apiClient.submitKycVerification()
                state.isLoading = false
                state.submitSuccess = response.status == "submitted" || response.status == "verified"
                if let newLevel = response.newLevel {
                    state.currentLevel = newLevel
                }
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    private func syncDocuments(with remoteDocs: [KycDocument]) {
        var updated = state.documents
        for remote in remoteDocs {
            if let index = updated.firstIndex(where: { $0.type == remote.type }) {
                updated[index].status = remote.status
                updated[index].fileName = remote.fileName
            }
        }
        state.documents = updated
        updateCanSubmit()
    }

    private func updateCanSubmit() {
        state.canSubmit = state.documents.allSatisfy { $0.status == "uploaded" || $0.status == "verified" }
    }
}
