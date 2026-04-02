import UIKit
import SnapKit

class KycVerificationViewController: UIViewController, UIImagePickerControllerDelegate, UINavigationControllerDelegate {

    private var viewModel = KycVerificationViewModel()
    private var selectedDocumentType: String?

    private let scrollView: UIScrollView = {
        let scrollView = UIScrollView()
        scrollView.showsVerticalScrollIndicator = false
        return scrollView
    }()

    private let contentView = UIView()

    private let levelTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "Verification Level"
        label.textColor = .white
        label.font = .systemFont(ofSize: 18, weight: .bold)
        return label
    }()

    private let levelProgressView: UIProgressView = {
        let progress = UIProgressView(progressViewStyle: .bar)
        progress.trackTintColor = UIColor(hex: "#1E1E3F")
        progress.progressTintColor = UIColor(hex: "#FF6B35")
        progress.layer.cornerRadius = 4
        progress.clipsToBounds = true
        return progress
    }()

    private let levelLabelsStack: UIStackView = {
        let stack = UIStackView()
        stack.distribution = .equalSpacing
        return stack
    }()

    private let documentsTitleLabel: UILabel = {
        let label = UILabel()
        label.text = "Required Documents"
        label.textColor = .white
        label.font = .systemFont(ofSize: 18, weight: .bold)
        return label
    }()

    private let documentStack: UIStackView = {
        let stack = UIStackView()
        stack.axis = .vertical
        stack.spacing = 12
        return stack
    }()

    private let submitButton: UIButton = {
        let button = UIButton(type: .system)
        button.setTitle("Submit for Verification", for: .normal)
        button.setTitleColor(.white, for: .normal)
        button.backgroundColor = UIColor(hex: "#FF6B35").withAlphaComponent(0.5)
        button.layer.cornerRadius = 12
        button.titleLabel?.font = .systemFont(ofSize: 16, weight: .semibold)
        button.isEnabled = false
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

    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        setupBindings()
        viewModel.loadStatus()
    }

    private func setupUI() {
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "KYC Verification"

        view.addSubview(scrollView)
        scrollView.addSubview(contentView)

        [levelTitleLabel, levelProgressView, levelLabelsStack,
         documentsTitleLabel, documentStack, submitButton, errorLabel].forEach {
            contentView.addSubview($0)
        }

        scrollView.snp.makeConstraints { make in
            make.edges.equalTo(view.safeAreaLayoutGuide)
        }
        contentView.snp.makeConstraints { make in
            make.edges.equalToSuperview()
            make.width.equalTo(view)
        }
        levelTitleLabel.snp.makeConstraints { make in
            make.top.equalToSuperview().offset(20)
            make.leading.trailing.equalToSuperview().inset(20)
        }
        levelProgressView.snp.makeConstraints { make in
            make.top.equalTo(levelTitleLabel.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview().inset(20)
            make.height.equalTo(8)
        }
        levelLabelsStack.snp.makeConstraints { make in
            make.top.equalTo(levelProgressView.snp.bottom).offset(8)
            make.leading.trailing.equalToSuperview().inset(20)
        }

        for i in 0...3 {
            let label = UILabel()
            label.text = "L\(i)"
            label.font = .systemFont(ofSize: 12, weight: .medium)
            label.textAlignment = .center
            label.textColor = .white.withAlphaComponent(0.5)
            label.tag = 100 + i
            levelLabelsStack.addArrangedSubview(label)
        }

        documentsTitleLabel.snp.makeConstraints { make in
            make.top.equalTo(levelLabelsStack.snp.bottom).offset(30)
            make.leading.trailing.equalToSuperview().inset(20)
        }
        documentStack.snp.makeConstraints { make in
            make.top.equalTo(documentsTitleLabel.snp.bottom).offset(16)
            make.leading.trailing.equalToSuperview().inset(20)
        }
        submitButton.snp.makeConstraints { make in
            make.top.equalTo(documentStack.snp.bottom).offset(24)
            make.leading.trailing.equalToSuperview().inset(20)
            make.height.equalTo(56)
        }
        errorLabel.snp.makeConstraints { make in
            make.top.equalTo(submitButton.snp.bottom).offset(12)
            make.leading.trailing.equalToSuperview().inset(20)
            make.bottom.equalToSuperview().offset(-20)
        }

        submitButton.addTarget(self, action: #selector(submitTapped), for: .touchUpInside)
    }

    private func buildDocumentRows(_ documents: [KycDocumentItem]) {
        documentStack.arrangedSubviews.forEach { $0.removeFromSuperview() }
        for doc in documents {
            documentStack.addArrangedSubview(createDocumentRow(doc))
        }
    }

    private func createDocumentRow(_ doc: KycDocumentItem) -> UIView {
        let container = UIView()
        container.backgroundColor = UIColor(hex: "#1E1E3F")
        container.layer.cornerRadius = 12

        let iconView = UIImageView(image: UIImage(systemName: doc.iconName))
        iconView.tintColor = UIColor(hex: "#FF6B35")
        iconView.contentMode = .scaleAspectFit

        let titleLabel = UILabel()
        titleLabel.text = doc.title
        titleLabel.textColor = .white
        titleLabel.font = .systemFont(ofSize: 15, weight: .medium)

        let statusLabel = UILabel()
        statusLabel.font = .systemFont(ofSize: 12, weight: .semibold)
        setKycStatusLabel(statusLabel, status: doc.status)

        let uploadButton = UIButton(type: .system)
        uploadButton.setTitle(doc.status == "not_uploaded" ? "Upload" : "Re-upload", for: .normal)
        uploadButton.setTitleColor(UIColor(hex: "#FF6B35"), for: .normal)
        uploadButton.titleLabel?.font = .systemFont(ofSize: 13, weight: .semibold)
        uploadButton.addAction(UIAction { [weak self] _ in
            self?.selectedDocumentType = doc.type
            self?.presentCamera()
        }, for: .touchUpInside)

        [iconView, titleLabel, statusLabel, uploadButton].forEach { container.addSubview($0) }
        container.snp.makeConstraints { make in make.height.equalTo(64) }
        iconView.snp.makeConstraints { make in
            make.leading.equalToSuperview().inset(16)
            make.centerY.equalToSuperview()
            make.size.equalTo(28)
        }
        titleLabel.snp.makeConstraints { make in
            make.leading.equalTo(iconView.snp.trailing).offset(12)
            make.centerY.equalToSuperview()
        }
        uploadButton.snp.makeConstraints { make in
            make.trailing.equalToSuperview().inset(16)
            make.centerY.equalToSuperview()
        }
        statusLabel.snp.makeConstraints { make in
            make.trailing.equalTo(uploadButton.snp.leading).offset(-8)
            make.centerY.equalToSuperview()
        }
        return container
    }

    private func setKycStatusLabel(_ label: UILabel, status: String) {
        switch status {
        case "uploaded", "verified":
            label.text = status == "verified" ? "Verified" : "Uploaded"
            label.textColor = UIColor(hex: "#4CAF50")
        case "rejected":
            label.text = "Rejected"
            label.textColor = UIColor(hex: "#FF5252")
        case "pending", "reviewing":
            label.text = "Reviewing"
            label.textColor = UIColor(hex: "#FFC107")
        default:
            label.text = "Required"
            label.textColor = .white.withAlphaComponent(0.5)
        }
    }

    private func updateLevelProgress() {
        let level = Float(viewModel.state.currentLevel)
        let max = Float(viewModel.state.maxLevel)
        levelProgressView.setProgress(max > 0 ? level / max : 0, animated: true)
        for i in 0...3 {
            if let label = levelLabelsStack.viewWithTag(100 + i) as? UILabel {
                label.textColor = i <= viewModel.state.currentLevel
                    ? UIColor(hex: "#FF6B35") : .white.withAlphaComponent(0.5)
            }
        }
    }

    private func presentCamera() {
        let picker = UIImagePickerController()
        picker.delegate = self
        picker.sourceType = .camera
        picker.allowsEditing = true
        present(picker, animated: true)
    }

    private func setupBindings() {
        viewModel.onStateChange = { [weak self] state in
            DispatchQueue.main.async {
                self?.updateLevelProgress()
                self?.buildDocumentRows(state.documents)
                self?.submitButton.isEnabled = state.canSubmit
                self?.submitButton.backgroundColor = state.canSubmit
                    ? UIColor(hex: "#FF6B35") : UIColor(hex: "#FF6B35").withAlphaComponent(0.5)
                if let error = state.error {
                    self?.errorLabel.text = error
                    self?.errorLabel.isHidden = false
                } else {
                    self?.errorLabel.isHidden = true
                }
                if state.submitSuccess {
                    let alert = UIAlertController(title: "Submitted", message: "Documents submitted for verification.", preferredStyle: .alert)
                    alert.addAction(UIAlertAction(title: "OK", style: .default))
                    self?.present(alert, animated: true)
                }
            }
        }
    }

    @objc private func submitTapped() {
        viewModel.submitVerification()
    }

    func imagePickerController(_ picker: UIImagePickerController, didFinishPickingMediaWithInfo info: [UIImagePickerController.InfoKey: Any]) {
        picker.dismiss(animated: true)
        guard let type = selectedDocumentType,
              let image = info[.editedImage] as? UIImage,
              let data = image.jpegData(compressionQuality: 0.8) else { return }
        viewModel.uploadDocument(type: type, fileName: "\(type)_\(Int(Date().timeIntervalSince1970)).jpg", data: data)
    }

    func imagePickerControllerDidCancel(_ picker: UIImagePickerController) {
        picker.dismiss(animated: true)
    }
}
