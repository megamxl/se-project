//
//  UserCreateView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 22.04.25.
//

import SwiftUI
import OpenAPIClient

struct UserCreateView: View {
    @Environment(\.dismiss) private var dismiss
    @ObservedObject var viewModel: UserManagementViewModel
    @State private var username = ""
    @State private var email = ""
    @State private var password = ""

    var body: some View {
        NavigationStack {
            Form {
                TextField("Username", text: $username)
                TextField("Email", text: $email)
                    .keyboardType(.emailAddress)
                SecureField("Password", text: $password)
            }
            .navigationTitle("Add User")
            .toolbar {
                ToolbarItem(placement: .cancellationAction) {
                    Button("Cancel") { dismiss() }
                }
                ToolbarItem(placement: .confirmationAction) {
                    Button("Save") {
                        let mutation = OpenAPIClientAPI.UserMutation(
                            username: username,
                            email: email,
                            password: password
                        )
                        viewModel.add(mutation: mutation) { dismiss() }
                    }
                }
            }
        }
    }
}

struct UserCreateView_Previews: PreviewProvider {
    static var previews: some View {
        UserCreateView(viewModel: UserManagementViewModel())
    }
}
