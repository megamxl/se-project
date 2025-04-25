//
//  UserManagementViewModel.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 22.04.25.
//

import Foundation
import Combine
import OpenAPIClient

class UserManagementViewModel: ObservableObject, ManagementViewModelProtocol {
    @Published var items: [OpenAPIClientAPI.User] = []
    @Published var isLoading = false
    @Published var errorMessage: String? = nil

    func fetchItems() {
        isLoading = true
        OpenAPIClientAPI.AdminAPI.getAllUsers(
            apiResponseQueue: DispatchQueue.main
        ) { [weak self] users, error in
            guard let self = self else { return }
            self.isLoading = false
            if let error = error {
                self.errorMessage = error.localizedDescription
            } else {
                self.items = users?.compactMap { $0 } ?? []
            }
        }
    }

    func delete(item: OpenAPIClientAPI.User) {
        guard let id = item.id else { return }
        isLoading = true
        OpenAPIClientAPI.UserAPI.deleteUser(
            id: id,
            apiResponseQueue: DispatchQueue.main
        ) { [weak self] _, _ in
            guard let self = self else { return }
            self.isLoading = false
            self.fetchItems()
        }
    }

    func add(mutation: OpenAPIClientAPI.UserMutation, completion: @escaping () -> Void) {
        isLoading = true
        OpenAPIClientAPI.UserAPI.registerUser(
            userMutation: mutation,
            apiResponseQueue: DispatchQueue.main
        ) { [weak self] _, _ in
            guard let self = self else { return }
            self.isLoading = false
            self.fetchItems()
            completion()
        }
    }
}
