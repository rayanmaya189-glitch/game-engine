package com.game_engine.casino

import android.content.Intent
import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.ui.Modifier
import com.game_engine.casino.security.SecurityManager
import com.game_engine.casino.ui.navigation.CasinoNavHost
import com.game_engine.casino.ui.SecurityBlockActivity
import com.game_engine.casino.ui.theme.CasinoGameTheme
import dagger.hilt.android.AndroidEntryPoint
import javax.inject.Inject

@AndroidEntryPoint
class MainActivity : ComponentActivity() {
    
    @Inject
    lateinit var securityManager: SecurityManager
    
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        
        // Perform security check before showing the app
        if (!performSecurityCheck()) {
            return
        }
        
        enableEdgeToEdge()
        setContent {
            CasinoGameTheme {
                Surface(
                    modifier = Modifier.fillMaxSize(),
                    color = MaterialTheme.colorScheme.background
                ) {
                    CasinoNavHost()
                }
            }
        }
    }
    
    private fun performSecurityCheck(): Boolean {
        val result = securityManager.performFullSecurityCheck()
        
        if (!result.isSecure) {
            // Block the app and show security block screen
            val intent = Intent(this, SecurityBlockActivity::class.java).apply {
                putExtra("security_issues", result.issues.map { it.name }.toTypedArray())
                putExtra("remote_apps", result.detectedRemoteApps.toTypedArray())
                flags = Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TASK
            }
            startActivity(intent)
            finish()
            return false
        }
        
        return true
    }
}
