import UIKit
import UserNotifications
import FirebaseCore
import FirebaseMessaging

/**
 * Notification Service
 * 
 * Handles push notifications via FCM (which wraps APNs on iOS).
 */
class NotificationService: NSObject, MessagingDelegate {
    
    static let shared = NotificationService()
    
    private override init() {
        super.init()
    }
    
    // MARK: - Setup
    
    func configure() {
        FirebaseApp.configure()
        
        // Set delegate
        Messaging.messaging().delegate = self
        
        // Request permission
        requestAuthorization()
        
        // Register for remote notifications
        UNUserNotificationCenter.current().delegate = self
    }
    
    private func requestAuthorization() {
        let options: UNAuthorizationOptions = [.alert, .badge, .sound]
        
        UNUserNotificationCenter.current().requestAuthorization(options: options) { granted, error in
            if let error = error {
                print("Notification authorization error: \(error)")
                return
            }
            
            if granted {
                DispatchQueue.main.async {
                    UIApplication.shared.registerForRemoteNotifications()
                }
            }
        }
    }
    
    // MARK: - Token Management
    
    func messaging(_ messaging: Messaging, didReceiveRegistrationToken fcmToken: String?) {
        guard let token = fcmToken else { return }
        
        // Send token to server
        sendTokenToServer(token)
    }
    
    private func sendTokenToServer(_ token: String) {
        // API call to register device token
        print("FCM Token: \(token)")
    }
    
    // MARK: - Handle Notification
    
    func handleNotification(userInfo: [AnyHashable: Any]) {
        guard let type = userInfo["type"] as? String else { return }
        
        switch type {
        case "game":
            handleGameNotification(userInfo)
        case "financial":
            handleFinancialNotification(userInfo)
        case "promotion":
            handlePromotionNotification(userInfo)
        case "tournament":
            handleTournamentNotification(userInfo)
        default:
            handleSystemNotification(userInfo)
        }
    }
    
    private func handleGameNotification(_ userInfo: [AnyHashable: Any]) {
        guard let gameId = userInfo["game_id"] as? String else { return }
        
        // Navigate to game
        NotificationCenter.default.post(
            name: .navigateToGame,
            object: nil,
            userInfo: ["gameId": gameId]
        )
    }
    
    private func handleFinancialNotification(_ userInfo: [AnyHashable: Any]) {
        // Navigate to wallet
        NotificationCenter.default.post(
            name: .navigateToWallet,
            object: nil
        )
    }
    
    private func handlePromotionNotification(_ userInfo: [AnyHashable: Any]) {
        guard let bonusId = userInfo["bonus_id"] as? String else { return }
        
        // Navigate to bonus
        NotificationCenter.default.post(
            name: .navigateToBonus,
            object: nil,
            userInfo: ["bonusId": bonusId]
        )
    }
    
    private func handleTournamentNotification(_ userInfo: [AnyHashable: Any]) {
        guard let tournamentId = userInfo["tournament_id"] as? String else { return }
        
        // Navigate to tournament
        NotificationCenter.default.post(
            name: .navigateToTournament,
            object: nil,
            userInfo: ["tournamentId": tournamentId]
        )
    }
    
    private func handleSystemNotification(_ userInfo: [AnyHashable: Any]) {
        // Handle system message
    }
    
    // MARK: - Local Notifications
    
    func scheduleLocalNotification(
        title: String,
        body: String,
        identifier: String,
        delay: TimeInterval = 0
    ) {
        let content = UNMutableNotificationContent()
        content.title = title
        content.body = body
        content.sound = .default
        
        let trigger = UNTimeIntervalNotificationTrigger(
            timeInterval: delay > 0 ? delay : 0.1,
            repeats: false
        )
        
        let request = UNNotificationRequest(
            identifier: identifier,
            content: content,
            trigger: trigger
        )
        
        UNUserNotificationCenter.current().add(request)
    }
}

// MARK: - UNUserNotificationCenterDelegate

extension NotificationService: UNUserNotificationCenterDelegate {
    
    func userNotificationCenter(
        _ center: UNUserNotificationCenter,
        willPresent notification: UNNotification,
        withCompletionHandler completionHandler: @escaping (UNNotificationPresentationOptions) -> Void
    ) {
        // Show notification even when app is in foreground
        completionHandler([.banner, .badge, .sound])
    }
    
    func userNotificationCenter(
        _ center: UNUserNotificationCenter,
        didReceive response: UNNotificationResponse,
        withCompletionHandler completionHandler: @escaping () -> Void
    ) {
        let userInfo = response.notification.request.content.userInfo
        handleNotification(userInfo: userInfo)
        
        completionHandler()
    }
}

// MARK: - Notification Names

extension Notification.Name {
    static let navigateToGame = Notification.Name("navigateToGame")
    static let navigateToWallet = Notification.Name("navigateToWallet")
    static let navigateToBonus = Notification.Name("navigateToBonus")
    static let navigateToTournament = Notification.Name("navigateToTournament")
}
