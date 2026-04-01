import Foundation

struct DepositState {
    var isLoading: Bool = false
    var isSuccess: Bool = false
    var balance: Double = 0
    var paymentMethods: [PaymentMethod] = []
    var selectedMethodIndex: Int = 0
    var amount: Double = 0
    var error: String?
}

class DepositViewModel {

    var onStateChange: ((DepositState) -> Void)?

    private(set) var state = DepositState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadData() {
        loadBalance()
        loadPaymentMethods()
    }

    private func loadBalance() {
        Task {
            do {
                let balance = try await apiClient.getBalance()
                state.balance = balance.balance
            } catch { }
        }
    }

    private func loadPaymentMethods() {
        state.isLoading = true
        Task {
            do {
                let response = try await apiClient.getPaymentMethods()
                state.isLoading = false
                state.paymentMethods = response.depositMethods
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func selectMethod(at index: Int) {
        guard index < state.paymentMethods.count else { return }
        state.selectedMethodIndex = index
    }

    func deposit(amount: Double) {
        guard amount > 0, !state.paymentMethods.isEmpty else { return }
        let method = state.paymentMethods[state.selectedMethodIndex]
        state.isLoading = true
        state.error = nil
        state.amount = amount

        Task {
            do {
                let response = try await apiClient.deposit(
                    amount: amount,
                    paymentMethod: method.type,
                    paymentId: method.id,
                    currency: "USD"
                )
                state.isLoading = false
                state.isSuccess = response.status == "completed" || response.status == "pending"
                loadBalance()
            } catch let error as APIError {
                state.isLoading = false
                state.error = error.errorDescription
            } catch {
                state.isLoading = false
                state.error = error.localizedDescription
            }
        }
    }

    func validateAmount(_ text: String) -> (valid: Bool, message: String?) {
        guard let value = Double(text), value > 0 else {
            return (false, "Enter a valid amount")
        }
        let method = state.paymentMethods.isEmpty ? nil : state.paymentMethods[state.selectedMethodIndex]
        if let min = method?.minAmount, value < min {
            return (false, "Minimum deposit is $\(String(format: "%.2f", min))")
        }
        if let max = method?.maxAmount, value > max {
            return (false, "Maximum deposit is $\(String(format: "%.2f", max))")
        }
        return (true, nil)
    }
}
