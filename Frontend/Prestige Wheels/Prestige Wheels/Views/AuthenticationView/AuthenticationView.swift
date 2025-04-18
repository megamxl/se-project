//
//  AuthenticationView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.2025.
//

import SwiftUI

struct AuthenticationView: View {
    @EnvironmentObject var authenticationViewModel: AuthenticationViewModel
    
    enum FocusedField {
        case username, password
    }
    
    @FocusState private var focusedField: FocusedField?
    
    var body: some View {
        VStack {
            switch authenticationViewModel.authenticationMethod {
            case .login:
                loginView
            case .register:
                registerView
            }
        }
        .onAppear {
            focusedField = .username
        }
    }
    
    // MARK: - Login View
    
    var loginView: some View {
        Form {
            Image(uiImage: Bundle.main.icon ?? UIImage())
                .resizable()
                .scaledToFill()
                .clipped()
                .clipShape(RoundedRectangle(cornerRadius: 22))
                .frame(width: 100, height: 100, alignment: .center)
                .listRowBackground(Color.clear)
                .hAlign(.center)
            
            Section {
                TextField("E-Mail", text: $authenticationViewModel.email)
                    .disableAutocorrection(true)
                    .textContentType(.username)
                    .focused($focusedField, equals: .username)
                    .onSubmit {
                        focusedField = .password
                    }
                    .submitLabel(.next)
                SecureField("Password", text: $authenticationViewModel.password)
                    .focused($focusedField, equals: .password)
                    .onSubmit {
                        if !authenticationViewModel.username.isEmpty || !authenticationViewModel.password.isEmpty {
                            authenticationViewModel.login()
                        }
                    }
                    .submitLabel(.go)
            } header: {
                Text("Welcome, please login...")
                    .font(.headline)
                    .listRowInsets(.init())
                    .textCase(nil)
                    .padding(.bottom, .spacingS)
            }
            
            Section {
                Button {
                    authenticationViewModel.login()
                } label: {
                    Text("Login")
                        .font(.headline)
                        .hAlign(.center)
                }
                .disabled(authenticationViewModel.email.isEmpty || authenticationViewModel.password.isEmpty)
            } footer: {
                Button {
                    authenticationViewModel.authenticationMethod = .register
                } label: {
                    Text("or register now")
                        .font(.callout)
                        .fontWeight(.medium)
                }
                .hAlign(.center)
                .padding(.top, .spacingS)
            }
        }
    }
    
    // MARK: - Register View
    
    var registerView: some View {
        Form {
            Image(uiImage: Bundle.main.icon ?? UIImage())
                .resizable()
                .scaledToFill()
                .clipped()
                .clipShape(RoundedRectangle(cornerRadius: 22))
                .frame(width: 100, height: 100, alignment: .center)
                .listRowBackground(Color.clear)
                .hAlign(.center)
            
            Section {
                TextField("Name", text: $authenticationViewModel.username)
                    .disableAutocorrection(true)
                    .textContentType(.username)
                    .focused($focusedField, equals: .username)
                    .onSubmit {
                        focusedField = .password
                    }
                    .submitLabel(.next)
                TextField("E-Mail", text: $authenticationViewModel.email)
                    .disableAutocorrection(true)
                    .textContentType(.username)
                    .focused($focusedField, equals: .username)
                    .onSubmit {
                        focusedField = .password
                    }
                    .submitLabel(.next)
                SecureField("Password", text: $authenticationViewModel.password)
                    .focused($focusedField, equals: .password)
                    .onSubmit {
                        if !authenticationViewModel.username.isEmpty || !authenticationViewModel.password.isEmpty {
                            authenticationViewModel.login()
                        }
                    }
                    .submitLabel(.go)
            } header: {
                Text("Welcome, please register...")
                    .font(.headline)
                    .listRowInsets(.init())
                    .textCase(nil)
                    .padding(.bottom, .spacingS)
            }
            
            Section {
                Button {
                    authenticationViewModel.register()
                } label: {
                    Text("Register")
                        .font(.headline)
                        .hAlign(.center)
                }
                .disabled(authenticationViewModel.username.isEmpty || authenticationViewModel.password.isEmpty ||
                          authenticationViewModel.username.isEmpty)
            } footer: {
                Button {
                    authenticationViewModel.authenticationMethod = .login
                } label: {
                    Text("or login now")
                        .font(.callout)
                        .fontWeight(.medium)
                }
                .hAlign(.center)
                .padding(.top, .spacingS)
            }
        }
    }
}

#Preview {
    AuthenticationView()
        .environmentObject(AuthenticationViewModel())
}
