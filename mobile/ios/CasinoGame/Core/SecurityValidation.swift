import UIKit

extension SecurityService {

    // MARK: - Enhanced Jailbreak Detection

    func isDeviceJailbroken() -> Bool {
        #if targetEnvironment(simulator)
        return false
        #else
        let suspiciousPaths = [
            "/Applications/Cydia.app",
            "/Library/MobileSubstrate/MobileSubstrate.dylib",
            "/bin/bash",
            "/usr/sbin/sshd",
            "/etc/apt",
            "/private/var/lib/apt/",
            "/usr/bin/ssh",
            "/private/var/stash",
            "/private/var/lib/cydia",
            "/private/var/tmp/cydia.log",
            "/var/cache/apt",
            "/var/lib/cydia"
        ]

        for path in suspiciousPaths {
            if FileManager.default.fileExists(atPath: path) {
                return true
            }
        }

        if let url = URL(string: "cydia://package/com.example.package"),
           UIApplication.shared.canOpenURL(url) {
            return true
        }

        let testPaths = [
            "/private/jailbreak_test_1.txt",
            "/private/var/mobile/Library/AddressBook/AddressBook.sqlitedb",
            "/Library/MobileSubstrate/DynamicLibraries"
        ]

        for testPath in testPaths {
            if FileManager.default.isWritableFile(atPath: testPath) {
                return true
            }
        }

        let suspiciousSymlinks = ["/Applications", "/Library/Ringtones", "/Library/Wallpaper"]
        for path in suspiciousSymlinks {
            var isSymlink: ObjCBool = false
            if FileManager.default.fileExists(atPath: path, isDirectory: &isSymlink) {
                do {
                    let _ = try FileManager.default.destinationOfSymbolicLink(atPath: path)
                    return true
                } catch {}
            }
        }

        let libraries = dlopen(nil, RTLD_NOW)
        defer { dlclose(libraries) }

        let suspiciousLibs = ["SubstrateLoader", "SubstrateInserter", "SubstrateBootstrap", "Cydia"]
        for lib in suspiciousLibs {
            if dlsym(libraries, lib) != nil {
                return true
            }
        }

        if canFork() {
            return true
        }

        return false
        #endif
    }

    private func canFork() -> Bool {
        #if targetEnvironment(simulator)
        return false
        #else
        let task = Process()
        task.executableURL = URL(fileURLWithPath: "/bin/ls")
        do {
            try task.run()
            return false
        } catch {
            return true
        }
        #endif
    }

    // MARK: - Remote Access App Detection

    func detectRemoteAccessApps() -> [String] {
        var detectedApps: [String] = []

        let remoteAppBundleIds: [String: String] = [
            "com.anydesk.anydeskandroid": "AnyDesk",
            "com.teamviewer.teamviewer.market.mobile": "TeamViewer",
            "com.teamviewer.host.mobile": "TeamViewer Host",
            "com.teamviewer.quicksupport.mobile": "TeamViewer QuickSupport",
            "com.sand.airdroid": "AirDroid",
            "com.airdroid": "AirDroid Classic",
            "com.google.android.apps.remotely": "Chrome Remote Desktop",
            "com.iiordanov.freebVNC": "bVNC",
            "com.iiordanov.bVNC": "bVNC Pro",
            "com.microsoft.rdc.android": "Microsoft Remote Desktop",
            "com.logmein.gotomypc.android": "GoToMyPC",
            "com.parsecgaming.parsec": "Parsec",
            "com.limelight": "Moonlight",
            "com.splashtop.remote.pad": "Splashtop",
            "com.zoho.assist": "Zoho Assist",
            "com.remoteutilities.viewer": "Remote Utilities",
            "com.philandro.anydesk": "AnyDesk",
            "com.mobizen.mirror": "Mobizen",
            "com.apowersoft.mirror": "ApowerMirror",
            "com.letsview": "LetsView",
            "com.screen.mirroring": "Screen Mirroring"
        ]

        for (bundleId, appName) in remoteAppBundleIds {
            if let url = URL(string: "\(bundleId)://"),
               UIApplication.shared.canOpenURL(url) {
                detectedApps.append(appName)
            }
        }

        return detectedApps
    }

    func hasRemoteAccessApps() -> Bool {
        return !detectRemoteAccessApps().isEmpty
    }

    func performFullSecurityCheck() -> FullSecurityResult {
        var issues: [SecurityIssue] = []

        if isDeviceJailbroken() {
            issues.append(.jailbreakDetected)
        }

        if hasRemoteAccessApps() {
            issues.append(.remoteAccessAppInstalled)
        }

        if isDebuggerAttached() {
            issues.append(.debuggerAttached)
        }

        if isRuntimeManipulated() {
            issues.append(.runtimeManipulation)
        }

        if UIScreen.main.isCaptured {
            issues.append(.screenRecordingDetected)
        }

        return FullSecurityResult(
            isSecure: issues.isEmpty,
            issues: issues,
            remoteApps: detectRemoteAccessApps(),
            checkedAt: Date()
        )
    }

    // MARK: - Debugger Detection

    private func isDebuggerAttached() -> Bool {
        var info = kinfo_proc()
        var size = MemoryLayout<kinfo_proc>.stride
        var mib: [Int32] = [CTL_KERN, KERN_PROC, KERN_PROC_PID, getpid()]

        let result = sysctl(&mib, UInt32(mib.count), &info, &size, nil, 0)

        return result == 0 && (info.kp_proc.p_flag & P_TRACED) != 0
    }

    // MARK: - Runtime Manipulation Detection

    private func isRuntimeManipulated() -> Bool {
        let loadedLibs = _dyld_image_count()

        for i in 0..<loadedLibs {
            if let name = _dyld_get_image_name(i) {
                let libName = String(cString: name)
                if libName.contains("Frida") || libName.contains("frida") ||
                   libName.contains("Substrate") || libName.contains("substrate") {
                    return true
                }
            }
        }

        return false
    }
}
