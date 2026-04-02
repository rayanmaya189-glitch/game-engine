import UIKit
import SnapKit

class ReferralViewController: UIViewController {

    private var viewModel = ReferralViewModel()

    private let codeLabel: UILabel = {
        let label = UILabel()
        label.textColor = UIColor(hex: "#FFD700")
        label.font = .systemFont(ofSize: 28, weight: .bold)
        label.textAlignment = .center
        return label
    }()

    private let copyButton: UIButton = {
        let btn = UIButton(type: .system)
        btn.setTitle("Copy Code", for: .normal)
        btn.setTitleColor(.white, for: .normal)
        btn.backgroundColor = UIColor(hex: "#FF6B35")
        btn.layer.cornerRadius = 12
        btn.titleLabel?.font = .systemFont(ofSize: 15, weight: .semibold)
        return btn
    }()

    private let shareButton: UIButton = {
        let btn = UIButton(type: .system)
        btn.setTitle("Share", for: .normal)
        btn.setTitleColor(.white, for: .normal)
        btn.backgroundColor = UIColor(hex: "#4CAF50")
        btn.layer.cornerRadius = 12
        btn.titleLabel?.font = .systemFont(ofSize: 15, weight: .semibold)
        return btn
    }()

    private let referralsLabel = UILabel()
    private let earningsLabel = UILabel()

    private let historyTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "Referral History"
        label.textColor = .white
        label.font = .systemFont(ofSize: 18, weight: .bold)
        return label
    }()

    private let tableView: UITableView = {
        let tv = UITableView()
        tv.backgroundColor = .clear
        tv.separatorStyle = .none
        tv.register(ReferralHistoryCell.self, forCellReuseIdentifier: "ReferralHistoryCell")
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
        viewModel.loadReferralData()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Refer & Earn"

        let statsStack = UIStackView(arrangedSubviews: [referralsLabel, earningsLabel])
        statsStack.axis = .horizontal
        statsStack.distribution = .fillEqually
        statsStack.spacing = 12

        let buttonStack = UIStackView(arrangedSubviews: [copyButton, shareButton])
        buttonStack.axis = .horizontal
        buttonStack.spacing = 12
        buttonStack.distribution = .fillEqually

        [codeLabel, buttonStack, statsStack, historyTitleLabel, tableView, activityIndicator].forEach {
            view.addSubview($0)
        }

        codeLabel.snp.makeConstraints { make in
            make.top.equalTo(view.safeAreaLayoutGuide).offset(24)
            make.centerX.equalToSuperview()
        }

        buttonStack.snp.makeConstraints { make in
            make.top.equalTo(codeLabel.snp.bottom).offset(16)
            make.leading.trailing.equalToSuperview().inset(16)
            make.height.equalTo(44)
        }

        statsStack.snp.makeConstraints { make in
            make.top.equalTo(buttonStack.snp.bottom).offset(20)
            make.leading.trailing.equalToSuperview().inset(16)
            make.height.equalTo(60)
        }

        historyTitleLabel.snp.makeConstraints { make in
            make.top.equalTo(statsStack.snp.bottom).offset(24)
            make.leading.equalToSuperview().inset(16)
        }

        tableView.snp.makeConstraints { make in
            make.top.equalTo(historyTitleLabel.snp.bottom).offset(12)
            make.leading.trailing.bottom.equalToSuperview()
        }

        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        configureStatsLabel(referralsLabel, title: "Referrals", value: "0")
        configureStatsLabel(earningsLabel, title: "Earnings", value: "$0.00")

        tableView.delegate = self
        tableView.dataSource = self

        copyButton.addTarget(self, action: #selector(copyTapped), for: .touchUpInside)
        shareButton.addTarget(self, action: #selector(shareTapped), for: .touchUpInside)
    }

    private func configureStatsLabel(_ label: UILabel, title: String, value: String) {
        label.numberOfLines = 2
        label.textAlignment = .center
        label.backgroundColor = UIColor(hex: "#1E1E3F")
        label.layer.cornerRadius = 12
        label.clipsToBounds = true
        updateStatsLabel(label, title: title, value: value)
    }

    private func updateStatsLabel(_ label: UILabel, title: String, value: String) {
        let attr = NSMutableAttributedString(
            string: "\(value)\n",
            attributes: [.font: UIFont.systemFont(ofSize: 22, weight: .bold), .foregroundColor: UIColor.white]
        )
        attr.append(NSAttributedString(
            string: title,
            attributes: [.font: UIFont.systemFont(ofSize: 13), .foregroundColor: UIColor.white.withAlphaComponent(0.6)]
        ))
        label.attributedText = attr
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.codeLabel.text = state.referralCode
                if let stats = state.stats {
                    self?.updateStatsLabel(self?.referralsLabel ?? UILabel(), title: "Referrals", value: "\(stats.totalReferrals)")
                    self?.updateStatsLabel(self?.earningsLabel ?? UILabel(), title: "Earnings", value: "$\(String(format: "%.2f", stats.totalEarnings))")
                }
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

    @objc private func copyTapped() {
        UIPasteboard.general.string = viewModel.state.referralCode
        let alert = UIAlertController(title: "Copied!", message: "Referral code copied to clipboard.", preferredStyle: .alert)
        alert.addAction(UIAlertAction(title: "OK", style: .cancel))
        present(alert, animated: true)
    }

    @objc private func shareTapped() {
        let shareUrl = viewModel.state.shareUrl.isEmpty
            ? "Join using my referral code: \(viewModel.state.referralCode)"
            : viewModel.state.shareUrl
        let activityVC = UIActivityViewController(activityItems: [shareUrl], applicationActivities: nil)
        present(activityVC, animated: true)
    }
}

// MARK: - UITableView

extension ReferralViewController: UITableViewDataSource, UITableViewDelegate {
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return viewModel.state.history.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "ReferralHistoryCell", for: indexPath) as! ReferralHistoryCell
        cell.configure(with: viewModel.state.history[indexPath.row])
        return cell
    }

    func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
        return 60
    }
}

// MARK: - ReferralHistoryCell

class ReferralHistoryCell: UITableViewCell {
    private let usernameLabel = UILabel()
    private let statusLabel = UILabel()
    private let amountLabel = UILabel()

    override init(style: UITableViewCell.CellStyle, reuseIdentifier: String?) {
        super.init(style: style, reuseIdentifier: reuseIdentifier)
        backgroundColor = UIColor(hex: "#1E1E3F")
        selectionStyle = .none
        layer.cornerRadius = 8

        usernameLabel.textColor = .white
        usernameLabel.font = .systemFont(ofSize: 14, weight: .medium)
        statusLabel.textColor = .white.withAlphaComponent(0.6)
        statusLabel.font = .systemFont(ofSize: 12)
        amountLabel.textColor = UIColor(hex: "#4CAF50")
        amountLabel.font = .systemFont(ofSize: 14, weight: .bold)
        amountLabel.textAlignment = .right

        [usernameLabel, statusLabel, amountLabel].forEach { contentView.addSubview($0) }

        usernameLabel.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(10)
            make.leading.equalToSuperview().inset(16)
        }

        statusLabel.snp.makeConstraints { make in
            make.top.equalTo(usernameLabel.snp.bottom).offset(4)
            make.leading.equalToSuperview().inset(16)
        }

        amountLabel.snp.makeConstraints { make in
            make.centerY.equalToSuperview()
            make.trailing.equalToSuperview().inset(16)
        }
    }

    required init?(coder: NSCoder) { fatalError() }

    func configure(with entry: ReferralEntry) {
        usernameLabel.text = entry.username
        statusLabel.text = entry.status.capitalized
        amountLabel.text = "+$\(String(format: "%.2f", entry.earnedAmount))"
    }
}
