import UIKit
import SnapKit

class JackpotViewController: UIViewController {

    private var viewModel = JackpotViewModel()

    private let jackpotCard: UIView = {
        let view = UIView()
        view.backgroundColor = UIColor(hex: "#FF6B35")
        view.layer.cornerRadius = 16
        return view
    }()

    private let jackpotTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "PROGRESSIVE JACKPOT"
        label.textColor = .white.withAlphaComponent(0.9)
        label.font = .systemFont(ofSize: 14, weight: .medium)
        return label
    }()

    private let jackpotAmountLabel: UILabel = {
        let label = UILabel()
        label.text = "$0"
        label.textColor = .white
        label.font = .systemFont(ofSize: 36, weight: .bold)
        return label
    }()

    private let winnersTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "Recent Winners"
        label.textColor = .white
        label.font = .systemFont(ofSize: 18, weight: .bold)
        return label
    }()

    private let tableView: UITableView = {
        let tv = UITableView()
        tv.backgroundColor = .clear
        tv.separatorStyle = .none
        tv.register(JackpotWinnerCell.self, forCellReuseIdentifier: "JackpotWinnerCell")
        return tv
    }()

    private let activityIndicator: UIActivityIndicatorView = {
        let indicator = UIActivityIndicatorView(style: .large)
        indicator.color = UIColor(hex: "#FF6B35")
        indicator.hidesWhenStopped = true
        return indicator
    }()

    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        setupBindings()
        viewModel.loadData()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Jackpot"

        [jackpotCard, winnersTitleLabel, tableView, activityIndicator].forEach { view.addSubview($0) }
        jackpotCard.addSubview(jackpotTitleLabel)
        jackpotCard.addSubview(jackpotAmountLabel)

        jackpotCard.snp.makeConstraints { make in
            make.top.equalTo(view.safeAreaLayoutGuide).offset(16)
            make.leading.trailing.equalToSuperview().inset(16)
            make.height.equalTo(120)
        }

        jackpotTitleLabel.snp.makeConstraints { make in
            make.top.leading.equalToSuperview().inset(20)
        }

        jackpotAmountLabel.snp.makeConstraints { make in
            make.top.equalTo(jackpotTitleLabel.snp.bottom).offset(8)
            make.leading.equalToSuperview().inset(20)
        }

        winnersTitleLabel.snp.makeConstraints { make in
            make.top.equalTo(jackpotCard.snp.bottom).offset(24)
            make.leading.equalToSuperview().inset(16)
        }

        tableView.snp.makeConstraints { make in
            make.top.equalTo(winnersTitleLabel.snp.bottom).offset(12)
            make.leading.trailing.bottom.equalToSuperview()
        }

        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        tableView.delegate = self
        tableView.dataSource = self
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.updateUI(with: state)
            }
        }
    }

    private func updateUI(with state: JackpotState) {
        if state.isLoading {
            activityIndicator.startAnimating()
        } else {
            activityIndicator.stopAnimating()
        }

        jackpotAmountLabel.text = "$\(String(format: "%,.0f", state.jackpotAmount))"
        tableView.reloadData()
    }
}

// MARK: - UITableView

extension JackpotViewController: UITableViewDataSource, UITableViewDelegate {
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return viewModel.state.winners.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "JackpotWinnerCell", for: indexPath) as! JackpotWinnerCell
        if indexPath.row < viewModel.state.winners.count {
            cell.configure(with: viewModel.state.winners[indexPath.row])
        }
        return cell
    }

    func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
        return 60
    }
}

// MARK: - JackpotWinnerCell

class JackpotWinnerCell: UITableViewCell {
    private let usernameLabel = UILabel()
    private let amountLabel = UILabel()
    private let gameLabel = UILabel()

    override init(style: UITableViewCell.CellStyle, reuseIdentifier: String?) {
        super.init(style: style, reuseIdentifier: reuseIdentifier)
        backgroundColor = UIColor(hex: "#1E1E3F")
        selectionStyle = .none
        layer.cornerRadius = 8

        usernameLabel.textColor = .white
        usernameLabel.font = .systemFont(ofSize: 14, weight: .medium)
        amountLabel.textColor = UIColor(hex: "#FFD700")
        amountLabel.font = .systemFont(ofSize: 14, weight: .bold)
        amountLabel.textAlignment = .right
        gameLabel.textColor = .white.withAlphaComponent(0.6)
        gameLabel.font = .systemFont(ofSize: 12)

        [usernameLabel, amountLabel, gameLabel].forEach { contentView.addSubview($0) }

        usernameLabel.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(10)
            make.leading.equalToSuperview().inset(16)
        }

        amountLabel.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(10)
            make.trailing.equalToSuperview().inset(16)
        }

        gameLabel.snp.makeConstraints { make in
            make.top.equalTo(usernameLabel.snp.bottom).offset(4)
            make.leading.equalToSuperview().inset(16)
        }
    }

    required init?(coder: NSCoder) { fatalError() }

    func configure(with winner: JackpotWinner) {
        usernameLabel.text = winner.username
        amountLabel.text = "+$\(String(format: "%.2f", winner.amount))"
        gameLabel.text = winner.gameName
    }
}
