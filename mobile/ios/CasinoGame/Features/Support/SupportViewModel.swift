import Foundation

struct FAQ: Codable {
    let id: String
    let question: String
    let answer: String
}

struct SupportTicket: Codable, Identifiable {
    let id: String
    let subject: String
    let status: String
    let createdAt: String

    enum CodingKeys: String, CodingKey {
        case id, subject, status
        case createdAt = "created_at"
    }
}

struct FaqResponse: Codable {
    let faqs: [FAQ]
}

struct TicketsResponse: Codable {
    let tickets: [SupportTicket]
}

struct SupportState {
    var isLoading: Bool = false
    var faqs: [FAQ] = []
    var tickets: [SupportTicket] = []
    var expandedFaqIds: Set<String> = []
    var error: String?
}

class SupportViewModel {

    var onStateChange: ((SupportState) -> Void)?

    private(set) var state = SupportState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadFaqs() {
        state.isLoading = true
        Task {
            do {
                let response = try await apiClient.getFaq()
                state.isLoading = false
                state.faqs = response.faqs
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func loadTickets() {
        Task {
            do {
                let response = try await apiClient.getTickets()
                state.tickets = response.tickets
            } catch { }
        }
    }

    func toggleFaq(_ id: String) {
        if state.expandedFaqIds.contains(id) {
            state.expandedFaqIds.remove(id)
        } else {
            state.expandedFaqIds.insert(id)
        }
    }

    func createTicket(subject: String, message: String) {
        Task {
            do {
                let ticket = try await apiClient.createTicket(subject: subject, message: message)
                state.tickets.insert(ticket, at: 0)
            } catch let error as APIError {
                state.error = error.errorDescription
            } catch {
                state.error = error.localizedDescription
            }
        }
    }
}
