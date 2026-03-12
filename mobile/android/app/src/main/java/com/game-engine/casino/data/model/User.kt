package com.game-engine.casino.data.model

import com.google.gson.annotations.SerializedName

data class User(
    @SerializedName("id")
    val id: String,
    @SerializedName("email")
    val email: String,
    @SerializedName("username")
    val username: String,
    @SerializedName("phone")
    val phone: String?,
    @SerializedName("first_name")
    val firstName: String?,
    @SerializedName("last_name")
    val lastName: String?,
    @SerializedName("avatar_url")
    val avatarUrl: String?,
    @SerializedName("kyc_level")
    val kycLevel: Int,
    @SerializedName("is_verified")
    val isVerified: Boolean,
    @SerializedName("currency")
    val currency: String,
    @SerializedName("language")
    val language: String,
    @SerializedName("created_at")
    val createdAt: String
)

data class LoginRequest(
    @SerializedName("email")
    val email: String,
    @SerializedName("password")
    val password: String,
    @SerializedName("device_id")
    val deviceId: String?,
    @SerializedName("device_name")
    val deviceName: String?
)

data class LoginResponse(
    @SerializedName("user")
    val user: User,
    @SerializedName("access_token")
    val accessToken: String,
    @SerializedName("refresh_token")
    val refreshToken: String,
    @SerializedName("expires_in")
    val expiresIn: Long
)

data class RegisterRequest(
    @SerializedName("email")
    val email: String,
    @SerializedName("username")
    val username: String,
    @SerializedName("password")
    val password: String,
    @SerializedName("phone")
    val phone: String?,
    @SerializedName("currency")
    val currency: String,
    @SerializedName("referral_code")
    val referralCode: String?
)

data class RegisterResponse(
    @SerializedName("user")
    val user: User,
    @SerializedName("access_token")
    val accessToken: String,
    @SerializedName("refresh_token")
    val refreshToken: String,
    @SerializedName("expires_in")
    val expiresIn: Long,
    @SerializedName("message")
    val message: String
)

data class RefreshTokenRequest(
    @SerializedName("refresh_token")
    val refreshToken: String
)

data class RefreshTokenResponse(
    @SerializedName("access_token")
    val accessToken: String,
    @SerializedName("refresh_token")
    val refreshToken: String,
    @SerializedName("expires_in")
    val expiresIn: Long
)

data class ChangePasswordRequest(
    @SerializedName("current_password")
    val currentPassword: String,
    @SerializedName("new_password")
    val newPassword: String
)

data class ResetPasswordRequest(
    @SerializedName("email")
    val email: String
)

data class VerifyEmailRequest(
    @SerializedName("code")
    val code: String
)

data class ProfileUpdateRequest(
    @SerializedName("first_name")
    val firstName: String?,
    @SerializedName("last_name")
    val lastName: String?,
    @SerializedName("phone")
    val phone: String?,
    @SerializedName("language")
    val language: String?,
    @SerializedName("currency")
    val currency: String?
)
