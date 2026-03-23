package com.game_engine.casino.ui

import android.content.Intent
import android.content.pm.PackageManager
import android.os.Bundle
import android.view.WindowManager
import androidx.appcompat.app.AlertDialog
import androidx.appcompat.app.AppCompatActivity
import com.game_engine.casino.databinding.ActivitySecurityBlockBinding
import com.game_engine.casino.security.SecurityIssue
import com.game_engine.casino.security.SecurityManager

/**
 * Security Block Activity
 * 
 * This activity is shown when the app detects:
 * - Rooted device
 * - Remote access apps (AnyDesk, TeamViewer, etc.)
 * - Debugger attached
 * - Other security threats
 * 
 * The app cannot proceed until these issues are resolved.
 */
class SecurityBlockActivity : AppCompatActivity() {

    private lateinit var binding: ActivitySecurityBlockBinding
    private val securityManager by lazy { SecurityManager(applicationContext) }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        
        // Prevent screenshots and screen recording
        window.setFlags(
            WindowManager.LayoutParams.FLAG_SECURE,
            WindowManager.LayoutParams.FLAG_SECURE
        )
        
        binding = ActivitySecurityBlockBinding.inflate(layoutInflater)
        setContentView(binding.root)
        
        performSecurityCheck()
    }

    private fun performSecurityCheck() {
        val result = securityManager.performFullSecurityCheck()
        
        if (result.isSecure) {
            // Security passed - finish and go to main app
            finish()
            return
        }
        
        // Show security blocked UI
        showSecurityBlocked(result.issues, result.remoteAppsDetected)
    }

    private fun showSecurityBlocked(
        issues: List<SecurityIssue>,
        remoteApps: List<String>
    ) {
        val messageBuilder = StringBuilder()
        
        if (issues.contains(SecurityIssue.REMOTE_ACCESS_APP_DETECTED) || 
            issues.contains(SecurityIssue.REMOTE_SERVICE_RUNNING)) {
            messageBuilder.append(\"⚠️ REMOTE ACCESS APP DETECTED\\n\\n\")
            messageBuilder.append(\"The following apps must be uninstalled:\\n\\n\")
            
            val appNames = remoteApps.mapNotNull { getAppName(it) }
            appNames.forEach { name ->
                messageBuilder.append(\"• $name\\n\")
            }
            
            messageBuilder.append(\"\\nPlease uninstall these apps and restart the app.\\n\\n\")
        }
        
        if (issues.contains(SecurityIssue.ROOT_DETECTED)) {
            messageBuilder.append(\"⚠️ ROOTED DEVICE DETECTED\\n\\n\")
            messageBuilder.append(\"This app cannot run on rooted devices for security reasons.\\n\\n\")
        }
        
        if (issues.contains(SecurityIssue.DEBUGGER_ATTACHED)) {
            messageBuilder.append(\"⚠️ DEBUGGER DETECTED\\n\\n\")
            messageBuilder.append(\"Debugging tools must be disabled.\\n\\n\")
        }
        
        if (issues.contains(SecurityIssue.APP_TAMPERED)) {
            messageBuilder.append(\"⚠️ APP TAMPERING DETECTED\\n\\n\")
            messageBuilder.append(\"The app has been modified and cannot run.\\n\\n\")
        }
        
        if (issues.contains(SecurityIssue.DEVELOPER_MODE_ENABLED)) {
            messageBuilder.append(\"⚠️ DEVELOPER MODE ENABLED\\n\\n\")
            messageBuilder.append(\"Please disable Developer Options in settings.\\n\\n\")
        }
        
        messageBuilder.append(\"If you believe this is an error, please contact support.\")
        
        binding.tvSecurityMessage.text = messageBuilder.toString()
        
        binding.btnRetry.setOnClickListener {
            // Re-check security
            performSecurityCheck()
        }
        
        binding.btnExit.setOnClickListener {
            // Force close the app
            finishAffinity()
            System.exit(0)
        }
    }
    
    private fun getAppName(packageName: String): String? {
        return try {
            val pm = packageManager
            pm.getApplicationLabel(
                pm.getApplicationInfo(packageName, 0)
            ).toString()
        } catch (e: PackageManager.NameNotFoundException) {
            packageName.substringAfterLast(".")
        }
    }

    override fun onBackPressed() {
        // Prevent going back - force user to address the issue
        AlertDialog.Builder(this)
            .setTitle(\"Cannot Go Back\")
            .setMessage(\"Please address the security issues to use this app.\")
            .setPositiveButton(\"OK\", null)
            .show()
    }
}
