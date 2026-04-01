import UIKit
import SwiftUI
import Combine

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
