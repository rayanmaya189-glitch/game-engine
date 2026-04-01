import UIKit
import SnapKit

class TransactionCell: UITableViewCell {

    private let iconView: UIView = {
        let view = UIView()
        view.layer.cornerRadius = 8
        return view
    }()

    private let iconImageView: UIImageView = {
        let imageView = UIImageView()
        imageView.tintColor = .white
        return imageView
    }()

    private let typeLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white
        label.font = .systemFont(ofSize: 14, weight: .medium)
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
        label.textColor = .white
        label.font = .systemFont(ofSize: 14, weight: .bold)
        label.textAlignment = .right
        return label
    }()

    private let statusLabel: UILabel = {
        let label = UILabel()
        label.font = .systemFont(ofSize: 12)
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

        contentView.addSubview(iconView)
        iconView.addSubview(iconImageView)
        contentView.addSubview(typeLabel)
        contentView.addSubview(dateLabel)
        contentView.addSubview(amountLabel)
        contentView.addSubview(statusLabel)

        iconView.snp.makeConstraints { make in
            make.leading.equalToSuperview().inset(16)
            make.centerY.equalToSuperview()
            make.size.equalTo(40)
        }

        iconImageView.snp.makeConstraints { make in
            make.center.equalToSuperview()
            make.size.equalTo(20)
        }

        typeLabel.snp.makeConstraints { make in
            make.leading.equalTo(iconView.snp.trailing).offset(12)
            make.top.equalToSuperview().inset(16)
        }

        dateLabel.snp.makeConstraints { make in
            make.leading.equalTo(iconView.snp.trailing).offset(12)
            make.top.equalTo(typeLabel.snp.bottom).offset(4)
        }

        amountLabel.snp.makeConstraints { make in
            make.trailing.equalToSuperview().inset(16)
            make.top.equalToSuperview().inset(16)
        }

        statusLabel.snp.makeConstraints { make in
            make.trailing.equalToSuperview().inset(16)
            make.top.equalTo(amountLabel.snp.bottom).offset(4)
        }
    }

    func configure(with transaction: Transaction) {
        typeLabel.text = transaction.type.capitalized
        dateLabel.text = String(transaction.createdAt.prefix(10))
        amountLabel.text = "\(transaction.type.lowercased() == "deposit" || transaction.type.lowercased() == "win" ? "+" : "-")$\(String(format: "%.2f", transaction.amount))"

        let color: UIColor
        switch transaction.type.lowercased() {
        case "deposit", "win":
            color = UIColor(hex: "#4CAF50")
            iconView.backgroundColor = UIColor(hex: "#4CAF50").withAlphaComponent(0.2)
            iconImageView.image = UIImage(systemName: "arrow.down")
        case "withdraw":
            color = UIColor(hex: "#FF6B35")
            iconView.backgroundColor = UIColor(hex: "#FF6B35").withAlphaComponent(0.2)
            iconImageView.image = UIImage(systemName: "arrow.up")
        default:
            color = .white
            iconView.backgroundColor = UIColor.white.withAlphaComponent(0.2)
            iconImageView.image = UIImage(systemName: "arrow.left.arrow.right")
        }

        amountLabel.textColor = color
        iconImageView.tintColor = color

        switch transaction.status.lowercased() {
        case "completed":
            statusLabel.text = "Completed"
            statusLabel.textColor = UIColor(hex: "#4CAF50")
        case "pending":
            statusLabel.text = "Pending"
            statusLabel.textColor = UIColor(hex: "#FFC107")
        case "failed":
            statusLabel.text = "Failed"
            statusLabel.textColor = UIColor(hex: "#FF5252")
        default:
            statusLabel.text = transaction.status.capitalized
            statusLabel.textColor = .white.withAlphaComponent(0.6)
        }
    }
}
