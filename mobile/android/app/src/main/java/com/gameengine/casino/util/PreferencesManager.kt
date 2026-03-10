package com.gameengine.casino.util

import android.content.Context
import androidx.datastore.core.DataStore
import androidx.datastore.preferences.core.*
import androidx.datastore.preferences.preferencesDataStore
import dagger.hilt.android.qualifiers.ApplicationContext
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.runBlocking
import javax.inject.Inject
import javax.inject.Singleton

private val Context.dataStore: DataStore<Preferences> by preferencesDataStore(name = "casino_prefs")

@Singleton
class PreferencesManager @Inject constructor(
    @ApplicationContext private val context: Context
) {
    companion object {
        private val ACCESS_TOKEN = stringPreferencesKey("access_token")
        private val REFRESH_TOKEN = stringPreferencesKey("refresh_token")
        private val USER_ID = stringPreferencesKey("user_id")
        private val USER_EMAIL = stringPreferencesKey("user_email")
        private val USERNAME = stringPreferencesKey("username")
        private val CURRENCY = stringPreferencesKey("currency")
        private val LANGUAGE = stringPreferencesKey("language")
        private val BIOMETRIC_ENABLED = booleanPreferencesKey("biometric_enabled")
        private val IS_FIRST_LAUNCH = booleanPreferencesKey("is_first_launch")
    }
    
    suspend fun saveAccessToken(token: String) {
        context.dataStore.edit { preferences ->
            preferences[ACCESS_TOKEN] = token
        }
    }
    
    suspend fun saveRefreshToken(token: String) {
        context.dataStore.edit { preferences ->
            preferences[REFRESH_TOKEN] = token
        }
    }
    
    suspend fun saveUserId(userId: String) {
        context.dataStore.edit { preferences ->
            preferences[USER_ID] = userId
        }
    }
    
    suspend fun saveUserEmail(email: String) {
        context.dataStore.edit { preferences ->
            preferences[USER_EMAIL] = email
        }
    }
    
    suspend fun saveUsername(username: String) {
        context.dataStore.edit { preferences ->
            preferences[USERNAME] = username
        }
    }
    
    suspend fun saveCurrency(currency: String) {
        context.dataStore.edit { preferences ->
            preferences[CURRENCY] = currency
        }
    }
    
    suspend fun saveLanguage(language: String) {
        context.dataStore.edit { preferences ->
            preferences[LANGUAGE] = language
        }
    }
    
    suspend fun setBiometricEnabled(enabled: Boolean) {
        context.dataStore.edit { preferences ->
            preferences[BIOMETRIC_ENABLED] = enabled
        }
    }
    
    suspend fun setFirstLaunch(isFirst: Boolean) {
        context.dataStore.edit { preferences ->
            preferences[IS_FIRST_LAUNCH] = isFirst
        }
    }
    
    suspend fun clearTokens() {
        context.dataStore.edit { preferences ->
            preferences.remove(ACCESS_TOKEN)
            preferences.remove(REFRESH_TOKEN)
            preferences.remove(USER_ID)
        }
    }
    
    fun getAccessToken(): String? = runBlocking {
        context.dataStore.data.first()[ACCESS_TOKEN]
    }
    
    fun getRefreshToken(): String? = runBlocking {
        context.dataStore.data.first()[REFRESH_TOKEN]
    }
    
    fun getUserId(): String? = runBlocking {
        context.dataStore.data.first()[USER_ID]
    }
    
    fun getUserEmail(): String? = runBlocking {
        context.dataStore.data.first()[USER_EMAIL]
    }
    
    fun getUsername(): String? = runBlocking {
        context.dataStore.data.first()[USERNAME]
    }
    
    fun getCurrency(): String? = runBlocking {
        context.dataStore.data.first()[CURRENCY]
    }
    
    fun getLanguage(): String? = runBlocking {
        context.dataStore.data.first()[LANGUAGE]
    }
    
    fun isBiometricEnabled(): Boolean = runBlocking {
        context.dataStore.data.first()[BIOMETRIC_ENABLED] ?: false
    }
    
    fun isFirstLaunch(): Boolean = runBlocking {
        context.dataStore.data.first()[IS_FIRST_LAUNCH] ?: true
    }
}
