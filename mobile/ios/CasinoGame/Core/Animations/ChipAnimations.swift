import UIKit

// MARK: - ChipView

final class ChipView: UIView {
    let label = UILabel()
    var denomination: Int = 0 {
        didSet { label.text = "\(denomination)" }
    }

    private let borderLayer = CAShapeLayer()

    override init(frame: CGRect) {
        let size = min(frame.width, frame.height)
        var f = frame
        f.size = CGSize(width: size, height: size)
        super.init(frame: f)

        layer.cornerRadius = size / 2
        backgroundColor = UIColor(red: 0.1, green: 0.4, blue: 0.8, alpha: 1)

        borderLayer.strokeColor = UIColor.white.cgColor
        borderLayer.lineWidth = 2
        borderLayer.lineDashPattern = [6, 4]
        borderLayer.fillColor = nil
        layer.addSublayer(borderLayer)

        label.textAlignment = .center
        label.font = UIFont.boldSystemFont(ofSize: size * 0.3)
        label.textColor = .white
        label.frame = bounds
        label.autoresizingMask = [.flexibleWidth, .flexibleHeight]
        addSubview(label)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    override func layoutSubviews() {
        super.layoutSubviews()
        let inset: CGFloat = 3
        borderLayer.path = UIBezierPath(ovalIn: bounds.insetBy(dx: inset, dy: inset)).cgPath
    }
}

// MARK: - ChipBetView

final class ChipBetView: UIView {
    private var chips: [ChipView] = []

    func placeBet(denominations: [Int], from origin: CGPoint, to target: CGPoint, completion: (() -> Void)? = nil) {
        chips.forEach { $0.removeFromSuperview() }
        chips = []

        for (i, value) in denominations.enumerated() {
            let chip = ChipView(frame: CGRect(x: 0, y: 0, width: 44, height: 44))
            chip.denomination = value
            chip.center = origin
            chip.alpha = 0
            addSubview(chip)
            chips.append(chip)

            let delay = Double(i) * 0.1
            UIView.animate(withDuration: 0.4, delay: delay, usingSpringWithDamping: 0.7, initialSpringVelocity: 0.8, options: [], animations: {
                chip.center = CGPoint(x: target.x + CGFloat(i) * 2, y: target.y - CGFloat(i) * 3)
                chip.alpha = 1
            })
        }

        let totalTime = Double(denominations.count) * 0.1 + 0.5
        DispatchQueue.main.asyncAfter(deadline: .now() + totalTime) { completion?() }
    }
}

// MARK: - ChipCollectView

final class ChipCollectView: UIView {
    func collect(chips: [ChipView], to destination: CGPoint, completion: (() -> Void)? = nil) {
        for (i, chip) in chips.enumerated() {
            let start = chip.center
            let mid = CGPoint(
                x: (start.x + destination.x) / 2,
                y: min(start.y, destination.y) - 80
            )

            let path = UIBezierPath()
            path.move(to: start)
            path.addQuadCurve(to: destination, controlPoint: mid)

            let animation = CAKeyframeAnimation(keyPath: "position")
            animation.path = path.cgPath
            animation.duration = 0.5
            animation.beginTime = CACurrentMediaTime() + Double(i) * 0.05
            animation.timingFunction = CAMediaTimingFunction(name: .easeIn)
            animation.fillMode = .forwards
            animation.isRemovedOnCompletion = false

            let shrink = CABasicAnimation(keyPath: "transform.scale")
            shrink.fromValue = 1
            shrink.toValue = 0.3
            shrink.duration = animation.duration
            shrink.beginTime = animation.beginTime
            shrink.fillMode = .forwards
            shrink.isRemovedOnCompletion = false

            chip.layer.add(animation, forKey: "flyPath")
            chip.layer.add(shrink, forKey: "shrink")
        }

        let totalTime = Double(chips.count) * 0.05 + 0.6
        DispatchQueue.main.asyncAfter(deadline: .now() + totalTime) {
            chips.forEach { $0.removeFromSuperview() }
            completion?()
        }
    }
}

// MARK: - ChipStackView

final class ChipStackView: UIView {
    private var stack: [ChipView] = []

    func buildStack(denominations: [Int], at position: CGPoint, completion: (() -> Void)? = nil) {
        stack.forEach { $0.removeFromSuperview() }
        stack = []

        for (i, value) in denominations.enumerated() {
            let chip = ChipView(frame: CGRect(x: 0, y: 0, width: 44, height: 44))
            chip.denomination = value
            let targetY = position.y - CGFloat(i) * 4
            chip.center = CGPoint(x: position.x, y: targetY - 50)
            chip.alpha = 0
            addSubview(chip)
            stack.append(chip)

            UIView.animate(
                withDuration: 0.35,
                delay: Double(i) * 0.08,
                usingSpringWithDamping: 0.5,
                initialSpringVelocity: 1.0,
                options: [],
                animations: {
                    chip.center = CGPoint(x: position.x, y: targetY)
                    chip.alpha = 1
                }
            )
        }

        let totalTime = Double(denominations.count) * 0.08 + 0.4
        DispatchQueue.main.asyncAfter(deadline: .now() + totalTime) { completion?() }
    }

    func clear() {
        stack.forEach { $0.removeFromSuperview() }
        stack = []
    }
}
