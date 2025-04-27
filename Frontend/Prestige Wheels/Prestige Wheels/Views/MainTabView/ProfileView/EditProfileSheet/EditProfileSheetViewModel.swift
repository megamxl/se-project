//
//  EditProfileSheetViewModel.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 20.04.25.
//

import Foundation
import OpenAPIClient

class EditProfileSheetViewModel: ObservableObject {
    @Published var user: OpenAPIClientAPI.User?
    @Published var username: String = ""
    @Published var email: String = ""
    @Published var password: String = ""
    @Published var isSaving = false
    @Published var showAlert = false
    @Published var alertMessage: String = ""
    
    public init(user: OpenAPIClientAPI.User?) {
        self.user = user
    }
    
    func saveChanges() {
        let request = OpenAPIClientAPI.UserMutation(username: username, email: email, password: password)
        
        OpenAPIClientAPI.UserAPI.updateUser(userMutation: request) { [weak self] result, error  in
            guard let self = self else { return }
            
            self.isSaving = false
            
            if let error = error {
                alertMessage = "Update failed: \(error.localizedDescription)"
            } else {
                alertMessage = "Update erfolgreich"
            }
            showAlert = true
        }
    }
}
