import UIKit

// MARK: - SlotReelSpinView

final class SlotReelSpinView: UIView {
    private let scrollView = UIScrollView()
    private let symbols = ["🍒", "🍋", "🍊", "🍇", "⭐", "7️⃣", "💎", "🔔"]
    private let symbolHeight: CGFloat = 80
    private var totalSymbols = 0

    override init(frame: CGRect) {
        super.init(frame: frame)
        scrollView.frame = bounds
        scrollView.showsVerticalScrollIndicator = false
        scrollView.isScrollEnabled = false
        scrollView.clipsToBounds = true
        addSubview(scrollView)
        totalSymbols = symbols.count * 5
        scrollView.contentSize = CGSize(width: bounds.width, height: CGFloat(totalSymbols) * symbolHeight)
        for i in 0..<totalSymbols {
            let label = UILabel()
            label.text = symbols[i % symbols.count]
            label.font = UIFont.systemFont(ofSize: 44)
            label.textAlignment = .center
            label.frame = CGRect(x: 0, y: CGFloat(i) * symbolHeight, width: bounds.width, height: symbolHeight)
            scrollView.addSubview(label)
        }
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func spin(toSymbolIndex targetIndex: Int, duration: CFTimeInterval = 2.0, completion: (() -> Void)?) {
        let targetOffset = CGFloat(totalSymbols / 2 + targetIndex) * symbolHeight - symbolHeight
        scrollView.contentOffset = .zero
        let anim = CABasicAnimation(keyPath: "bounds.origin.y")
        anim.fromValue = 0
        anim.toValue = targetOffset
        anim.duration = duration
        anim.timingFunction = CAMediaTimingFunction(controlPoints: 0.2, 0.8, 0.3, 1.0)
        anim.fillMode = .forwards
        anim.isRemovedOnCompletion = false
        CATransaction.begin()
        CATransaction.setCompletionBlock { [weak self] in
            self?.scrollView.contentOffset = CGPoint(x: 0, y: targetOffset)
            completion?()
        }
        scrollView.layer.add(anim, forKey: "reelSpin")
        CATransaction.commit()
    }

    func currentSymbolIndex() -> Int {
        let index = Int(round(scrollView.contentOffset.y / symbolHeight)) % symbols.count
        return index >= 0 ? index : index + symbols.count
    }
}

// MARK: - SlotMachineSpinView

final class SlotMachineSpinView: UIView {
    private var reels: [SlotReelSpinView] = []
    private let reelCount: Int

    init(frame: CGRect, reelCount: Int = 3) {
        self.reelCount = reelCount
        super.init(frame: frame)
        let reelWidth = (bounds.width - CGFloat(reelCount - 1) * 4) / CGFloat(reelCount)
        for i in 0..<reelCount {
            let reel = SlotReelSpinView(frame: CGRect(x: CGFloat(i) * (reelWidth + 4), y: 0, width: reelWidth, height: bounds.height))
            reel.clipsToBounds = true
            reel.layer.cornerRadius = 6
            addSubview(reel)
            reels.append(reel)
        }
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func spinAllReels(results: [Int], baseDuration: CFTimeInterval = 2.0, staggerDelay: CFTimeInterval = 0.4, completion: (() -> Void)?) {
        var completedCount = 0
        for (i, reel) in reels.enumerated() {
            DispatchQueue.main.asyncAfter(deadline: .now() + staggerDelay * Double(i)) {
                reel.spin(toSymbolIndex: i < results.count ? results[i] : 0, duration: baseDuration + Double(i) * 0.3) {
                    completedCount += 1
                    if completedCount == self.reels.count { completion?() }
                }
            }
        }
    }
}

// MARK: - SlotSymbolLandView

final class SlotSymbolLandView: UIView {
    private let symbolLabel = UILabel()

    override init(frame: CGRect) {
        super.init(frame: frame)
        symbolLabel.font = UIFont.systemFont(ofSize: 44)
        symbolLabel.textAlignment = .center
        symbolLabel.frame = bounds
        symbolLabel.autoresizingMask = [.flexibleWidth, .flexibleHeight]
        addSubview(symbolLabel)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func land(symbol: String, completion: (() -> Void)?) {
        symbolLabel.text = symbol
        symbolLabel.transform = CGAffineTransform(translationX: 0, y: -60).scaledBy(x: 0.7, y: 0.7)
        symbolLabel.alpha = 0.5
        UIView.animate(withDuration: 0.4, delay: 0, usingSpringWithDamping: 0.45, initialSpringVelocity: 1.5, options: [], animations: {
            self.symbolLabel.transform = .identity
            self.symbolLabel.alpha = 1.0
        }, completion: { _ in
            let bounce = CASpringAnimation(keyPath: "transform.translation.y")
            bounce.fromValue = 0; bounce.toValue = -8
            bounce.damping = 3; bounce.initialVelocity = 5
            bounce.duration = bounce.settlingDuration
            CATransaction.begin()
            CATransaction.setCompletionBlock { completion?() }
            self.symbolLabel.layer.add(bounce, forKey: "symbolBounce")
            CATransaction.commit()
        })
    }
}

// MARK: - SlotPaylineWinView

final class SlotPaylineWinView: UIView {
    private let paylineLayer = CAShapeLayer()

    override init(frame: CGRect) {
        super.init(frame: frame)
        isUserInteractionEnabled = false
        paylineLayer.strokeColor = UIColor.systemYellow.cgColor
        paylineLayer.lineWidth = 3
        paylineLayer.lineCap = .round
        paylineLayer.fillColor = nil
        layer.addSublayer(paylineLayer)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func showPayline(from start: CGPoint, to end: CGPoint, duration: CFTimeInterval = 2.5, completion: (() -> Void)?) {
        let path = UIBezierPath()
        path.move(to: start); path.addLine(to: end)
        paylineLayer.path = path.cgPath
        let draw = CABasicAnimation(keyPath: "strokeEnd")
        draw.fromValue = 0; draw.toValue = 1; draw.duration = 0.3
        let glow = CABasicAnimation(keyPath: "shadowRadius")
        glow.fromValue = 2; glow.toValue = 12; glow.duration = 0.3
        glow.autoreverses = true; glow.repeatCount = 4; glow.beginTime = 0.3
        let color = CABasicAnimation(keyPath: "strokeColor")
        color.fromValue = UIColor.systemYellow.cgColor; color.toValue = UIColor.white.cgColor
        color.duration = 0.3; color.autoreverses = true; color.repeatCount = 4; color.beginTime = 0.3
        let fade = CABasicAnimation(keyPath: "opacity")
        fade.fromValue = 1; fade.toValue = 0; fade.beginTime = duration - 0.4; fade.duration = 0.4
        fade.fillMode = .forwards; fade.isRemovedOnCompletion = false
        let group = CAAnimationGroup()
        group.animations = [draw, glow, color, fade]
        group.duration = duration; group.fillMode = .forwards; group.isRemovedOnCompletion = false
        CATransaction.begin()
        CATransaction.setCompletionBlock { completion?() }
        paylineLayer.add(group, forKey: "paylineWin")
        CATransaction.commit()
    }

    func hidePayline() { paylineLayer.removeAllAnimations(); paylineLayer.opacity = 0 }
}

// MARK: - SlotBigWinView

final class SlotBigWinView: UIView {
    private let emitterLayer = CAEmitterLayer()
    private let winLabel = UILabel()

    override init(frame: CGRect) {
        super.init(frame: frame)
        isUserInteractionEnabled = false; isHidden = true
        winLabel.font = UIFont.boldSystemFont(ofSize: 48)
        winLabel.textAlignment = .center; winLabel.textColor = .systemYellow
        winLabel.shadowColor = .orange; winLabel.shadowOffset = CGSize(width: 2, height: 2)
        winLabel.frame = bounds; winLabel.autoresizingMask = [.flexibleWidth, .flexibleHeight]
        addSubview(winLabel)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func showBigWin(amount: String, completion: (() -> Void)?) {
        isHidden = false; winLabel.text = amount
        winLabel.alpha = 0; winLabel.transform = CGAffineTransform(scaleX: 3, y: 3)
        setupCoinBurst()
        let shake = CAKeyframeAnimation(keyPath: "transform.translation.x")
        shake.values = [-8, 8, -6, 6, -3, 3, 0]; shake.duration = 0.5
        shake.timingFunction = CAMediaTimingFunction(name: .easeInEaseOut)
        UIView.animate(withDuration: 0.6, delay: 0, usingSpringWithDamping: 0.5, initialSpringVelocity: 2, options: [], animations: {
            self.winLabel.alpha = 1; self.winLabel.transform = .identity
        })
        layer.add(shake, forKey: "screenShake")
        DispatchQueue.main.asyncAfter(deadline: .now() + 3.5) { [weak self] in
            self?.emitterLayer.birthRate = 0
            UIView.animate(withDuration: 0.5, animations: { self?.alpha = 0 }, completion: { _ in
                self?.isHidden = true; self?.alpha = 1; self?.emitterLayer.removeFromSuperlayer(); completion?()
            })
        }
    }

    private func setupCoinBurst() {
        emitterLayer.emitterPosition = CGPoint(x: bounds.midX, y: bounds.midY)
        emitterLayer.emitterSize = bounds.size; emitterLayer.emitterShape = .circle
        let cell = CAEmitterCell()
        cell.contents = createCoinImage()?.cgImage
        cell.birthRate = 80; cell.lifetime = 2.5; cell.velocity = 300
        cell.velocityRange = 100; cell.emissionRange = .pi * 2
        cell.yAcceleration = 200; cell.scale = 0.15; cell.scaleRange = 0.05
        cell.spin = 2; cell.spinRange = 4
        emitterLayer.emitterCells = [cell]
        layer.addSublayer(emitterLayer)
    }

    private func createCoinImage() -> UIImage? {
        let size = CGSize(width: 30, height: 30)
        UIGraphicsBeginImageContextWithOptions(size, false, 0)
        let ctx = UIGraphicsGetCurrentContext()
        ctx?.setFillColor(UIColor.systemYellow.cgColor)
        ctx?.fillEllipse(in: CGRect(origin: .zero, size: size))
        ctx?.setStrokeColor(UIColor.orange.cgColor); ctx?.setLineWidth(2)
        ctx?.strokeEllipse(in: CGRect(origin: .zero, size: size))
        let image = UIGraphicsGetImageFromCurrentImageContext()
        UIGraphicsEndImageContext()
        return image
    }
}

// MARK: - SlotFreeSpinsView

final class SlotFreeSpinsView: UIView {
    private let frameLayer = CAShapeLayer()
    private let counterLabel = UILabel()

    override init(frame: CGRect) {
        super.init(frame: frame)
        isUserInteractionEnabled = false; isHidden = true
        frameLayer.path = UIBezierPath(roundedRect: bounds.insetBy(dx: 2, dy: 2), cornerRadius: 12).cgPath
        frameLayer.fillColor = nil; frameLayer.strokeColor = UIColor.systemYellow.cgColor
        frameLayer.lineWidth = 3; layer.addSublayer(frameLayer)
        let spinsLabel = UILabel()
        spinsLabel.text = "FREE SPINS"; spinsLabel.font = UIFont.boldSystemFont(ofSize: 14)
        spinsLabel.textColor = .systemYellow; spinsLabel.textAlignment = .center
        spinsLabel.frame = CGRect(x: 0, y: 4, width: bounds.width, height: 20)
        spinsLabel.autoresizingMask = [.flexibleWidth]; addSubview(spinsLabel)
        counterLabel.font = UIFont.boldSystemFont(ofSize: 28); counterLabel.textColor = .white
        counterLabel.textAlignment = .center
        counterLabel.frame = CGRect(x: 0, y: 22, width: bounds.width, height: 36)
        counterLabel.autoresizingMask = [.flexibleWidth]; addSubview(counterLabel)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func show(remainingSpins: Int, completion: (() -> Void)?) {
        isHidden = false; counterLabel.text = "\(remainingSpins)"
        let glow = CABasicAnimation(keyPath: "strokeColor")
        glow.fromValue = UIColor.systemYellow.cgColor; glow.toValue = UIColor.white.cgColor
        glow.duration = 0.6; glow.autoreverses = true; glow.repeatCount = .infinity
        frameLayer.add(glow, forKey: "frameGlow")
        let pulse = CABasicAnimation(keyPath: "transform.scale")
        pulse.fromValue = 1.0; pulse.toValue = 1.05
        pulse.duration = 0.6; pulse.autoreverses = true; pulse.repeatCount = .infinity
        layer.add(pulse, forKey: "framePulse")
        completion?()
    }

    func updateCounter(_ count: Int) {
        counterLabel.text = "\(count)"
        UIView.animate(withDuration: 0.15, animations: {
            self.counterLabel.transform = CGAffineTransform(scaleX: 1.3, y: 1.3)
        }, completion: { _ in UIView.animate(withDuration: 0.15) { self.counterLabel.transform = .identity } })
    }

    func hide() { frameLayer.removeAllAnimations(); layer.removeAllAnimations(); isHidden = true }
}
