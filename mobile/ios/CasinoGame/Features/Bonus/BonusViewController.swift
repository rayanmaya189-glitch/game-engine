import UIKit
import SnapKit

class BonusViewController: UIViewController {

    private var viewModel = BonusViewModel()

    private let segmentedControl: UISegmentedControl = {
        let control = UISegmentedControl(items: ["Available", "Active"])
        control.selectedSegmentIndex = 0
        control.backgroundColor = UIColor(hex: "#1E1E3F")
        control.selectedSegmentTintColor = UIColor(hex: "#FF6B35")
        control.setTitleTextAttributes([.foregroundColor: UIColor.white], for: .selected)
        control.setTitleTextAttributes([.foregroundColor: UIColor.white.withAlphaComponent(0.7)], for: .normal)
        return control
    }()

    private lazy var collectionView: UICollectionView = {
        let layout = UICollectionViewFlowLayout()
        layout.scrollDirection = .vertical
        layout.minimumInteritemSpacing = 12
        layout.minimumLineSpacing = 12
        layout.sectionInset = UIEdgeInsets(top: 12, left: 16, bottom: 12, right: 16)
        let cv = UICollectionView(frame: .zero, collectionViewLayout: layout)
        cv.backgroundColor = .clear
        cv.register(BonusCardCell.self, forCellWithReuseIdentifier: "BonusCardCell")
        return cv
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
        viewModel.loadBonuses()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Bonuses"

        navigationItem.titleView = segmentedControl

        view.addSubview(collectionView)
        view.addSubview(activityIndicator)

        collectionView.snp.makeConstraints { make in
            make.edges.equalTo(view.safeAreaLayoutGuide)
        }

        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        collectionView.delegate = self
        collectionView.dataSource = self
        segmentedControl.addTarget(self, action: #selector(segmentChanged), for: .valueChanged)
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.collectionView.reloadData()
                if state.isLoading {
                    self?.activityIndicator.startAnimating()
                } else {
                    self?.activityIndicator.stopAnimating()
                }
                if let error = state.error {
                    let alert = UIAlertController(title: "Error", message: error, preferredStyle: .alert)
                    alert.addAction(UIAlertAction(title: "OK", style: .cancel))
                    self?.present(alert, animated: true)
                }
            }
        }
    }

    private var currentBonuses: [Bonus] {
        segmentedControl.selectedSegmentIndex == 0
            ? viewModel.state.availableBonuses
            : viewModel.state.activeBonuses
    }

    @objc private func segmentChanged() {
        collectionView.reloadData()
    }
}

// MARK: - UICollectionView

extension BonusViewController: UICollectionViewDataSource, UICollectionViewDelegateFlowLayout {
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return currentBonuses.count
    }

    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let cell = collectionView.dequeueReusableCell(withReuseIdentifier: "BonusCardCell", for: indexPath) as! BonusCardCell
        let bonus = currentBonuses[indexPath.item]
        cell.configure(with: bonus, isClaimable: segmentedControl.selectedSegmentIndex == 0)
        cell.onClaimTapped = { [weak self] in
            self?.viewModel.claimBonus(bonus.id)
        }
        return cell
    }

    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, sizeForItemAt indexPath: IndexPath) -> CGSize {
        let width = collectionView.bounds.width - 32
        return CGSize(width: width, height: 140)
    }
}

// MARK: - BonusCardCell

class BonusCardCell: UICollectionViewCell {
    var onClaimTapped: (() -> Void)?

    private let nameLabel = UILabel()
    private let descLabel = UILabel()
    private let amountLabel = UILabel()
    private let progressBar = UIProgressView(progressViewStyle: .default)
    private let progressLabel = UILabel()
    private let claimButton: UIButton = {
        let btn = UIButton(type: .system)
        btn.setTitle("CLAIM", for: .normal)
        btn.setTitleColor(.white, for: .normal)
        btn.backgroundColor = UIColor(hex: "#4CAF50")
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

        nameLabel.textColor = .white
        nameLabel.font = .systemFont(ofSize: 16, weight: .bold)
        descLabel.textColor = .white.withAlphaComponent(0.7)
        descLabel.font = .systemFont(ofSize: 13)
        descLabel.numberOfLines = 2
        amountLabel.textColor = UIColor(hex: "#FFD700")
        amountLabel.font = .systemFont(ofSize: 20, weight: .bold)
        progressBar.tintColor = UIColor(hex: "#FF6B35")
        progressLabel.textColor = .white.withAlphaComponent(0.7)
        progressLabel.font = .systemFont(ofSize: 12)

        [nameLabel, descLabel, amountLabel, progressBar, progressLabel, claimButton].forEach {
            contentView.addSubview($0)
        }

        nameLabel.snp.makeConstraints { make in
            make.top.leading.equalToSuperview().inset(16)
        }

        amountLabel.snp.makeConstraints { make in
            make.top.equalToSuperview().inset(16)
            make.trailing.equalToSuperview().inset(16)
        }

        descLabel.snp.makeConstraints { make in
            make.top.equalTo(nameLabel.snp.bottom).offset(6)
            make.leading.equalToSuperview().inset(16)
            make.trailing.equalTo(claimButton.snp.leading).offset(-8)
        }

        progressBar.snp.makeConstraints { make in
            make.bottom.equalToSuperview().inset(44)
            make.leading.equalToSuperview().inset(16)
            make.trailing.equalToSuperview().inset(16)
            make.height.equalTo(6)
        }

        progressLabel.snp.makeConstraints { make in
            make.bottom.equalTo(progressBar.snp.top).offset(-4)
            make.leading.equalToSuperview().inset(16)
        }

        claimButton.snp.makeConstraints { make in
            make.bottom.equalToSuperview().inset(16)
            make.trailing.equalToSuperview().inset(16)
            make.width.equalTo(80)
            make.height.equalTo(36)
        }

        claimButton.addTarget(self, action: #selector(claimTapped), for: .touchUpInside)
    }

    func configure(with bonus: Bonus, isClaimable: Bool) {
        nameLabel.text = bonus.name
        descLabel.text = bonus.description
        amountLabel.text = "$\(String(format: "%.0f", bonus.amount))"
        claimButton.isHidden = !isClaimable

        if bonus.wageringRequirement > 0 {
            let progress = Float(bonus.wageringProgress / bonus.wageringRequirement)
            progressBar.progress = progress
            progressLabel.text = "Wagering: \(Int(progress * 100))%"
            progressBar.isHidden = false
            progressLabel.isHidden = false
        } else {
            progressBar.isHidden = true
            progressLabel.isHidden = true
        }
    }

    @objc private func claimTapped() {
        onClaimTapped?()
    }
}
