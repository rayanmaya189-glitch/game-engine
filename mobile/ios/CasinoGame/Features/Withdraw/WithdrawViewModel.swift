import Foundation

struct WithdrawState {
    var isLoading: Bool = false
    var isSuccess: Bool = false
    var balance: Double = 0
    var withdrawMethods: [PaymentMethod] = []
    var selectedMethodIndex: Int = 0
    var amount: Double = 0
    var fee: Double = 0
    var error: String?
}

class WithdrawViewModel {

    var onStateChange: ((WithdrawState) -> Void)?

    private(set) var state = WithdrawState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared

    func loadData() {
        loadBalance()
        loadWithdrawMethods()
    }

    private func loadBalance() {
        Task {
            do {
                let balance = try await apiClient.getBalance()
                state.balance = balance.balance
            } catch { }
        }
    }

    private func loadWithdrawMethods() {
        state.isLoading = true
        Task {
            do {
                let response = try await apiClient.getPaymentMethods()
                state.isLoading = false
                state.withdrawMethods = response.withdrawMethods
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
        guard index < state.withdrawMethods.count else { return }
        state.selectedMethodIndex = index
    }

    func withdraw(amount: Double, paymentDetails: String) {
        guard amount > 0, !state.withdrawMethods.isEmpty else { return }
        let method = state.withdrawMethods[state.selectedMethodIndex]
        state.isLoading = true
        state.error = nil
        state.amount = amount

        Task {
            do {
                let response = try await apiClient.withdraw(
                    amount: amount,
                    paymentMethod: method.type,
                    paymentDetails: paymentDetails,
                    currency: "USD"
                )
                state.isLoading = false
                state.isSuccess = response.status == "completed" || response.status == "pending"
                state.fee = response.fee
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
        if value > state.balance {
            return (false, "Insufficient balance")
        }
        let method = state.withdrawMethods.isEmpty ? nil : state.withdrawMethods[state.selectedMethodIndex]
        if let min = method?.minAmount, value < min {
            return (false, "Minimum withdrawal is $\(String(format: "%.2f", min))")
        }
        if let max = method?.maxAmount, value > max {
            return (false, "Maximum withdrawal is $\(String(format: "%.2f", max))")
        }
        return (true, nil)
    }
}
