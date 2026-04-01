import UIKit
import Combine

class TournamentCell: UITableViewCell {
    static let identifier = "TournamentCell"

    var onRegisterTapped: (() -> Void)?

    private let nameLabel = UILabel()
    private let gameTypeLabel = UILabel()
    private let buyInLabel = UILabel()
    private let prizePoolLabel = UILabel()
    private let statusBadge = UILabel()
    private let registerButton = UIButton(type: .system)

    override init(style: UITableViewCell.CellStyle, reuseIdentifier: String?) {
        super.init(style: style, reuseIdentifier: reuseIdentifier)
        setupUI()
    }

    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    private func setupUI() {
        contentView.addSubview(nameLabel)
        contentView.addSubview(gameTypeLabel)
        contentView.addSubview(buyInLabel)
        contentView.addSubview(prizePoolLabel)
        contentView.addSubview(statusBadge)
        contentView.addSubview(registerButton)

        nameLabel.font = .boldSystemFont(ofSize: 17)
        gameTypeLabel.font = .systemFont(ofSize: 14)
        gameTypeLabel.textColor = .secondaryLabel
        buyInLabel.font = .systemFont(ofSize: 14)
        prizePoolLabel.font = .boldSystemFont(ofSize: 16)

        statusBadge.font = .systemFont(ofSize: 12, weight: .medium)
        statusBadge.textAlignment = .center
        statusBadge.layer.cornerRadius = 10
        statusBadge.clipsToBounds = true

        registerButton.addTarget(self, action: #selector(registerTapped), for: .touchUpInside)

        nameLabel.translatesAutoresizingMaskIntoConstraints = false
        statusBadge.translatesAutoresizingMaskIntoConstraints = false
        gameTypeLabel.translatesAutoresizingMaskIntoConstraints = false
        buyInLabel.translatesAutoresizingMaskIntoConstraints = false
        prizePoolLabel.translatesAutoresizingMaskIntoConstraints = false
        registerButton.translatesAutoresizingMaskIntoConstraints = false

        NSLayoutConstraint.activate([
            nameLabel.topAnchor.constraint(equalTo: contentView.topAnchor, constant: 12),
            nameLabel.leadingAnchor.constraint(equalTo: contentView.leadingAnchor, constant: 16),

            statusBadge.centerYAnchor.constraint(equalTo: nameLabel.centerYAnchor),
            statusBadge.trailingAnchor.constraint(equalTo: contentView.trailingAnchor, constant: -16),
            statusBadge.widthAnchor.constraint(greaterThanOrEqualToConstant: 60),
            statusBadge.heightAnchor.constraint(equalToConstant: 20),

            gameTypeLabel.topAnchor.constraint(equalTo: nameLabel.bottomAnchor, constant: 4),
            gameTypeLabel.leadingAnchor.constraint(equalTo: nameLabel.leadingAnchor),

            buyInLabel.topAnchor.constraint(equalTo: gameTypeLabel.bottomAnchor, constant: 8),
            buyInLabel.leadingAnchor.constraint(equalTo: nameLabel.leadingAnchor),

            prizePoolLabel.topAnchor.constraint(equalTo: buyInLabel.bottomAnchor, constant: 4),
            prizePoolLabel.leadingAnchor.constraint(equalTo: nameLabel.leadingAnchor),

            registerButton.centerYAnchor.constraint(equalTo: contentView.centerYAnchor),
            registerButton.trailingAnchor.constraint(equalTo: contentView.trailingAnchor, constant: -16),
            registerButton.bottomAnchor.constraint(lessThanOrEqualTo: contentView.bottomAnchor, constant: -12)
        ])
    }

    func configure(with tournament: Tournament) {
        nameLabel.text = tournament.name
        gameTypeLabel.text = tournament.gameType
        buyInLabel.text = "Buy-in: \(tournament.buyIn.currencyFormatted())"
        prizePoolLabel.text = "Prize Pool: \(tournament.prizePool.currencyFormatted())"

        switch tournament.status {
        case .upcoming:
            statusBadge.text = "Upcoming"
            statusBadge.backgroundColor = .systemBlue.withAlphaComponent(0.2)
            statusBadge.textColor = .systemBlue
            registerButton.setTitle("Register", for: .normal)
        case .registrationOpen:
            statusBadge.text = "Register"
            statusBadge.backgroundColor = .systemGreen.withAlphaComponent(0.2)
            statusBadge.textColor = .systemGreen
            registerButton.setTitle("Register", for: .normal)
        case .inProgress:
            statusBadge.text = "Live"
            statusBadge.backgroundColor = .systemRed.withAlphaComponent(0.2)
            statusBadge.textColor = .systemRed
            registerButton.setTitle("Play", for: .normal)
        case .completed:
            statusBadge.text = "Completed"
            statusBadge.backgroundColor = .systemGray.withAlphaComponent(0.2)
            statusBadge.textColor = .systemGray
            registerButton.isHidden = true
        default:
            statusBadge.text = tournament.status.rawValue.capitalized
            registerButton.isHidden = true
        }
    }

    @objc private func registerTapped() {
        onRegisterTapped?()
    }
}

// MARK: - Tournament Detail

class TournamentDetailViewController: UIViewController {
    private let tournamentId: String
    private let viewModel = TournamentViewModel()
    private var cancellables = Set<AnyCancellable>()

    init(tournamentId: String) {
        self.tournamentId = tournamentId
        super.init(nibName: nil, bundle: nil)
    }

    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    override func viewDidLoad() {
        super.viewDidLoad()
        viewModel.selectTournament(tournamentId)
    }
}

// MARK: - Extensions

extension Double {
    func currencyFormatted() -> String {
        let formatter = NumberFormatter()
        formatter.numberStyle = .currency
        formatter.currencyCode = "USD"
        return formatter.string(from: NSNumber(value: self)) ?? "$\(self)"
    }
}
