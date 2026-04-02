import UIKit

// MARK: - SlotSymbolView

final class SlotSymbolView: UIView {
    private let label = UILabel()
    var symbol: String = "" {
        didSet { label.text = symbol }
    }

    override init(frame: CGRect) {
        super.init(frame: frame)
        label.font = UIFont.systemFont(ofSize: 44)
        label.textAlignment = .center
        label.frame = bounds
        label.autoresizingMask = [.flexibleWidth, .flexibleHeight]
        addSubview(label)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func animateScale(to scale: CGFloat, duration: CFTimeInterval = 0.2) {
        UIView.animate(withDuration: duration, delay: 0, usingSpringWithDamping: 0.6, initialSpringVelocity: 1, options: [], animations: {
            self.label.transform = CGAffineTransform(scaleX: scale, y: scale)
        })
    }

    func animatePosition(to offset: CGFloat, duration: CFTimeInterval = 0.15) {
        UIView.animate(withDuration: duration, animations: {
            self.label.transform = CGAffineTransform(translationX: 0, y: offset)
        }, completion: { _ in
            UIView.animate(withDuration: 0.15) {
                self.label.transform = .identity
            }
        })
    }

    func reset() {
        label.transform = .identity
        label.alpha = 1
    }
}

// MARK: - SlotReelView

final class SlotReelView: UIView, UITableViewDataSource, UITableViewDelegate {
    private let tableView = UITableView()
    private var symbols: [String] = []
    private let allSymbols = ["🍒", "🍋", "🍊", "🍇", "⭐", "7️⃣", "💎", "🔔"]
    private let symbolHeight: CGFloat = 80
    private let repetitions = 10
    private let cellID = "SymbolCell"

    override init(frame: CGRect) {
        super.init(frame: frame)
        setupSymbols()
        setupTableView()
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    private func setupSymbols() {
        symbols = []
        for _ in 0..<repetitions {
            symbols.append(contentsOf: allSymbols)
        }
    }

    private func setupTableView() {
        tableView.frame = bounds
        tableView.dataSource = self
        tableView.delegate = self
        tableView.isScrollEnabled = false
        tableView.showsVerticalScrollIndicator = false
        tableView.separatorStyle = .none
        tableView.register(UITableViewCell.self, forCellReuseIdentifier: cellID)
        tableView.autoresizingMask = [.flexibleWidth, .flexibleHeight]
        addSubview(tableView)
    }

    func spin(toSymbol symbol: String, duration: CFTimeInterval = 2.0, completion: (() -> Void)?) {
        guard let baseIndex = allSymbols.firstIndex(of: symbol) else { return }
        let targetIndex = (symbols.count / 2) + baseIndex
        let targetOffset = CGFloat(targetIndex) * symbolHeight

        tableView.contentOffset = CGPoint(x: 0, y: 0)

        let animation = CABasicAnimation(keyPath: "bounds.origin.y")
        animation.fromValue = 0
        animation.toValue = targetOffset
        animation.duration = duration
        animation.timingFunction = CAMediaTimingFunction(controlPoints: 0.1, 0.9, 0.25, 1.0)
        animation.fillMode = .forwards
        animation.isRemovedOnCompletion = false

        CATransaction.begin()
        CATransaction.setCompletionBlock { [weak self] in
            self?.tableView.contentOffset = CGPoint(x: 0, y: targetOffset)
            completion?()
        }
        tableView.layer.add(animation, forKey: "reelScroll")
        CATransaction.commit()
    }

    func currentSymbol() -> String {
        let index = Int(round(tableView.contentOffset.y / symbolHeight)) % allSymbols.count
        return allSymbols[index >= 0 ? index : index + allSymbols.count]
    }

    // MARK: UITableViewDataSource

    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return symbols.count
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: cellID, for: indexPath)
        cell.textLabel?.text = symbols[indexPath.row]
        cell.textLabel?.font = UIFont.systemFont(ofSize: 44)
        cell.textLabel?.textAlignment = .center
        cell.selectionStyle = .none
        cell.backgroundColor = .clear
        return cell
    }

    func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
        return symbolHeight
    }
}

// MARK: - SlotPaylineView

final class SlotPaylineView: UIView {
    private let lineLayer = CAShapeLayer()

    override init(frame: CGRect) {
        super.init(frame: frame)
        isUserInteractionEnabled = false
        lineLayer.strokeColor = UIColor.systemYellow.cgColor
        lineLayer.lineWidth = 4
        lineLayer.lineCap = .round
        lineLayer.fillColor = nil
        lineLayer.shadowColor = UIColor.systemYellow.cgColor
        lineLayer.shadowOpacity = 0.8
        lineLayer.shadowRadius = 6
        lineLayer.opacity = 0
        layer.addSublayer(lineLayer)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func drawLine(through points: [CGPoint], animated: Bool, duration: CFTimeInterval = 2.0) {
        guard points.count >= 2 else { return }
        let path = UIBezierPath()
        path.move(to: points[0])
        for i in 1..<points.count {
            path.addLine(to: points[i])
        }
        lineLayer.path = path.cgPath

        if animated {
            let draw = CABasicAnimation(keyPath: "strokeEnd")
            draw.fromValue = 0
            draw.toValue = 1
            draw.duration = 0.4

            let glow = CABasicAnimation(keyPath: "shadowRadius")
            glow.fromValue = 2
            glow.toValue = 10
            glow.duration = 0.4
            glow.autoreverses = true
            glow.repeatCount = 3
            glow.beginTime = 0.4

            let fade = CABasicAnimation(keyPath: "opacity")
            fade.fromValue = 1
            fade.toValue = 0
            fade.beginTime = duration - 0.3
            fade.duration = 0.3
            fade.fillMode = .forwards
            fade.isRemovedOnCompletion = false

            let group = CAAnimationGroup()
            group.animations = [draw, glow, fade]
            group.duration = duration
            group.fillMode = .forwards
            group.isRemovedOnCompletion = false

            lineLayer.add(group, forKey: "payline")
        } else {
            lineLayer.opacity = 1
        }
    }

    func clear() {
        lineLayer.removeAllAnimations()
        lineLayer.opacity = 0
    }
}

// MARK: - SlotMachineFrameView

final class SlotMachineFrameView: UIView {
    private var lightBulbs: [CALayer] = []
    private let frameLayer = CAShapeLayer()
    private let frameColor = UIColor(red: 0.6, green: 0.4, blue: 0.1, alpha: 1)

    override init(frame: CGRect) {
        super.init(frame: frame)
        setupFrame()
        setupLights()
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    private func setupFrame() {
        frameLayer.path = UIBezierPath(roundedRect: bounds.insetBy(dx: 4, dy: 4), cornerRadius: 16).cgPath
        frameLayer.fillColor = UIColor.clear.cgColor
        frameLayer.strokeColor = frameColor.cgColor
        frameLayer.lineWidth = 6
        layer.addSublayer(frameLayer)

        let innerFrame = CAShapeLayer()
        innerFrame.path = UIBezierPath(roundedRect: bounds.insetBy(dx: 10, dy: 10), cornerRadius: 12).cgPath
        innerFrame.fillColor = nil
        innerFrame.strokeColor = UIColor(red: 0.8, green: 0.6, blue: 0.2, alpha: 1).cgColor
        innerFrame.lineWidth = 2
        layer.addSublayer(innerFrame)
    }

    private func setupLights() {
        let bulbSize: CGFloat = 10
        let spacing: CGFloat = 30
        let inset: CGFloat = 16

        func addBulb(at point: CGPoint) {
            let bulb = CALayer()
            bulb.frame = CGRect(x: point.x - bulbSize / 2, y: point.y - bulbSize / 2, width: bulbSize, height: bulbSize)
            bulb.cornerRadius = bulbSize / 2
            bulb.backgroundColor = UIColor.systemYellow.cgColor
            bulb.shadowColor = UIColor.systemYellow.cgColor
            bulb.shadowOpacity = 0.8
            bulb.shadowRadius = 4
            layer.addSublayer(bulb)
            lightBulbs.append(bulb)
        }

        var x = inset + spacing
        while x < bounds.width - inset {
            addBulb(at: CGPoint(x: x, y: inset))
            addBulb(at: CGPoint(x: x, y: bounds.height - inset))
            x += spacing
        }

        var y = inset + spacing
        while y < bounds.height - inset {
            addBulb(at: CGPoint(x: inset, y: y))
            addBulb(at: CGPoint(x: bounds.width - inset, y: y))
            y += spacing
        }
    }

    func startLightAnimation() {
        for (i, bulb) in lightBulbs.enumerated() {
            let anim = CABasicAnimation(keyPath: "opacity")
            anim.fromValue = 1.0
            anim.toValue = 0.3
            anim.duration = 0.5
            anim.autoreverses = true
            anim.repeatCount = .infinity
            anim.beginTime = CACurrentMediaTime() + Double(i) * 0.08
            bulb.add(anim, forKey: "blink")
        }
    }

    func stopLightAnimation() {
        for bulb in lightBulbs {
            bulb.removeAllAnimations()
            bulb.opacity = 1
        }
    }
}
