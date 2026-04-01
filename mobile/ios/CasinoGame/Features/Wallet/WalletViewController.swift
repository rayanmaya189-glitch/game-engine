import UIKit
import SnapKit

class WalletViewController: UIViewController {

    private var viewModel = WalletViewModel()

    private let scrollView: UIScrollView = {
        let scrollView = UIScrollView()
        scrollView.showsVerticalScrollIndicator = false
        return scrollView
    }()

    private let contentView = UIView()

    private let balanceCard: UIView = {
        let view = UIView()
        view.backgroundColor = UIColor(hex: "#FF6B35")
        view.layer.cornerRadius = 16
        return view
    }()

    private let balanceTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "Total Balance"
        label.textColor = .white.withAlphaComponent(0.9)
        label.font = .systemFont(ofSize: 14, weight: .medium)
        return label
    }()

    private let balanceLabel: UILabel = {
        let label = UILabel()
        label.text = "$0.00"
        label.textColor = .white
        label.font = .systemFont(ofSize: 36, weight: .bold)
        return label
    }()

    private let depositButton: UIButton = {
        let button = UIButton(type: .system)
        button.setTitle("Deposit", for: .normal)
        button.setTitleColor(.white, for: .normal)
        button.backgroundColor = UIColor.white.withAlphaComponent(0.2)
        button.layer.cornerRadius = 12
        return button
    }()

    private let withdrawButton: UIButton = {
        let button = UIButton(type: .system)
        button.setTitle("Withdraw", for: .normal)
        button.setTitleColor(.white, for: .normal)
        button.backgroundColor = UIColor.white.withAlphaComponent(0.2)
        button.layer.cornerRadius = 12
        return button
    }()

    private let transactionsTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "Recent Transactions"
        label.textColor = .white
        label.font = .systemFont(ofSize: 18, weight: .bold)
        return label
    }()

    private let transactionsTableView: UITableView = {
        let tableView = UITableView()
        tableView.backgroundColor = .clear
        tableView.separatorStyle = .none
        tableView.register(TransactionCell.self, forCellReuseIdentifier: "TransactionCell")
        return tableView
    }()

    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        setupBindings()
        viewModel.loadData()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Wallet"

        view.addSubview(scrollView)
        scrollView.addSubview(contentView)

        contentView.addSubview(balanceCard)
        balanceCard.addSubview(balanceTitleLabel)
        balanceCard.addSubview(balanceLabel)
        balanceCard.addSubview(depositButton)
        balanceCard.addSubview(withdrawButton)

        contentView.addSubview(transactionsTitleLabel)
        contentView.addSubview(transactionsTableView)

        scrollView.snp.makeConstraints { make in
            make.edges.equalTo(view.safeAreaLayoutGuide)
        }

        contentView.snp.makeConstraints { make in
            make.edges.equalToSuperview()
            make.width.equalTo(view)
        }

        balanceCard.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(16)
            make.leading.trailing.equalToSuperview().inset(16)
            make.height.equalTo(180)
        }

        balanceTitleLabel.snp.makeConstraints { make in
            make.top.leading.equalToSuperview().inset(20)
        }

        balanceLabel.snp.makeConstraints { make in
            make.top.equalTo(balanceTitleLabel.snp.bottom).offset(8)
            make.leading.equalToSuperview().inset(20)
        }

        depositButton.snp.makeConstraints { make in
            make.top.equalTo(balanceLabel.snp.bottom).offset(20)
            make.leading.equalToSuperview().inset(20)
            make.height.equalTo(44)
        }

        withdrawButton.snp.makeConstraints { make in
            make.top.equalTo(balanceLabel.snp.bottom).offset(20)
            make.leading.equalTo(depositButton.snp.trailing).offset(12)
            make.trailing.equalToSuperview().inset(20)
            make.width.equalTo(depositButton)
            make.height.equalTo(44)
        }

        transactionsTitleLabel.snp.makeConstraints { make in
            make.top.equalTo(balanceCard.snp.bottom).offset(24)
            make.leading.equalToSuperview().inset(16)
        }

        transactionsTableView.snp.makeConstraints { make in
            make.top.equalTo(transactionsTitleLabel.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview()
            make.height.equalTo(400)
            make.bottom.equalToSuperview().offset(-16)
        }

        transactionsTableView.delegate = self
        transactionsTableView.dataSource = self
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.updateUI(with: state)
            }
        }
    }

    private func updateUI(with state: WalletState) {
        if let balance = state.balance {
            balanceLabel.text = "$\(String(format: "%.2f", balance.balance))"
        }
        transactionsTableView.reloadData()
    }
}

extension WalletViewController: UITableViewDataSource, UITableViewDelegate {
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return viewModel.state.transactions.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "TransactionCell", for: indexPath) as! TransactionCell
        if indexPath.row < viewModel.state.transactions.count {
            cell.configure(with: viewModel.state.transactions[indexPath.row])
        }
        return cell
    }

    func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
        return 70
    }
}
