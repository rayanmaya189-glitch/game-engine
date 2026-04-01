package com.game_engine.casino.security

import android.content.Context
import android.os.Build
import dagger.hilt.android.qualifiers.ApplicationContext
import java.security.MessageDigest
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class SecurityManager @Inject constructor(
    @ApplicationContext private val context: Context
) {
    private val crypto = SecurityCrypto(context)

    fun isDeviceRooted(): Boolean = SecurityChecks.isDeviceRooted(context)

    fun hasRemoteAccessApps(): List<String> = SecurityChecks.hasRemoteAccessApps(context)

    fun isRemoteAccessAppDetected(): Boolean = hasRemoteAccessApps().isNotEmpty()

    fun performFullSecurityCheck(): SecurityCheckResult {
        val issues = mutableListOf<SecurityIssue>()
        val remoteApps = hasRemoteAccessApps()

        if (isDeviceRooted()) issues.add(SecurityIssue.ROOT_DETECTED)
        if (remoteApps.isNotEmpty()) issues.add(SecurityIssue.REMOTE_ACCESS_APP_DETECTED)
        if (SecurityChecks.hasRunningRemoteServices()) issues.add(SecurityIssue.REMOTE_SERVICE_RUNNING)
        if (SecurityChecks.isDebuggerAttached()) issues.add(SecurityIssue.DEBUGGER_ATTACHED)
        if (!SecurityChecks.verifyApkSignature(context)) issues.add(SecurityIssue.APP_TAMPERED)
        if (SecurityChecks.isDeveloperModeEnabled(context)) issues.add(SecurityIssue.DEVELOPER_MODE_ENABLED)

        return SecurityCheckResult(
            isSecure = issues.isEmpty(),
            issues = issues,
            remoteAppsDetected = remoteApps,
            checkedAt = System.currentTimeMillis()
        )
    }

    fun verifyPlayIntegrity(): IntegrityResult {
        return IntegrityResult(
            isGenuine = !isDeviceRooted(),
            isVerified = true,
            isDeveloperModeEnabled = SecurityChecks.isDeveloperModeEnabled(context),
            lastChecked = System.currentTimeMillis()
        )
    }

    fun performIntegrityCheck(): Boolean {
        if (isDeviceRooted()) return false
        if (SecurityChecks.isDebuggerAttached()) return false
        if (!SecurityChecks.verifyApkSignature(context)) return false
        return true
    }

    fun getDeviceFingerprint(): DeviceFingerprint {
        val sb = StringBuilder()
        sb.append(Build.MANUFACTURER)
        sb.append(Build.MODEL)
        sb.append(Build.DEVICE)
        sb.append(Build.HARDWARE)
        sb.append(Build.getSerial())
        sb.append(Build.FINGERPRINT)

        val hash = sha256(sb.toString())

        return DeviceFingerprint(
            deviceId = hash,
            canvasHash = sha256("${Build.MANUFACTURER}_${Build.MODEL}_canvas_signature"),
            model = Build.MODEL,
            manufacturer = Build.MANUFACTURER,
            osVersion = Build.VERSION.SDK_INT,
            isRooted = isDeviceRooted(),
            isVerified = verifyPlayIntegrity().isGenuine
        )
    }

    private fun sha256(input: String): String {
        val bytes = MessageDigest.getInstance("SHA-256").digest(input.toByteArray())
        return bytes.joinToString("") { "%02x".format(it) }
    }

    fun generateBiometricProtectedKey(keyAlias: String = DEFAULT_KEY_ALIAS) =
        crypto.generateBiometricProtectedKey(keyAlias)

    fun encryptWithBiometric(data: String, keyAlias: String = DEFAULT_KEY_ALIAS) =
        crypto.encryptWithBiometric(data, keyAlias)

    fun decryptWithBiometric(encryptedData: EncryptedData, keyAlias: String = DEFAULT_KEY_ALIAS) =
        crypto.decryptWithBiometric(encryptedData, keyAlias)

    fun encryptData(data: String, keyAlias: String = DEFAULT_KEY_ALIAS) =
        crypto.encryptData(data, keyAlias)

    fun decryptData(encryptedData: String, keyAlias: String = DEFAULT_KEY_ALIAS) =
        crypto.decryptData(encryptedData, keyAlias)

    fun storeToken(key: String, value: String) = crypto.storeToken(key, value)

    fun getToken(key: String) = crypto.getToken(key)

    fun clearSensitiveData() = crypto.clearSensitiveData()

    companion object {
        private const val DEFAULT_KEY_ALIAS = "casino_secure_key"
    }
}
