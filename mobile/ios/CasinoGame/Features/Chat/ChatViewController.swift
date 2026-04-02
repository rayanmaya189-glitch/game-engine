import UIKit
import SnapKit

class ChatViewController: UIViewController {

    private var viewModel = ChatViewModel()

    private let onlineUsersBar: UICollectionView = {
        let layout = UICollectionViewFlowLayout()
        layout.scrollDirection = .horizontal
        layout.itemSize = CGSize(width: 60, height: 70)
        layout.minimumInteritemSpacing = 12
        layout.sectionInset = UIEdgeInsets(top: 0, left: 16, bottom: 0, right: 16)
        let cv = UICollectionView(frame: .zero, collectionViewLayout: layout)
        cv.backgroundColor = UIColor(hex: "#1E1E3F")
        cv.showsHorizontalScrollIndicator = false
        cv.register(OnlineUserCell.self, forCellWithReuseIdentifier: "OnlineUserCell")
        return cv
    }()

    private let tableView: UITableView = {
        let tv = UITableView()
        tv.backgroundColor = .clear
        tv.separatorStyle = .none
        tv.register(MessageCell.self, forCellReuseIdentifier: "MessageCell")
        tv.keyboardDismissMode = .interactive
        return tv
    }()

    private let inputContainer: UIView = {
        let view = UIView()
        view.backgroundColor = UIColor(hex: "#1E1E3F")
        return view
    }()

    private let messageField: UITextField = {
        let tf = UITextField()
        tf.placeholder = "Type a message..."
        tf.textColor = .white
        tf.attributedPlaceholder = NSAttributedString(
            string: "Type a message...",
            attributes: [.foregroundColor: UIColor.white.withAlphaComponent(0.4)]
        )
        tf.backgroundColor = UIColor(hex: "#0F0F23")
        tf.layer.cornerRadius = 20
        tf.leftView = UIView(frame: CGRect(x: 0, y: 0, width: 16, height: 0))
        tf.leftViewMode = .always
        tf.rightView = UIView(frame: CGRect(x: 0, y: 0, width: 16, height: 0))
        tf.rightViewMode = .always
        return tf
    }()

    private let sendButton: UIButton = {
        let btn = UIButton(type: .system)
        btn.setImage(UIImage(systemName: "paperplane.fill"), for: .normal)
        btn.tintColor = UIColor(hex: "#FF6B35")
        return btn
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
        viewModel.loadMessages()
        viewModel.loadOnlineUsers()
    }


    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Chat"

        [onlineUsersBar, tableView, inputContainer, activityIndicator].forEach { view.addSubview($0) }
        inputContainer.addSubview(messageField)
        inputContainer.addSubview(sendButton)

        onlineUsersBar.snp.makeConstraints { make in
            make.top.equalTo(view.safeAreaLayoutGuide)
            make.leading.trailing.equalToSuperview()
            make.height.equalTo(70)
        }

        tableView.snp.makeConstraints { make in
            make.top.equalTo(onlineUsersBar.snp.bottom)
            make.leading.trailing.equalToSuperview()
            make.bottom.equalTo(inputContainer.snp.top)
        }

        inputContainer.snp.makeConstraints { make in
            make.leading.trailing.equalToSuperview()
            make.bottom.equalTo(view.safeAreaLayoutGuide)
            make.height.equalTo(56)
        }

        messageField.snp.makeConstraints { make in
            make.leading.equalToSuperview().inset(16)
            make.centerY.equalToSuperview()
            make.height.equalTo(40)
            make.trailing.equalTo(sendButton.snp.leading).offset(-8)
        }

        sendButton.snp.makeConstraints { make in
            make.trailing.equalToSuperview().inset(16)
            make.centerY.equalToSuperview()
            make.size.equalTo(36)
        }

        activityIndicator.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        tableView.delegate = self
        tableView.dataSource = self
        onlineUsersBar.delegate = self
        onlineUsersBar.dataSource = self

        sendButton.addTarget(self, action: #selector(sendTapped), for: .touchUpInside)
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.tableView.reloadData()
                self?.onlineUsersBar.reloadData()
                self?.activityIndicator.stopAnimating()
                self?.scrollToBottom()
                if let error = state.error {
                    let alert = UIAlertController(title: "Error", message: error, preferredStyle: .alert)
                    alert.addAction(UIAlertAction(title: "OK", style: .cancel))
                    self?.present(alert, animated: true)
                }
            }
        }
    }

    private func scrollToBottom() {
        guard !viewModel.state.messages.isEmpty else { return }
        let indexPath = IndexPath(row: viewModel.state.messages.count - 1, section: 0)
        tableView.scrollToRow(at: indexPath, at: .bottom, animated: true)
    }

    @objc private func sendTapped() {
        guard let text = messageField.text, !text.isEmpty else { return }
        viewModel.sendMessage(text)
        messageField.text = ""
    }
}

// MARK: - UITableViewDataSource & Delegate

extension ChatViewController: UITableViewDataSource, UITableViewDelegate {
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return viewModel.state.messages.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "MessageCell", for: indexPath) as! MessageCell
        cell.configure(with: viewModel.state.messages[indexPath.row])
        return cell
    }

    func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
        return UITableView.automaticDimension
    }

    func tableView(_ tableView: UITableView, estimatedHeightForRowAt indexPath: IndexPath) -> CGFloat {
        return 70
    }
}

// MARK: - Online Users Collection
extension ChatViewController: UICollectionViewDataSource, UICollectionViewDelegate {
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return viewModel.state.onlineUsers.count
    }

    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let cell = collectionView.dequeueReusableCell(withReuseIdentifier: "OnlineUserCell", for: indexPath) as! OnlineUserCell
        cell.configure(with: viewModel.state.onlineUsers[indexPath.row])
        return cell
    }
}

// MARK: - MessageCell
class MessageCell: UITableViewCell {
    private let senderLabel = UILabel()
    private let messageLabel = UILabel()
    private let timeLabel = UILabel()
    private let bubbleView = UIView()

    override init(style: UITableViewCell.CellStyle, reuseIdentifier: String?) {
        super.init(style: style, reuseIdentifier: reuseIdentifier)
        backgroundColor = .clear
        selectionStyle = .none

        bubbleView.backgroundColor = UIColor(hex: "#1E1E3F")
        bubbleView.layer.cornerRadius = 12
        senderLabel.textColor = UIColor(hex: "#FF6B35")
        senderLabel.font = .systemFont(ofSize: 12, weight: .semibold)
        messageLabel.textColor = .white
        messageLabel.font = .systemFont(ofSize: 15)
        messageLabel.numberOfLines = 0
        timeLabel.textColor = .white.withAlphaComponent(0.4)
        timeLabel.font = .systemFont(ofSize: 10)

        [bubbleView, senderLabel, messageLabel, timeLabel].forEach { contentView.addSubview($0) }

        bubbleView.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(4)
            make.leading.equalToSuperview().inset(16)
            make.trailing.lessThanOrEqualToSuperview().inset(60)
            make.bottom.equalToSuperview().offset(-4)
        }

        senderLabel.snp.makeConstraints { make in
            make.top.equalTo(bubbleView).offset(8)
            make.leading.equalTo(bubbleView).inset(12)
        }

        messageLabel.snp.makeConstraints { make in
            make.top.equalTo(senderLabel.snp.bottom).offset(4)
            make.leading.trailing.equalTo(bubbleView).inset(12)
        }

        timeLabel.snp.makeConstraints { make in
            make.top.equalTo(messageLabel.snp.bottom).offset(4)
            make.trailing.equalTo(bubbleView).inset(12)
            make.bottom.equalTo(bubbleView).inset(8)
        }
    }

    required init?(coder: NSCoder) { fatalError() }

    func configure(with message: ChatMessage) {
        senderLabel.text = message.senderName
        messageLabel.text = message.text
        let formatter = ISO8601DateFormatter()
        if let date = formatter.date(from: message.createdAt) {
            let display = DateFormatter()
            display.dateFormat = "HH:mm"
            timeLabel.text = display.string(from: date)
        } else {
            timeLabel.text = message.createdAt
        }
    }
}

// MARK: - OnlineUserCell

class OnlineUserCell: UICollectionViewCell {
    private let avatarView = UIView()
    private let avatarLabel = UILabel()
    private let nameLabel = UILabel()

    override init(frame: CGRect) {
        super.init(frame: frame)

        avatarView.backgroundColor = UIColor(hex: "#FF6B35")
        avatarView.layer.cornerRadius = 24
        avatarLabel.textColor = .white
        avatarLabel.font = .systemFont(ofSize: 14, weight: .bold)
        avatarLabel.textAlignment = .center
        nameLabel.textColor = .white.withAlphaComponent(0.7)
        nameLabel.font = .systemFont(ofSize: 10)
        nameLabel.textAlignment = .center

        [avatarView, nameLabel].forEach { contentView.addSubview($0) }
        avatarView.addSubview(avatarLabel)

        avatarView.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(4)
            make.centerX.equalToSuperview()
            make.size.equalTo(48)
        }

        avatarLabel.snp.makeConstraints { make in
            make.center.equalToSuperview()
        }

        nameLabel.snp.makeConstraints { make in
            make.top.equalTo(avatarView.snp.bottom).offset(4)
            make.centerX.equalToSuperview()
            make.leading.trailing.equalToSuperview().inset(2)
        }
    }

    required init?(coder: NSCoder) { fatalError() }

    func configure(with user: OnlineUser) {
        nameLabel.text = user.username
        avatarLabel.text = String(user.username.prefix(2)).uppercased()
    }
}
