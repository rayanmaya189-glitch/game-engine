package com.gameengine.casino.data.api

import com.gameengine.casino.data.model.*
import retrofit2.Response
import retrofit2.http.*

interface AuthApi {
    
    @POST("auth/login")
    suspend fun login(@Body request: LoginRequest): Response<LoginResponse>
    
    @POST("auth/register")
    suspend fun register(@Body request: RegisterRequest): Response<RegisterResponse>
    
    @POST("auth/refresh")
    suspend fun refreshToken(@Body request: RefreshTokenRequest): Response<RefreshTokenResponse>
    
    @POST("auth/logout")
    suspend fun logout(): Response<Unit>
    
    @POST("auth/change-password")
    suspend fun changePassword(@Body request: ChangePasswordRequest): Response<Unit>
    
    @POST("auth/reset-password")
    suspend fun resetPassword(@Body request: ResetPasswordRequest): Response<Unit>
    
    @POST("auth/verify-email")
    suspend fun verifyEmail(@Body request: VerifyEmailRequest): Response<Unit>
    
    @GET("auth/me")
    suspend fun getCurrentUser(): Response<User>
    
    @PUT("auth/profile")
    suspend fun updateProfile(@Body request: ProfileUpdateRequest): Response<User>
    
    @GET("auth/2fa/status")
    suspend fun get2FAStatus(): Response<Map<String, Boolean>>
    
    @POST("auth/2fa/enable")
    suspend fun enable2FA(): Response<Map<String, String>>
    
    @POST("auth/2fa/verify")
    suspend fun verify2FA(@Body request: Map<String, String>): Response<Unit>
    
    @POST("auth/2fa/disable")
    suspend fun disable2FA(@Body request: Map<String, String>): Response<Unit>
}
