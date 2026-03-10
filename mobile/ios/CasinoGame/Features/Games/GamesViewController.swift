import UIKit
import SnapKit

class GamesViewController: UIViewController {
    
    private var viewModel = GamesViewModel()
    
    private let searchBar: UISearchBar = {
        let searchBar = UISearchBar()
        searchBar.placeholder = "Search games..."
        searchBar.searchBarStyle = .minimal
        searchBar.barTintColor = UIColor(hex: "#0F0F23")
        searchBar.searchTextField.textColor = .white
        searchBar.searchTextField.backgroundColor = UIColor(hex: "#1E1E3F")
        return searchBar
    }()
    
    private lazy var collectionView: UICollectionView = {
        let layout = UICollectionViewFlowLayout()
        layout.scrollDirection = .vertical
        layout.minimumInteritemSpacing = 12
        layout.minimumLineSpacing = 12
        layout.sectionInset = UIEdgeInsets(top: 12, left: 16, bottom: 12, right: 16)
        
        let collectionView = UICollectionView(frame: .zero, collectionViewLayout: layout)
        collectionView.backgroundColor = .clear
        collectionView.register(GameGridCell.self, forCellWithReuseIdentifier: "GameGridCell")
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
        viewModel.loadGames()
    }
    
    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Games"
        
        view.addSubview(searchBar)
        view.addSubview(collectionView)
        view.addSubview(activityIndicator)
        
        searchBar.snp.makeConstraints { make in
            make.top.equalTo(view.safeAreaLayoutGuide)
            make.leading.trailing.equalToSuperview().inset(8)
        }
        
        collectionView.snp.makeConstraints { make in
            make.top.equalTo(searchBar.snp.bottom)
            make.leading.trailing.bottom.equalToSuperview()
        }
        
        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }
        
        collectionView.delegate = self
        collectionView.dataSource = self
        searchBar.delegate = self
    }
    
    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.updateUI(with: state)
            }
        }
    }
    
    private func updateUI(with state: GamesState) {
        if state.isLoading {
            activityIndicator.startAnimating()
        } else {
            activityIndicator.stopAnimating()
        }
        
        collectionView.reloadData()
    }
}

extension GamesViewController: UICollectionViewDataSource, UICollectionViewDelegateFlowLayout {
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return viewModel.state.games.count
    }
    
    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let cell = collectionView.dequeueReusableCell(withReuseIdentifier: "GameGridCell", for: indexPath) as! GameGridCell
        if indexPath.item < viewModel.state.games.count {
            cell.configure(with: viewModel.state.games[indexPath.item])
        }
        return cell
    }
    
    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, sizeForItemAt indexPath: IndexPath) -> CGSize {
        let width = (collectionView.bounds.width - 44) / 2
        return CGSize(width: width, height: width * 1.3)
    }
    
    func collectionView(_ collectionView: UICollectionView, didSelectItemAt indexPath: IndexPath) {
        if indexPath.item < viewModel.state.games.count {
            let game = viewModel.state.games[indexPath.item]
            let gameDetailVC = GameDetailViewController(gameId: game.id)
            navigationController?.pushViewController(gameDetailVC, animated: true)
        }
    }
}

extension GamesViewController: UISearchBarDelegate {
    func searchBar(_ searchBar: UISearchBar, textDidChange searchText: String) {
        viewModel.search(query: searchText)
    }
    
    func searchBarSearchButtonClicked(_ searchBar: UISearchBar) {
        searchBar.resignFirstResponder()
    }
}

// MARK: - GameGridCell

class GameGridCell: UICollectionViewCell {
    
    private let thumbnailImageView: UIImageView = {
        let imageView = UIImageView()
        imageView.contentMode = .scaleAspectFill
        imageView.clipsToBounds = true
        imageView.backgroundColor = UIColor(hex: "#1E1E3F")
        imageView.layer.cornerRadius = 12
        return imageView
    }()
    
    private let nameabel = {
        let label = UILabel()
        label.textColor =Label: UIL .white
        label.font = .systemFont(ofSize: 14, weight: .medium)
        label.textAlignment = .center
        return label
    }()
    
    private let providerLabel: UILabel = {
        let label = UILabel()
        label.textColor = .white.withAlphaComponent(0.6)
        label.font = .systemFont(ofSize: 12)
        label.textAlignment = .center
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
        contentView.addSubview(providerLabel)
        
        thumbnailImageView.snp.makeConstraints { make in
            make.top.leading.trailing.equalToSuperview()
            make.height.equalTo(contentView.snp.height).multipliedBy(0.75)
        }
        
        nameLabel.snp.makeConstraints { make in
            make.top.equalTo(thumbnailImageView.snp.bottom).offset(8)
            make.leading.trailing.equalToSuperview()
        }
        
        providerLabel.snp.makeConstraints { make in
            make.top.equalTo(nameLabel.snp.bottom).offset(2)
            make.leading.trailing.equalToSuperview()
        }
    }
    
    func configure(with game: Game) {
        nameLabel.text = game.name
        providerLabel.text = game.provider?.name
    }
}
