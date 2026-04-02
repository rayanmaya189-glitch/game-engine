import UIKit

// MARK: - CardView

final class CardView: UIView {
    private let frontLayer = CALayer()
    private let backLayer = CALayer()

    var isFaceUp: Bool = false { didSet { updateFacing() } }

    init(frontImage: UIImage?, backImage: UIImage?, size: CGSize = CGSize(width: 70, height: 100)) {
        super.init(frame: CGRect(origin: .zero, size: size))
        layer.cornerRadius = 6
        layer.masksToBounds = true
        backLayer.contents = backImage?.cgImage
        backLayer.frame = bounds
        backLayer.contentsGravity = .resizeAspectFill
        layer.addSublayer(backLayer)

        frontLayer.contents = frontImage?.cgImage
        frontLayer.frame = bounds
        frontLayer.contentsGravity = .resizeAspectFill
        frontLayer.isHidden = true
        layer.addSublayer(frontLayer)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    private func updateFacing() {
        frontLayer.isHidden = !isFaceUp
        backLayer.isHidden = isFaceUp
    }
}

// MARK: - CardShuffleView

final class CardShuffleView: UIView {
    private var cards: [CardView] = []

    func configure(with cardImages: [(front: UIImage?, back: UIImage?)]) {
        cards.forEach { $0.removeFromSuperview() }
        cards = cardImages.map { CardView(frontImage: $0.front, backImage: $0.back) }
        cards.forEach { addSubview($0) }
    }

    func performShuffle(completion: (() -> Void)? = nil) {
        let count = cards.count
        guard count > 1 else { completion?(); return }
        let mid = bounds.midX
        let fanAngle: CGFloat = .pi / CGFloat(count)

        for (i, card) in cards.enumerated() {
            let angle = fanAngle * CGFloat(i) - fanAngle * CGFloat(count) / 2
            let x = mid + sin(angle) * 120
            let y = bounds.midY - cos(angle) * 120
            card.center = CGPoint(x: mid, y: bounds.midY)

            let fanOut = CAKeyframeAnimation(keyPath: "position")
            fanOut.values = [
                CGPoint(x: mid, y: bounds.midY),
                CGPoint(x: x, y: y)
            ].map { NSValue(cgPoint: $0) }
            fanOut.keyTimes = [0, 1]
            fanOut.duration = 0.5
            fanOut.beginTime = CACurrentMediaTime() + Double(i) * 0.05
            fanOut.timingFunction = CAMediaTimingFunction(name: .easeOut)
            fanOut.fillMode = .forwards
            fanOut.isRemovedOnCompletion = false

            let rotate = CABasicAnimation(keyPath: "transform.rotation.z")
            rotate.fromValue = 0
            rotate.toValue = angle
            rotate.duration = 0.5
            rotate.beginTime = fanOut.beginTime
            rotate.timingFunction = CAMediaTimingFunction(name: .easeOut)
            rotate.fillMode = .forwards
            rotate.isRemovedOnCompletion = false

            card.layer.add(fanOut, forKey: "fanOut")
            card.layer.add(rotate, forKey: "rotate")
        }

        let collapseDelay = 0.7
        for (i, card) in cards.enumerated() {
            let collapse = CABasicAnimation(keyPath: "position")
            collapse.toValue = NSValue(cgPoint: CGPoint(x: mid, y: bounds.midY))
            collapse.duration = 0.4
            collapse.beginTime = CACurrentMediaTime() + collapseDelay + Double(i) * 0.03
            collapse.timingFunction = CAMediaTimingFunction(name: .easeIn)
            collapse.fillMode = .forwards
            collapse.isRemovedOnCompletion = false

            let resetRotation = CABasicAnimation(keyPath: "transform.rotation.z")
            resetRotation.toValue = 0
            resetRotation.duration = 0.4
            resetRotation.beginTime = collapse.beginTime

            card.layer.add(collapse, forKey: "collapse")
            card.layer.add(resetRotation, forKey: "resetRotation")
        }

        DispatchQueue.main.asyncAfter(deadline: .now() + collapseDelay + 0.6) {
            self.cards.forEach { card in
                card.layer.removeAllAnimations()
                card.center = CGPoint(x: mid, y: self.bounds.midY)
                card.transform = .identity
            }
            completion?()
        }
    }
}

// MARK: - CardDealView

final class CardDealView: UIView {
    private let card: CardView

    init(card: CardView) {
        self.card = card
        super.init(frame: .zero)
        addSubview(card)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func deal(from origin: CGPoint, to destination: CGPoint, delay: TimeInterval = 0, completion: (() -> Void)? = nil) {
        card.center = origin
        card.alpha = 0
        card.layer.transform = CATransform3DMakeRotation(.pi / 2, 0, 1, 0)

        UIView.animate(
            withDuration: 0.5,
            delay: delay,
            usingSpringWithDamping: 0.8,
            initialSpringVelocity: 0.5,
            options: .curveEaseOut,
            animations: {
                self.card.center = destination
                self.card.alpha = 1
                self.card.layer.transform = CATransform3DIdentity
            },
            completion: { _ in completion?() }
        )
    }
}

// MARK: - CardFlipView

final class CardFlipView: UIView {
    private let card: CardView

    init(card: CardView) {
        self.card = card
        super.init(frame: card.bounds)
        addSubview(card)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func flipToReveal(completion: (() -> Void)? = nil) {
        let flip = CABasicAnimation(keyPath: "transform.rotation.y")
        flip.fromValue = 0
        flip.toValue = Double.pi
        flip.duration = 0.5
        flip.timingFunction = CAMediaTimingFunction(name: .easeInEaseOut)

        CATransaction.begin()
        CATransaction.setCompletionQueue(.main)
        CATransaction.setCompletionBlock {
            self.card.isFaceUp = true
            completion?()
        }
        card.layer.add(flip, forKey: "flip")
        CATransaction.commit()
    }
}

// MARK: - CardFanView

final class CardFanView: UIView {
    private var cards: [CardView] = []

    func configure(with cardImages: [(front: UIImage?, back: UIImage?)]) {
        cards.forEach { $0.removeFromSuperview() }
        cards = cardImages.map { CardView(frontImage: $0.front, backImage: $0.back) }
        cards.forEach { addSubview($0) }
    }

    func fanOut(completion: (() -> Void)? = nil) {
        let count = cards.count
        guard count > 0 else { completion?(); return }
        let spreadAngle: CGFloat = .pi / 3.0
        let radius: CGFloat = 140
        let startAngle = -spreadAngle / 2

        for (i, card) in cards.enumerated() {
            let fraction = count > 1 ? CGFloat(i) / CGFloat(count - 1) : 0.5
            let angle = startAngle + spreadAngle * fraction
            let x = bounds.midX + sin(angle) * radius
            let y = bounds.midY - cos(angle) * radius

            card.center = CGPoint(x: bounds.midX, y: bounds.midY)

            let path = CAKeyframeAnimation(keyPath: "position")
            let midX = bounds.midX + sin(angle) * radius * 0.5
            let midY = bounds.midY - cos(angle) * radius * 0.5
            path.values = [
                NSValue(cgPoint: card.center),
                NSValue(cgPoint: CGPoint(x: midX, y: midY)),
                NSValue(cgPoint: CGPoint(x: x, y: y))
            ]
            path.keyTimes = [0, 0.6, 1]
            path.duration = 0.6
            path.beginTime = CACurrentMediaTime() + Double(i) * 0.08
            path.timingFunction = CAMediaTimingFunction(name: .easeOut)
            path.fillMode = .forwards
            path.isRemovedOnCompletion = false

            let rotation = CABasicAnimation(keyPath: "transform.rotation.z")
            rotation.fromValue = 0
            rotation.toValue = angle
            rotation.duration = 0.6
            rotation.beginTime = path.beginTime
            rotation.fillMode = .forwards
            rotation.isRemovedOnCompletion = false

            card.layer.add(path, forKey: "fanPath")
            card.layer.add(rotation, forKey: "fanRotation")
        }

        let totalTime = Double(count) * 0.08 + 0.7
        DispatchQueue.main.asyncAfter(deadline: .now() + totalTime) {
            completion?()
        }
    }
}
