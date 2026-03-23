import UIKit
import SwiftUI
import Combine

/**
 * Tournament View Controller
 * 
 * Displays tournament lobby with browsing, registration, and in-play functionality.
 */

class TournamentViewController: UIViewController {
    
    private let viewModel = TournamentViewModel()
    private var cancellables = Set<AnyCancellable>()
    
    // MARK: - UI Components
    
    private lazy var filterSegmentedControl: UISegmentedControl = {
        let items = ["All", "Upcoming", "Active", "My Tournaments"]
        let control = UISegmentedControl(items: items)
        control.selectedSegmentIndex = 0
        control.addTarget(self, action: #selector(filterChanged), for: .valueChanged)
        return control
    }()
    
    private lazy var tableView: UITableView = {
        let table = UITableView(frame: .zero, style: .insetGrouped)
        table.register(TournamentCell.self, forCellReuseIdentifier: TournamentCell.identifier)
        table.delegate = self
        table.dataSource = self
        table.refreshControl = UIRefreshControl()
        table.refreshControl?.addTarget(self, action: #selector(refreshData), for: .valueChanged)
        return table
    }()
    
    private lazy var loadingIndicator: UIActivityIndicatorView = {
        let indicator = UIActivityIndicatorView(style: .large)
        indicator.hidesWhenStopped = true
        return indicator
    }()
    
    private lazy var emptyStateLabel: UILabel = {
        let label = UILabel()
        label.text = "No tournaments available"
        label.textAlignment = .center
        label.textColor = .secondaryLabel
        label.isHidden = true
        return label
    }()
    
    // MARK: - Lifecycle
    
    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        bindViewModel()
        viewModel.loadTournaments()
    }
    
    // MARK: - Setup
    
    private func setupUI() {
        title = "Tournaments"
        view.backgroundColor = .systemBackground
        
        navigationItem.titleView = filterSegmentedControl
        
        view.addSubview(tableView)
        view.addSubview(loadingIndicator)
        view.addSubview(emptyStateLabel)
        
        tableView.translatesAutoresizingMaskIntoConstraints = false
        loadingIndicator.translatesAutoresizingMaskIntoConstraints = false
        emptyStateLabel.translatesAutoresizingMaskIntoConstraints = false
        
        NSLayoutConstraint.activate([
            tableView.topAnchor.constraint(equalTo: view.safeAreaLayoutGuide.topAnchor),
            tableView.leadingAnchor.constraint(equalTo: view.leadingAnchor),
            tableView.trailingAnchor.constraint(equalTo: view.trailingAnchor),
            tableView.bottomAnchor.constraint(equalTo: view.bottomAnchor),
            
            loadingIndicator.centerXAnchor.constraint(equalTo: view.centerXAnchor),
            loadingIndicator.centerYAnchor.constraint(equalTo: view.centerYAnchor),
            
            emptyStateLabel.centerXAnchor.constraint(equalTo: view.centerXAnchor),
            emptyStateLabel.centerYAnchor.constraint(equalTo: view.centerYAnchor)
        ])
    }
    
    private func bindViewModel() {
        viewModel.$uiState
            .receive(on: DispatchQueue.main)
            .sink { [weak self] state in
                self?.handleState(state)
            }
            .store(in: &cancellables)
    }
    
    private func handleState(_ state: TournamentUiState) {
        tableView.refreshControl?.endRefreshing()
        
        switch state {
        case .loading:
            loadingIndicator.startAnimating()
            emptyStateLabel.isHidden = true
        case .success(let tournaments):
            loadingIndicator.stopAnimating()
            emptyStateLabel.isHidden = !tournaments.isEmpty
            tableView.reloadData()
        case .error(let message):
            loadingIndicator.stopAnimating()
            showError(message)
        }
    }
    
    // MARK: - Actions
    
    @objc private func filterChanged() {
        let filters: [TournamentFilter] = [.all, .upcoming, .active, .myTournaments]
        viewModel.filterTournaments(filters[filterSegmentedControl.selectedSegmentIndex])
    }
    
    @objc private func refreshData() {
        viewModel.loadTournaments()
    }
    
    private func showError(_ message: String) {
        let alert = UIAlertController(title: "Error", message: message, preferredStyle: .alert)
        alert.addAction(UIAlertAction(title: "Retry", style: .default) { [weak self] _ in
            self?.viewModel.loadTournaments()
        })
        alert.addAction(UIAlertAction(title: "Cancel", style: .cancel))
        present(alert, animated: true)
    }
}

// MARK: - UITableViewDataSource & Delegate

extension TournamentViewController: UITableViewDataSource, UITableViewDelegate {
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        if case .success(let tournaments) = viewModel.uiState {
            return tournaments.count
        }
        return 0
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        guard let cell = tableView.dequeueReusableCell(withIdentifier: TournamentCell.identifier, for: indexPath) as? TournamentCell,
              case .success(let tournaments) = viewModel.uiState else {
            return UITableViewCell()
        }
        
        let tournament = tournaments[indexPath.row]
        cell.configure(with: tournament)
        cell.onRegisterTapped = { [weak self] in
            self?.viewModel.registerForTournament(tournament.id)
        }
        
        return cell
    }
    
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
        
        if case .success(let tournaments) = viewModel.uiState {
            let tournament = tournaments[indexPath.row]
            let detailVC = TournamentDetailViewController(tournamentId: tournament.id)
            navigationController?.pushViewController(detailVC, animated: true)
        }
    }
}

// MARK: - Tournament Cell

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
        // Setup detail UI
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
