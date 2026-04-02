import UIKit
import SnapKit

// MARK: - PaymentTransactionCell

class PaymentTransactionCell: UITableViewCell {

    private let iconView: UIView = {
        let view = UIView()
        view.layer.cornerRadius = 20
        return view
    }()

    private let iconImageView: UIImageView = {
        let imageView = UIImageView()
        imageView.tintColor = .white
        imageView.contentMode = .scaleAspectFit
        return imageView
    }()

    private let typeLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white
        label.font = .systemFont(ofSize: 15, weight: .medium)
        return label
    }()

    private let dateLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white.withAlphaComponent(0.6)
        label.font = .systemFont(ofSize: 12)
        return label
    }()

    private let amountLabel: UILabel = {
        let label = UILabel()
        label.font = .systemFont(ofSize: 15, weight: .bold)
        label.textAlignment = .right
        return label
    }()

    private let statusBadge: UILabel = {
        let label = UILabel()
        label.font = .systemFont(ofSize: 11, weight: .semibold)
        label.textAlignment = .right
        label.layer.cornerRadius = 4
        label.clipsToBounds = true
        return label
    }()

    override init(style: UITableViewCell.CellStyle, reuseIdentifier: String?) {
        super.init(style: style, reuseIdentifier: reuseIdentifier)
        setupUI()
    }

    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    private func setupUI() {
        backgroundColor = UIColor(hex: "#1E1E3F")
        selectionStyle = .none
        layer.cornerRadius = 12

        contentView.addSubview(iconView)
        iconView.addSubview(iconImageView)
        contentView.addSubview(typeLabel)
        contentView.addSubview(dateLabel)
        contentView.addSubview(amountLabel)
        contentView.addSubview(statusBadge)

        iconView.snp.makeConstraints { make in
            make.leading.equalToSuperview().inset(12)
            make.centerY.equalToSuperview()
            make.size.equalTo(40)
        }

        iconImageView.snp.makeConstraints { make in
            make.center.equalToSuperview()
            make.size.equalTo(20)
        }

        typeLabel.snp.makeConstraints { make in
            make.leading.equalTo(iconView.snp.trailing).offset(12)
            make.top.equalToSuperview().inset(14)
        }

        dateLabel.snp.makeConstraints { make in
            make.leading.equalTo(iconView.snp.trailing).offset(12)
            make.top.equalTo(typeLabel.snp.bottom).offset(4)
        }

        amountLabel.snp.makeConstraints { make in
            make.trailing.equalToSuperview().inset(12)
            make.top.equalToSuperview().inset(14)
        }

        statusBadge.snp.makeConstraints { make in
            make.trailing.equalToSuperview().inset(12)
            make.top.equalTo(amountLabel.snp.bottom).offset(6)
            make.width.greaterThanOrEqualTo(60)
        }
    }

    func configure(with transaction: Transaction) {
        typeLabel.text = transaction.type.capitalized
        dateLabel.text = String(transaction.createdAt.prefix(10))
        amountLabel.text = "\(transaction.type.lowercased() == "deposit" ? "+" : "-")$\(String(format: "%.2f", transaction.amount))"

        let isDeposit = transaction.type.lowercased() == "deposit"
        let color: UIColor = isDeposit ? UIColor(hex: "#4CAF50") : UIColor(hex: "#FF6B35")
        iconView.backgroundColor = color.withAlphaComponent(0.2)
        iconImageView.image = UIImage(systemName: isDeposit ? "arrow.down.circle.fill" : "arrow.up.circle.fill")
        iconImageView.tintColor = color
        amountLabel.textColor = color

        switch transaction.status.lowercased() {
        case "completed":
            statusBadge.text = "  Completed  "
            statusBadge.textColor = UIColor(hex: "#4CAF50")
            statusBadge.backgroundColor = UIColor(hex: "#4CAF50").withAlphaComponent(0.15)
        case "pending":
            statusBadge.text = "  Pending  "
            statusBadge.textColor = UIColor(hex: "#FFC107")
            statusBadge.backgroundColor = UIColor(hex: "#FFC107").withAlphaComponent(0.15)
        case "failed":
            statusBadge.text = "  Failed  "
            statusBadge.textColor = UIColor(hex: "#FF5252")
            statusBadge.backgroundColor = UIColor(hex: "#FF5252").withAlphaComponent(0.15)
        default:
            statusBadge.text = "  \(transaction.status.capitalized)  "
            statusBadge.textColor = .white.withAlphaComponent(0.6)
            statusBadge.backgroundColor = UIColor.white.withAlphaComponent(0.1)
        }
    }
}

// MARK: - BetHistoryCell

class BetHistoryCell: UITableViewCell {

    private let containerView: UIView = {
        let view = UIView()
        view.backgroundColor = UIColor(hex: "#1E1E3F")
        view.layer.cornerRadius = 12
        return view
    }()

    private let resultIndicator: UIView = {
        let view = UIView()
        view.layer.cornerRadius = 4
        return view
    }()

    private let gameNameLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white
        label.font = .systemFont(ofSize: 15, weight: .medium)
        return label
    }()

    private let dateLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white.withAlphaComponent(0.6)
        label.font = .systemFont(ofSize: 12)
        return label
    }()

    private let stakeLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white.withAlphaComponent(0.7)
        label.font = .systemFont(ofSize: 13)
        label.textAlignment = .right
        return label
    }()

    private let resultLabel: UILabel = {
        let label = UILabel()
        label.font = .systemFont(ofSize: 14, weight: .bold)
        label.textAlignment = .right
        return label
    }()

    override init(style: UITableViewCell.CellStyle, reuseIdentifier: String?) {
        super.init(style: style, reuseIdentifier: reuseIdentifier)
        setupUI()
    }

    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    private func setupUI() {
        backgroundColor = .clear
        selectionStyle = .none

        contentView.addSubview(containerView)
        containerView.addSubview(resultIndicator)
        containerView.addSubview(gameNameLabel)
        containerView.addSubview(dateLabel)
        containerView.addSubview(stakeLabel)
        containerView.addSubview(resultLabel)

        containerView.snp.makeConstraints { make in
            make.edges.equalToSuperview().inset(UIEdgeInsets(top: 4, left: 16, bottom: 4, right: 16))
        }

        resultIndicator.snp.makeConstraints { make in
            make.leading.equalToSuperview().inset(12)
            make.centerY.equalToSuperview()
            make.width.equalTo(6)
            make.height.equalTo(40)
        }

        gameNameLabel.snp.makeConstraints { make in
            make.leading.equalTo(resultIndicator.snp.trailing).offset(12)
            make.top.equalToSuperview().inset(14)
        }

        dateLabel.snp.makeConstraints { make in
            make.leading.equalTo(resultIndicator.snp.trailing).offset(12)
            make.top.equalTo(gameNameLabel.snp.bottom).offset(4)
        }

        resultLabel.snp.makeConstraints { make in
            make.trailing.equalToSuperview().inset(12)
            make.top.equalToSuperview().inset(12)
        }

        stakeLabel.snp.makeConstraints { make in
            make.trailing.equalToSuperview().inset(12)
            make.top.equalTo(resultLabel.snp.bottom).offset(4)
        }
    }

    func configure(with bet: Bet) {
        gameNameLabel.text = bet.gameName
        dateLabel.text = String(bet.placedAt.prefix(10))
        stakeLabel.text = "Stake: $\(String(format: "%.2f", bet.stake))"

        let resultColor: UIColor
        switch bet.result.lowercased() {
        case "won":
            resultLabel.text = "+$\(String(format: "%.2f", bet.winAmount ?? 0))"
            resultColor = UIColor(hex: "#4CAF50")
            resultIndicator.backgroundColor = UIColor(hex: "#4CAF50")
        case "lost":
            resultLabel.text = "-$\(String(format: "%.2f", bet.stake))"
            resultColor = UIColor(hex: "#FF5252")
            resultIndicator.backgroundColor = UIColor(hex: "#FF5252")
        case "pending":
            resultLabel.text = "Pending"
            resultColor = UIColor(hex: "#FFC107")
            resultIndicator.backgroundColor = UIColor(hex: "#FFC107")
        default:
            resultLabel.text = bet.result.capitalized
            resultColor = .white.withAlphaComponent(0.6)
            resultIndicator.backgroundColor = .white.withAlphaComponent(0.3)
        }

        resultLabel.textColor = resultColor
    }
}
