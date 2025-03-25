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
            Section {
                Text("Prestige Wheels")
                    .font(.largeTitle)
                    .fontWeight(.semibold)
                    .hAlign(.center)
            }
            .listRowBackground(Color.clear)

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
                Button {} label: {
                    Text("Problems with the registration?")
                        .font(.caption)
                        .foregroundStyle(.gray)
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
