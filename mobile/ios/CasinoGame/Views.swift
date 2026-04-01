import SwiftUI

struct LoginView: View {
    @Binding var isLoggedIn: Bool
    @State private var email = ""
    @State private var password = ""
    @State private var rememberMe = false
    @State private var showError = false
    @State private var isLoading = false

    private func loginUser() {
        guard !email.isEmpty, !password.isEmpty else {
            showError = true
            return
        }

        guard email.contains("@"), password.count >= 4 else {
            showError = true
            return
        }

        isLoading = true

        Task {
            do {
                let _ = try await APIClient.shared.login(email: email, password: password)
                await MainActor.run {
                    isLoading = false
                    isLoggedIn = true
                    if rememberMe {
                        UserDefaults.standard.set(true, forKey: "rememberMe")
                    }
                }
            } catch {
                await MainActor.run {
                    isLoading = false
                    showError = true
                }
            }
        }
    }

    var body: some View {
        NavigationStack {
            VStack(spacing: 30) {
                Spacer()

                VStack(spacing: 10) {
                    Image(systemName: "gamecontroller.fill")
                        .font(.system(size: 60))
                        .foregroundColor(.orange)

                    Text("Casino Game")
                        .font(.largeTitle)
                        .fontWeight(.bold)

                    Text("Play & Win Big")
                        .font(.subheadline)
                        .foregroundColor(.secondary)
                }

                VStack(spacing: 20) {
                    VStack(alignment: .leading, spacing: 8) {
                        Text("Email")
                            .font(.caption)
                            .foregroundColor(.secondary)

                        TextField("Enter your email", text: $email)
                            .textFieldStyle(RoundedTextFieldStyle())
                            .textContentType(.emailAddress)
                            .autocapitalization(.none)
                    }

                    VStack(alignment: .leading, spacing: 8) {
                        Text("Password")
                            .font(.caption)
                            .foregroundColor(.secondary)

                        SecureField("Enter your password", text: $password)
                            .textFieldStyle(RoundedTextFieldStyle())
                            .textContentType(.password)
                    }

                    HStack {
                        Toggle("Remember me", isOn: $rememberMe)
                            .tint(.orange)

                        Spacer()

                        Button("Forgot Password?") {
                        }
                        .font(.subheadline)
                        .foregroundColor(.orange)
                    }

                    Button {
                        loginUser()
                    } label: {
                        Text("Login")
                            .font(.headline)
                            .foregroundColor(.white)
                            .frame(maxWidth: .infinity)
                            .padding()
                            .background(Color.orange)
                            .cornerRadius(12)
                    }

                    HStack {
                        Text("Don't have an account?")
                            .foregroundColor(.secondary)

                        Button("Register") {
                        }
                        .foregroundColor(.orange)
                        .fontWeight(.semibold)
                    }
                }
                .padding(.horizontal, 30)

                Spacer()
            }
            .alert("Login Failed", isPresented: $showError) {
                Button("OK", role: .cancel) { }
            } message: {
                Text("Invalid email or password. Please try again.")
            }
        }
    }
}

struct RoundedTextFieldStyle: TextFieldStyle {
    func _body(configuration: TextField<Self._Label>) -> some View {
        configuration
            .padding()
            .background(Color(.systemGray6))
            .cornerRadius(12)
    }
}

struct GamesView: View {
    @State private var searchText = ""
    @State private var selectedCategory = "All"

    let categories = ["All", "Slots", "Card Games", "Table Games", "Live Casino", "Jackpots"]

    var body: some View {
        NavigationStack {
            VStack(spacing: 0) {
                HStack {
                    Image(systemName: "magnifyingglass")
                        .foregroundColor(.secondary)

                    TextField("Search games", text: $searchText)
                }
                .padding()
                .background(Color(.systemGray6))
                .cornerRadius(12)
                .padding()

                ScrollView(.horizontal, showsIndicators: false) {
                    HStack(spacing: 12) {
                        ForEach(categories, id: \.self) { category in
                            CategoryButton(
                                title: category,
                                isSelected: selectedCategory == category
                            ) {
                                selectedCategory = category
                            }
                        }
                    }
                    .padding(.horizontal)
                }

                ScrollView {
                    LazyVGrid(columns: [
                        GridItem(.flexible(), spacing: 15),
                        GridItem(.flexible(), spacing: 15)
                    ], spacing: 15) {
                        GameCard(name: "Starburst", provider: "NetEnt", imageColor: .purple)
                        GameCard(name: "Gonzo's Quest", provider: "NetEnt", imageColor: .blue)
                        GameCard(name: "Book of Dead", provider: "Play'n GO", imageColor: .orange)
                        GameCard(name: "Mega Moolah", provider: "Microgaming", imageColor: .yellow)
                        GameCard(name: "Blackjack", provider: "Evolution", imageColor: .green)
                        GameCard(name: "Roulette", provider: "Evolution", imageColor: .red)
                    }
                    .padding()
                }
            }
            .navigationTitle("Games")
        }
    }
}
