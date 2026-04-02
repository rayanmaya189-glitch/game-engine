import UIKit

// MARK: - RouletteWheelSpinView

final class RouletteWheelSpinView: UIView {
    private let indicatorLayer = CAShapeLayer()
    private let segmentCount = 37
    private let segmentColors: [UIColor] = [.systemGreen,
        .systemRed, .black, .systemRed, .black, .systemRed, .black,
        .systemRed, .black, .systemRed, .black, .systemRed, .black,
        .systemRed, .black, .systemRed, .black, .systemRed, .black,
        .systemRed, .black, .systemRed, .black, .systemRed, .black,
        .systemRed, .black, .systemRed, .black, .systemRed, .black,
        .systemRed, .black, .systemRed, .black, .systemRed, .black]

    override init(frame: CGRect) {
        super.init(frame: frame)
        setupWheel()
        setupIndicator()
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    private func setupWheel() {
        let radius = min(bounds.width, bounds.height) / 2 - 4
        let center = CGPoint(x: bounds.midX, y: bounds.midY)
        let anglePerSegment = (2 * CGFloat.pi) / CGFloat(segmentCount)
        for i in 0..<segmentCount {
            let startAngle = CGFloat(i) * anglePerSegment - CGFloat.pi / 2
            let endAngle = startAngle + anglePerSegment
            let path = UIBezierPath()
            path.move(to: center)
            path.addArc(withCenter: center, radius: radius, startAngle: startAngle, endAngle: endAngle, clockwise: true)
            path.close()
            let segLayer = CAShapeLayer()
            segLayer.path = path.cgPath
            segLayer.fillColor = segmentColors[i].cgColor
            segLayer.strokeColor = UIColor.white.withAlphaComponent(0.3).cgColor
            segLayer.lineWidth = 0.5
            layer.addSublayer(segLayer)
        }
        let innerCircle = UIBezierPath(arcCenter: center, radius: radius * 0.35, startAngle: 0, endAngle: 2 * .pi, clockwise: true)
        let innerLayer = CAShapeLayer()
        innerLayer.path = innerCircle.cgPath
        innerLayer.fillColor = UIColor.darkGray.cgColor
        layer.addSublayer(innerLayer)
    }

    private func setupIndicator() {
        let tipY: CGFloat = 6
        let path = UIBezierPath()
        path.move(to: CGPoint(x: bounds.midX - 8, y: tipY))
        path.addLine(to: CGPoint(x: bounds.midX + 8, y: tipY))
        path.addLine(to: CGPoint(x: bounds.midX, y: tipY + 16))
        path.close()
        indicatorLayer.path = path.cgPath
        indicatorLayer.fillColor = UIColor.yellow.cgColor
        indicatorLayer.strokeColor = UIColor.white.cgColor
        indicatorLayer.lineWidth = 1
        layer.addSublayer(indicatorLayer)
    }

    func spin(toWinningNumber number: Int, duration: CFTimeInterval = 4.0, completion: (() -> Void)?) {
        let totalRotation = 5 * 2 * CGFloat.pi - CGFloat(number) * (2 * CGFloat.pi) / CGFloat(segmentCount)
        let animation = CABasicAnimation(keyPath: "transform.rotation.z")
        animation.fromValue = 0
        animation.toValue = totalRotation
        animation.duration = duration
        animation.timingFunction = CAMediaTimingFunction(controlPoints: 0.15, 0.85, 0.2, 1.0)
        animation.fillMode = .forwards
        animation.isRemovedOnCompletion = false
        CATransaction.begin()
        CATransaction.setCompletionBlock { [weak self] in
            self?.layer.transform = CATransform3DMakeRotation(totalRotation, 0, 0, 1)
            completion?()
        }
        layer.add(animation, forKey: "wheelSpin")
        CATransaction.commit()
    }
}

// MARK: - RouletteBallBounceView

final class RouletteBallBounceView: UIView {
    private let ballLayer = CAShapeLayer()
    private let trackRadius: CGFloat

    init(frame: CGRect, trackRadius: CGFloat) {
        self.trackRadius = trackRadius
        super.init(frame: frame)
        let ballSize: CGFloat = 10
        ballLayer.path = UIBezierPath(ovalIn: CGRect(x: -ballSize / 2, y: -ballSize / 2, width: ballSize, height: ballSize)).cgPath
        ballLayer.fillColor = UIColor.white.cgColor
        ballLayer.strokeColor = UIColor.lightGray.cgColor
        ballLayer.lineWidth = 1
        ballLayer.shadowColor = UIColor.black.cgColor
        ballLayer.shadowOpacity = 0.5
        ballLayer.shadowRadius = 3
        ballLayer.position = CGPoint(x: bounds.midX, y: bounds.midY)
        layer.addSublayer(ballLayer)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func animateBall(toWinningAngle angle: CGFloat, duration: CFTimeInterval = 3.5, completion: (() -> Void)?) {
        let center = CGPoint(x: bounds.midX, y: bounds.midY)
        let outerRadius = trackRadius
        let innerRadius = trackRadius * 0.55
        var values: [NSValue] = []
        var keyTimes: [NSNumber] = []
        let steps = 120
        for i in 0...steps {
            let t = CGFloat(i) / CGFloat(steps)
            let easedT = pow(t, 0.4)
            let r = outerRadius - (outerRadius - innerRadius) * easedT
            let a = 8 * 2 * .pi * (1 - t) + angle
            values.append(NSValue(cgPoint: CGPoint(x: center.x + r * cos(a), y: center.y + r * sin(a))))
            keyTimes.append(NSNumber(value: Double(t)))
        }
        let pathAnim = CAKeyframeAnimation(keyPath: "position")
        pathAnim.values = values
        pathAnim.keyTimes = keyTimes
        pathAnim.duration = duration
        pathAnim.timingFunction = CAMediaTimingFunction(name: .easeIn)
        pathAnim.fillMode = .forwards
        pathAnim.isRemovedOnCompletion = false
        CATransaction.begin()
        CATransaction.setCompletionBlock { completion?() }
        ballLayer.add(pathAnim, forKey: "ballBounce")
        CATransaction.commit()
    }
}

// MARK: - RouletteNumberRevealView

final class RouletteNumberRevealView: UIView {
    private let numberLabel = UILabel()
    private let glowLayer = CALayer()

    override init(frame: CGRect) {
        super.init(frame: frame)
        isHidden = true
        glowLayer.backgroundColor = UIColor.yellow.cgColor
        glowLayer.cornerRadius = 40
        glowLayer.shadowColor = UIColor.yellow.cgColor
        glowLayer.shadowOpacity = 0.8
        glowLayer.shadowRadius = 20
        glowLayer.frame = CGRect(x: bounds.midX - 40, y: bounds.midY - 40, width: 80, height: 80)
        layer.addSublayer(glowLayer)
        numberLabel.font = UIFont.boldSystemFont(ofSize: 32)
        numberLabel.textAlignment = .center
        numberLabel.textColor = .white
        numberLabel.frame = bounds
        numberLabel.autoresizingMask = [.flexibleWidth, .flexibleHeight]
        addSubview(numberLabel)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func reveal(number: Int, color: UIColor, completion: (() -> Void)?) {
        numberLabel.text = "\(number)"
        numberLabel.textColor = color
        glowLayer.backgroundColor = color.withAlphaComponent(0.4).cgColor
        glowLayer.shadowColor = color.cgColor
        isHidden = false
        alpha = 0
        transform = CGAffineTransform(scaleX: 0.5, y: 0.5)
        UIView.animate(withDuration: 0.5, delay: 0, usingSpringWithDamping: 0.6, initialSpringVelocity: 1.2, options: [], animations: {
            self.alpha = 1
            self.transform = .identity
        }, completion: { _ in
            let glowAnim = CABasicAnimation(keyPath: "opacity")
            glowAnim.fromValue = 0.3; glowAnim.toValue = 1.0
            glowAnim.duration = 0.8; glowAnim.autoreverses = true; glowAnim.repeatCount = 2
            let scaleAnim = CABasicAnimation(keyPath: "transform.scale")
            scaleAnim.fromValue = 1.0; scaleAnim.toValue = 1.15
            scaleAnim.duration = 0.8; scaleAnim.autoreverses = true; scaleAnim.repeatCount = 2
            let group = CAAnimationGroup()
            group.animations = [glowAnim, scaleAnim]
            group.duration = 1.6
            CATransaction.begin()
            CATransaction.setCompletionBlock { completion?() }
            self.glowLayer.add(group, forKey: "glowPulse")
            CATransaction.commit()
        })
    }
}

// MARK: - RouletteBoardHighlightView

final class RouletteBoardHighlightView: UIView {
    private let borderLayer = CAShapeLayer()

    init(frame: CGRect, color: UIColor = .systemYellow) {
        super.init(frame: frame)
        isUserInteractionEnabled = false
        borderLayer.path = UIBezierPath(roundedRect: bounds.insetBy(dx: 1, dy: 1), cornerRadius: 4).cgPath
        borderLayer.fillColor = color.withAlphaComponent(0.15).cgColor
        borderLayer.strokeColor = color.cgColor
        borderLayer.lineWidth = 2
        borderLayer.opacity = 0
        layer.addSublayer(borderLayer)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func highlight(duration: CFTimeInterval = 2.0, completion: (() -> Void)?) {
        borderLayer.path = UIBezierPath(roundedRect: bounds.insetBy(dx: 1, dy: 1), cornerRadius: 4).cgPath
        let opacityIn = CABasicAnimation(keyPath: "opacity")
        opacityIn.fromValue = 0; opacityIn.toValue = 1; opacityIn.duration = 0.3
        let strokeAnim = CABasicAnimation(keyPath: "strokeEnd")
        strokeAnim.fromValue = 0; strokeAnim.toValue = 1; strokeAnim.duration = 0.5
        let pulse = CABasicAnimation(keyPath: "lineWidth")
        pulse.fromValue = 2; pulse.toValue = 4; pulse.duration = 0.4
        pulse.autoreverses = true; pulse.repeatCount = 3; pulse.beginTime = 0.5
        let fadeOut = CABasicAnimation(keyPath: "opacity")
        fadeOut.fromValue = 1; fadeOut.toValue = 0
        fadeOut.beginTime = duration - 0.5; fadeOut.duration = 0.5
        fadeOut.fillMode = .forwards; fadeOut.isRemovedOnCompletion = false
        let group = CAAnimationGroup()
        group.animations = [opacityIn, strokeAnim, pulse, fadeOut]
        group.duration = duration
        group.fillMode = .forwards
        group.isRemovedOnCompletion = false
        CATransaction.begin()
        CATransaction.setCompletionBlock { completion?() }
        borderLayer.add(group, forKey: "highlight")
        CATransaction.commit()
    }

    func clearHighlight() {
        borderLayer.removeAllAnimations()
        borderLayer.opacity = 0
    }
}
