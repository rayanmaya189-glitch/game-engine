import UIKit

class SecurityBlockViewController: UIViewController {
    private let securityResult: FullSecurityResult

    private let scrollView = UIScrollView()
    private let contentView = UIView()

    private let iconImageView: UIImageView = {
        let imageView = UIImageView()
        imageView.image = UIImage(systemName: "exclamationmark.shield.fill")
        imageView.tintColor = .systemRed
        imageView.contentMode = .scaleAspectFit
        return imageView
    }()

    private let titleLabel: UILabel = {
        let label = UILabel()
        label.text = "Access Blocked"
        label.font = .systemFont(ofSize: 28, weight: .bold)
        label.textColor = .systemRed
        label.textAlignment = .center
        return label
    }()

    private let messageLabel: UILabel = {
        let label = UILabel()
        label.text = "This app cannot be opened due to security concerns."
        label.font = .systemFont(ofSize: 16)
        label.textColor = .secondaryLabel
        label.textAlignment = .center
        label.numberOfLines = 0
        return label
    }()

    private let issuesStackView: UIStackView = {
        let stackView = UIStackView()
        stackView.axis = .vertical
        stackView.spacing = 12
        stackView.alignment = .fill
        return stackView
    }()

    private let uninstallLabel: UILabel = {
        let label = UILabel()
        label.text = "Please uninstall the following apps to continue:"
        label.font = .systemFont(ofSize: 14, weight: .medium)
        label.textColor = .label
        label.textAlignment = .center
        return label
    }()

    private let appsStackView: UIStackView = {
        let stackView = UIStackView()
        stackView.axis = .vertical
        stackView.spacing = 8
        stackView.alignment = .fill
        return stackView
    }()

    private let instructionLabel: UILabel = {
        let label = UILabel()
        label.text = "1. Uninstall the apps listed above\n2. Restart your device\n3. Reopen this app"
        label.font = .systemFont(ofSize: 14)
        label.textColor = .secondaryLabel
        label.textAlignment = .left
        label.numberOfLines = 0
        return label
    }()

    private let contactLabel: UILabel = {
        let label = UILabel()
        label.text = "If you believe this is an error, please contact support."
        label.font = .systemFont(ofSize: 12)
        label.textColor = .tertiaryLabel
        label.textAlignment = .center
        label.numberOfLines = 0
        return label
    }()

    init(securityResult: FullSecurityResult) {
        self.securityResult = securityResult
        super.init(nibName: nil, bundle: nil)
    }

    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }

    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        displaySecurityIssues()
    }

    private func setupUI() {
        view.backgroundColor = .systemBackground

        view.addSubview(scrollView)
        scrollView.translatesAutoresizingMaskIntoConstraints = false
        scrollView.addSubview(contentView)
        contentView.translatesAutoresizingMaskIntoConstraints = false

        NSLayoutConstraint.activate([
            scrollView.topAnchor.constraint(equalTo: view.safeAreaLayoutGuide.topAnchor),
            scrollView.leadingAnchor.constraint(equalTo: view.leadingAnchor),
            scrollView.trailingAnchor.constraint(equalTo: view.trailingAnchor),
            scrollView.bottomAnchor.constraint(equalTo: view.safeAreaLayoutGuide.bottomAnchor),
            contentView.topAnchor.constraint(equalTo: scrollView.topAnchor),
            contentView.leadingAnchor.constraint(equalTo: scrollView.leadingAnchor),
            contentView.trailingAnchor.constraint(equalTo: scrollView.trailingAnchor),
            contentView.bottomAnchor.constraint(equalTo: scrollView.bottomAnchor),
            contentView.widthAnchor.constraint(equalTo: scrollView.widthAnchor)
        ])

        contentView.addSubview(iconImageView)
        iconImageView.translatesAutoresizingMaskIntoConstraints = false
        NSLayoutConstraint.activate([
            iconImageView.topAnchor.constraint(equalTo: contentView.topAnchor, constant: 40),
            iconImageView.centerXAnchor.constraint(equalTo: contentView.centerXAnchor),
            iconImageView.widthAnchor.constraint(equalToConstant: 80),
            iconImageView.heightAnchor.constraint(equalToConstant: 80)
        ])

        contentView.addSubview(titleLabel)
        titleLabel.translatesAutoresizingMaskIntoConstraints = false
        NSLayoutConstraint.activate([
            titleLabel.topAnchor.constraint(equalTo: iconImageView.bottomAnchor, constant: 20),
            titleLabel.leadingAnchor.constraint(equalTo: contentView.leadingAnchor, constant: 20),
            titleLabel.trailingAnchor.constraint(equalTo: contentView.trailingAnchor, constant: -20)
        ])

        contentView.addSubview(messageLabel)
        messageLabel.translatesAutoresizingMaskIntoConstraints = false
        NSLayoutConstraint.activate([
            messageLabel.topAnchor.constraint(equalTo: titleLabel.bottomAnchor, constant: 12),
            messageLabel.leadingAnchor.constraint(equalTo: contentView.leadingAnchor, constant: 20),
            messageLabel.trailingAnchor.constraint(equalTo: contentView.trailingAnchor, constant: -20)
        ])

        contentView.addSubview(issuesStackView)
        issuesStackView.translatesAutoresizingMaskIntoConstraints = false
        NSLayoutConstraint.activate([
            issuesStackView.topAnchor.constraint(equalTo: messageLabel.bottomAnchor, constant: 24),
            issuesStackView.leadingAnchor.constraint(equalTo: contentView.leadingAnchor, constant: 20),
            issuesStackView.trailingAnchor.constraint(equalTo: contentView.trailingAnchor, constant: -20)
        ])

        if !securityResult.remoteApps.isEmpty {
            contentView.addSubview(uninstallLabel)
            uninstallLabel.translatesAutoresizingMaskIntoConstraints = false
            NSLayoutConstraint.activate([
                uninstallLabel.topAnchor.constraint(equalTo: issuesStackView.bottomAnchor, constant: 24),
                uninstallLabel.leadingAnchor.constraint(equalTo: contentView.leadingAnchor, constant: 20),
                uninstallLabel.trailingAnchor.constraint(equalTo: contentView.trailingAnchor, constant: -20)
            ])

            contentView.addSubview(appsStackView)
            appsStackView.translatesAutoresizingMaskIntoConstraints = false
            NSLayoutConstraint.activate([
                appsStackView.topAnchor.constraint(equalTo: uninstallLabel.bottomAnchor, constant: 12),
                appsStackView.leadingAnchor.constraint(equalTo: contentView.leadingAnchor, constant: 20),
                appsStackView.trailingAnchor.constraint(equalTo: contentView.trailingAnchor, constant: -20)
            ])

            contentView.addSubview(instructionLabel)
            instructionLabel.translatesAutoresizingMaskIntoConstraints = false
            NSLayoutConstraint.activate([
                instructionLabel.topAnchor.constraint(equalTo: appsStackView.bottomAnchor, constant: 24),
                instructionLabel.leadingAnchor.constraint(equalTo: contentView.leadingAnchor, constant: 20),
                instructionLabel.trailingAnchor.constraint(equalTo: contentView.trailingAnchor, constant: -20)
            ])

            contentView.addSubview(contactLabel)
            contactLabel.translatesAutoresizingMaskIntoConstraints = false
            NSLayoutConstraint.activate([
                contactLabel.topAnchor.constraint(equalTo: instructionLabel.bottomAnchor, constant: 24),
                contactLabel.leadingAnchor.constraint(equalTo: contentView.leadingAnchor, constant: 20),
                contactLabel.trailingAnchor.constraint(equalTo: contentView.trailingAnchor, constant: -20),
                contactLabel.bottomAnchor.constraint(equalTo: contentView.bottomAnchor, constant: -40)
            ])
        } else {
            contentView.addSubview(contactLabel)
            contactLabel.translatesAutoresizingMaskIntoConstraints = false
            NSLayoutConstraint.activate([
                contactLabel.topAnchor.constraint(equalTo: issuesStackView.bottomAnchor, constant: 24),
                contactLabel.leadingAnchor.constraint(equalTo: contentView.leadingAnchor, constant: 20),
                contactLabel.trailingAnchor.constraint(equalTo: contentView.trailingAnchor, constant: -20),
                contactLabel.bottomAnchor.constraint(equalTo: contentView.bottomAnchor, constant: -40)
            ])
        }
    }

    private func displaySecurityIssues() {
        for issue in securityResult.issues {
            let issueView = createIssueView(issue: issue)
            issuesStackView.addArrangedSubview(issueView)
        }

        if !securityResult.remoteApps.isEmpty {
            for appName in securityResult.remoteApps {
                let appView = createAppView(appName: appName)
                appsStackView.addArrangedSubview(appView)
            }
        }
    }

    private func createIssueView(issue: SecurityIssue) -> UIView {
        let container = UIView()
        container.backgroundColor = .systemRed.withAlphaComponent(0.1)
        container.layer.cornerRadius = 8

        let iconView = UIImageView()
        iconView.image = UIImage(systemName: "xmark.circle.fill")
        iconView.tintColor = .systemRed
        iconView.translatesAutoresizingMaskIntoConstraints = false
        container.addSubview(iconView)

        let label = UILabel()
        label.text = issue.rawValue
        label.font = .systemFont(ofSize: 14, weight: .medium)
        label.textColor = .systemRed
        label.numberOfLines = 0
        label.translatesAutoresizingMaskIntoConstraints = false
        container.addSubview(label)

        NSLayoutConstraint.activate([
            iconView.leadingAnchor.constraint(equalTo: container.leadingAnchor, constant: 12),
            iconView.centerYAnchor.constraint(equalTo: container.centerYAnchor),
            iconView.widthAnchor.constraint(equalToConstant: 20),
            iconView.heightAnchor.constraint(equalToConstant: 20),
            label.leadingAnchor.constraint(equalTo: iconView.trailingAnchor, constant: 8),
            label.trailingAnchor.constraint(equalTo: container.trailingAnchor, constant: -12),
            label.topAnchor.constraint(equalTo: container.topAnchor, constant: 12),
            label.bottomAnchor.constraint(equalTo: container.bottomAnchor, constant: -12)
        ])

        return container
    }

    private func createAppView(appName: String) -> UIView {
        let container = UIView()
        container.backgroundColor = .systemOrange.withAlphaComponent(0.1)
        container.layer.cornerRadius = 8

        let iconView = UIImageView()
        iconView.image = UIImage(systemName: "app.badge.fill")
        iconView.tintColor = .systemOrange
        iconView.translatesAutoresizingMaskIntoConstraints = false
        container.addSubview(iconView)

        let label = UILabel()
        label.text = appName
        label.font = .systemFont(ofSize: 14, weight: .medium)
        label.textColor = .systemOrange
        label.translatesAutoresizingMaskIntoConstraints = false
        container.addSubview(label)

        NSLayoutConstraint.activate([
            iconView.leadingAnchor.constraint(equalTo: container.leadingAnchor, constant: 12),
            iconView.centerYAnchor.constraint(equalTo: container.centerYAnchor),
            iconView.widthAnchor.constraint(equalToConstant: 20),
            iconView.heightAnchor.constraint(equalToConstant: 20),
            label.leadingAnchor.constraint(equalTo: iconView.trailingAnchor, constant: 8),
            label.trailingAnchor.constraint(equalTo: container.trailingAnchor, constant: -12),
            label.topAnchor.constraint(equalTo: container.topAnchor, constant: 12),
            label.bottomAnchor.constraint(equalTo: container.bottomAnchor, constant: -12)
        ])

        return container
    }
}
