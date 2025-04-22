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
    
    var body: some View {
        NavigationStack {
            Form {
                
                // MARK: - User Info
                
                if let user = userViewModel.user {
                    Section {
                        HStack {
                            Image(systemName: "person.circle.fill")
                                .resizable()
                                .clipShape(Circle())
                                .scaledToFill()
                                .frame(width: 60, height: 60)
                                .foregroundStyle(.gray.gradient)
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
                        userViewModel.deleteUser()
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
                }
                
                // MARK: - Admin
                
                if true {
                    Section("Admin") {
                        Button {
                            
                        } label: {
                            Label {
                                Text("Car Management")
                            } icon: {
                                Text("")
                                    .faDuotoneThin(size: 20)
                                    .foregroundStyle(Color(uiColor: .darkGray))
                            }
                        }
                        .tint(.primary)
                        Button {
                            
                        } label: {
                            Label {
                                Text("Booking Management")
                            } icon: {
                                Text("")
                                    .faDuotoneThin(size: 20)
                                    .foregroundStyle(Color(uiColor: .darkGray))
                            }
                        }
                        .tint(.primary)
                        Button {
                            
                        } label: {
                            Label {
                                Text("User Management")
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
                    Button{
                    } label: {
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
