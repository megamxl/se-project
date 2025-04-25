//
//  ProfileView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.25.
//

import SwiftUI

struct ProfileView: View {

    @EnvironmentObject var authenticationViewModel: AuthenticationViewModel
    @EnvironmentObject var userViewModel: UserViewModel
    @EnvironmentObject var route: RouteObject

    @State private var isEditSheetPresented = false
    @State private var showDeleteDialog = false

    var body: some View {
        NavigationStack {
            Form {

                // MARK: - User Info

                if let user = userViewModel.user {
                    Section {
                        HStack {
                            if authenticationViewModel.isAdmin {
                                VStack(alignment: .center, spacing: 0){
                                    Text("")
                                        .faDuotoneSolid(size: 40)
                                    Text("ADMIN")
                                        .font(.caption)
                                        .fontWeight(.semibold)
                                        .foregroundStyle(.secondary)
                                }
                                    .frame(width: 60, height: 60)
                                    .foregroundStyle(.gray.gradient)
                            } else {
                                Image(systemName: "person.circle.fill")
                                    .resizable()
                                    .clipShape(Circle())
                                    .scaledToFill()
                                    .frame(width: 60, height: 60)
                                    .foregroundStyle(.gray.gradient)
                            }
                            VStack(alignment: .leading) {
                                Text(user.username ?? "")
                                Text(verbatim: user.email ?? "")
                                    .font(.caption)
                                    .foregroundStyle(.gray)
                            }
                            .padding(.leading)
                        }
                        .listRowSeparator(.hidden)
                    }
                }

                // MARK: - Profile

                Section("Profile") {
                    Button {
                        isEditSheetPresented = true
                    } label: {
                        Label {
                            Text("Edit Profile")
                        } icon: {
                            Text("")
                                .faDuotoneThin(size: 20)
                                .foregroundStyle(.blue)
                                .offset(x: .spacing2XS)
                        }
                    }
                    .tint(.primary)

                    Button {
                        showDeleteDialog = true
                    } label: {
                        Label {
                            Text("Delete Account")
                        } icon: {
                            Text("")
                                .faDuotoneThin(size: 20)
                                .foregroundStyle(.red)
                        }
                    }
                    .tint(.primary)
                    .confirmationDialog("Are you sure?", isPresented: $showDeleteDialog) {
                        Button("Delete your account?", role: .destructive) {
                            userViewModel.deleteUser()
                        }
                    }
                }
                
                // MARK: - Admin Management
                if authenticationViewModel.isAdmin {
                    Section("Admin") {
                        // Buttons push onto route.pathAdmin
                        Button { route.pathAdmin.append(.manageCars) } label: {
                            Label {
                                HStack {
                                    Text("Car Management")
                                    Spacer()
                                    Image(systemName: "chevron.forward")
                                        .font(.caption)
                                        .fontWeight(.bold)
                                        .foregroundStyle(.gray.opacity(0.6))
                                }
                            } icon: {
                                Text("")
                                    .faDuotoneThin(size: 20)
                                    .foregroundStyle(Color(uiColor: .darkGray))
                            }
                        }
                        .tint(.primary)

                        Button { route.pathAdmin.append(.manageBookings) } label: {
                            Label {
                                HStack {
                                    Text("Booking Management")
                                    Spacer()
                                    Image(systemName: "chevron.forward")
                                        .font(.caption)
                                        .fontWeight(.bold)
                                        .foregroundStyle(.gray.opacity(0.6))
                                }
                            } icon: {
                                Text("")
                                    .faDuotoneThin(size: 20)
                                    .foregroundStyle(Color(uiColor: .darkGray))
                            }
                        }
                        .tint(.primary)

                        Button { route.pathAdmin.append(.manageUsers) } label: {
                            Label {
                                HStack {
                                    Text("User Management")
                                    Spacer()
                                    Image(systemName: "chevron.forward")
                                        .font(.caption)
                                        .fontWeight(.bold)
                                        .foregroundStyle(.gray.opacity(0.6))
                                }
                            } icon: {
                                Text("")
                                    .faDuotoneThin(size: 20)
                                    .foregroundStyle(Color(uiColor: .darkGray))
                            }
                        }
                        .tint(.primary)
                    }
                }

                // MARK: - Licensing

                Section("Licensing") {
                    NavigationLink(destination: LicenseListView()) {
                        Label {
                            Text("Licensing")
                        } icon: {
                            Text("")
                                .faDuotoneThin(size: 20)
                                .foregroundStyle(.gray)
                        }
                    }
                    .tint(.primary)
                }

                // MARK: - Logout

                Section {
                    Button(role: .destructive) {
                        authenticationViewModel.logout()
                    } label: {
                        Label {
                            Text("Logout")
                                .fontWeight(.medium)
                        } icon: {
                            Text("")
                                .faDuotoneRegular(size: 20)
                                .foregroundStyle(.red)
                        }
                    }
                }
            }
            .navigationDestination(for: RouteAdmin.self) { routeCase in
                switch routeCase {
                case .manageCars:
                    AdminManagementView(
                        viewModel: CarManagementViewModel(),
                        title: "Cars",
                        rowContent: { car in
                            VStack(alignment: .leading) {
                                Text(car.brand ?? "")
                                Text(car.model ?? "")
                            }
                        },
                        createContent: { vm in CarCreateView(viewModel: vm) }
                    )
                case .manageBookings:
                    AdminManagementView(
                        viewModel: BookingManagementViewModel(),
                        title: "Bookings",
                        rowContent: { booking in
                            VStack(alignment: .leading) {
                                Text(booking.bookingId ?? "")
                                Text(booking.status ?? "")
                            }
                        },
                        createContent: { vm in BookingCreateView(viewModel: vm) }
                    )
                case .manageUsers:
                    AdminManagementView(
                        viewModel: UserManagementViewModel(),
                        title: "Users",
                        rowContent: { user in
                            VStack(alignment: .leading) {
                                Text(user.username ?? "")
                                Text(user.email ?? "")
                            }
                        },
                        createContent: { vm in UserCreateView(viewModel: vm) }
                    )
                }
            }
            .alert("Profile delete", isPresented: $userViewModel.showAlert) {
                Button("OK", role: .cancel) {
                    authenticationViewModel.logout()
                }
            } message: {
                Text(userViewModel.alertMessage)
            }
            .onAppear {
                userViewModel.getUserInfo()
            }
            .sheet(isPresented: $isEditSheetPresented) {
                let viewModel = EditProfileSheetViewModel(user: userViewModel.user)
                EditProfileSheetView(viewModel: viewModel)
                    .presentationDetents([.medium, .large])
                    .presentationDragIndicator(.visible)
            }
            
        }
    }
}

#Preview {
    ProfileView()
        .environmentObject(AuthenticationViewModel())
        .environmentObject(UserViewModel())
}
