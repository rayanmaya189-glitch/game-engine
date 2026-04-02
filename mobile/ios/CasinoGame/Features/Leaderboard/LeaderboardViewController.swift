import UIKit

class LeaderboardViewController: UIViewController {

    private var viewModel = LeaderboardViewModel()

    private lazy var periodControl: UISegmentedControl = {
        let items = LeaderboardPeriod.allCases.map { $0.rawValue.capitalized }
        let control = UISegmentedControl(items: items)
        control.selectedSegmentIndex = 0
        control.backgroundColor = UIColor(hex: "#1E1E3F")
        control.selectedSegmentTintColor = UIColor(hex: "#FF6B35")
        control.setTitleTextAttributes([.foregroundColor: UIColor.white], for: .normal)
        control.setTitleTextAttributes([.foregroundColor: UIColor.white], for: .selected)
        control.addTarget(self, action: #selector(periodChanged), for: .valueChanged)
        return control
    }()

    private lazy var tableView: UITableView = {
        let tv = UITableView(frame: .zero, style: .insetGrouped)
        tv.backgroundColor = .clear
        tv.register(UITableViewCell.self, forCellReuseIdentifier: "LeaderboardCell")
        tv.refreshControl = UIRefreshControl()
        tv.refreshControl?.tintColor = UIColor(hex: "#FF6B35")
        tv.refreshControl?.addTarget(self, action: #selector(refreshData), for: .valueChanged)
        return tv
    }()

    private let activityIndicator: UIActivityIndicatorView = {
        let indicator = UIActivityIndicatorView(style: .large)
        indicator.color = UIColor(hex: "#FF6B35")
        indicator.hidesWhenStopped = true
        return indicator
    }()

    private let medalIcons = ["🥇", "🥈", "🥉"]

    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        setupBindings()
        viewModel.loadLeaderboard()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Leaderboard"

        view.addSubview(periodControl)
        view.addSubview(tableView)
        view.addSubview(activityIndicator)

        tableView.delegate = self
        tableView.dataSource = self

        periodControl.snp.makeConstraints { make in
            make.top.equalTo(view.safeAreaLayoutGuide).offset(8)
            make.leading.trailing.equalToSuperview().inset(16)
            make.height.equalTo(36)
        }
        tableView.snp.makeConstraints { make in
            make.top.equalTo(periodControl.snp.bottom).offset(8)
            make.leading.trailing.bottom.equalToSuperview()
        }
        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.tableView.refreshControl?.endRefreshing()
                self?.tableView.reloadData()
                self?.activityIndicator.stopAnimating()
                if let error = state.error {
                    let alert = UIAlertController(title: "Error", message: error, preferredStyle: .alert)
                    alert.addAction(UIAlertAction(title: "OK", style: .cancel))
                    self?.present(alert, animated: true)
                }
            }
        }
    }

    @objc private func periodChanged() {
        let period = LeaderboardPeriod.allCases[periodControl.selectedSegmentIndex]
        viewModel.changePeriod(period)
    }

    @objc private func refreshData() {
        viewModel.loadLeaderboard()
    }
}

// MARK: - UITableViewDataSource & Delegate

extension LeaderboardViewController: UITableViewDataSource, UITableViewDelegate {

    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return viewModel.state.entries.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "LeaderboardCell", for: indexPath)
        let entry = viewModel.state.entries[indexPath.row]

        var config = cell.defaultContentConfiguration()
        let rankPrefix = entry.rank <= 3 ? "\(medalIcons[entry.rank - 1]) " : "#\(entry.rank) "
        config.text = "\(rankPrefix)\(entry.username)"
        config.secondaryText = "Score: \(formatNumber(entry.score))  •  Prize: $\(formatNumber(entry.prize))"
        config.textProperties.color = entry.isCurrentUser ? UIColor(hex: "#FF6B35") : .white
        config.secondaryTextProperties.color = .white.withAlphaComponent(0.6)

        if entry.isCurrentUser {
            cell.backgroundColor = UIColor(hex: "#FF6B35").withAlphaComponent(0.15)
            cell.layer.borderWidth = 1
            cell.layer.borderColor = UIColor(hex: "#FF6B35").cgColor
            cell.layer.cornerRadius = 8
        } else {
            cell.backgroundColor = UIColor(hex: "#1E1E3F")
            cell.layer.borderWidth = 0
        }

        cell.contentConfiguration = config
        cell.selectionStyle = .none
        return cell
    }

    private func formatNumber(_ value: Double) -> String {
        if value >= 1_000_000 {
            return String(format: "%.1fM", value / 1_000_000)
        } else if value >= 1_000 {
            return String(format: "%.1fK", value / 1_000)
        }
        return String(format: "%.0f", value)
    }
}
