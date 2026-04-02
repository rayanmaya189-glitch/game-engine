import UIKit

class SupportViewController: UIViewController {

    private var viewModel = SupportViewModel()

    private enum Section: Int, CaseIterable {
        case faq
        case contact
        case tickets
    }

    private lazy var tableView: UITableView = {
        let tv = UITableView(frame: .zero, style: .insetGrouped)
        tv.backgroundColor = .clear
        tv.register(UITableViewCell.self, forCellReuseIdentifier: "SupportCell")
        return tv
    }()

    private let activityIndicator: UIActivityIndicatorView = {
        let indicator = UIActivityIndicatorView(style: .large)
        indicator.color = UIColor(hex: "#FF6B35")
        indicator.hidesWhenStopped = true
        return indicator
    }()

    private let sectionTitles = ["FAQ", "Contact Us", "My Tickets"]

    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        setupBindings()
        viewModel.loadFaqs()
        viewModel.loadTickets()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Support"

        view.addSubview(tableView)
        view.addSubview(activityIndicator)

        tableView.delegate = self
        tableView.dataSource = self

        tableView.snp.makeConstraints { make in
            make.edges.equalTo(view.safeAreaLayoutGuide)
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
                if let error = state.error {
                    let alert = UIAlertController(title: "Error", message: error, preferredStyle: .alert)
                    alert.addAction(UIAlertAction(title: "OK", style: .cancel))
                    self?.present(alert, animated: true)
                }
            }
        }
    }

    @objc private func createTicketTapped() {
        let alert = UIAlertController(title: "New Ticket", message: nil, preferredStyle: .alert)
        alert.addTextField { $0.placeholder = "Subject" }
        alert.addTextField { $0.placeholder = "Message"; $0.clearButtonMode = .whileEditing }
        alert.addAction(UIAlertAction(title: "Cancel", style: .cancel))
        alert.addAction(UIAlertAction(title: "Submit", style: .default) { [weak self] _ in
            guard let subject = alert.textFields?[0].text, !subject.isEmpty,
                  let message = alert.textFields?[1].text, !message.isEmpty else { return }
            self?.viewModel.createTicket(subject: subject, message: message)
        })
        present(alert, animated: true)
    }

    private func contactAction(_ type: String) {
        let alert = UIAlertController(title: type, message: "Opening \(type.lowercased())...", preferredStyle: .alert)
        alert.addAction(UIAlertAction(title: "OK", style: .cancel))
        present(alert, animated: true)
    }
}

// MARK: - UITableViewDataSource & Delegate

extension SupportViewController: UITableViewDataSource, UITableViewDelegate {

    func numberOfSections(in tableView: UITableView) -> Int {
        return Section.allCases.count
    }

    func tableView(_ tableView: UITableView, titleForHeaderInSection section: Int) -> String? {
        return sectionTitles[section]
    }

    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        guard let sec = Section(rawValue: section) else { return 0 }
        switch sec {
        case .faq: return viewModel.state.faqs.count
        case .contact: return 3
        case .tickets: return viewModel.state.tickets.count + 1
        }
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: "SupportCell", for: indexPath)
        guard let sec = Section(rawValue: indexPath.section) else { return cell }

        var config = cell.defaultContentConfiguration()
        config.textProperties.color = .white
        config.secondaryTextProperties.color = .white.withAlphaComponent(0.6)
        config.imageProperties.tintColor = UIColor(hex: "#FF6B35")
        cell.accessoryView = nil
        cell.accessoryType = .none

        switch sec {
        case .faq:
            let faq = viewModel.state.faqs[indexPath.row]
            let isExpanded = viewModel.state.expandedFaqIds.contains(faq.id)
            config.text = faq.question
            if isExpanded {
                config.secondaryText = faq.answer
            }
            cell.accessoryType = isExpanded ? .none : .disclosureIndicator

        case .contact:
            let contacts = [
                ("Live Chat", "message.fill"),
                ("Email Support", "envelope.fill"),
                ("Phone Support", "phone.fill")
            ]
            config.text = contacts[indexPath.row].0
            config.image = UIImage(systemName: contacts[indexPath.row].1)

        case .tickets:
            if indexPath.row == viewModel.state.tickets.count {
                config.text = "Create New Ticket"
                config.textProperties.color = UIColor(hex: "#FF6B35")
                config.image = UIImage(systemName: "plus.circle.fill")
            } else {
                let ticket = viewModel.state.tickets[indexPath.row]
                config.text = ticket.subject
                config.secondaryText = ticket.createdAt
                config.image = UIImage(systemName: "ticket.fill")
                let badge = UILabel()
                badge.text = " \(ticket.status) "
                badge.font = .systemFont(ofSize: 12)
                badge.textColor = .white
                badge.backgroundColor = ticket.status == "open" ? .systemGreen : .systemGray
                badge.layer.cornerRadius = 4
                badge.clipsToBounds = true
                badge.sizeToFit()
                cell.accessoryView = badge
            }
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
        case .faq:
            let faq = viewModel.state.faqs[indexPath.row]
            viewModel.toggleFaq(faq.id)
            tableView.reloadRows(at: [indexPath], with: .automatic)

        case .contact:
            let actions = ["Live Chat", "Email", "Phone"]
            contactAction(actions[indexPath.row])

        case .tickets:
            if indexPath.row == viewModel.state.tickets.count {
                createTicketTapped()
            }
        }
    }
}
