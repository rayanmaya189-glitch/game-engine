import Foundation

struct SpinResult: Codable {
    let reels: [[String]]
    let winAmount: Double
    let multiplier: Double
    let isJackpot: Bool
    let newBalance: Double

    enum CodingKeys: String, CodingKey {
        case reels, multiplier, jackpot
        case winAmount = "win_amount"
        case isJackpot = "is_jackpot"
        case newBalance = "new_balance"
    }

    init(from decoder: Decoder) throws {
        let c = try decoder.container(keyedBy: CodingKeys.self)
        reels = try c.decode([[String]].self, forKey: .reels)
        winAmount = try c.decode(Double.self, forKey: .winAmount)
        multiplier = try c.decodeIfPresent(Double.self, forKey: .multiplier) ?? 0
        isJackpot = try c.decodeIfPresent(Bool.self, forKey: .isJackpot) ?? false
        newBalance = try c.decode(Double.self, forKey: .newBalance)
    }
}

struct GamePlayState {
    var isLoading: Bool = false
    var isSpinning: Bool = false
    var balance: Double = 0
    var betAmount: Double = 1.0
    var winAmount: Double = 0
    var reels: [[String]] = Array(repeating: Array(repeating: "🍒", count: 3), count: 5)
    var error: String?
}

class GamePlayViewModel {

    var onStateChange: ((GamePlayState) -> Void)?

    private(set) var state = GamePlayState() {
        didSet { onStateChange?(state) }
    }

    private let apiClient = APIClient.shared
    let minBet: Double = 0.10
    let maxBet: Double = 100.0
    let betSteps: [Double] = [0.10, 0.25, 0.50, 1.0, 2.5, 5.0, 10.0, 25.0, 50.0, 100.0]

    func loadBalance() {
        Task {
            do {
                let balance = try await apiClient.getBalance()
                state.balance = balance.balance
            } catch { }
        }
    }

    func updateBet(_ amount: Double) {
        state.betAmount = amount
        state.winAmount = 0
    }

    func spin() {
        guard !state.isSpinning, state.betAmount <= state.balance else { return }
        state.isSpinning = true
        state.winAmount = 0
        state.error = nil

        Task {
            do {
                let result: SpinResult = try await apiClient.spinGame(betAmount: state.betAmount)
                state.isSpinning = false
                state.reels = result.reels
                state.winAmount = result.winAmount
                state.balance = result.newBalance
            } catch let error as APIError {
                state.isSpinning = false
                state.error = error.errorDescription
            } catch {
                state.isSpinning = false
                state.error = error.localizedDescription
            }
        }
    }
}
