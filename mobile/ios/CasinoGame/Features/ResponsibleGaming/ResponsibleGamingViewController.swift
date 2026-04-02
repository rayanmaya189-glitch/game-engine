import UIKit

class ResponsibleGamingViewController: UIViewController {

    private var viewModel = ResponsibleGamingViewModel()

    private enum Section: Int, CaseIterable {
        case limits
        case selfExclusion
        case realityCheck
    }

    private lazy var tableView: UITableView = {
        let tv = UITableView(frame: .zero, style: .insetGrouped)
        tv.backgroundColor = .clear
        tv.register(UITableViewCell.self, forCellReuseIdentifier: "RGCell")
        return tv
    }()

    private lazy var saveButton: UIButton = {
        let button = UIButton(type: .system)
        button.setTitle("Save Settings", for: .normal)
        button.setTitleColor(.white, for: .normal)
        button.backgroundColor = UIColor(hex: "#FF6B35")
        button.layer.cornerRadius = 12
        button.titleLabel?.font = .systemFont(ofSize: 16, weight: .semibold)
        button.addTarget(self, action: #selector(saveTapped), for: .touchUpInside)
        return button
    }()

    private let activityIndicator: UIActivityIndicatorView = {
        let indicator = UIActivityIndicatorView(style: .large)
        indicator.color = UIColor(hex: "#FF6B35")
        indicator.hidesWhenStopped = true
        return indicator
    }()

    private let sectionTitles = ["Limits", "Self-Exclusion", "Reality Check"]
    private let exclusionOptions = [("None", 0), ("1 Day", 1), ("7 Days", 7), ("30 Days", 30), ("Permanent", 36500)]

    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        setupBindings()
        viewModel.loadLimits()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Responsible Gaming"

        view.addSubview(tableView)
        view.addSubview(saveButton)
        view.addSubview(activityIndicator)

        tableView.delegate = self
        tableView.dataSource = self

        tableView.snp.makeConstraints { make in
            make.top.leading.trailing.equalTo(view.safeAreaLayoutGuide)
            make.bottom.equalTo(saveButton.snp.top).offset(-12)
        }
        saveButton.snp.makeConstraints { make in
            make.leading.trailing.equalToSuperview().inset(16)
            make.bottom.equalTo(view.safeAreaLayoutGuide).offset(-8)
            make.height.equalTo(50)
        }
        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.tableView.reloadData()
                self?.activityIndicator.stopAnimating()
                if state.isSaved {
                    let alert = UIAlertController(title: "Saved", message: "Your settings have been updated.", preferredStyle: .alert)
                    alert.addAction(UIAlertAction(title: "OK", style: .cancel))
                    self?.present(alert, animated: true)
                }
                if let error = state.error {
                    let alert = UIAlertController(title: "Error", message: error, preferredStyle: .alert)
                    alert.addAction(UIAlertAction(title: "OK", style: .cancel))
                    self?.present(alert, animated: true)
                }
            }
        }
    }

    @objc private func saveTapped() {
        viewModel.saveLimits()
    }

    @objc private func depositSliderChanged(_ sender: UISlider) {
        viewModel.updateDepositLimit(Double(Int(sender.value)))
    }

    @objc private func lossSliderChanged(_ sender: UISlider) {
        viewModel.updateLossLimit(Double(Int(sender.value)))
    }

    @objc private func sessionSliderChanged(_ sender: UISlider) {
        viewModel.updateSessionLimit(Int(sender.value))
    }

    @objc private func coolOffToggled(_ sender: UISwitch) {
        viewModel.toggleCoolOff(sender.isOn)
    }

    private func makeSlider(min: Float, max: Float, value: Float, action: Selector) -> UISlider {
        let slider = UISlider()
        slider.minimumValue = min
        slider.maximumValue = max
        slider.value = value
        slider.tintColor = UIColor(hex: "#FF6B35")
        slider.addTarget(self, action: action, for: .valueChanged)
        return slider
    }
}

// MARK: - UITableViewDataSource & Delegate

extension ResponsibleGamingViewController: UITableViewDataSource, UITableViewDelegate {

    func numberOfSections(in tableView: UITableView) -> Int {
        return Section.allCases.count
    }

    func tableView(_ tableView: UITableView, titleForHeaderInSection section: Int) -> String? {
        return sectionTitles[section]
    }

    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        guard let sec = Section(rawValue: section) else { return 0 }
        switch sec {
        case .limits: return 3
        case .selfExclusion: return exclusionOptions.count
        case .realityCheck: return 1
        }
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "RGCell", for: indexPath)
        guard let sec = Section(rawValue: indexPath.section) else { return cell }

        var config = cell.defaultContentConfiguration()
        config.textProperties.color = .white
        config.secondaryTextProperties.color = .white.withAlphaComponent(0.6)
        config.imageProperties.tintColor = UIColor(hex: "#FF6B35")
        cell.accessoryView = nil
        cell.accessoryType = .none
        cell.backgroundColor = UIColor(hex: "#1E1E3F")
        cell.selectionStyle = .none

        let limits = viewModel.state.limits

        switch sec {
        case .limits:
            switch indexPath.row {
            case 0:
                config.text = "Deposit Limit"
                config.secondaryText = "$\(Int(limits.depositLimit))"
                let slider = makeSlider(min: 50, max: 50000, value: Float(limits.depositLimit), action: #selector(depositSliderChanged(_:)))
                slider.frame = CGRect(x: 0, y: 0, width: 150, height: 30)
                cell.accessoryView = slider
            case 1:
                config.text = "Loss Limit"
                config.secondaryText = "$\(Int(limits.lossLimit))"
                let slider = makeSlider(min: 50, max: 50000, value: Float(limits.lossLimit), action: #selector(lossSliderChanged(_:)))
                slider.frame = CGRect(x: 0, y: 0, width: 150, height: 30)
                cell.accessoryView = slider
            default:
                config.text = "Session Limit"
                config.secondaryText = "\(limits.sessionLimit) minutes"
                let slider = makeSlider(min: 15, max: 480, value: Float(limits.sessionLimit), action: #selector(sessionSliderChanged(_:)))
                slider.frame = CGRect(x: 0, y: 0, width: 150, height: 30)
                cell.accessoryView = slider
            }

        case .selfExclusion:
            let option = exclusionOptions[indexPath.row]
            config.text = option.0
            if limits.selfExclusionDays == option.1 {
                cell.accessoryType = .checkmark
                cell.tintColor = UIColor(hex: "#FF6B35")
            }

        case .realityCheck:
            config.text = "Cool-Off Period"
            config.secondaryText = "Temporarily pause your account"
            let toggle = UISwitch()
            toggle.isOn = limits.coolOffEnabled
            toggle.onTintColor = UIColor(hex: "#FF6B35")
            toggle.addTarget(self, action: #selector(coolOffToggled(_:)), for: .valueChanged)
            cell.accessoryView = toggle
        }

        cell.contentConfiguration = config
        return cell
    }

    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
        guard let sec = Section(rawValue: indexPath.section), sec == .selfExclusion else { return }
        let option = exclusionOptions[indexPath.row]
        viewModel.setSelfExclusion(days: option.1)
        tableView.reloadSections(IndexSet(integer: indexPath.section), with: .automatic)
    }
}
