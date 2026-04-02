import UIKit

// MARK: - RouletteWheelView

final class RouletteWheelView: UIView {
    private let segmentCount = 37
    private let numbers = [0, 32, 15, 19, 4, 21, 2, 25, 17, 34, 6, 27, 13, 36, 11, 30, 8, 23, 10,
                           5, 24, 16, 33, 1, 20, 14, 31, 9, 22, 18, 29, 7, 28, 12, 35, 3, 26]

    private let redNumbers: Set<Int> = [1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36]

    override init(frame: CGRect) {
        super.init(frame: frame)
        backgroundColor = .clear
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    override func draw(_ rect: CGRect) {
        guard let ctx = UIGraphicsGetCurrentContext() else { return }
        let center = CGPoint(x: rect.midX, y: rect.midY)
        let outerRadius = min(rect.width, rect.height) / 2 - 2
        let innerRadius = outerRadius * 0.65
        let angleStep = (2 * CGFloat.pi) / CGFloat(segmentCount)

        for i in 0..<segmentCount {
            let startAngle = CGFloat(i) * angleStep - CGFloat.pi / 2
            let endAngle = startAngle + angleStep

            let number = numbers[i]
            let color: UIColor = number == 0 ? .systemGreen : (redNumbers.contains(number) ? .systemRed : .black)

            ctx.setFillColor(color.cgColor)
            ctx.move(to: CGPoint(x: center.x + innerRadius * cos(startAngle), y: center.y + innerRadius * sin(startAngle)))
            ctx.addArc(center: center, radius: outerRadius, startAngle: startAngle, endAngle: endAngle, clockwise: false)
            ctx.addArc(center: center, radius: innerRadius, startAngle: endAngle, endAngle: startAngle, clockwise: true)
            ctx.closePath()
            ctx.fillPath()

            ctx.setStrokeColor(UIColor.white.withAlphaComponent(0.4).cgColor)
            ctx.setLineWidth(0.5)
            ctx.move(to: CGPoint(x: center.x + innerRadius * cos(startAngle), y: center.y + innerRadius * sin(startAngle)))
            ctx.addLine(to: CGPoint(x: center.x + outerRadius * cos(startAngle), y: center.y + outerRadius * sin(startAngle)))
            ctx.strokePath()

            let midAngle = startAngle + angleStep / 2
            let textRadius = (outerRadius + innerRadius) / 2
            let textPoint = CGPoint(x: center.x + textRadius * cos(midAngle) - 8,
                                     y: center.y + textRadius * sin(midAngle) - 7)

            let attrs: [NSAttributedString.Key: Any] = [
                .font: UIFont.boldSystemFont(ofSize: 11),
                .foregroundColor: UIColor.white
            ]
            "\(number)".draw(at: textPoint, withAttributes: attrs)
        }

        let hubRect = CGRect(x: center.x - innerRadius, y: center.y - innerRadius, width: innerRadius * 2, height: innerRadius * 2)
        ctx.setFillColor(UIColor.darkGray.cgColor)
        ctx.fillEllipse(in: hubRect)
        ctx.setStrokeColor(UIColor.lightGray.cgColor)
        ctx.setLineWidth(2)
        ctx.strokeEllipse(in: hubRect)
    }
}

// MARK: - RouletteBallView

final class RouletteBallView: UIView {
    override init(frame: CGRect) {
        let ballSize: CGFloat = 12
        super.init(frame: CGRect(x: 0, y: 0, width: ballSize, height: ballSize))
        backgroundColor = .white
        layer.cornerRadius = ballSize / 2
        layer.shadowColor = UIColor.black.cgColor
        layer.shadowOpacity = 0.6
        layer.shadowRadius = 4
        layer.shadowOffset = CGSize(width: 1, height: 1)
        layer.borderColor = UIColor.lightGray.cgColor
        layer.borderWidth = 0.5
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func animateAlongPath(_ path: UIBezierPath, duration: CFTimeInterval, completion: (() -> Void)?) {
        let anim = CAKeyframeAnimation(keyPath: "position")
        anim.path = path.cgPath
        anim.duration = duration
        anim.timingFunction = CAMediaTimingFunction(name: .easeIn)
        anim.rotationMode = .none
        anim.fillMode = .forwards
        anim.isRemovedOnCompletion = false

        CATransaction.begin()
        CATransaction.setCompletionBlock { completion?() }
        layer.add(anim, forKey: "ballPath")
        CATransaction.commit()
    }

    func moveToCenter(of parent: UIView, animated: Bool) {
        let target = CGPoint(x: parent.bounds.midX, y: parent.bounds.midY)
        if animated {
            UIView.animate(withDuration: 0.3, delay: 0, usingSpringWithDamping: 0.5, initialSpringVelocity: 2, options: [], animations: {
                self.center = target
            })
        } else {
            center = target
        }
    }
}

// MARK: - RouletteBettingBoardView

final class RouletteBettingBoardView: UIView, UICollectionViewDataSource, UICollectionViewDelegateFlowLayout {
    private var collectionView: UICollectionView!
    private let numbers = Array(0...36)
    private var selectionHandler: ((Int) -> Void)?
    private let cellID = "RouletteCell"

    private let redNumbers: Set<Int> = [1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36]

    override init(frame: CGRect) {
        super.init(frame: frame)
        let layout = UICollectionViewFlowLayout()
        layout.minimumInteritemSpacing = 2
        layout.minimumLineSpacing = 2

        collectionView = UICollectionView(frame: bounds, collectionViewLayout: layout)
        collectionView.backgroundColor = .darkGray
        collectionView.dataSource = self
        collectionView.delegate = self
        collectionView.register(UICollectionViewCell.self, forCellWithReuseIdentifier: cellID)
        collectionView.autoresizingMask = [.flexibleWidth, .flexibleHeight]
        addSubview(collectionView)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    func configure(onSelect: @escaping (Int) -> Void) {
        selectionHandler = onSelect
    }

    func numberOfSections(in collectionView: UICollectionView) -> Int { 1 }

    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return numbers.count
    }

    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let cell = collectionView.dequeueReusableCell(withReuseIdentifier: cellID, for: indexPath)
        let number = numbers[indexPath.item]
        cell.contentView.subviews.forEach { $0.removeFromSuperview() }

        let bgColor: UIColor = number == 0 ? .systemGreen : (redNumbers.contains(number) ? .systemRed : .black)
        cell.backgroundColor = bgColor
        cell.layer.cornerRadius = 4
        cell.layer.borderWidth = 0.5
        cell.layer.borderColor = UIColor.white.withAlphaComponent(0.3).cgColor

        let label = UILabel(frame: cell.contentView.bounds)
        label.text = "\(number)"
        label.textAlignment = .center
        label.textColor = .white
        label.font = UIFont.boldSystemFont(ofSize: 14)
        label.autoresizingMask = [.flexibleWidth, .flexibleHeight]
        cell.contentView.addSubview(label)

        return cell
    }

    func collectionView(_ collectionView: UICollectionView, layout: UICollectionViewLayout, sizeForItemAt indexPath: IndexPath) -> CGSize {
        let columns: CGFloat = 13
        let spacing: CGFloat = 2 * (columns - 1)
        let width = (collectionView.bounds.width - spacing) / columns
        return CGSize(width: width, height: width)
    }

    func collectionView(_ collectionView: UICollectionView, didSelectItemAt indexPath: IndexPath) {
        let number = numbers[indexPath.item]
        selectionHandler?(number)

        if let cell = collectionView.cellForItem(at: indexPath) {
            UIView.animate(withDuration: 0.1, animations: {
                cell.transform = CGAffineTransform(scaleX: 1.15, y: 1.15)
            }, completion: { _ in
                UIView.animate(withDuration: 0.1) {
                    cell.transform = .identity
                }
            })
        }
    }
}

// MARK: - RouletteChipPlacementView

final class RouletteChipPlacementView: UIView {
    private let chipLayer = CAShapeLayer()
    let chipColor: UIColor
    let betNumber: Int

    init(center: CGPoint, chipColor: UIColor, betNumber: Int) {
        self.chipColor = chipColor
        self.betNumber = betNumber
        let size: CGFloat = 28
        super.init(frame: CGRect(x: center.x - size / 2, y: center.y - size / 2, width: size, height: size))

        backgroundColor = chipColor
        layer.cornerRadius = size / 2
        layer.borderColor = UIColor.white.cgColor
        layer.borderWidth = 2

        chipLayer.strokeColor = UIColor.white.cgColor
        chipLayer.lineWidth = 1
        chipLayer.lineDashPattern = [4, 3]
        chipLayer.fillColor = nil
        layer.addSublayer(chipLayer)

        let label = UILabel(frame: bounds)
        label.text = "\(betNumber)"
        label.textAlignment = .center
        label.textColor = .white
        label.font = UIFont.boldSystemFont(ofSize: 10)
        addSubview(label)

        alpha = 0
        transform = CGAffineTransform(scaleX: 0.3, y: 0.3)
    }

    @available(*, unavailable)
    required init?(coder: NSCoder) { fatalError() }

    override func layoutSubviews() {
        super.layoutSubviews()
        chipLayer.path = UIBezierPath(ovalIn: bounds.insetBy(dx: 3, dy: 3)).cgPath
    }

    func placeAnimated(completion: (() -> Void)?) {
        UIView.animate(withDuration: 0.3, delay: 0, usingSpringWithDamping: 0.5, initialSpringVelocity: 1.5, options: [], animations: {
            self.alpha = 1
            self.transform = .identity
        }, completion: { _ in
            completion?()
        })
    }

    func removeAnimated(completion: (() -> Void)?) {
        UIView.animate(withDuration: 0.2, animations: {
            self.alpha = 0
            self.transform = CGAffineTransform(scaleX: 0.3, y: 0.3)
        }, completion: { _ in
            self.removeFromSuperview()
            completion?()
        })
    }
}
