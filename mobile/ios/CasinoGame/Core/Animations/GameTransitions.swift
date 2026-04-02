import UIKit

// MARK: - GameSlideTransition

final class GameSlideTransition: NSObject, UIViewControllerAnimatedTransitioning {
    let isPush: Bool

    init(isPush: Bool = true) {
        self.isPush = isPush
        super.init()
    }

    func transitionDuration(using transitionContext: UIViewControllerContextTransitioning?) -> TimeInterval {
        return 0.4
    }

    func animateTransition(using transitionContext: UIViewControllerContextTransitioning) {
        guard let toView = transitionContext.view(forKey: .to),
              let fromView = transitionContext.view(forKey: .from) else {
            transitionContext.completeTransition(false)
            return
        }

        let container = transitionContext.containerView
        let width = container.bounds.width
        toView.frame = container.bounds
        container.addSubview(toView)

        let offset = isPush ? width : -width
        toView.transform = CGAffineTransform(translationX: offset, y: 0)

        UIView.animate(withDuration: transitionDuration(using: transitionContext),
                       delay: 0,
                       usingSpringWithDamping: 0.85,
                       initialSpringVelocity: 0.5,
                       options: .curveEaseInOut,
                       animations: {
            fromView.transform = CGAffineTransform(translationX: -offset, y: 0)
            toView.transform = .identity
        }, completion: { finished in
            fromView.transform = .identity
            transitionContext.completeTransition(!transitionContext.transitionWasCancelled)
        })
    }
}

// MARK: - GameFadeTransition

final class GameFadeTransition: NSObject, UIViewControllerAnimatedTransitioning {
    let duration: TimeInterval

    init(duration: TimeInterval = 0.35) {
        self.duration = duration
        super.init()
    }

    func transitionDuration(using transitionContext: UIViewControllerContextTransitioning?) -> TimeInterval {
        return duration
    }

    func animateTransition(using transitionContext: UIViewControllerContextTransitioning) {
        guard let toView = transitionContext.view(forKey: .to),
              let fromView = transitionContext.view(forKey: .from) else {
            transitionContext.completeTransition(false)
            return
        }

        let container = transitionContext.containerView
        toView.frame = container.bounds
        toView.alpha = 0
        container.addSubview(toView)

        UIView.animate(withDuration: duration, animations: {
            fromView.alpha = 0
            toView.alpha = 1
        }, completion: { finished in
            fromView.alpha = 1
            transitionContext.completeTransition(!transitionContext.transitionWasCancelled)
        })
    }
}

// MARK: - LoadingSpinnerView

final class LoadingSpinnerView: UIView {
    private let chipLayer = CAShapeLayer()
    private var isAnimating = false

    override init(frame: CGRect) {
        super.init(frame: frame)
        setupChip()
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    private func setupChip() {
        let size = min(bounds.width, bounds.height) * 0.6
        chipLayer.frame = CGRect(
            x: (bounds.width - size) / 2,
            y: (bounds.height - size) / 2,
            width: size, height: size
        )
        chipLayer.path = UIBezierPath(ovalIn: chipLayer.bounds).cgPath
        chipLayer.fillColor = UIColor(red: 0.1, green: 0.4, blue: 0.8, alpha: 1).cgColor
        chipLayer.strokeColor = UIColor.white.cgColor
        chipLayer.lineWidth = 2
        chipLayer.lineDashPattern = [4, 3]
        layer.addSublayer(chipLayer)
    }

    func startAnimating() {
        guard !isAnimating else { return }
        isAnimating = true

        let spin = CAKeyframeAnimation(keyPath: "transform.rotation.z")
        spin.values = [0, Double.pi, Double.pi * 2]
        spin.keyTimes = [0, 0.5, 1]
        spin.duration = 1.0
        spin.repeatCount = .infinity
        spin.timingFunction = CAMediaTimingFunction(name: .linear)
        chipLayer.add(spin, forKey: "spin")

        let pulse = CAKeyframeAnimation(keyPath: "transform.scale")
        pulse.values = [1.0, 1.15, 1.0]
        pulse.keyTimes = [0, 0.5, 1]
        pulse.duration = 1.0
        pulse.repeatCount = .infinity
        layer.add(pulse, forKey: "pulse")
    }

    func stopAnimating() {
        isAnimating = false
        chipLayer.removeAllAnimations()
        layer.removeAllAnimations()
    }
}
