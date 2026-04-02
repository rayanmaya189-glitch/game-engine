import UIKit

// MARK: - WinView

final class WinView: UIView {
    private let glowLayer = CAGradientLayer()
    private let amountLabel = UILabel()
    private var confettiLayer: CALayer?

    override init(frame: CGRect) {
        super.init(frame: frame)
        isUserInteractionEnabled = false

        glowLayer.colors = [
            UIColor(red: 1, green: 0.84, blue: 0, alpha: 0.8).cgColor,
            UIColor(red: 1, green: 0.65, blue: 0, alpha: 0.3).cgColor,
            UIColor.clear.cgColor
        ]
        glowLayer.locations = [0, 0.5, 1]
        glowLayer.opacity = 0
        layer.addSublayer(glowLayer)

        amountLabel.textAlignment = .center
        amountLabel.font = UIFont.boldSystemFont(ofSize: 48)
        amountLabel.textColor = UIColor(red: 1, green: 0.84, blue: 0, alpha: 1)
        amountLabel.shadowColor = .black
        amountLabel.shadowOffset = CGSize(width: 2, height: 2)
        amountLabel.alpha = 0
        addSubview(amountLabel)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    override func layoutSubviews() {
        super.layoutSubviews()
        glowLayer.frame = bounds
        amountLabel.frame = CGRect(x: 0, y: bounds.midY - 30, width: bounds.width, height: 60)
    }

    func showWin(amount: Int, completion: (() -> Void)? = nil) {
        glowLayer.opacity = 0
        let glowAnim = CABasicAnimation(keyPath: "opacity")
        glowAnim.fromValue = 0
        glowAnim.toValue = 1
        glowAnim.duration = 0.6
        glowAnim.autoreverses = true
        glowAnim.repeatCount = 2

        let pulse = CABasicAnimation(keyPath: "transform.scale")
        pulse.fromValue = 1.0
        pulse.toValue = 1.15
        pulse.duration = 0.5
        pulse.autoreverses = true
        pulse.repeatCount = 2

        glowLayer.add(glowAnim, forKey: "glow")
        layer.add(pulse, forKey: "pulse")

        UIView.animate(withDuration: 0.3) { self.amountLabel.alpha = 1 }

        let steps = 20
        let stepDuration = 0.8 / Double(steps)
        for i in 0...steps {
            DispatchQueue.main.asyncAfter(deadline: .now() + stepDuration * Double(i)) {
                let current = Int(Double(amount) * Double(i) / Double(steps))
                self.amountLabel.text = "+\(current)"
            }
        }

        DispatchQueue.main.asyncAfter(deadline: .now() + 1.2) { completion?() }
    }
}

// MARK: - LoseView

final class LoseView: UIView {
    private let flashOverlay = UIView()

    override init(frame: CGRect) {
        super.init(frame: frame)
        isUserInteractionEnabled = false
        flashOverlay.backgroundColor = UIColor.red.withAlphaComponent(0.4)
        flashOverlay.frame = bounds
        flashOverlay.autoresizingMask = [.flexibleWidth, .flexibleHeight]
        flashOverlay.alpha = 0
        addSubview(flashOverlay)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func showLose(on parentView: UIView, completion: (() -> Void)? = nil) {
        parentView.addSubview(self)

        let shake = CAKeyframeAnimation(keyPath: "position.x")
        shake.values = [
            parentView.center.x,
            parentView.center.x - 12,
            parentView.center.x + 12,
            parentView.center.x - 8,
            parentView.center.x + 8,
            parentView.center.x
        ]
        shake.keyTimes = [0, 0.2, 0.4, 0.6, 0.8, 1]
        shake.duration = 0.5

        parentView.layer.add(shake, forKey: "screenShake")

        UIView.animate(withDuration: 0.15, animations: {
            self.flashOverlay.alpha = 1
        }, completion: { _ in
            UIView.animate(withDuration: 0.4, delay: 0.2, options: [], animations: {
                self.flashOverlay.alpha = 0
            }, completion: { _ in
                self.removeFromSuperview()
                completion?()
            })
        })
    }
}

// MARK: - PushView

final class PushView: UIView {
    private let indicatorLabel = UILabel()

    override init(frame: CGRect) {
        super.init(frame: frame)
        isUserInteractionEnabled = false
        indicatorLabel.text = "PUSH"
        indicatorLabel.textAlignment = .center
        indicatorLabel.font = UIFont.boldSystemFont(ofSize: 36)
        indicatorLabel.textColor = .white
        indicatorLabel.alpha = 0
        addSubview(indicatorLabel)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    override func layoutSubviews() {
        super.layoutSubviews()
        indicatorLabel.frame = bounds
    }

    func showPush(completion: (() -> Void)? = nil) {
        indicatorLabel.transform = CGAffineTransform(scaleX: 0.5, y: 0.5)
        UIView.animate(withDuration: 0.4, delay: 0, usingSpringWithDamping: 0.6, initialSpringVelocity: 0.8, options: [], animations: {
            self.indicatorLabel.alpha = 1
            self.indicatorLabel.transform = .identity
        }, completion: { _ in
            UIView.animate(withDuration: 0.3, delay: 1.0, options: [], animations: {
                self.indicatorLabel.alpha = 0
            }, completion: { _ in completion?() })
        })
    }
}

// MARK: - JackpotView

final class JackpotView: UIView {
    private let banner = UIView()
    private let titleLabel = UILabel()
    private let amountLabel = UILabel()

    override init(frame: CGRect) {
        super.init(frame: frame)
        isUserInteractionEnabled = false

        banner.backgroundColor = UIColor(red: 0.6, green: 0, blue: 0, alpha: 0.95)
        banner.layer.cornerRadius = 12

        titleLabel.text = "JACKPOT!"
        titleLabel.textAlignment = .center
        titleLabel.font = UIFont.boldSystemFont(ofSize: 40)
        titleLabel.textColor = UIColor(red: 1, green: 0.84, blue: 0, alpha: 1)

        amountLabel.textAlignment = .center
        amountLabel.font = UIFont.boldSystemFont(ofSize: 32)
        amountLabel.textColor = .white

        banner.addSubview(titleLabel)
        banner.addSubview(amountLabel)
        addSubview(banner)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    override func layoutSubviews() {
        super.layoutSubviews()
        banner.frame = CGRect(x: 20, y: bounds.midY - 60, width: bounds.width - 40, height: 120)
        titleLabel.frame = CGRect(x: 0, y: 10, width: banner.bounds.width, height: 50)
        amountLabel.frame = CGRect(x: 0, y: 65, width: banner.bounds.width, height: 40)
    }

    func showJackpot(amount: Int, completion: (() -> Void)? = nil) {
        banner.transform = CGAffineTransform(translationX: 0, y: -200)
        banner.alpha = 0
        amountLabel.text = "$\(amount)"

        UIView.animate(withDuration: 0.6, delay: 0, usingSpringWithDamping: 0.7, initialSpringVelocity: 1.0, options: [], animations: {
            self.banner.transform = .identity
            self.banner.alpha = 1
        })

        if let emitter = GoldEmitter.create(in: bounds) {
            layer.addSublayer(emitter)
            DispatchQueue.main.asyncAfter(deadline: .now() + 3.0) {
                emitter.removeFromSuperlayer()
                completion?()
            }
        } else {
            DispatchQueue.main.asyncAfter(deadline: .now() + 3.0) { completion?() }
        }
    }
}
