//
//  ContentView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 24.03.25.
//

import SwiftUI

struct ContentView: View {
    @StateObject private var authenticationViewModel = AuthenticationViewModel()
    @StateObject private var userViewModel = UserViewModel()
    @StateObject private var bookingViewModel = BookingViewModel()
    
    var body: some View {
        MainTabView()
            .environmentObject(authenticationViewModel)
            .environmentObject(userViewModel)
            .environmentObject(bookingViewModel)
    }
}

#Preview {
    ContentView()
}
