import UIKit

class MainTabBarController: UITabBarController {
    
    override func viewDidLoad() {
        super.viewDidLoad()
        setupAppearance()
    }
    
    private func setupAppearance() {
        tabBar.backgroundColor = UIColor(hex: "#1A1A2E")
        tabBar.tintColor = UIColor(hex: "#FF6B35")
        tabBar.unselectedItemTintColor = UIColor.white.withAlphaComponent(0.6)
        tabBar.barTintColor = UIColor(hex: "#1A1A2E")
        tabBar.isTranslucent = false
        
        if #available(iOS 15.0, *) {
            let appearance = UITabBarAppearance()
            appearance.configureWithOpaqueBackground()
            appearance.backgroundColor = UIColor(hex: "#1A1A2E")
            
            tabBar.standardAppearance = appearance
            tabBar.scrollEdgeAppearance = appearance
        }
    }
}
