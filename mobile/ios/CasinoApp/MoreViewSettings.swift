import SwiftUI

// MARK: - Profile View
struct ProfileView: View {
    @EnvironmentObject var authManager: AuthManager
    @State private var profile: UserProfile?
    @State private var isLoading = true

    var body: some View {
        NavigationStack {
            List {
                Section {
                    HStack {
                        Image(systemName: "person.circle.fill")
                            .font(.system(size: 60))
                            .foregroundColor(.purple)

                        VStack(alignment: .leading, spacing: 4) {
                            Text(profile?.username ?? "Player")
                                .font(.headline)
                            Text(profile?.email ?? "")
                                .font(.subheadline)
                                .foregroundColor(.secondary)
                        }
                    }
                    .padding(.vertical, 8)
                }

                Section("Account") {
                    NavigationLink("Edit Profile") { }
                    NavigationLink("Security") { }
                    NavigationLink("Notifications") { }
                }

                Section("Support") {
                    NavigationLink("Help & Support") { }
                    NavigationLink("About") { }
                }

                Section {
                    Button("Logout") {
                        Task {
                            await authManager.logout()
                        }
                    }
                    .foregroundColor(.red)
                }
            }
            .navigationTitle("Profile")
        }
        .task {
            await loadData()
        }
    }

    private func loadData() async {
        isLoading = true
        do {
            profile = try await APIClient.shared.getProfile()
        } catch {
            print(error)
        }
        isLoading = false
    }
}
