//
//  LoginView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.2025.
//

import SwiftUI

struct LoginView: View {
    @EnvironmentObject var loginViewModel: LoginViewModel

    enum FocusedField {
        case username, password
    }

    @FocusState private var focusedField: FocusedField?

    var body: some View {
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
                TextField("Username", text: $loginViewModel.username)
                    .disableAutocorrection(true)
                    .textContentType(.username)
                    .focused($focusedField, equals: .username)
                    .onSubmit {
                        focusedField = .password
                    }
                    .submitLabel(.next)
                SecureField("Password", text: $loginViewModel.password)
                    .focused($focusedField, equals: .password)
                    .onSubmit {
                        if !loginViewModel.username.isEmpty || !loginViewModel.password.isEmpty {
                            loginViewModel.login()
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
                    loginViewModel.login()
                } label: {
                    Text("Login")
                        .font(.headline)
                        .hAlign(.center)
                }
                .disabled(loginViewModel.username.isEmpty || loginViewModel.password.isEmpty)
            } footer: {
                Button {
                    
                } label: {
                    Text("or register now")
                        .font(.callout)
                        .fontWeight(.medium)
                }
                .hAlign(.center)
                .padding(.top, .spacingS)
            }
        }
        .onAppear {
            focusedField = .username
        }
    }
}

#Preview {
    LoginView()
        .environmentObject(LoginViewModel())
}
