import UIKit
import SnapKit

class PaymentHistoryViewController: UIViewController {

    private var viewModel = PaymentHistoryViewModel()

    private let filterSegment: UISegmentedControl = {
        let items = PaymentFilter.allCases.map { $0.title }
        let control = UISegmentedControl(items: items)
        control.selectedSegmentIndex = 0
        control.backgroundColor = UIColor(hex: "#1E1E3F")
        control.selectedSegmentTintColor = UIColor(hex: "#FF6B35")
        control.setTitleTextAttributes([.foregroundColor: UIColor.white], for: .selected)
        control.setTitleTextAttributes([.foregroundColor: UIColor.white.withAlphaComponent(0.7)], for: .normal)
        return control
    }()

    private let tableView: UITableView = {
        let tableView = UITableView()
        tableView.backgroundColor = .clear
        tableView.separatorStyle = .none
        tableView.register(PaymentTransactionCell.self, forCellReuseIdentifier: "PaymentTransactionCell")
        return tableView
    }()

    private let emptyLabel: UILabel = {
        let label = UILabel()
        label.text = "No transactions found"
        label.textColor = .white.withAlphaComponent(0.5)
        label.font = .systemFont(ofSize: 16)
        label.textAlignment = .center
        label.isHidden = true
        return label
    }()

    private let activityIndicator: UIActivityIndicatorView = {
        let indicator = UIActivityIndicatorView(style: .large)
        indicator.color = UIColor(hex: "#FF6B35")
        indicator.hidesWhenStopped = true
        return indicator
    }()

    private let refreshControl = UIRefreshControl()

    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        setupBindings()
        viewModel.loadTransactions()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Payment History"

        view.addSubview(filterSegment)
        view.addSubview(tableView)
        view.addSubview(emptyLabel)
        view.addSubview(activityIndicator)

        filterSegment.snp.makeConstraints { make in
            make.top.equalTo(view.safeAreaLayoutGuide).offset(16)
            make.leading.trailing.equalToSuperview().inset(16)
            make.height.equalTo(40)
        }

        tableView.snp.makeConstraints { make in
            make.top.equalTo(filterSegment.snp.bottom).offset(16)
            make.leading.trailing.bottom.equalToSuperview()
        }

        emptyLabel.snp.makeConstraints { make in
            make.center.equalTo(tableView)
        }

        activityIndicator.snp.makeConstraints { make in
            make.center.equalTo(tableView)
        }

        tableView.delegate = self
        tableView.dataSource = self

        filterSegment.addTarget(self, action: #selector(filterChanged), for: .valueChanged)

        refreshControl.tintColor = UIColor(hex: "#FF6B35")
        refreshControl.addTarget(self, action: #selector(refreshPulled), for: .valueChanged)
        tableView.refreshControl = refreshControl
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.updateUI(with: state)
            }
        }
    }

    private func updateUI(with state: PaymentHistoryState) {
        if state.isLoading && !refreshControl.isRefreshing {
            activityIndicator.startAnimating()
        } else {
            activityIndicator.stopAnimating()
            refreshControl.endRefreshing()
        }

        emptyLabel.isHidden = !state.transactions.isEmpty
        tableView.reloadData()
    }

    @objc private func filterChanged() {
        let filters = PaymentFilter.allCases
        guard filterSegment.selectedSegmentIndex < filters.count else { return }
        viewModel.setFilter(filters[filterSegment.selectedSegmentIndex])
    }

    @objc private func refreshPulled() {
        viewModel.loadTransactions()
    }
}

extension PaymentHistoryViewController: UITableViewDataSource, UITableViewDelegate {
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return viewModel.state.transactions.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "PaymentTransactionCell", for: indexPath) as! PaymentTransactionCell
        if indexPath.row < viewModel.state.transactions.count {
            cell.configure(with: viewModel.state.transactions[indexPath.row])
        }
        return cell
    }

    func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
        return 72
    }
}
