//
//  LoginViewModel.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.2025.
//

import Combine
import Foundation
import OSLog

@MainActor
class LoginViewModel: ObservableObject {
    static let shared = LoginViewModel()

    // MARK: - Input Fields

    @Published var username: String = ""
    @Published var password: String = ""

    // MARK: - Output Properties

    @Published var errorMessage: String?
    @Published var isLoggedIn: Bool = false


    init() {
        
    }

    // MARK: - Login Method

    func login() {
        errorMessage = nil
    }

    // MARK: - Logout Method

    func logout() {
    }
}
