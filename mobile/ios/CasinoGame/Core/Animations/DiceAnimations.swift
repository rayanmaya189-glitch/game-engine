import UIKit

// MARK: - DiceFaceView

final class DiceFaceView: UIView {
    var value: Int = 1 { didSet { setNeedsDisplay() } }

    override func draw(_ rect: CGRect) {
        guard let ctx = UIGraphicsGetCurrentContext() else { return }
        UIColor.white.setFill()
        UIBezierPath(roundedRect: rect.insetBy(dx: 2, dy: 2), cornerRadius: 8).fill()

        UIColor.black.setFill()
        let dotSize: CGFloat = min(rect.width, rect.height) * 0.14
        let positions = dotPositions(in: rect)
        for point in positions {
            let dot = CGRect(x: point.x - dotSize / 2, y: point.y - dotSize / 2, width: dotSize, height: dotSize)
            UIBezierPath(ovalIn: dot).fill()
        }
    }

    private func dotPositions(in rect: CGRect) -> [CGPoint] {
        let inset = rect.width * 0.28
        let tl = CGPoint(x: inset, y: inset)
        let tr = CGPoint(x: rect.width - inset, y: inset)
        let ml = CGPoint(x: rect.midX, y: rect.midY)
        let bl = CGPoint(x: inset, y: rect.height - inset)
        let br = CGPoint(x: rect.width - inset, y: rect.height - inset)

        switch value {
        case 1: return [ml]
        case 2: return [tl, br]
        case 3: return [tl, ml, br]
        case 4: return [tl, tr, bl, br]
        case 5: return [tl, tr, ml, bl, br]
        case 6: return [tl, tr, CGPoint(x: inset, y: rect.midY), CGPoint(x: rect.width - inset, y: rect.midY), bl, br]
        default: return [ml]
        }
    }
}

// MARK: - DiceRollView

final class DiceRollView: UIView {
    private let dieA = DiceFaceView()
    private let dieB = DiceFaceView()

    override init(frame: CGRect) {
        super.init(frame: frame)
        [dieA, dieB].forEach { v in
            v.frame = CGRect(x: 0, y: 0, width: 50, height: 50)
            addSubview(v)
        }
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func roll(toA a: Int, toB b: Int, completion: (() -> Void)? = nil) {
        let offset: CGFloat = 70
        dieA.center = CGPoint(x: bounds.midX - offset / 2, y: bounds.midY)
        dieB.center = CGPoint(x: bounds.midX + offset / 2, y: bounds.midY)

        let duration: TimeInterval = 0.8
        let steps = 8
        let stepDuration = duration / Double(steps)

        func animateStep(_ step: Int) {
            guard step < steps else {
                self.dieA.value = a
                self.dieB.value = b
                self.dieA.layer.transform = CATransform3DIdentity
                self.dieB.layer.transform = CATransform3DIdentity
                completion?()
                return
            }

            let randA = CGFloat.random(in: -.pi ... .pi)
            let randB = CGFloat.random(in: -.pi ... .pi)

            UIView.animate(withDuration: stepDuration, delay: 0, options: .curveLinear, animations: {
                self.dieA.layer.transform = CATransform3DMakeRotation(randA, 1, 1, 0)
                self.dieB.layer.transform = CATransform3DMakeRotation(randB, 0, 1, 1)
                self.dieA.value = Int.random(in: 1...6)
                self.dieB.value = Int.random(in: 1...6)
            }, completion: { _ in
                animateStep(step + 1)
            })
        }

        animateStep(0)
    }
}

// MARK: - DiceShakeView

final class DiceShakeView: UIView {
    private let diceContainer = UIView()

    override init(frame: CGRect) {
        super.init(frame: frame)
        diceContainer.frame = bounds
        diceContainer.autoresizingMask = [.flexibleWidth, .flexibleHeight]
        addSubview(diceContainer)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func shake(duration: TimeInterval = 1.0, completion: (() -> Void)? = nil) {
        let shake = CAKeyframeAnimation(keyPath: "position.x")
        shake.values = [center.x, center.x - 15, center.x + 15, center.x - 10, center.x + 10, center.x]
        shake.keyTimes = [0, 0.2, 0.4, 0.6, 0.8, 1]
        shake.duration = duration

        let yShake = CAKeyframeAnimation(keyPath: "position.y")
        yShake.values = [center.y, center.y - 10, center.y + 8, center.y - 5, center.y + 5, center.y]
        yShake.keyTimes = shake.keyTimes
        yShake.duration = duration

        let group = CAAnimationGroup()
        group.animations = [shake, yShake]
        group.duration = duration
        group.timingFunction = CAMediaTimingFunction(name: .easeInEaseOut)

        CATransaction.begin()
        CATransaction.setCompletionQueue(.main)
        CATransaction.setCompletionBlock { completion?() }
        layer.add(group, forKey: "shake")
        CATransaction.commit()
    }
}

// MARK: - DiceRevealView

final class DiceRevealView: UIView {
    private let dieA = DiceFaceView()
    private let dieB = DiceFaceView()

    override init(frame: CGRect) {
        super.init(frame: frame)
        [dieA, dieB].forEach { addSubview($0) }
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func reveal(a: Int, b: Int, completion: (() -> Void)? = nil) {
        dieA.value = a
        dieB.value = b
        let offset: CGFloat = 60
        dieA.frame = CGRect(x: bounds.midX - offset - 25, y: -60, width: 50, height: 50)
        dieB.frame = CGRect(x: bounds.midX + offset - 25, y: -60, width: 50, height: 50)

        UIView.animate(
            withDuration: 0.6,
            delay: 0,
            usingSpringWithDamping: 0.5,
            initialSpringVelocity: 1.2,
            options: .curveEaseOut,
            animations: {
                self.dieA.center = CGPoint(x: self.bounds.midX - offset, y: self.bounds.midY)
                self.dieB.center = CGPoint(x: self.bounds.midX + offset, y: self.bounds.midY)
            },
            completion: { _ in completion?() }
        )
    }
}
