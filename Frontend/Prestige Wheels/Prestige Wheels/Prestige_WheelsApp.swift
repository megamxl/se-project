//
//  Prestige_WheelsApp.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 24.03.25.
//

import SwiftUI
import OpenAPIClient

@main
struct Prestige_WheelsApp: App {
    init() {
        configureSessionWithCookies()
        
        let customDateFormatter = DateFormatter()
        customDateFormatter.calendar = Calendar(identifier: .iso8601)
        customDateFormatter.locale = Locale(identifier: "en_US_POSIX")
        customDateFormatter.timeZone = TimeZone(secondsFromGMT: 0)
        customDateFormatter.dateFormat = "yyyy-MM-dd"

        CodableHelper.dateFormatter = customDateFormatter
    }

    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }

    func configureSessionWithCookies() {
        let config = URLSessionConfiguration.default
        config.httpCookieStorage = HTTPCookieStorage.shared
        config.httpShouldSetCookies = true
        config.httpCookieAcceptPolicy = .always

        let session = URLSession(configuration: config)

        OpenAPIClientAPI.requestBuilderFactory = CustomRequestBuilderFactory(session: session)
    }
}
