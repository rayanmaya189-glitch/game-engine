import UIKit
import SnapKit

class RegisterViewController: UIViewController {
    var onRegisterSuccess: (() -> Void)?
    var onLoginTapped: (() -> Void)?

    private var viewModel = AuthViewModel()

    private let scrollView: UIScrollView = {
        let scrollView = UIScrollView()
        scrollView.showsVerticalScrollIndicator = false
        return scrollView
    }()

    private let contentView = UIView()

    private let titleLabel: UILabel = {
        let label = UILabel()
        label.text = "CREATE ACCOUNT"
        label.font = .systemFont(ofSize: 28, weight: .bold)
        label.textColor = .white
        label.textAlignment = .center
        return label
    }()

    private let emailTextField: UITextField = {
        let textField = UITextField()
        textField.placeholder = "Email"
        textField.borderStyle = .none
        textField.backgroundColor = UIColor(hex: "#1E1E3F")
        textField.textColor = .white
        textField.keyboardType = .emailAddress
        textField.autocapitalizationType = .none
        textField.layer.cornerRadius = 12
        textField.leftView = UIView(frame: CGRect(x: 0, y: 0, width: 16, height: 0))
        textField.leftViewMode = .always
        return textField
    }()

    private let usernameTextField: UITextField = {
        let textField = UITextField()
        textField.placeholder = "Username"
        textField.borderStyle = .none
        textField.backgroundColor = UIColor(hex: "#1E1E3F")
        textField.textColor = .white
        textField.autocapitalizationType = .none
        textField.layer.cornerRadius = 12
        textField.leftView = UIView(frame: CGRect(x: 0, y: 0, width: 16, height: 0))
        textField.leftViewMode = .always
        return textField
    }()

    private let phoneTextField: UITextField = {
        let textField = UITextField()
        textField.placeholder = "Phone (Optional)"
        textField.borderStyle = .none
        textField.backgroundColor = UIColor(hex: "#1E1E3F")
        textField.textColor = .white
        textField.keyboardType = .phonePad
        textField.layer.cornerRadius = 12
        textField.leftView = UIView(frame: CGRect(x: 0, y: 0, width: 16, height: 0))
        textField.leftViewMode = .always
        return textField
    }()

    private let passwordTextField: UITextField = {
        let textField = UITextField()
        textField.placeholder = "Password"
        textField.borderStyle = .none
        textField.backgroundColor = UIColor(hex: "#1E1E3F")
        textField.textColor = .white
        textField.isSecureTextEntry = true
        textField.layer.cornerRadius = 12
        textField.leftView = UIView(frame: CGRect(x: 0, y: 0, width: 16, height: 0))
        textField.leftViewMode = .always
        return textField
    }()

    private let confirmPasswordTextField: UITextField = {
        let textField = UITextField()
        textField.placeholder = "Confirm Password"
        textField.borderStyle = .none
        textField.backgroundColor = UIColor(hex: "#1E1E3F")
        textField.textColor = .white
        textField.isSecureTextEntry = true
        textField.layer.cornerRadius = 12
        textField.leftView = UIView(frame: CGRect(x: 0, y: 0, width: 16, height: 0))
        textField.leftViewMode = .always
        return textField
    }()

    private let referralCodeTextField: UITextField = {
        let textField = UITextField()
        textField.placeholder = "Referral Code (Optional)"
        textField.borderStyle = .none
        textField.backgroundColor = UIColor(hex: "#1E1E3F")
        textField.textColor = .white
        textField.autocapitalizationType = .none
        textField.layer.cornerRadius = 12
        textField.leftView = UIView(frame: CGRect(x: 0, y: 0, width: 16, height: 0))
        textField.leftViewMode = .always
        return textField
    }()

    private let termsCheckbox: UIButton = {
        let button = UIButton(type: .system)
        button.setImage(UIImage(systemName: "square"), for: .normal)
        button.setImage(UIImage(systemName: "checkmark.square.fill"), for: .selected)
        button.tintColor = UIColor(hex: "#FF6B35")
        return button
    }()

    private let termsLabel: UILabel = {
        let label = UILabel()
        label.text = "I agree to Terms & Conditions"
        label.textColor = .white
        label.font = .systemFont(ofSize: 14)
        return label
    }()

    private let registerButton: UIButton = {
        let button = UIButton(type: .system)
        button.setTitle("REGISTER", for: .normal)
        button.setTitleColor(.white, for: .normal)
        button.backgroundColor = UIColor(hex: "#FF6B35")
        button.layer.cornerRadius = 12
        button.titleLabel?.font = .systemFont(ofSize: 16, weight: .semibold)
        return button
    }()

    private let loginButton: UIButton = {
        let button = UIButton(type: .system)
        let text = "Already have an account? Sign In"
        button.setTitle(text, for: .normal)
        button.setTitleColor(.white, for: .normal)
        button.titleLabel?.font = .systemFont(ofSize: 14)
        return button
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
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        navigationController?.setNavigationBarHidden(true, animated: false)

        view.addSubview(scrollView)
        scrollView.addSubview(contentView)

        [titleLabel, emailTextField, usernameTextField, phoneTextField,
         passwordTextField, confirmPasswordTextField, referralCodeTextField,
         termsCheckbox, termsLabel, registerButton, loginButton].forEach {
            contentView.addSubview($0)
        }

        registerButton.addSubview(activityIndicator)

        scrollView.snp.makeConstraints { make in
            make.edges.equalToSuperview()
        }

        contentView.snp.makeConstraints { make in
            make.edges.equalToSuperview()
            make.width.equalTo(view)
        }

        titleLabel.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(60)
            make.centerX.equalToSuperview()
        }

        emailTextField.snp.makeConstraints { make in
            make.top.equalTo(titleLabel.snp.bottom).offset(32)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }

        usernameTextField.snp.makeConstraints { make in
            make.top.equalTo(emailTextField.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }

        phoneTextField.snp.makeConstraints { make in
            make.top.equalTo(usernameTextField.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }

        passwordTextField.snp.makeConstraints { make in
            make.top.equalTo(phoneTextField.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }

        confirmPasswordTextField.snp.makeConstraints { make in
            make.top.equalTo(passwordTextField.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }

        referralCodeTextField.snp.makeConstraints { make in
            make.top.equalTo(confirmPasswordTextField.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }

        termsCheckbox.snp.makeConstraints { make in
            make.top.equalTo(referralCodeTextField.snp.bottom).offset(16)
            make.leading.equalToSuperview().inset(24)
            make.size.equalTo(24)
        }

        termsLabel.snp.makeConstraints { make in
            make.centerY.equalTo(termsCheckbox)
            make.leading.equalTo(termsCheckbox.snp.trailing).offset(8)
        }

        registerButton.snp.makeConstraints { make in
            make.top.equalTo(termsCheckbox.snp.bottom).offset(24)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }

        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        loginButton.snp.makeConstraints { make in
            make.top.equalTo(registerButton.snp.bottom).offset(24)
            make.centerX.equalToSuperview()
            make.bottom.equalToSuperview().offset(-32)
        }

        termsCheckbox.addTarget(self, action: #selector(termsTapped), for: .touchUpInside)
        registerButton.addTarget(self, action: #selector(registerTapped), for: .touchUpInside)
        loginButton.addTarget(self, action: #selector(loginTapped), for: .touchUpInside)
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.updateUI(with: state)
            }
        }
    }

    private func updateUI(with state: AuthState) {
        registerButton.isEnabled = !state.isLoading
        registerButton.setTitle(state.isLoading ? "" : "REGISTER", for: .normal)

        if state.isLoading {
            activityIndicator.startAnimating()
        } else {
            activityIndicator.stopAnimating()
        }

        if state.isLoggedIn {
            onRegisterSuccess?()
        }
    }

    @objc private func termsTapped() {
        termsCheckbox.isSelected.toggle()
    }

    @objc private func registerTapped() {
        guard let email = emailTextField.text, !email.isEmpty,
              let username = usernameTextField.text, !username.isEmpty,
              let password = passwordTextField.text, !password.isEmpty,
              let confirmPassword = confirmPasswordTextField.text, confirmPassword == password,
              termsCheckbox.isSelected else {
            return
        }

        viewModel.register(
            email: email,
            username: username,
            password: password,
            phone: phoneTextField.text?.isEmpty == false ? phoneTextField.text : nil,
            currency: "USD",
            referralCode: referralCodeTextField.text?.isEmpty == false ? referralCodeTextField.text : nil
        )
    }

    @objc private func loginTapped() {
        onLoginTapped?()
    }
}
