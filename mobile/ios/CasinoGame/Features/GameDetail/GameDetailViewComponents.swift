import UIKit
import SnapKit

class BigWinCell: UICollectionViewCell {
    private let usernameLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white
        label.font = .systemFont(ofSize: 12, weight: .medium)
        return label
    }()

    private let amountLabel: UILabel = {
        let label = UILabel()
        label.textColor = UIColor(hex: "#FFD700")
        label.font = .systemFont(ofSize: 16, weight: .bold)
        return label
    }()

    private let multiplierLabel: UILabel = {
        let label = UILabel()
        label.textColor = UIColor(hex: "#4CAF50")
        label.font = .systemFont(ofSize: 11)
        return label
    }()

    override init(frame: CGRect) {
        super.init(frame: frame)
        backgroundColor = UIColor(hex: "#1E1E3F")
        layer.cornerRadius = 10
        addSubview(usernameLabel)
        addSubview(amountLabel)
        addSubview(multiplierLabel)

        amountLabel.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(10)
            make.leading.trailing.equalToSuperview().inset(10)
        }
        usernameLabel.snp.makeConstraints { make in
            make.top.equalTo(amountLabel.snp.bottom).offset(4)
            make.leading.trailing.equalToSuperview().inset(10)
        }
        multiplierLabel.snp.makeConstraints { make in
            make.top.equalTo(usernameLabel.snp.bottom).offset(2)
            make.leading.trailing.equalToSuperview().inset(10)
        }
    }

    required init?(coder: NSCoder) { fatalError() }

    func configure(with win: RecentWin) {
        usernameLabel.text = "***\(win.username.suffix(3))"
        amountLabel.text = "$\(String(format: "%.0f", win.amount))"
        multiplierLabel.text = "\(String(format: "%.0f", win.multiplier))x"
    }
}

class RelatedGameCell: UICollectionViewCell {
    private let thumbnailView: UIImageView = {
        let imageView = UIImageView()
        imageView.contentMode = .scaleAspectFill
        imageView.clipsToBounds = true
        imageView.backgroundColor = UIColor(hex: "#1E1E3F")
        imageView.layer.cornerRadius = 10
        return imageView
    }()

    private let nameLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white
        label.font = .systemFont(ofSize: 11, weight: .medium)
        label.textAlignment = .center
        label.numberOfLines = 2
        return label
    }()

    override init(frame: CGRect) {
        super.init(frame: frame)
        addSubview(thumbnailView)
        addSubview(nameLabel)
        thumbnailView.snp.makeConstraints { make in
            make.top.leading.trailing.equalToSuperview()
            make.height.equalTo(80)
        }
        nameLabel.snp.makeConstraints { make in
            make.top.equalTo(thumbnailView.snp.bottom).offset(4)
            make.leading.trailing.equalToSuperview()
        }
    }

    required init?(coder: NSCoder) { fatalError() }

    func configure(with game: Game) {
        nameLabel.text = game.name
    }
}
