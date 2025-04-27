//
//  EditProfileSheetView.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 20.04.25.
//

import SwiftUI

struct EditProfileSheetView: View {
    @ObservedObject var viewModel: EditProfileSheetViewModel
    @Environment(\.dismiss) private var dismiss
    
    var body: some View {
        NavigationStack {
            Form {
                Section(header: Text("Username")) {
                    TextField("Username", text: $viewModel.username)
                        .autocapitalization(.none)
                }
                
                Section(header: Text("Email")) {
                    TextField("Email", text: $viewModel.email)
                        .keyboardType(.emailAddress)
                        .autocapitalization(.none)
                }
                
                Section(header: Text("Password")) {
                    SecureField("New Password", text: $viewModel.password)
                }
            }
            .navigationTitle("Edit Profile")
            .toolbar {
                ToolbarItem(placement: .cancellationAction) {
                    Button("Cancel") {
                        dismiss()
                    }
                }
                ToolbarItem(placement: .confirmationAction) {
                    if viewModel.isSaving {
                        ProgressView()
                    } else {
                        Button {
                            viewModel.saveChanges()
                        } label: {
                            Text("Save")
                        }
                        .disabled(viewModel.username.isEmpty || viewModel.email.isEmpty || viewModel.password.isEmpty)
                    }
                }
            }
            .alert("Update failed", isPresented: $viewModel.showAlert) {
                Button("OK", role: .cancel) {
                    dismiss()
                }
            } message: {
                Text(viewModel.alertMessage)
            }
        }
        .presentationDetents([.medium, .large])
        .presentationDragIndicator(.visible)
    }
}

#Preview {
    EditProfileSheetView(viewModel: .init(user: .init()))
}
