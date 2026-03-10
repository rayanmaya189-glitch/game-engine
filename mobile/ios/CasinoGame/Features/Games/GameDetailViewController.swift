import UIKit
import SnapKit

class GameDetailViewController: UIViewController {
    
    private let gameId: String
    private var viewModel: GameDetailViewModel?
    
    init(gameId: String) {
        self.gameId = gameId
        super.init(nibName: nil, bundle: nil)
    }
    
    required init?(coder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        view.backgroundColor = UIColor(hex: "#0F0F23")
        title = "Game Details"
    }
}
