import UIKit
import SnapKit

class NotificationViewController: UIViewController {

    private var viewModel = NotificationViewModel()

    private let tableView: UITableView = {
        let tv = UITableView(frame: .zero, style: .plain)
        tv.backgroundColor = .clear
        tv.separatorStyle = .none
        tv.register(NotificationCell.self, forCellReuseIdentifier: "NotificationCell")
        tv.refreshControl = UIRefreshControl()
        return tv
    }()

    private let emptyLabel: UILabel = {
        let label = UILabel()
        label.text = "No notifications yet"
        label.textAlignment = .center
        label.textColor = .white.withAlphaComponent(0.5)
        label.font = .systemFont(ofSize: 16)
        label.isHidden = true
        return label
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
        viewModel.loadNotifications()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Notifications"

        [tableView, emptyLabel, activityIndicator].forEach { view.addSubview($0) }

        tableView.snp.makeConstraints { make in
            make.edges.equalTo(view.safeAreaLayoutGuide)
        }

        emptyLabel.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        tableView.delegate = self
        tableView.dataSource = self
        tableView.refreshControl?.addTarget(self, action: #selector(refreshData), for: .valueChanged)
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.tableView.reloadData()
                self?.tableView.refreshControl?.endRefreshing()
                self?.activityIndicator.stopAnimating()
                self?.emptyLabel.isHidden = !state.notifications.isEmpty
                if let error = state.error {
                    let alert = UIAlertController(title: "Error", message: error, preferredStyle: .alert)
                    alert.addAction(UIAlertAction(title: "OK", style: .cancel))
                    self?.present(alert, animated: true)
                }
            }
        }
    }

    @objc private func refreshData() {
        viewModel.loadNotifications()
    }
}

// MARK: - UITableViewDataSource & Delegate

extension NotificationViewController: UITableViewDataSource, UITableViewDelegate {
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return viewModel.state.notifications.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "NotificationCell", for: indexPath) as! NotificationCell
        cell.configure(with: viewModel.state.notifications[indexPath.row])
        return cell
    }

    func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
        return UITableView.automaticDimension
    }

    func tableView(_ tableView: UITableView, estimatedHeightForRowAt indexPath: IndexPath) -> CGFloat {
        return 80
    }

    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
        let notification = viewModel.state.notifications[indexPath.row]
        if !notification.isRead {
            viewModel.markRead(notification.id)
        }
    }

    func tableView(_ tableView: UITableView, trailingSwipeActionsConfigurationForRowAt indexPath: IndexPath) -> UISwipeActionsConfiguration? {
        let notification = viewModel.state.notifications[indexPath.row]

        let deleteAction = UIContextualAction(style: .destructive, title: "Delete") { [weak self] _, _, completion in
            self?.viewModel.delete(notification.id)
            completion(true)
        }
        deleteAction.backgroundColor = .systemRed

        let readAction = UIContextualAction(style: .normal, title: notification.isRead ? "Unread" : "Read") { [weak self] _, _, completion in
            if !notification.isRead {
                self?.viewModel.markRead(notification.id)
            }
            completion(true)
        }
        readAction.backgroundColor = UIColor(hex: "#4CAF50")

        return UISwipeActionsConfiguration(actions: [deleteAction, readAction])
    }
}

// MARK: - NotificationCell

class NotificationCell: UITableViewCell {
    private let iconView = UIView()
    private let iconLabel = UILabel()
    private let titleLabel = UILabel()
    private let bodyLabel = UILabel()
    private let timeLabel = UILabel()
    private let unreadDot = UIView()

    override init(style: UITableViewCell.CellStyle, reuseIdentifier: String?) {
        super.init(style: style, reuseIdentifier: reuseIdentifier)
        backgroundColor = .clear
        selectionStyle = .none

        iconView.backgroundColor = UIColor(hex: "#1E1E3F")
        iconView.layer.cornerRadius = 24
        iconLabel.font = .systemFont(ofSize: 20)
        iconLabel.textAlignment = .center
        titleLabel.textColor = .white
        titleLabel.font = .systemFont(ofSize: 15, weight: .semibold)
        bodyLabel.textColor = .white.withAlphaComponent(0.7)
        bodyLabel.font = .systemFont(ofSize: 13)
        bodyLabel.numberOfLines = 2
        timeLabel.textColor = .white.withAlphaComponent(0.4)
        timeLabel.font = .systemFont(ofSize: 11)
        unreadDot.backgroundColor = UIColor(hex: "#FF6B35")
        unreadDot.layer.cornerRadius = 5

        [iconView, titleLabel, bodyLabel, timeLabel, unreadDot].forEach { contentView.addSubview($0) }
        iconView.addSubview(iconLabel)

        iconView.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(8)
            make.leading.equalToSuperview().inset(16)
            make.size.equalTo(48)
        }

        iconLabel.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        titleLabel.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(10)
            make.leading.equalTo(iconView.snp.trailing).offset(12)
            make.trailing.equalTo(unreadDot.snp.leading).offset(-8)
        }

        bodyLabel.snp.makeConstraints { make in
            make.top.equalTo(titleLabel.snp.bottom).offset(4)
            make.leading.equalTo(iconView.snp.trailing).offset(12)
            make.trailing.equalToSuperview().inset(16)
        }

        timeLabel.snp.makeConstraints { make in
            make.top.equalTo(bodyLabel.snp.bottom).offset(4)
            make.leading.equalTo(iconView.snp.trailing).offset(12)
            make.bottom.equalToSuperview().inset(8)
        }

        unreadDot.snp.makeConstraints { make in
            make.centerY.equalTo(titleLabel)
            make.trailing.equalToSuperview().inset(16)
            make.size.equalTo(10)
        }
    }

    required init?(coder: NSCoder) { fatalError() }

    func configure(with notification: AppNotification) {
        titleLabel.text = notification.title
        bodyLabel.text = notification.body
        unreadDot.isHidden = notification.isRead
        timeLabel.text = formatTime(notification.createdAt)

        switch notification.type {
        case .bonus:
            iconLabel.text = "🎁"
            iconView.backgroundColor = UIColor(hex: "#4CAF50").withAlphaComponent(0.2)
        case .tournament:
            iconLabel.text = "🏆"
            iconView.backgroundColor = UIColor(hex: "#FFD700").withAlphaComponent(0.2)
        case .jackpot:
            iconLabel.text = "💰"
            iconView.backgroundColor = UIColor(hex: "#FF6B35").withAlphaComponent(0.2)
        case .system:
            iconLabel.text = "ℹ️"
            iconView.backgroundColor = UIColor(hex: "#2196F3").withAlphaComponent(0.2)
        }
    }

    private func formatTime(_ dateString: String) -> String {
        let formatter = ISO8601DateFormatter()
        guard let date = formatter.date(from: dateString) else { return dateString }
        let display = DateFormatter()
        display.dateStyle = .medium
        display.timeStyle = .short
        return display.string(from: date)
    }
}
