package com.game_engine.casino.security

import android.content.Context
import android.os.Build

data class EncryptedData(val ciphertext: String, val iv: String)

data class DeviceFingerprint(
    val deviceId: String,
    val canvasHash: String,
    val model: String,
    val manufacturer: String,
    val osVersion: Int,
    val isRooted: Boolean,
    val isVerified: Boolean = false
)

data class IntegrityResult(
    val isGenuine: Boolean,
    val isVerified: Boolean,
    val isDeveloperModeEnabled: Boolean,
    val lastChecked: Long
)

enum class SecurityIssue {
    ROOT_DETECTED,
    REMOTE_ACCESS_APP_DETECTED,
    REMOTE_SERVICE_RUNNING,
    DEBUGGER_ATTACHED,
    APP_TAMPERED,
    DEVELOPER_MODE_ENABLED,
    UNKNOWN_SOURCES_ENABLED
}

data class SecurityCheckResult(
    val isSecure: Boolean,
    val issues: List<SecurityIssue>,
    val remoteAppsDetected: List<String>,
    val checkedAt: Long
)

object SecurityChecks {

    fun isDeviceRooted(context: Context): Boolean {
        if (checkRootApps(context)) return true
        if (checkRootBinaries()) return true
        if (checkDangerousPaths()) return true
        if (checkSuBinary()) return true
        if (checkTestKeys()) return true
        if (checkMagisk(context)) return true
        return false
    }

    fun hasRemoteAccessApps(context: Context): List<String> {
        val remoteApps = mutableListOf<String>()

        val dangerousApps = listOf(
            "com.anydesk.anydeskandroid",
            "com.philandro.anydesk",
            "com.teamviewer.teamviewer.market.mobile",
            "com.teamviewer.host.mobile",
            "com.teamviewer.quicksupport.mobile",
            "com.sand.airdroid",
            "com.airdroid",
            "com.google.android.apps.remotely",
            "com.iiordanov.freebVNC",
            "com.iiordanov.bVNC",
            "com.iiordanov.proVNC",
            "com.microsoft.rdc.android",
            "com.royaltek.bluedvnc",
            "com.teamviewer.teamviewer",
            "com.airmirror",
            "com.zoho.assist",
            "com.splashtop.remote.pad",
            "com.remoteutilities.viewer",
            "com.logmein.gotomypc.android",
            "com.parsecgaming.parsec",
            "com.limelight",
            "com.limelight.sunshine",
            "com.screen.mirroring",
            "com.mobizen.mirror",
            "com.mirrorop.pcv",
            "com.airsquirrels.reflector",
            "com.apowersoft.mirror",
            "com.letsview",
            "com.genymobile.scrcpy",
            "com.touchmouse.mobilemouse",
            "com.teslariustefan.remoteMouse",
            "com.hidmouse.wifimouse"
        )

        val pm = context.packageManager
        for (app in dangerousApps) {
            try {
                val info = pm.getPackageInfo(app, 0)
                if (info != null) {
                    remoteApps.add(app)
                }
            } catch (e: Exception) {
                // Not installed
            }
        }

        return remoteApps
    }

    fun hasRunningRemoteServices(): Boolean {
        val runningProcesses = android.os.ProcessManager.getRunningProcesses(100)
            .mapNotNull { it.processName }
            .toList()

        val remoteServices = listOf(
            "anydesk", "teamviewer", "airdroid", "airmirror",
            "remoteviewing", "vncserver", "rdp", "splashtop"
        )

        return runningProcesses.any { process ->
            remoteServices.any { service ->
                process.lowercase().contains(service)
            }
        }
    }

    fun isDeveloperModeEnabled(context: Context): Boolean {
        return try {
            if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.JELLY_BEAN_MR1) {
                android.provider.Settings.Secure.getInt(
                    context.contentResolver,
                    android.provider.Settings.Global.DEVELOPMENT_SETTINGS_ENABLED,
                    0
                ) == 1
            } else false
        } catch (e: Exception) { false }
    }

    fun isDebuggerAttached(): Boolean {
        return android.os.Debug.isDebuggerConnected()
    }

    fun verifyApkSignature(context: Context): Boolean {
        return try {
            val pm = context.packageManager
            val packageInfo = pm.getPackageInfo(
                context.packageName,
                android.content.pm.PackageManager.GET_SIGNATURES
            )
            packageInfo.signatures != null && packageInfo.signatures.isNotEmpty()
        } catch (e: Exception) { false }
    }

    private fun checkRootApps(context: Context): Boolean {
        val rootApps = listOf(
            "com.topjohnwu.magisk",
            "com.noshufou.android.su",
            "com.noshufou.android.su.elite",
            "eu.chainfire.supersu",
            "com.koushikdutta.superuser",
            "com.thirdparty.superuser",
            "com.yellowes.su",
            "com.kingroot.kinguser",
            "com.kingo.root",
            "com.smedialink.oneclickroot",
            "com.zhiqupk.root.global",
            "com.termux"
        )

        val pm = context.packageManager
        for (app in rootApps) {
            try {
                pm.getPackageInfo(app, 0)
                return true
            } catch (e: Exception) { }
        }
        return false
    }

    private fun checkRootBinaries(): Boolean {
        val paths = listOf(
            "/system/app/Superuser.apk", "/sbin/su", "/system/bin/su",
            "/system/xbin/su", "/data/local/xbin/su", "/data/local/bin/su",
            "/system/sdcard/su", "/data/local/su", "/su/bin/su", "/magisk/.core/bin/su"
        )

        for (path in paths) {
            if (java.io.File(path).exists()) return true
        }
        return false
    }

    private fun checkDangerousPaths(): Boolean {
        val paths = listOf(
            "/data/adb", "/data/adb/modules", "/data/dalvik-cache", "/data/local/tmp"
        )

        for (path in paths) {
            if (java.io.File(path).exists()) return true
        }
        return false
    }

    private fun checkSuBinary(): Boolean {
        return try {
            val process = Runtime.getRuntime().exec(arrayOf("su", "-c", "id"))
            val output = java.io.BufferedReader(
                java.io.InputStreamReader(process.inputStream)
            ).readLine()
            output != null && output.contains("uid=0")
        } catch (e: Exception) {
            false
        }
    }

    private fun checkTestKeys(): Boolean {
        val buildTags = Build.TAGS
        return buildTags != null && buildTags.contains("test-keys")
    }

    private fun checkMagisk(context: Context): Boolean {
        val magiskFiles = listOf(
            "/sbin/.magisk", "/sbin/.core", "/data/adb/magisk",
            "/data/adb/magisk.img", "/data/adb/modules"
        )

        for (file in magiskFiles) {
            if (java.io.File(file).exists()) return true
        }

        try {
            context.packageManager.getPackageInfo("com.topjohnwu.magisk", 0)
            return true
        } catch (e: Exception) { }

        return false
    }
}
