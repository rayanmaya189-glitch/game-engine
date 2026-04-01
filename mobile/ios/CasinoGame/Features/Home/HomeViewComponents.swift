import UIKit
import SnapKit

class GameCollectionCell: UICollectionViewCell {

    private let thumbnailImageView: UIImageView = {
        let imageView = UIImageView()
        imageView.contentMode = .scaleAspectFill
        imageView.clipsToBounds = true
        imageView.backgroundColor = UIColor(hex: "#1E1E3F")
        imageView.layer.cornerRadius = 12
        return imageView
    }()

    private let nameLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white
        label.font = .systemFont(ofSize: 14, weight: .medium)
        label.textAlignment = .center
        return label
    }()

    private let badgeLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white
        label.font = .systemFont(ofSize: 10, weight: .bold)
        label.textAlignment = .center
        label.backgroundColor = UIColor(hex: "#FF6B35")
        label.layer.cornerRadius = 4
        label.clipsToBounds = true
        label.isHidden = true
        return label
    }()

    override init(frame: CGRect) {
        super.init(frame: frame)
        setupUI()
    }

    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    private func setupUI() {
        contentView.addSubview(thumbnailImageView)
        contentView.addSubview(nameLabel)
        contentView.addSubview(badgeLabel)

        thumbnailImageView.snp.makeConstraints { make in
            make.top.leading.trailing.equalToSuperview()
            make.height.equalTo(140)
        }

        nameLabel.snp.makeConstraints { make in
            make.top.equalTo(thumbnailImageView.snp.bottom).offset(8)
            make.leading.trailing.equalToSuperview()
        }

        badgeLabel.snp.makeConstraints { make in
            make.top.trailing.equalTo(thumbnailImageView).inset(8)
            make.width.equalTo(50)
            make.height.equalTo(20)
        }
    }

    func configure(with game: Game) {
        nameLabel.text = game.name

        if let urlString = game.thumbnailUrl, let url = URL(string: urlString) {
        }

        if game.isFeatured {
            badgeLabel.text = "FEATURED"
            badgeLabel.backgroundColor = UIColor(hex: "#FFD700")
            badgeLabel.isHidden = false
        } else if game.isNew {
            badgeLabel.text = "NEW"
            badgeLabel.backgroundColor = UIColor(hex: "#4CAF50")
            badgeLabel.isHidden = false
        } else if game.isHot {
            badgeLabel.text = "HOT"
            badgeLabel.backgroundColor = UIColor(hex: "#FF6B35")
            badgeLabel.isHidden = false
        } else {
            badgeLabel.isHidden = true
        }
    }
}
