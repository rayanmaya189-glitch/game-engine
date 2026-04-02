import UIKit
import SnapKit

class LiveDealerViewController: UIViewController {

    private var viewModel = LiveDealerViewModel()

    private lazy var collectionView: UICollectionView = {
        let layout = UICollectionViewFlowLayout()
        layout.scrollDirection = .vertical
        layout.minimumInteritemSpacing = 12
        layout.minimumLineSpacing = 12
        layout.sectionInset = UIEdgeInsets(top: 16, left: 16, bottom: 16, right: 16)
        let cv = UICollectionView(frame: .zero, collectionViewLayout: layout)
        cv.backgroundColor = .clear
        cv.register(LiveDealerCell.self, forCellWithReuseIdentifier: "LiveDealerCell")
        return cv
    }()

    private let emptyLabel: UILabel = {
        let label = UILabel()
        label.text = "No live tables available"
        label.textAlignment = .center
        label.textColor = .white.withAlphaComponent(0.5)
        label.font = .systemFont(ofSize: 16)
        label.isHidden = true
        return label
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
        viewModel.loadTables()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Live Dealer"

        [collectionView, emptyLabel, activityIndicator].forEach { view.addSubview($0) }

        collectionView.snp.makeConstraints { make in
            make.edges.equalTo(view.safeAreaLayoutGuide)
        }

        emptyLabel.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        collectionView.delegate = self
        collectionView.dataSource = self
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.collectionView.reloadData()
                self?.activityIndicator.stopAnimating()
                self?.emptyLabel.isHidden = !state.tables.isEmpty
                if let error = state.error {
                    let alert = UIAlertController(title: "Error", message: error, preferredStyle: .alert)
                    alert.addAction(UIAlertAction(title: "OK", style: .cancel))
                    self?.present(alert, animated: true)
                }
            }
        }
    }
}

// MARK: - UICollectionView

extension LiveDealerViewController: UICollectionViewDataSource, UICollectionViewDelegateFlowLayout {
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return viewModel.state.tables.count
    }

    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let cell = collectionView.dequeueReusableCell(withReuseIdentifier: "LiveDealerCell", for: indexPath) as! LiveDealerCell
        let table = viewModel.state.tables[indexPath.row]
        cell.configure(with: table)
        cell.onJoinTapped = { [weak self] in
            self?.viewModel.joinTable(table.id)
        }
        return cell
    }

    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, sizeForItemAt indexPath: IndexPath) -> CGSize {
        let width = collectionView.bounds.width - 32
        return CGSize(width: width, height: 160)
    }
}

// MARK: - LiveDealerCell

class LiveDealerCell: UICollectionViewCell {
    var onJoinTapped: (() -> Void)?

    private let previewView: UIView = {
        let view = UIView()
        view.backgroundColor = UIColor(hex: "#1A1A3E")
        view.layer.cornerRadius = 12
        view.clipsToBounds = true
        return view
    }()

    private let previewIcon: UILabel = {
        let label = UILabel()
        label.text = "📹"
        label.font = .systemFont(ofSize: 36)
        label.textAlignment = .center
        return label
    }()

    private let liveIndicator: UIView = {
        let view = UIView()
        view.backgroundColor = .systemRed
        view.layer.cornerRadius = 6
        return view
    }()

    private let liveLabel: UILabel = {
        let label = UILabel()
        label.text = "LIVE"
        label.textColor = .white
        label.font = .systemFont(ofSize: 10, weight: .bold)
        return label
    }()

    private let gameTypeLabel = UILabel()
    private let dealerLabel = UILabel()
    private let playersLabel = UILabel()
    private let betRangeLabel = UILabel()

    private let joinButton: UIButton = {
        let btn = UIButton(type: .system)
        btn.setTitle("JOIN TABLE", for: .normal)
        btn.setTitleColor(.white, for: .normal)
        btn.backgroundColor = UIColor(hex: "#FF6B35")
        btn.layer.cornerRadius = 8
        btn.titleLabel?.font = .systemFont(ofSize: 14, weight: .semibold)
        return btn
    }()

    override init(frame: CGRect) {
        super.init(frame: frame)
        setupUI()
    }

    required init?(coder: NSCoder) { fatalError() }

    private func setupUI() {
        backgroundColor = UIColor(hex: "#1E1E3F")
        layer.cornerRadius = 12

        gameTypeLabel.textColor = .white
        gameTypeLabel.font = .systemFont(ofSize: 18, weight: .bold)
        dealerLabel.textColor = .white.withAlphaComponent(0.7)
        dealerLabel.font = .systemFont(ofSize: 13)
        playersLabel.textColor = UIColor(hex: "#4CAF50")
        playersLabel.font = .systemFont(ofSize: 12)
        betRangeLabel.textColor = .white.withAlphaComponent(0.6)
        betRangeLabel.font = .systemFont(ofSize: 12)

        [previewView, gameTypeLabel, dealerLabel, playersLabel, betRangeLabel, joinButton].forEach {
            contentView.addSubview($0)
        }
        previewView.addSubview(previewIcon)
        previewView.addSubview(liveIndicator)
        previewView.addSubview(liveLabel)

        previewView.snp.makeConstraints { make in
            make.top.leading.equalToSuperview().inset(12)
            make.width.equalTo(100)
            make.height.equalTo(70)
        }

        previewIcon.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        liveIndicator.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(6)
            make.leading.equalToSuperview().offset(6)
            make.size.equalTo(12)
        }

        liveLabel.snp.makeConstraints { make in
            make.centerY.equalTo(liveIndicator)
            make.leading.equalTo(liveIndicator.snp.trailing).offset(4)
        }

        gameTypeLabel.snp.makeConstraints { make in
            make.top.equalToSuperview().inset(12)
            make.leading.equalTo(previewView.snp.trailing).offset(12)
            make.trailing.equalToSuperview().inset(12)
        }

        dealerLabel.snp.makeConstraints { make in
            make.top.equalTo(gameTypeLabel.snp.bottom).offset(4)
            make.leading.equalTo(previewView.snp.trailing).offset(12)
            make.trailing.equalToSuperview().inset(12)
        }

        playersLabel.snp.makeConstraints { make in
            make.top.equalTo(dealerLabel.snp.bottom).offset(4)
            make.leading.equalTo(previewView.snp.trailing).offset(12)
        }

        betRangeLabel.snp.makeConstraints { make in
            make.top.equalTo(playersLabel.snp.bottom).offset(2)
            make.leading.equalTo(previewView.snp.trailing).offset(12)
        }

        joinButton.snp.makeConstraints { make in
            make.leading.trailing.equalToSuperview().inset(12)
            make.bottom.equalToSuperview().inset(12)
            make.height.equalTo(40)
        }

        joinButton.addTarget(self, action: #selector(joinTapped), for: .touchUpInside)
    }

    func configure(with table: LiveDealerTable) {
        gameTypeLabel.text = table.gameType.capitalized
        dealerLabel.text = "Dealer: \(table.dealerName)"
        playersLabel.text = "\(table.currentPlayers)/\(table.maxPlayers) players"
        betRangeLabel.text = "$\(String(format: "%.0f", table.minBet)) - $\(String(format: "%.0f", table.maxBet))"
        joinButton.isEnabled = table.isActive
        joinButton.alpha = table.isActive ? 1.0 : 0.5
        previewIcon.alpha = table.isActive ? 1.0 : 0.4
    }

    @objc private func joinTapped() {
        onJoinTapped?()
    }
}
