package com.gameengine.casino.ui.auth

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.*
import androidx.compose.material3.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.text.input.PasswordVisualTransformation
import androidx.compose.ui.text.input.VisualTransformation
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.hilt.navigation.compose.hiltViewModel
import com.gameengine.casino.ui.theme.Gold
import com.gameengine.casino.ui.theme.Primary

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun RegisterScreen(
    onNavigateToLogin: () -> Unit,
    onRegisterSuccess: () -> Unit,
    viewModel: AuthViewModel = hiltViewModel()
) {
    val uiState by viewModel.uiState.collectAsState()
    
    var email by remember { mutableStateOf("") }
    var username by remember { mutableStateOf("") }
    var password by remember { mutableStateOf("") }
    var confirmPassword by remember { mutableStateOf("") }
    var phone by remember { mutableStateOf("") }
    var referralCode by remember { mutableStateOf("") }
    var passwordVisible by remember { mutableStateOf(false) }
    var confirmPasswordVisible by remember { mutableStateOf(false) }
    var agreedToTerms by remember { mutableStateOf(false) }
    
    LaunchedEffect(uiState.isLoggedIn) {
        if (uiState.isLoggedIn) {
            onRegisterSuccess()
        }
    }
    
    Box(
        modifier = Modifier
            .fillMaxSize()
            .background(MaterialTheme.colorScheme.background)
    ) {
        Column(
            modifier = Modifier
                .fillMaxSize()
                .verticalScroll(rememberScrollState())
                .padding(24.dp),
            horizontalAlignment = Alignment.CenterHorizontally
        ) {
            Spacer(modifier = Modifier.height(40.dp))
            
            // Logo/Title
            Text(
                text = "CREATE",
                style = MaterialTheme.typography.displaySmall,
                color = Gold
            )
            Text(
                text = "ACCOUNT",
                style = MaterialTheme.typography.headlineMedium,
                color = MaterialTheme.colorScheme.onBackground
            )
            
            Spacer(modifier = Modifier.height(32.dp))
            
            // Email field
            OutlinedTextField(
                value = email,
                onValueChange = { email = it },
                label = { Text("Email") },
                leadingIcon = {
                    Icon(Icons.Default.Email, contentDescription = null)
                },
                modifier = Modifier.fillMaxWidth(),
                keyboardOptions = KeyboardOptions(
                    keyboardType = KeyboardType.Email,
                    imeAction = ImeAction.Next
                ),
                singleLine = true,
                shape = RoundedCornerShape(12.dp)
            )
            
            Spacer(modifier = Modifier.height(12.dp))
            
            // Username field
            OutlinedTextField(
                value = username,
                onValueChange = { username = it },
                label = { Text("Username") },
                leadingIcon = {
                    Icon(Icons.Default.Person, contentDescription = null)
                },
                modifier = Modifier.fillMaxWidth(),
                keyboardOptions = KeyboardOptions(imeAction = ImeAction.Next),
                singleLine = true,
                shape = RoundedCornerShape(12.dp)
            )
            
            Spacer(modifier = Modifier.height(12.dp))
            
            // Phone field
            OutlinedTextField(
                value = phone,
                onValueChange = { phone = it },
                label = { Text("Phone (Optional)") },
                leadingIcon = {
                    Icon(Icons.Default.Phone, contentDescription = null)
                },
                modifier = Modifier.fillMaxWidth(),
                keyboardOptions = KeyboardOptions(
                    keyboardType = KeyboardType.Phone,
                    imeAction = ImeAction.Next
                ),
                singleLine = true,
                shape = RoundedCornerShape(12.dp)
            )
            
            Spacer(modifier = Modifier.height(12.dp))
            
            // Password field
            OutlinedTextField(
                value = password,
                onValueChange = { password = it },
                label = { Text("Password") },
                leadingIcon = {
                    Icon(Icons.Default.Lock, contentDescription = null)
                },
                trailingIcon = {
                    IconButton(onClick = { passwordVisible = !passwordVisible }) {
                        Icon(
                            imageVector = if (passwordVisible) Icons.Default.Visibility else Icons.Default.VisibilityOff,
                            contentDescription = null
                        )
                    }
                },
                visualTransformation = if (passwordVisible) VisualTransformation.None else PasswordVisualTransformation(),
                modifier = Modifier.fillMaxWidth(),
                keyboardOptions = KeyboardOptions(
                    keyboardType = KeyboardType.Password,
                    imeAction = ImeAction.Next
                ),
                singleLine = true,
                shape = RoundedCornerShape(12.dp)
            )
            
            Spacer(modifier = Modifier.height(12.dp))
            
            // Confirm Password field
            OutlinedTextField(
                value = confirmPassword,
                onValueChange = { confirmPassword = it },
                label = { Text("Confirm Password") },
                leadingIcon = {
                    Icon(Icons.Default.Lock, contentDescription = null)
                },
                trailingIcon = {
                    IconButton(onClick = { confirmPasswordVisible = !confirmPasswordVisible }) {
                        Icon(
                            imageVector = if (confirmPasswordVisible) Icons.Default.Visibility else Icons.Default.VisibilityOff,
                            contentDescription = null
                        )
                    }
                },
                visualTransformation = if (confirmPasswordVisible) VisualTransformation.None else PasswordVisualTransformation(),
                modifier = Modifier.fillMaxWidth(),
                keyboardOptions = KeyboardOptions(
                    keyboardType = KeyboardType.Password,
                    imeAction = ImeAction.Next
                ),
                singleLine = true,
                shape = RoundedCornerShape(12.dp)
            )
            
            Spacer(modifier = Modifier.height(12.dp))
            
            // Referral Code field
            OutlinedTextField(
                value = referralCode,
                onValueChange = { referralCode = it },
                label = { Text("Referral Code (Optional)") },
                leadingIcon = {
                    Icon(Icons.Default.CardGiftcard, contentDescription = null)
                },
                modifier = Modifier.fillMaxWidth(),
                keyboardOptions = KeyboardOptions(imeAction = ImeAction.Done),
                singleLine = true,
                shape = RoundedCornerShape(12.dp)
            )
            
            Spacer(modifier = Modifier.height(16.dp))
            
            // Terms checkbox
            Row(
                modifier = Modifier.fillMaxWidth(),
                verticalAlignment = Alignment.CenterVertically
            ) {
                Checkbox(
                    checked = agreedToTerms,
                    onCheckedChange = { agreedToTerms = it }
                )
                Text(
                    text = "I agree to Terms & Conditions",
                    style = MaterialTheme.typography.bodyMedium,
                    color = MaterialTheme.colorScheme.onBackground
                )
            }
            
            Spacer(modifier = Modifier.height(24.dp))
            
            // Register Button
            Button(
                onClick = {
                    viewModel.register(
                        email = email,
                        username = username,
                        password = password,
                        phone = phone.ifEmpty { null },
                        currency = "USD",
                        referralCode = referralCode.ifEmpty { null }
                    )
                },
                modifier = Modifier
                    .fillMaxWidth()
                    .height(56.dp),
                enabled = !uiState.isLoading &&
                        email.isNotEmpty() &&
                        username.isNotEmpty() &&
                        password.isNotEmpty() &&
                        password == confirmPassword &&
                        agreedToTerms,
                shape = RoundedCornerShape(12.dp),
                colors = ButtonDefaults.buttonColors(containerColor = Primary)
            ) {
                if (uiState.isLoading) {
                    CircularProgressIndicator(
                        modifier = Modifier.size(24.dp),
                        color = MaterialTheme.colorScheme.onPrimary
                    )
                } else {
                    Text("REGISTER", style = MaterialTheme.typography.titleMedium)
                }
            }
            
            // Error Message
            if (uiState.error != null) {
                Spacer(modifier = Modifier.height(16.dp))
                Text(
                    text = uiState.error!!,
                    color = MaterialTheme.colorScheme.error,
                    style = MaterialTheme.typography.bodyMedium,
                    textAlign = TextAlign.Center
                )
            }
            
            Spacer(modifier = Modifier.height(24.dp))
            
            // Login link
            Row(
                modifier = Modifier.padding(vertical = 16.dp),
                verticalAlignment = Alignment.CenterVertically
            ) {
                Text(
                    text = "Already have an account?",
                    color = MaterialTheme.colorScheme.onBackground
                )
                TextButton(onClick = onNavigateToLogin) {
                    Text(text = "Sign In", color = Primary)
                }
            }
        }
    }
}
