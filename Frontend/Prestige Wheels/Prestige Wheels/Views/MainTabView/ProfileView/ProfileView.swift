//
//  ProfileView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.25.
//

import SwiftUI

struct ProfileView: View {
    
    @EnvironmentObject var loginViewModel: LoginViewModel
    @EnvironmentObject var userViewModel: UserViewModel
    
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
                    .swipeActions(edge: .trailing) {
                        Button {
                            isEditSheetPresented = true
                        } label: {
                            Label("Edit", systemImage: "slider.horizontal.3")
                        }
                    }
                    .contextMenu {
                        Button {
                            isEditSheetPresented = true
                        } label: {
                            Label("Edit", systemImage: "slider.horizontal.3")
                        }
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
                        loginViewModel.logout()
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
            .sheet(isPresented: $isEditSheetPresented) {
                // reload user Data?
            } content: {
                Text("Edit Profile")
                    .presentationDetents([.medium, .large])
                    .presentationDragIndicator(.visible)
            }
            .onAppear {
                userViewModel.getUserInfo()
            }
        }
    }
}

#Preview {
    ProfileView()
        .environmentObject(LoginViewModel())
        .environmentObject(UserViewModel())
}
