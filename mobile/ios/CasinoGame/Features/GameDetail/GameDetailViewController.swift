import UIKit
import SnapKit

class GameDetailViewController: UIViewController {

    private let gameId: String
    private lazy var viewModel = GameDetailViewModel(gameId: gameId)

    init(gameId: String) {
        self.gameId = gameId
        super.init(nibName: nil, bundle: nil)
    }

    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    private let scrollView: UIScrollView = {
        let scrollView = UIScrollView()
        scrollView.showsVerticalScrollIndicator = false
        return scrollView
    }()

    private let contentView = UIView()

    private let bannerImageView: UIImageView = {
        let imageView = UIImageView()
        imageView.contentMode = .scaleAspectFill
        imageView.clipsToBounds = true
        imageView.backgroundColor = UIColor(hex: "#1E1E3F")
        imageView.layer.cornerRadius = 16
        return imageView
    }()

    private let nameLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white
        label.font = .systemFont(ofSize: 22, weight: .bold)
        return label
    }()

    private let providerLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white.withAlphaComponent(0.7)
        label.font = .systemFont(ofSize: 14)
        return label
    }()

    private let infoStack: UIStackView = {
        let stack = UIStackView()
        stack.distribution = .fillEqually
        stack.spacing = 8
        return stack
    }()

    private let winsTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "Recent Big Wins"
        label.textColor = .white
        label.font = .systemFont(ofSize: 18, weight: .bold)
        return label
    }()

    private lazy var winsCollectionView: UICollectionView = {
        let layout = UICollectionViewFlowLayout()
        layout.scrollDirection = .horizontal
        layout.minimumInteritemSpacing = 12
        layout.sectionInset = UIEdgeInsets(top: 0, left: 16, bottom: 0, right: 16)
        let cv = UICollectionView(frame: .zero, collectionViewLayout: layout)
        cv.backgroundColor = .clear
        cv.showsHorizontalScrollIndicator = false
        cv.register(BigWinCell.self, forCellWithReuseIdentifier: "BigWinCell")
        return cv
    }()

    private let relatedTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "Related Games"
        label.textColor = .white
        label.font = .systemFont(ofSize: 18, weight: .bold)
        return label
    }()

    private lazy var relatedCollectionView: UICollectionView = {
        let layout = UICollectionViewFlowLayout()
        layout.scrollDirection = .horizontal
        layout.minimumInteritemSpacing = 12
        layout.sectionInset = UIEdgeInsets(top: 0, left: 16, bottom: 0, right: 16)
        let cv = UICollectionView(frame: .zero, collectionViewLayout: layout)
        cv.backgroundColor = .clear
        cv.showsHorizontalScrollIndicator = false
        cv.register(RelatedGameCell.self, forCellWithReuseIdentifier: "RelatedGameCell")
        return cv
    }()

    private let playButton: UIButton = {
        let button = UIButton(type: .system)
        button.setTitle("Play Now", for: .normal)
        button.setTitleColor(.white, for: .normal)
        button.backgroundColor = UIColor(hex: "#FF6B35")
        button.layer.cornerRadius = 14
        button.titleLabel?.font = .systemFont(ofSize: 18, weight: .bold)
        return button
    }()

    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        setupBindings()
        viewModel.loadGameDetail()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Game Details"

        view.addSubview(scrollView)
        scrollView.addSubview(contentView)

        [bannerImageView, nameLabel, providerLabel, infoStack,
         winsTitleLabel, winsCollectionView, relatedTitleLabel,
         relatedCollectionView, playButton].forEach { contentView.addSubview($0) }

        scrollView.snp.makeConstraints { make in
            make.edges.equalTo(view.safeAreaLayoutGuide)
        }

        contentView.snp.makeConstraints { make in
            make.edges.equalToSuperview()
            make.width.equalTo(view)
        }

        bannerImageView.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(16)
            make.leading.trailing.equalToSuperview().inset(16)
            make.height.equalTo(200)
        }

        nameLabel.snp.makeConstraints { make in
            make.top.equalTo(bannerImageView.snp.bottom).offset(16)
            make.leading.trailing.equalToSuperview().inset(16)
        }

        providerLabel.snp.makeConstraints { make in
            make.top.equalTo(nameLabel.snp.bottom).offset(4)
            make.leading.trailing.equalToSuperview().inset(16)
        }

        infoStack.snp.makeConstraints { make in
            make.top.equalTo(providerLabel.snp.bottom).offset(16)
            make.leading.trailing.equalToSuperview().inset(16)
            make.height.equalTo(60)
        }

        winsTitleLabel.snp.makeConstraints { make in
            make.top.equalTo(infoStack.snp.bottom).offset(24)
            make.leading.equalToSuperview().inset(16)
        }

        winsCollectionView.snp.makeConstraints { make in
            make.top.equalTo(winsTitleLabel.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview()
            make.height.equalTo(80)
        }

        relatedTitleLabel.snp.makeConstraints { make in
            make.top.equalTo(winsCollectionView.snp.bottom).offset(24)
            make.leading.equalToSuperview().inset(16)
        }

        relatedCollectionView.snp.makeConstraints { make in
            make.top.equalTo(relatedTitleLabel.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview()
            make.height.equalTo(120)
        }

        playButton.snp.makeConstraints { make in
            make.top.equalTo(relatedCollectionView.snp.bottom).offset(24)
            make.leading.trailing.equalToSuperview().inset(16)
            make.height.equalTo(56)
            make.bottom.equalToSuperview().offset(-20)
        }

        winsCollectionView.delegate = self
        winsCollectionView.dataSource = self
        relatedCollectionView.delegate = self
        relatedCollectionView.dataSource = self

        playButton.addTarget(self, action: #selector(playTapped), for: .touchUpInside)
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.updateUI(with: state)
            }
        }
    }

    private func updateUI(with state: GameDetailState) {
        guard let game = state.game else { return }
        nameLabel.text = game.name
        providerLabel.text = game.provider?.name

        configureInfoCards(game)
        winsCollectionView.reloadData()
        relatedCollectionView.reloadData()
    }

    private func configureInfoCards(_ game: Game) {
        infoStack.arrangedSubviews.forEach { $0.removeFromSuperview() }
        let items: [(String, String)] = [
            ("RTP", game.rtp.map { "\($0)%" } ?? "N/A"),
            ("Volatility", game.volatility?.capitalized ?? "N/A"),
            ("Min Bet", game.minBet.map { "$\(String(format: "%.2f", $0))" } ?? "N/A"),
            ("Max Bet", game.maxBet.map { "$\(String(format: "%.2f", $0))" } ?? "N/A")
        ]
        for (title, value) in items { infoStack.addArrangedSubview(createInfoCard(title: title, value: value)) }
    }

    private func createInfoCard(title: String, value: String) -> UIView {
        let view = UIView()
        view.backgroundColor = UIColor(hex: "#1E1E3F")
        view.layer.cornerRadius = 10

        let titleLabel = UILabel()
        titleLabel.text = title
        titleLabel.textColor = .white.withAlphaComponent(0.6)
        titleLabel.font = .systemFont(ofSize: 11)
        titleLabel.textAlignment = .center

        let valueLabel = UILabel()
        valueLabel.text = value
        valueLabel.textColor = .white
        valueLabel.font = .systemFont(ofSize: 14, weight: .bold)
        valueLabel.textAlignment = .center

        view.addSubview(titleLabel)
        view.addSubview(valueLabel)

        titleLabel.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(10)
            make.centerX.equalToSuperview()
        }

        valueLabel.snp.makeConstraints { make in
            make.top.equalTo(titleLabel.snp.bottom).offset(4)
            make.centerX.equalToSuperview()
        }

        return view
    }

    @objc private func playTapped() {
        Task {
            do {
                let session = try await viewModel.startGame()
                let gamePlayVC = GamePlayViewController(sessionId: session.sessionId)
                navigationController?.pushViewController(gamePlayVC, animated: true)
            } catch {
                let alert = UIAlertController(title: "Error", message: "Unable to start game", preferredStyle: .alert)
                alert.addAction(UIAlertAction(title: "OK", style: .default))
                present(alert, animated: true)
            }
        }
    }
}

// MARK: - UICollectionView DataSource & Delegate

extension GameDetailViewController: UICollectionViewDataSource, UICollectionViewDelegateFlowLayout {
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        if collectionView == winsCollectionView {
            return viewModel.state.recentWins.count
        }
        return viewModel.state.relatedGames.count
    }

    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        if collectionView == winsCollectionView {
            let cell = collectionView.dequeueReusableCell(withReuseIdentifier: "BigWinCell", for: indexPath) as! BigWinCell
            cell.configure(with: viewModel.state.recentWins[indexPath.item])
            return cell
        }
        let cell = collectionView.dequeueReusableCell(withReuseIdentifier: "RelatedGameCell", for: indexPath) as! RelatedGameCell
        cell.configure(with: viewModel.state.relatedGames[indexPath.item])
        return cell
    }

    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, sizeForItemAt indexPath: IndexPath) -> CGSize {
        if collectionView == winsCollectionView {
            return CGSize(width: 160, height: 70)
        }
        return CGSize(width: 100, height: 110)
    }
}
