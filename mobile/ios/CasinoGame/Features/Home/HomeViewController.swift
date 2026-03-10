import UIKit
import SnapKit

class HomeViewController: UIViewController {
    
    private var viewModel = HomeViewModel()
    
    private let scrollView: UIScrollView = {
        let scrollView = UIScrollView()
        scrollView.showsVerticalScrollIndicator = false
        return scrollView
    }()
    
    private let contentView = UIView()
    
    private let balanceButton: UIButton = {
        let button = UIButton(type: .system)
        button.backgroundColor = UIColor(hex: "#1E1E3F")
        button.layer.cornerRadius = 20
        button.setTitleColor(.white, for: .normal)
        button.titleLabel?.font = .systemFont(ofSize: 16, weight: .semibold)
        
        let image = UIImage(systemName: "wallet.pass.fill")
        button.setImage(image, for: .normal)
        button.tintColor = UIColor(hex: "#FFD700")
        button.imageEdgeInsets = UIEdgeInsets(top: 0, left: 0, bottom: 0, right: 8)
        
        return button
    }()
    
    private let jackpotCard: UIView = {
        let view = UIView()
        view.backgroundColor = UIColor(hex: "#FF6B35")
        view.layer.cornerRadius = 16
        return view
    }()
    
    private let jackpotTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "JACKPOT"
        label.textColor = .white.withAlphaComponent(0.9)
        label.font = .systemFont(ofSize: 14, weight: .medium)
        return label
    }()
    
    private let jackpotAmountLabel: UILabel = {
        let label = UILabel()
        label.text = "$0"
        label.textColor = .white
        label.font = .systemFont(ofSize: 32, weight: .bold)
        return label
    }()
    
    private let featuredTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "Featured Games"
        label.textColor = .white
        label.font = .systemFont(ofSize: 18, weight: .bold)
        return label
    }()
    
    private let featuredCollectionView: UICollectionView = {
        let layout = UICollectionViewFlowLayout()
        layout.scrollDirection = .horizontal
        layout.itemSize = CGSize(width: 140, height: 180)
        layout.minimumInteritemSpacing = 12
        layout.sectionInset = UIEdgeInsets(top: 0, left: 16, bottom: 0, right: 16)
        
        let collectionView = UICollectionView(frame: .zero, collectionViewLayout: layout)
        collectionView.backgroundColor = .clear
        collectionView.showsHorizontalScrollIndicator = false
        collectionView.register(GameCollectionCell.self, forCellWithReuseIdentifier: "GameCell")
        return collectionView
    }()
    
    private let popularTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "Popular Games"
        label.textColor = .white
        label.font = .systemFont(ofSize: 18, weight: .bold)
        return label
    }()
    
    private let popularCollectionView: UICollectionView = {
        let layout = UICollectionViewFlowLayout()
        layout.scrollDirection = .horizontal
        layout.itemSize = CGSize(width: 140, height: 180)
        layout.minimumInteritemSpacing = 12
        layout.sectionInset = UIEdgeInsets(top: 0, left: 16, bottom: 0, right: 16)
        
        let collectionView = UICollectionView(frame: .zero, collectionViewLayout: layout)
        collectionView.backgroundColor = .clear
        collectionView.showsHorizontalScrollIndicator = false
        collectionView.register(GameCollectionCell.self, forCellWithReuseIdentifier: "GameCell")
        return collectionView
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
        
        navigationItem.titleView = {
            let titleLabel = UILabel()
            titleLabel.text = "CASINO"
            titleLabel.font = .systemFont(ofSize: 20, weight: .bold)
            titleLabel.textColor = UIColor(hex: "#FFD700")
            
            let gameLabel = UILabel()
            gameLabel.text = "GAME"
            gameLabel.font = .systemFont(ofSize: 20, weight: .bold)
            gameLabel.textColor = .white
            
            let stack = UIStackView(arrangedSubviews: [titleLabel, gameLabel])
            stack.axis = .horizontal
            stack.spacing = 4
            return stack
        }()
        
        navigationItem.rightBarButtonItem = UIBarButtonItem(customView: balanceButton)
        
        view.addSubview(scrollView)
        scrollView.addSubview(contentView)
        
        contentView.addSubview(jackpotCard)
        jackpotCard.addSubview(jackpotTitleLabel)
        jackpotCard.addSubview(jackpotAmountLabel)
        contentView.addSubview(featuredTitleLabel)
        contentView.addSubview(featuredCollectionView)
        contentView.addSubview(popularTitleLabel)
        contentView.addSubview(popularCollectionView)
        view.addSubview(activityIndicator)
        
        scrollView.snp.makeConstraints { make in
            make.edges.equalTo(view.safeAreaLayoutGuide)
        }
        
        contentView.snp.makeConstraints { make in
            make.edges.equalToSuperview()
            make.width.equalTo(view)
        }
        
        jackpotCard.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(16)
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
        
        featuredTitleLabel.snp.makeConstraints { make in
            make.top.equalTo(jackpotCard.snp.bottom).offset(24)
            make.leading.equalToSuperview().inset(16)
        }
        
        featuredCollectionView.snp.makeConstraints { make in
            make.top.equalTo(featuredTitleLabel.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview()
            make.height.equalTo(180)
        }
        
        popularTitleLabel.snp.makeConstraints { make in
            make.top.equalTo(featuredCollectionView.snp.bottom).offset(24)
            make.leading.equalToSuperview().inset(16)
        }
        
        popularCollectionView.snp.makeConstraints { make in
            make.top.equalTo(popularTitleLabel.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview()
            make.height.equalTo(180)
            make.bottom.equalToSuperview().offset(-16)
        }
        
        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }
        
        featuredCollectionView.delegate = self
        featuredCollectionView.dataSource = self
        popularCollectionView.delegate = self
        popularCollectionView.dataSource = self
    }
    
    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.updateUI(with: state)
            }
        }
    }
    
    private func updateUI(with state: HomeState) {
        if state.isLoading {
            activityIndicator.startAnimating()
        } else {
            activityIndicator.stopAnimating()
        }
        
        if let balance = state.balance {
            balanceButton.setTitle(" $\(String(format: "%.2f", balance.balance))", for: .normal)
        }
        
        if let jackpotAmount = state.jackpotGames.first?.jackpotAmount {
            jackpotAmountLabel.text = "$\(String(format: "%,.0f", jackpotAmount))"
        }
        
        featuredCollectionView.reloadData()
        popularCollectionView.reloadData()
    }
}

// MARK: - UICollectionViewDataSource & Delegate

extension HomeViewController: UICollectionViewDataSource, UICollectionViewDelegate {
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        if collectionView == featuredCollectionView {
            return viewModel.state.featuredGames.count
        } else {
            return viewModel.state.popularGames.count
        }
    }
    
    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let cell = collectionView.dequeueReusableCell(withReuseIdentifier: "GameCell", for: indexPath) as! GameCollectionCell
        
        let games = collectionView == featuredCollectionView ? viewModel.state.featuredGames : viewModel.state.popularGames
        if indexPath.item < games.count {
            cell.configure(with: games[indexPath.item])
        }
        
        return cell
    }
    
    func collectionView(_ collectionView: UICollectionView, didSelectItemAt indexPath: IndexPath) {
        let games = collectionView == featuredCollectionView ? viewModel.state.featuredGames : viewModel.state.popularGames
        if indexPath.item < games.count {
            let game = games[indexPath.item]
            let gameDetailVC = GameDetailViewController(gameId: game.id)
            navigationController?.pushViewController(gameDetailVC, animated: true)
        }
    }
}

// MARK: - GameCollectionCell

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
            // Using Kingfisher or similar for image loading
            // thumbnailImageView.kf.setImage(with: url)
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
