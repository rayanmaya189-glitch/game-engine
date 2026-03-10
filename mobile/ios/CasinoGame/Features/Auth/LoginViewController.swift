import UIKit
import SnapKit

class LoginViewController: UIViewController {
    
    // MARK: - Properties
    
    var onLoginSuccess: (() -> Void)?
    var onRegisterTapped: (() -> Void)?
    
    private var viewModel = AuthViewModel()
    
    // MARK: - UI Components
    
    private let scrollView: UIScrollView = {
        let scrollView = UIScrollView()
        scrollView.showsVerticalScrollIndicator = false
        return scrollView
    }()
    
    private let contentView = UIView()
    
    private let titleLabel: UILabel = {
        let label = UILabel()
        label.text = "CASINO"
        label.font = .systemFont(ofSize: 42, weight: .bold)
        label.textColor = UIColor(hex: "#FFD700")
        label.textAlignment = .center
        return label
    }()
    
    private let subtitleLabel: UILabel = {
        let label = UILabel()
        label.text = "GAME"
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
    
    private let forgotPasswordButton: UIButton = {
        let button = UIButton(type: .system)
        button.setTitle("Forgot Password?", for: .normal)
        button.setTitleColor(UIColor(hex: "#FF6B35"), for: .normal)
        button.titleLabel?.font = .systemFont(ofSize: 14)
        return button
    }()
    
    private let loginButton: UIButton = {
        let button = UIButton(type: .system)
        button.setTitle("LOGIN", for: .normal)
        button.setTitleColor(.white, for: .normal)
        button.backgroundColor = UIColor(hex: "#FF6B35")
        button.layer.cornerRadius = 12
        button.titleLabel?.font = .systemFont(ofSize: 16, weight: .semibold)
        return button
    }()
    
    private let biometricButton: UIButton = {
        let button = UIButton(type: .system)
        button.setTitle("  Login with Biometric", for: .normal)
        button.setTitleColor(.white, for: .normal)
        button.backgroundColor = UIColor(hex: "#1E1E3F")
        button.layer.cornerRadius = 12
        button.titleLabel?.font = .systemFont(ofSize: 16, weight: .medium)
        return button
    }()
    
    private let registerButton: UIButton = {
        let button = UIButton(type: .system)
        let text = "Don't have an account? Sign Up"
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
    
    private let errorLabel: UILabel = {
        let label = UILabel()
        label.textColor = .systemRed
        label.font = .systemFont(ofSize: 14)
        label.textAlignment = .center
        label.numberOfLines = 0
        label.isHidden = true
        return label
    }()
    
    // MARK: - Lifecycle
    
    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        setupBindings()
    }
    
    // MARK: - Setup
    
    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        
        view.addSubview(scrollView)
        scrollView.addSubview(contentView)
        
        contentView.addSubview(titleLabel)
        contentView.addSubview(subtitleLabel)
        contentView.addSubview(emailTextField)
        contentView.addSubview(passwordTextField)
        contentView.addSubview(forgotPasswordButton)
        contentView.addSubview(loginButton)
        contentView.addSubview(biometricButton)
        contentView.addSubview(registerButton)
        contentView.addSubview(errorLabel)
        loginButton.addSubview(activityIndicator)
        
        scrollView.snp.makeConstraints { make in
            make.edges.equalToSuperview()
        }
        
        contentView.snp.makeConstraints { make in
            make.edges.equalToSuperview()
            make.width.equalTo(view)
        }
        
        titleLabel.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(80)
            make.centerX.equalToSuperview()
        }
        
        subtitleLabel.snp.makeConstraints { make in
            make.top.equalTo(titleLabel.snp.bottom).offset(4)
            make.centerX.equalToSuperview()
        }
        
        emailTextField.snp.makeConstraints { make in
            make.top.equalTo(subtitleLabel.snp.bottom).offset(48)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }
        
        passwordTextField.snp.makeConstraints { make in
            make.top.equalTo(emailTextField.snp.bottom).offset(16)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }
        
        forgotPasswordButton.snp.makeConstraints { make in
            make.top.equalTo(passwordTextField.snp.bottom).offset(8)
            make.trailing.equalToSuperview().inset(24)
        }
        
        loginButton.snp.makeConstraints { make in
            make.top.equalTo(forgotPasswordButton.snp.bottom).offset(24)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }
        
        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }
        
        errorLabel.snp.makeConstraints { make in
            make.top.equalTo(loginButton.snp.bottom).offset(16)
            make.leading.trailing.equalToSuperview().inset(24)
        }
        
        biometricButton.snp.makeConstraints { make in
            make.top.equalTo(errorLabel.snp.bottom).offset(16)
            make.leading.trailing.equalToSuperview().inset(24)
            make.height.equalTo(56)
        }
        
        registerButton.snp.makeConstraints { make in
            make.top.equalTo(biometricButton.snp.bottom).offset(32)
            make.centerX.equalToSuperview()
            make.bottom.equalToSuperview().offset(-32)
        }
        
        // Actions
        loginButton.addTarget(self, action: #selector(loginTapped), for: .touchUpInside)
        registerButton.addTarget(self, action: #selector(registerTapped), for: .touchUpInside)
        biometricButton.addTarget(self, action: #selector(biometricTapped), for: .touchUpInside)
    }
    
    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.updateUI(with: state)
            }
        }
    }
    
    private func updateUI(with state: AuthState) {
        loginButton.isEnabled = !state.isLoading
        loginButton.setTitle(state.isLoading ? "" : "LOGIN", for: .normal)
        
        if state.isLoading {
            activityIndicator.startAnimating()
        } else {
            activityIndicator.stopAnimating()
        }
        
        if let error = state.error {
            errorLabel.text = error
            errorLabel.isHidden = false
        } else {
            errorLabel.isHidden = true
        }
        
        if state.isLoggedIn {
            onLoginSuccess?()
        }
    }
    
    // MARK: - Actions
    
    @objc private func loginTapped() {
        guard let email = emailTextField.text, !email.isEmpty,
              let password = passwordTextField.text, !password.isEmpty else {
            errorLabel.text = "Please enter email and password"
            errorLabel.isHidden = false
            return
        }
        
        viewModel.login(email: email, password: password)
    }
    
    @objc private func registerTapped() {
        onRegisterTapped?()
    }
    
    @objc private func biometricTapped() {
        // Implement biometric login
    }
}
