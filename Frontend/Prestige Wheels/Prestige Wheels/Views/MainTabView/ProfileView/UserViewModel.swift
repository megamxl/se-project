//
//  UserViewModel.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 10.04.25.
//

import Foundation
import OpenAPIClient
import OSLog

class UserViewModel: ObservableObject {
    @Published var user: OpenAPIClientAPI.User?
    @Published var showAlert = false
    @Published var alertMessage = ""
    
    func getUserInfo() {
        OpenAPIClientAPI.UserAPI.getUsers(apiResponseQueue: DispatchQueue.main) { [weak self] (user, error) in
            guard let self = self else { return }
            
            if let error = error {
                Logger.backgroundProcessing.error("\(error.localizedDescription)")
            } else if let user = user {
                self.user = user
            }
        }
    }
    
    func getAllUsers() {

    }
    
    func updateUser() {

    }
    
    func deleteUser() {
        guard let userId = user?.id else { return }
        
        OpenAPIClientAPI.UserAPI.deleteUser(id: userId) { [weak self] (result, error)  in
            guard let self = self else { return }
            
            if let error {
                alertMessage = "Delete failed: \(error.localizedDescription)"
            } else {
                alertMessage = "Delete success"
            }
            showAlert = true
        }
    }
}
