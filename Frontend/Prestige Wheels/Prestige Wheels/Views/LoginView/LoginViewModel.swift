//
//  LoginViewModel.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.2025.
//

import Combine
import Foundation
import OSLog
import OpenAPIClient

@MainActor
class LoginViewModel: ObservableObject {
    // MARK: - Input Fields

    @Published var username: String = "john.doe@example.com"
    @Published var password: String = "password123"

    // MARK: - Output Properties

    @Published var errorMessage: String?
    @Published var isLoggedIn: Bool = false

    // MARK: - Init

    init() {
        validateSession()
    }

    // MARK: - Login Method

    func login() {
        errorMessage = nil

        let login = OpenAPIClientAPI.LoginRequest(email: username, password: password)

        OpenAPIClientAPI.UserAPI.login(loginRequest: login) { _, error in
            if let error = error {
                Logger.authentication.info("❌ Login failed: \(error.localizedDescription)")
                self.isLoggedIn = false
                self.errorMessage = "Login fehlgeschlagen: \(error.localizedDescription)"
            } else {
                Logger.authentication.info("✅ Login success")
                self.isLoggedIn = true
                // Session-Cookie wird automatisch gespeichert
            }
        }
    }

    // MARK: - Logout Method

    func logout() {
        HTTPCookieStorage.shared.cookies?.forEach {
            HTTPCookieStorage.shared.deleteCookie($0)
        }
        self.isLoggedIn = false
        self.username = ""
        self.password = ""
    }

    // MARK: - Session Validierung

    func validateSession() {
        #warning("change endpoint to use for session validation")
        OpenAPIClientAPI.UserAPI.getUsers { user, error in
            if let _ = user {
                Logger.authentication.info("✅ Gültige Session erkannt")
                self.isLoggedIn = true
            } else {
                Logger.authentication.info("❌ Keine gültige Session")
                self.isLoggedIn = false
            }
        }
    }
}
