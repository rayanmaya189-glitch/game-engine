import UIKit

// MARK: - ConfettiEmitter

enum ConfettiEmitter {
    static func create(in rect: CGRect) -> CAEmitterLayer {
        let emitter = CAEmitterLayer()
        emitter.emitterPosition = CGPoint(x: rect.midX, y: -10)
        emitter.emitterSize = CGSize(width: rect.width, height: 1)
        emitter.emitterShape = .line

        let colors: [UIColor] = [
            .red, .blue, .green, .yellow, .magenta, .orange, .cyan
        ]

        emitter.emitterCells = colors.map { color in
            let cell = CAEmitterCell()
            cell.birthRate = 8
            cell.lifetime = 4.0
            cell.velocity = 200
            cell.velocityRange = 50
            cell.emissionLongitude = .pi
            cell.emissionRange = .pi / 4
            cell.spin = 4
            cell.spinRange = 8
            cell.color = color.cgColor
            cell.scale = 0.08
            cell.scaleRange = 0.04
            cell.contents = makeRectImage()?.cgImage
            return cell
        }

        return emitter
    }

    private static func makeRectImage() -> UIImage? {
        let size = CGSize(width: 12, height: 12)
        UIGraphicsBeginImageContext(size)
        guard let ctx = UIGraphicsGetCurrentContext() else { return nil }
        ctx.setFillColor(UIColor.white.cgColor)
        ctx.fill(CGRect(origin: .zero, size: size))
        let image = UIGraphicsGetImageFromCurrentImageContext()
        UIGraphicsEndImageContext()
        return image
    }
}

// MARK: - GoldEmitter

enum GoldEmitter {
    static func create(in rect: CGRect) -> CAEmitterLayer? {
        let emitter = CAEmitterLayer()
        emitter.emitterPosition = CGPoint(x: rect.midX, y: rect.midY)
        emitter.emitterSize = CGSize(width: rect.width * 0.5, height: rect.height * 0.5)
        emitter.emitterShape = .sphere

        let cell = CAEmitterCell()
        cell.birthRate = 40
        cell.lifetime = 2.5
        cell.velocity = 120
        cell.velocityRange = 60
        cell.emissionRange = .pi * 2
        cell.spin = 3
        cell.spinRange = 6
        cell.color = UIColor(red: 1, green: 0.84, blue: 0, alpha: 1).cgColor
        cell.scale = 0.06
        cell.scaleRange = 0.03
        cell.alphaSpeed = -0.4
        cell.contents = makeCircleImage()?.cgImage

        emitter.emitterCells = [cell]
        return emitter
    }

    private static func makeCircleImage() -> UIImage? {
        let size = CGSize(width: 20, height: 20)
        UIGraphicsBeginImageContext(size)
        guard let ctx = UIGraphicsGetCurrentContext() else { return nil }
        ctx.setFillColor(UIColor.white.cgColor)
        ctx.fillEllipse(in: CGRect(origin: .zero, size: size))
        let image = UIGraphicsGetImageFromCurrentImageContext()
        UIGraphicsEndImageContext()
        return image
    }
}

// MARK: - CoinEmitter

enum CoinEmitter {
    static func create(in rect: CGRect, coinImage: UIImage?) -> CAEmitterLayer {
        let emitter = CAEmitterLayer()
        emitter.emitterPosition = CGPoint(x: rect.midX, y: rect.maxY)
        emitter.emitterSize = CGSize(width: rect.width * 0.6, height: 1)
        emitter.emitterShape = .line

        let cell = CAEmitterCell()
        cell.birthRate = 15
        cell.lifetime = 3.0
        cell.velocity = 300
        cell.velocityRange = 100
        cell.emissionLongitude = -(.pi / 2)
        cell.emissionRange = .pi / 6
        cell.spin = 6
        cell.spinRange = 10
        cell.scale = 0.12
        cell.scaleRange = 0.04
        cell.alphaSpeed = -0.3
        cell.contents = coinImage?.cgImage ?? makeCircleImage()?.cgImage

        emitter.emitterCells = [cell]
        return emitter
    }

    private static func makeCircleImage() -> UIImage? {
        let size = CGSize(width: 24, height: 24)
        UIGraphicsBeginImageContext(size)
        guard let ctx = UIGraphicsGetCurrentContext() else { return nil }
        ctx.setFillColor(UIColor(red: 1, green: 0.84, blue: 0, alpha: 1).cgColor)
        ctx.fillEllipse(in: CGRect(origin: .zero, size: size))
        let image = UIGraphicsGetImageFromCurrentImageContext()
        UIGraphicsEndImageContext()
        return image
    }
}

// MARK: - UIView Extension

extension UIView {
    func startConfetti(duration: TimeInterval = 3.0) {
        let emitter = ConfettiEmitter.create(in: bounds)
        layer.addSublayer(emitter)
        DispatchQueue.main.asyncAfter(deadline: .now() + duration) {
            emitter.birthRate = 0
            DispatchQueue.main.asyncAfter(deadline: .now() + 4) {
                emitter.removeFromSuperlayer()
            }
        }
    }

    func startGoldEffect(duration: TimeInterval = 2.0) {
        guard let emitter = GoldEmitter.create(in: bounds) else { return }
        layer.addSublayer(emitter)
        DispatchQueue.main.asyncAfter(deadline: .now() + duration) {
            emitter.birthRate = 0
            DispatchQueue.main.asyncAfter(deadline: .now() + 3) {
                emitter.removeFromSuperlayer()
            }
        }
    }

    func startCoinEffect(image: UIImage? = nil, duration: TimeInterval = 2.5) {
        let emitter = CoinEmitter.create(in: bounds, coinImage: image)
        layer.addSublayer(emitter)
        DispatchQueue.main.asyncAfter(deadline: .now() + duration) {
            emitter.birthRate = 0
            DispatchQueue.main.asyncAfter(deadline: .now() + 3) {
                emitter.removeFromSuperlayer()
            }
        }
    }
}
