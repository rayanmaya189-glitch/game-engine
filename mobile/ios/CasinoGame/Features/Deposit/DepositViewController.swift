import UIKit
import SnapKit

class DepositViewController: UIViewController {

    private var viewModel = DepositViewModel()
    var onDepositComplete: (() -> Void)?

    private let balanceTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "Current Balance"
        label.textColor = .white.withAlphaComponent(0.7)
        label.font = .systemFont(ofSize: 14)
        return label
    }()

    private let balanceLabel: UILabel = {
        let label = UILabel()
        label.text = "$0.00"
        label.textColor = UIColor(hex: "#FFD700")
        label.font = .systemFont(ofSize: 28, weight: .bold)
        return label
    }()

    private let amountTextField: UITextField = {
        let textField = UITextField()
        textField.placeholder = "Enter amount"
        textField.borderStyle = .none
        textField.backgroundColor = UIColor(hex: "#1E1E3F")
        textField.textColor = .white
        textField.keyboardType = .decimalPad
        textField.layer.cornerRadius = 12
        textField.leftView = UIView(frame: CGRect(x: 0, y: 0, width: 16, height: 0))
        textField.leftViewMode = .always
        return textField
    }()

    private let methodSegment: UISegmentedControl = {
        let control = UISegmentedControl(items: ["Credit Card", "Crypto", "Bank"])
        control.selectedSegmentIndex = 0
        control.backgroundColor = UIColor(hex: "#1E1E3F")
        control.selectedSegmentTintColor = UIColor(hex: "#FF6B35")
        control.setTitleTextAttributes([.foregroundColor: UIColor.white], for: .selected)
        control.setTitleTextAttributes([.foregroundColor: UIColor.white.withAlphaComponent(0.7)], for: .normal)
        return control
    }()

    private let depositButton: UIButton = {
        let button = UIButton(type: .system)
        button.setTitle("DEPOSIT", for: .normal)
        button.setTitleColor(.white, for: .normal)
        button.backgroundColor = UIColor(hex: "#4CAF50")
        button.layer.cornerRadius = 12
        button.titleLabel?.font = .systemFont(ofSize: 16, weight: .semibold)
        return button
    }()

    private let errorLabel: UILabel = {
        let label = UILabel()
        label.textColor = .systemRed
        label.font = .systemFont(ofSize: 14)
        label.textAlignment = .center
        label.numberOfLines = 0
        label.isHidden = true
        return label
    }()

    private let activityIndicator: UIActivityIndicatorView = {
        let indicator = UIActivityIndicatorView(style: .medium)
        indicator.color = .white
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
        title = "Deposit"

        [balanceTitleLabel, balanceLabel, amountTextField, methodSegment,
         depositButton, errorLabel].forEach { view.addSubview($0) }
        depositButton.addSubview(activityIndicator)

        balanceTitleLabel.snp.makeConstraints { make in
            make.top.equalTo(view.safeAreaLayoutGuide).offset(24)
            make.centerX.equalToSuperview()
        }

        balanceLabel.snp.makeConstraints { make in
            make.top.equalTo(balanceTitleLabel.snp.bottom).offset(4)
            make.centerX.equalToSuperview()
        }

        amountTextField.snp.makeConstraints { make in
            make.top.equalTo(balanceLabel.snp.bottom).offset(32)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }

        methodSegment.snp.makeConstraints { make in
            make.top.equalTo(amountTextField.snp.bottom).offset(20)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(44)
        }

        depositButton.snp.makeConstraints { make in
            make.top.equalTo(methodSegment.snp.bottom).offset(24)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }

        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        errorLabel.snp.makeConstraints { make in
            make.top.equalTo(depositButton.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview().inset(24)
        }

        depositButton.addTarget(self, action: #selector(depositTapped), for: .touchUpInside)
        methodSegment.addTarget(self, action: #selector(methodChanged), for: .valueChanged)
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.updateUI(with: state)
            }
        }
    }

    private func updateUI(with state: DepositState) {
        balanceLabel.text = "$\(String(format: "%.2f", state.balance))"

        if state.isLoading {
            activityIndicator.startAnimating()
            depositButton.setTitle("", for: .normal)
        } else {
            activityIndicator.stopAnimating()
            depositButton.setTitle("DEPOSIT", for: .normal)
        }

        if state.isSuccess {
            showSuccess(amount: state.amount)
        }

        if let error = state.error {
            errorLabel.text = error
            errorLabel.isHidden = false
        } else {
            errorLabel.isHidden = true
        }
    }

    private func showSuccess(amount: Double) {
        let alert = UIAlertController(
            title: "Deposit Initiated",
            message: "Your deposit of $\(String(format: "%.2f", amount)) has been initiated.",
            preferredStyle: .alert
        )
        alert.addAction(UIAlertAction(title: "OK", style: .default) { [weak self] _ in
            self?.onDepositComplete?()
            self?.navigationController?.popViewController(animated: true)
        })
        present(alert, animated: true)
    }

    @objc private func depositTapped() {
        guard let text = amountTextField.text else { return }
        let validation = viewModel.validateAmount(text)
        if !validation.valid {
            errorLabel.text = validation.message
            errorLabel.isHidden = false
            return
        }
        errorLabel.isHidden = true
        viewModel.deposit(amount: Double(text) ?? 0)
    }

    @objc private func methodChanged() {
        viewModel.selectMethod(at: methodSegment.selectedSegmentIndex)
    }
}
