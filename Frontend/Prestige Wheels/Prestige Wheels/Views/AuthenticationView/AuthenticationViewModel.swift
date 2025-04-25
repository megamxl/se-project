//
//  AuthenticationViewModel.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.2025.
//

import Combine
import Foundation
import OSLog
import OpenAPIClient
import JWTDecode

@MainActor
class AuthenticationViewModel: ObservableObject {
    // MARK: - Input Fields

    @Published var username: String = ""
    @Published var email: String = "john@example.com"
    @Published var password: String = "securePass123"

    // MARK: - Output Properties

    @Published var errorMessage: String?
    @Published var isLoggedIn: Bool = false
    @Published var role: String?         // e.g. "user" or "admin"
    @Published var isAdmin: Bool = false // drives UI

    // MARK: - Properties

    @Published var authenticationMethod: AuthenticationMethod = .login

    enum AuthenticationMethod {
        case login
        case register
    }

    // MARK: - Init

    init() {
        validateSession()
    }

    // MARK: - Login Method

    func login() {
        errorMessage = nil
        let loginReq = OpenAPIClientAPI.LoginRequest(email: email, password: password)

        OpenAPIClientAPI.UserAPI.login(loginRequest: loginReq) { _, error in
            if let error = error {
                Logger.authentication.info("❌ Login failed: \(error.localizedDescription)")
                self.isLoggedIn = false
                self.errorMessage = "Login fehlgeschlagen: \(error.localizedDescription)"
            } else {
                Logger.authentication.info("✅ Login success")
                self.isLoggedIn = true

                // 1) Pull the JWT out of your cookies:
                guard
                    let cookies = HTTPCookieStorage.shared.cookies,
                    let jwtCookie = cookies.first(where: { $0.name == "jwt" })
                else {
                    Logger.authentication.error("❌ JWT cookie not found")
                    return
                }
                let token = jwtCookie.value

                // 2) Decode with JWTDecode.swift:
                do {
                    let jwt = try decode(jwt: token)
                    let r = jwt.claim(name: "roles").string
                    self.role = r
                    self.isAdmin = (r == "admin")
                    Logger.authentication.info("User roles: \(r ?? "nil")")
                } catch {
                    Logger.authentication.error("JWT decode error: \(error)")
                }
            }
        }
    }

    // MARK: - Register Method

    func register() {
        errorMessage = nil
        let registerReq = OpenAPIClientAPI.UserMutation(username: username, email: email, password: password)

        OpenAPIClientAPI.UserAPI.registerUser(userMutation: registerReq) { _, error in
            if let error = error {
                Logger.authentication.info("❌ Register failed: \(error.localizedDescription)")
                self.isLoggedIn = false
                self.errorMessage = "Registrierung fehlgeschlagen: \(error.localizedDescription)"
            } else {
                Logger.authentication.info("✅ Registration success")
                self.isLoggedIn = true
            }
        }
    }

    // MARK: - Logout Method

    func logout() {
        HTTPCookieStorage.shared.cookies?.forEach {
            HTTPCookieStorage.shared.deleteCookie($0)
        }
        self.isLoggedIn = false
        self.email = ""
        self.password = ""
        self.role = nil
        self.isAdmin = false
    }

    // MARK: - Session Validation

    func validateSession() {
        #warning("change endpoint to use for session validation")
        OpenAPIClientAPI.UserAPI.getUsers { user, error in
            if let _ = user {
                Logger.authentication.info("✅ Gültige Session erkannt")
                self.isLoggedIn = true

                // — now also decode the JWT from the cookie —
                guard
                    let cookies = HTTPCookieStorage.shared.cookies,
                    let jwtCookie = cookies.first(where: { $0.name == "jwt" })
                else {
                    Logger.authentication.error("❌ JWT cookie not found during session validate")
                    return
                }
                let token = jwtCookie.value

                do {
                    let jwt = try decode(jwt: token)
                    let r = jwt.claim(name: "roles").string
                    self.role = r
                    self.isAdmin = (r == "admin")
                    Logger.authentication.info("Session role: \(r ?? "nil")")
                } catch {
                    Logger.authentication.error("JWT decode error during session validate: \(error)")
                    // if decode fails, you may want to forcibly log out:
                    self.role = nil
                    self.isAdmin = false
                    self.isLoggedIn = false
                }

            } else {
                Logger.authentication.info("❌ Keine gültige Session")
                self.isLoggedIn = false
                self.role = nil
                self.isAdmin = false
            }
        }
    }
}
