import SwiftUI

// MARK: - Content View (Main Navigation)
struct ContentView: View {
    @EnvironmentObject var authManager: AuthManager

    var body: some View {
        if authManager.isLoggedIn {
            MainTabView()
        } else {
            LoginView()
        }
    }
}

// MARK: - Login View
struct LoginView: View {
    @EnvironmentObject var authManager: AuthManager
    @State private var email = ""
    @State private var password = ""
    @State private var showRegister = false

    var body: some View {
        NavigationStack {
            VStack(spacing: 24) {
                Image(systemName: "casino.fill")
                    .font(.system(size: 60))
                    .foregroundColor(.purple)

                Text("Casino Games")
                    .font(.largeTitle)
                    .fontWeight(.bold)

                Text("Sign in to play")
                    .foregroundColor(.secondary)

                if let error = authManager.error {
                    Text(error)
                        .foregroundColor(.red)
                        .font(.caption)
                }

                TextField("Email", text: $email)
                    .textFieldStyle(.roundedBorder)
                    .textContentType(.emailAddress)
                    .autocapitalization(.none)

                SecureField("Password", text: $password)
                    .textFieldStyle(.roundedBorder)
                    .textContentType(.password)

                Button(action: {
                    Task {
                        await authManager.login(email: email, password: password)
                    }
                }) {
                    if authManager.isLoading {
                        ProgressView()
                            .frame(maxWidth: .infinity)
                    } else {
                        Text("Sign In")
                            .frame(maxWidth: .infinity)
                    }
                }
                .buttonStyle(.borderedProminent)
                .disabled(email.isEmpty || password.isEmpty || authManager.isLoading)

                NavigationLink(destination: RegisterView()) {
                    Text("Don't have an account? Sign Up")
                }
            }
            .padding()
            .navigationBarHidden(true)
        }
    }
}

// MARK: - Register View
struct RegisterView: View {
    @EnvironmentObject var authManager: AuthManager
    @Environment(\.dismiss) var dismiss
    @State private var email = ""
    @State private var username = ""
    @State private var phone = ""
    @State private var password = ""
    @State private var confirmPassword = ""

    var body: some View {
        VStack(spacing: 16) {
            Text("Create Account")
                .font(.title)
                .fontWeight(.bold)

            if let error = authManager.error {
                Text(error)
                    .foregroundColor(.red)
                    .font(.caption)
            }

            TextField("Username", text: $username)
                .textFieldStyle(.roundedBorder)

            TextField("Email", text: $email)
                .textFieldStyle(.roundedBorder)
                .textContentType(.emailAddress)
                .autocapitalization(.none)

            TextField("Phone (Optional)", text: $phone)
                .textFieldStyle(.roundedBorder)
                .textContentType(.telephoneNumber)

            SecureField("Password", text: $password)
                .textFieldStyle(.roundedBorder)

            SecureField("Confirm Password", text: $confirmPassword)
                .textFieldStyle(.roundedBorder)

            Button(action: {
                Task {
                    await authManager.register(
                        email: email,
                        password: password,
                        username: username,
                        phone: phone.isEmpty ? nil : phone
                    )
                }
            }) {
                if authManager.isLoading {
                    ProgressView()
                        .frame(maxWidth: .infinity)
                } else {
                    Text("Create Account")
                        .frame(maxWidth: .infinity)
                }
            }
            .buttonStyle(.borderedProminent)
            .disabled(email.isEmpty || password.isEmpty || username.isEmpty || password != confirmPassword || authManager.isLoading)

            Button("Already have an account? Sign In") {
                dismiss()
            }
        }
        .padding()
        .navigationBarTitle("Register")
    }
}

// MARK: - Main Tab View
struct MainTabView: View {
    @EnvironmentObject var appState: AppState

    var body: some View {
        TabView(selection: $appState.selectedTab) {
            HomeView()
                .tabItem {
                    Label("Home", systemImage: "house.fill")
                }
                .tag(0)

            GamesView()
                .tabItem {
                    Label("Games", systemImage: "casino.fill")
                }
                .tag(1)

            TournamentsView()
                .tabItem {
                    Label("Tournaments", systemImage: "trophy.fill")
                }
                .tag(2)

            JackpotsView()
                .tabItem {
                    Label("Jackpots", systemImage: "star.fill")
                }
                .tag(3)

            WalletView()
                .tabItem {
                    Label("Wallet", systemImage: "creditcard.fill")
                }
                .tag(4)

            ProfileView()
                .tabItem {
                    Label("Profile", systemImage: "person.fill")
                }
                .tag(5)
        }
    }
}
