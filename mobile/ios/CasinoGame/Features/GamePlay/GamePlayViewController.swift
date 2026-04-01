import UIKit
import SnapKit

class GamePlayViewController: UIViewController {

    private var viewModel = GamePlayViewModel()

    private let balanceLabel: UILabel = {
        let label = UILabel()
        label.text = "$0.00"
        label.textColor = UIColor(hex: "#FFD700")
        label.font = .systemFont(ofSize: 18, weight: .bold)
        return label
    }()

    private let winLabel: UILabel = {
        let label = UILabel()
        label.text = "WIN: $0.00"
        label.textColor = UIColor(hex: "#4CAF50")
        label.font = .systemFont(ofSize: 22, weight: .bold)
        label.textAlignment = .center
        label.isHidden = true
        return label
    }()

    private lazy var collectionView: UICollectionView = {
        let layout = UICollectionViewFlowLayout()
        layout.scrollDirection = .horizontal
        layout.minimumInteritemSpacing = 4
        layout.minimumLineSpacing = 4
        layout.sectionInset = UIEdgeInsets(top: 8, left: 8, bottom: 8, right: 8)
        let cv = UICollectionView(frame: .zero, collectionViewLayout: layout)
        cv.backgroundColor = UIColor(hex: "#1E1E3F")
        cv.layer.cornerRadius = 12
        cv.isScrollEnabled = false
        cv.register(ReelSymbolCell.self, forCellWithReuseIdentifier: "ReelSymbolCell")
        return cv
    }()

    private let betSlider: UISlider = {
        let slider = UISlider()
        slider.minimumValue = 0
        slider.maximumValue = 9
        slider.value = 3
        slider.minimumTrackTintColor = UIColor(hex: "#FF6B35")
        slider.thumbTintColor = UIColor(hex: "#FFD700")
        return slider
    }()

    private let betLabel: UILabel = {
        let label = UILabel()
        label.text = "Bet: $1.00"
        label.textColor = .white
        label.font = .systemFont(ofSize: 16, weight: .medium)
        return label
    }()

    private let spinButton: UIButton = {
        let button = UIButton(type: .system)
        button.setTitle("SPIN", for: .normal)
        button.setTitleColor(.white, for: .normal)
        button.backgroundColor = UIColor(hex: "#FF6B35")
        button.titleLabel?.font = .systemFont(ofSize: 22, weight: .bold)
        button.layer.cornerRadius = 30
        return button
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
        viewModel.loadBalance()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Lucky Slots"

        navigationItem.rightBarButtonItem = UIBarButtonItem(customView: balanceLabel)

        view.addSubview(collectionView)
        view.addSubview(winLabel)
        view.addSubview(betLabel)
        view.addSubview(betSlider)
        view.addSubview(spinButton)
        view.addSubview(activityIndicator)

        collectionView.delegate = self
        collectionView.dataSource = self

        collectionView.snp.makeConstraints { make in
            make.top.equalTo(view.safeAreaLayoutGuide).offset(20)
            make.leading.trailing.equalToSuperview().inset(16)
            make.height.equalTo(200)
        }

        winLabel.snp.makeConstraints { make in
            make.top.equalTo(collectionView.snp.bottom).offset(16)
            make.centerX.equalToSuperview()
        }

        betLabel.snp.makeConstraints { make in
            make.top.equalTo(winLabel.snp.bottom).offset(24)
            make.leading.equalToSuperview().inset(24)
        }

        betSlider.snp.makeConstraints { make in
            make.top.equalTo(betLabel.snp.bottom).offset(8)
            make.leading.trailing.equalToSuperview().inset(24)
        }

        spinButton.snp.makeConstraints { make in
            make.top.equalTo(betSlider.snp.bottom).offset(24)
            make.centerX.equalToSuperview()
            make.size.equalTo(60)
        }

        activityIndicator.snp.makeConstraints { make in
            make.center.equalTo(spinButton)
        }

        spinButton.addTarget(self, action: #selector(spinTapped), for: .touchUpInside)
        betSlider.addTarget(self, action: #selector(betChanged), for: .valueChanged)
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.updateUI(with: state)
            }
        }
    }

    private func updateUI(with state: GamePlayState) {
        balanceLabel.text = "$\(String(format: "%.2f", state.balance))"
        spinButton.isEnabled = !state.isSpinning
        spinButton.alpha = state.isSpinning ? 0.5 : 1.0

        if state.isSpinning {
            activityIndicator.startAnimating()
            startReelAnimation()
        } else {
            activityIndicator.stopAnimating()
            stopReelAnimation()
            collectionView.reloadData()
        }

        if state.winAmount > 0 {
            winLabel.text = "WIN: $\(String(format: "%.2f", state.winAmount))"
            winLabel.isHidden = false
            animateWinLabel()
        } else {
            winLabel.isHidden = true
        }

        if let error = state.error {
            showError(error)
        }
    }

    private func startReelAnimation() {
        spinButton.setTitle("", for: .normal)
        Timer.scheduledTimer(withTimeInterval: 0.1, repeats: true) { [weak self] timer in
            guard let self = self, self.viewModel.state.isSpinning else {
                timer.invalidate()
                return
            }
            var symbols = self.viewModel.state.reels
            for i in 0..<symbols.count {
                for j in 0..<symbols[i].count {
                    symbols[i][j] = Self.allSymbols.randomElement() ?? "🍒"
                }
            }
            self.collectionView.reloadData()
        }
    }

    private func stopReelAnimation() {
        spinButton.setTitle("SPIN", for: .normal)
    }

    private func animateWinLabel() {
        winLabel.transform = CGAffineTransform(scaleX: 0.5, y: 0.5)
        UIView.animate(withDuration: 0.4, delay: 0, usingSpringWithDamping: 0.5, initialSpringVelocity: 0.5) {
            self.winLabel.transform = .identity
        }
    }

    private func showError(_ message: String) {
        let alert = UIAlertController(title: "Error", message: message, preferredStyle: .alert)
        alert.addAction(UIAlertAction(title: "OK", style: .cancel))
        present(alert, animated: true)
    }

    @objc private func spinTapped() {
        viewModel.spin()
    }

    @objc private func betChanged() {
        let index = Int(betSlider.value)
        let steps = viewModel.betSteps
        let bet = steps[min(index, steps.count - 1)]
        betLabel.text = "Bet: $\(String(format: "%.2f", bet))"
        viewModel.updateBet(bet)
    }

    static let allSymbols = ["🍒", "🍋", "🍊", "🍇", "⭐", "💎", "7️⃣", "🔔"]
}

// MARK: - UICollectionView

extension GamePlayViewController: UICollectionViewDataSource, UICollectionViewDelegateFlowLayout {
    func numberOfSections(in collectionView: UICollectionView) -> Int { return 5 }

    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int { return 3 }

    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let cell = collectionView.dequeueReusableCell(withReuseIdentifier: "ReelSymbolCell", for: indexPath) as! ReelSymbolCell
        let reels = viewModel.state.reels
        if indexPath.section < reels.count, indexPath.item < reels[indexPath.section].count {
            cell.configure(symbol: reels[indexPath.section][indexPath.item])
        }
        return cell
    }

    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, sizeForItemAt indexPath: IndexPath) -> CGSize {
        let width = (collectionView.bounds.width - 32 - 16) / 5
        let height = (collectionView.bounds.height - 16 - 8) / 3
        return CGSize(width: width, height: height)
    }
}

// MARK: - ReelSymbolCell

class ReelSymbolCell: UICollectionViewCell {
    private let symbolLabel: UILabel = {
        let label = UILabel()
        label.font = .systemFont(ofSize: 32)
        label.textAlignment = .center
        return label
    }()

    override init(frame: CGRect) {
        super.init(frame: frame)
        contentView.backgroundColor = UIColor(hex: "#0F0F23")
        contentView.layer.cornerRadius = 8
        contentView.addSubview(symbolLabel)
        symbolLabel.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }
    }

    required init?(coder: NSCoder) { fatalError() }

    func configure(symbol: String) {
        symbolLabel.text = symbol
    }
}
