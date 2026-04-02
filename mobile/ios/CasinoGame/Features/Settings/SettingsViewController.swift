import UIKit
import SnapKit

class SettingsViewController: UIViewController {

    private var viewModel = SettingsViewModel()

    private enum Section: Int, CaseIterable {
        case profile
        case preferences
        case security
        case account
    }

    private let tableView: UITableView = {
        let tv = UITableView(frame: .zero, style: .insetGrouped)
        tv.backgroundColor = .clear
        tv.register(UITableViewCell.self, forCellReuseIdentifier: "SettingsCell")
        return tv
    }()

    private let activityIndicator: UIActivityIndicatorView = {
        let indicator = UIActivityIndicatorView(style: .large)
        indicator.color = UIColor(hex: "#FF6B35")
        indicator.hidesWhenStopped = true
        return indicator
    }()

    private let sectionTitles = ["Profile", "Preferences", "Security", "Account"]

    private var profileRows: [(String, String)] {
        let user = viewModel.state.user
        return [
            ("person.fill", user?.username ?? "Username"),
            ("envelope.fill", user?.email ?? "Email"),
            ("camera.fill", "Change Avatar")
        ]
    }

    private var preferencesRows: [(String, String)] {
        [
            ("globe", "Language (\(viewModel.state.settings.language.uppercased()))"),
            ("bell.fill", "Push Notifications")
        ]
    }

    private var securityRows: [(String, String)] {
        [
            ("lock.shield.fill", "Two-Factor Authentication"),
            ("key.fill", "Change Password")
        ]
    }

    private var accountRows: [(String, String)] {
        [
            ("arrow.right.square.fill", "Logout"),
            ("trash.fill", "Delete Account")
        ]
    }

    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        setupBindings()
        viewModel.loadSettings()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Settings"

        view.addSubview(tableView)
        view.addSubview(activityIndicator)

        tableView.snp.makeConstraints { make in
            make.edges.equalTo(view.safeAreaLayoutGuide)
        }

        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        tableView.delegate = self
        tableView.dataSource = self
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.tableView.reloadData()
                self?.activityIndicator.stopAnimating()
                if state.isLoggedOut {
                    self?.navigateToLogin()
                }
                if let error = state.error {
                    let alert = UIAlertController(title: "Error", message: error, preferredStyle: .alert)
                    alert.addAction(UIAlertAction(title: "OK", style: .cancel))
                    self?.present(alert, animated: true)
                }
            }
        }
    }

    private func rows(for section: Section) -> [(String, String)] {
        switch section {
        case .profile: return profileRows
        case .preferences: return preferencesRows
        case .security: return securityRows
        case .account: return accountRows
        }
    }

    private func navigateToLogin() {
        if let scene = UIApplication.shared.connectedScenes.first as? UIWindowScene,
           let window = scene.windows.first {
            let coordinator = AppCoordinator(window: window)
            coordinator.start()
        }
    }
}

// MARK: - UITableViewDataSource & Delegate

extension SettingsViewController: UITableViewDataSource, UITableViewDelegate {
    func numberOfSections(in tableView: UITableView) -> Int {
        return Section.allCases.count
    }

    func tableView(_ tableView: UITableView, titleForHeaderInSection section: Int) -> String? {
        return sectionTitles[section]
    }

    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        guard let sec = Section(rawValue: section) else { return 0 }
        return rows(for: sec).count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "SettingsCell", for: indexPath)
        guard let sec = Section(rawValue: indexPath.section) else { return cell }
        let item = rows(for: sec)[indexPath.row]

        var config = cell.defaultContentConfiguration()
        config.text = item.1
        config.image = UIImage(systemName: item.0)
        config.textProperties.color = .white
        config.imageProperties.tintColor = UIColor(hex: "#FF6B35")

        if sec == .preferences && indexPath.row == 1 {
            let toggle = UISwitch()
            toggle.isOn = viewModel.state.settings.notificationsEnabled
            toggle.onTintColor = UIColor(hex: "#FF6B35")
            toggle.addTarget(self, action: #selector(notificationsToggled(_:)), for: .valueChanged)
            cell.accessoryView = toggle
        } else if sec == .security && indexPath.row == 0 {
            let toggle = UISwitch()
            toggle.isOn = viewModel.state.settings.twoFactorEnabled
            toggle.onTintColor = UIColor(hex: "#FF6B35")
            toggle.addTarget(self, action: #selector(twoFactorToggled(_:)), for: .valueChanged)
            cell.accessoryView = toggle
        } else {
            cell.accessoryView = nil
            cell.accessoryType = .disclosureIndicator
        }

        if sec == .account {
            config.textProperties.color = indexPath.row == 0 ? .systemRed : UIColor.systemRed.withAlphaComponent(0.7)
        }

        cell.contentConfiguration = config
        cell.backgroundColor = UIColor(hex: "#1E1E3F")
        cell.selectionStyle = .none
        return cell
    }

    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
        guard let sec = Section(rawValue: indexPath.section) else { return }

        switch sec {
        case .account:
            if indexPath.row == 0 {
                showConfirmation(title: "Logout", message: "Are you sure you want to logout?") {
                    self.viewModel.logout()
                }
            } else {
                showConfirmation(title: "Delete Account", message: "This action cannot be undone. Delete your account?") {
                    self.viewModel.deleteAccount()
                }
            }
        case .preferences where indexPath.row == 0:
            showLanguagePicker()
        default: break
        }
    }

    private func showConfirmation(title: String, message: String, onConfirm: @escaping () -> Void) {
        let alert = UIAlertController(title: title, message: message, preferredStyle: .alert)
        alert.addAction(UIAlertAction(title: "Cancel", style: .cancel))
        alert.addAction(UIAlertAction(title: title, style: .destructive) { _ in onConfirm() })
        present(alert, animated: true)
    }

    private func showLanguagePicker() {
        let alert = UIAlertController(title: "Language", message: nil, preferredStyle: .actionSheet)
        [("en", "English"), ("es", "Spanish"), ("fr", "French"), ("de", "German")].forEach { code, name in
            alert.addAction(UIAlertAction(title: name, style: .default) { [weak self] _ in
                self?.viewModel.updateLanguage(code)
            })
        }
        alert.addAction(UIAlertAction(title: "Cancel", style: .cancel))
        present(alert, animated: true)
    }

    @objc private func notificationsToggled(_ sender: UISwitch) {
        viewModel.toggleNotifications(sender.isOn)
    }

    @objc private func twoFactorToggled(_ sender: UISwitch) {
        viewModel.toggleTwoFactor(sender.isOn)
    }
}
