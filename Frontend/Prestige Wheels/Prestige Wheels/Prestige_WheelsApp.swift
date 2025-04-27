//
//  Prestige_WheelsApp.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 24.03.25.
//

import SwiftUI
import OpenAPIClient
import GoogleMaps

private func loadRocketSimConnect() {
    #if DEBUG
    guard (Bundle(path: "/Applications/RocketSim.app/Contents/Frameworks/RocketSimConnectLinker.nocache.framework")?.load() == true) else {
        print("Failed to load linker framework")
        return
    }
    print("RocketSim Connect successfully linked")
    #endif
}

@main
struct Prestige_WheelsApp: App {
    init() {
        loadRocketSimConnect()
        configureSessionWithCookies()
        
        let customDateFormatter = DateFormatter()
        customDateFormatter.calendar = Calendar(identifier: .iso8601)
        customDateFormatter.locale = Locale(identifier: "en_US_POSIX")
        customDateFormatter.timeZone = TimeZone(secondsFromGMT: 0)
        customDateFormatter.dateFormat = "yyyy-MM-dd"

        CodableHelper.dateFormatter = customDateFormatter
        
        GMSServices.provideAPIKey("API_KEY")
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
