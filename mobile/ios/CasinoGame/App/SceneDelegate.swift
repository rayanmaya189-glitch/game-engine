import UIKit

class SceneDelegate: UIResponder, UIWindowSceneDelegate {
    
    var window: UIWindow?
    private var appCoordinator: AppCoordinator?
    
    func scene(
        _ scene: UIScene,
        willConnectTo session: UISceneSession,
        options connectionOptions: UIScene.ConnectionOptions
    ) {
        guard let windowScene = (scene as? UIWindowScene) else { return }
        
        // Perform security check before starting the app
        let securityService = SecurityService.shared
        let securityResult = securityService.performFullSecurityCheck()
        
        if !securityResult.isSecure {
            // Show security block screen
            showSecurityBlockScreen(windowScene: windowScene, result: securityResult)
            return
        }
        
        let window = UIWindow(windowScene: windowScene)
        self.window = window
        
        appCoordinator = AppCoordinator(window: window)
        appCoordinator?.start()
        
        window.makeKeyAndVisible()
    }
    
    private func showSecurityBlockScreen(windowScene: UIWindowScene, result: FullSecurityResult) {
        let window = UIWindow(windowScene: windowScene)
        self.window = window
        
        let securityBlockVC = SecurityBlockViewController(securityResult: result)
        window.rootViewController = securityBlockVC
        window.makeKeyAndVisible()
    }
    
    func sceneDidDisconnect(_ scene: UIScene) {
        // Called when the scene is being released by the system
    }
    
    func sceneDidBecomeActive(_ scene: UIScene) {
        // Called when the scene has moved from an inactive state to an active state
    }
    
    func sceneWillResignActive(_ scene: UIScene) {
        // Called when the scene will move from an active state to an inactive state
    }
    
    func sceneWillEnterForeground(_ scene: UIScene) {
        // Called as the scene transitions from the background to the foreground
    }
    
    func sceneDidEnterBackground(_ scene: UIScene) {
        // Called as the scene transitions from the foreground to the background
    }
}

// MARK: - AppCoordinator

class AppCoordinator {
    
    private let window: UIWindow
    private let navigationController: UINavigationController
    
    init(window: UIWindow) {
        self.window = window
        self.navigationController = UINavigationController()
    }
    
    func start() {
        // Check if user is logged in
        if UserDefaultsManager.shared.isLoggedIn {
            showMainScreen()
        } else {
            showLoginScreen()
        }
        
        window.rootViewController = navigationController
    }
    
    func showLoginScreen() {
        let loginVC = LoginViewController()
        loginVC.onLoginSuccess = { [weak self] in
            self?.showMainScreen()
        }
        loginVC.onRegisterTapped = { [weak self] in
            self?.showRegisterScreen()
        }
        
        navigationController.setViewControllers([loginVC], animated: false)
    }
    
    func showRegisterScreen() {
        let registerVC = RegisterViewController()
        registerVC.onRegisterSuccess = { [weak self] in
            self?.showMainScreen()
        }
        registerVC.onLoginTapped = { [weak self] in
            self?.navigationController.popViewController(animated: true)
        }
        
        navigationController.pushViewController(registerVC, animated: true)
    }
    
    func showMainScreen() {
        let mainTabBar = MainTabBarController()
        
        // Setup tab bar items
        let homeVC = HomeViewController()
        homeVC.tabBarItem = UITabBarItem(title: "Home", image: UIImage(systemName: "house"), selectedImage: UIImage(systemName: "house.fill"))
        
        let gamesVC = GamesViewController()
        gamesVC.tabBarItem = UITabBarItem(title: "Games", image: UIImage(systemName: "gamecontroller"), selectedImage: UIImage(systemName: "gamecontroller.fill"))
        
        let walletVC = WalletViewController()
        walletVC.tabBarItem = UITabBarItem(title: "Wallet", image: UIImage(systemName: "wallet.pass"), selectedImage: UIImage(systemName: "wallet.pass.fill"))
        
        let profileVC = ProfileViewController()
        profileVC.tabBarItem = UITabBarItem(title: "Profile", image: UIImage(systemName: "person"), selectedImage: UIImage(systemName: "person.fill"))
        
        mainTabBar.viewControllers = [
            UINavigationController(rootViewController: homeVC),
            UINavigationController(rootViewController: gamesVC),
            UINavigationController(rootViewController: walletVC),
            UINavigationController(rootViewController: profileVC)
        ]
        
        window.rootViewController = mainTabBar
    }
}
