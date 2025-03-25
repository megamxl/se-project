//
//  ContentView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 24.03.25.
//

import SwiftUI

struct ContentView: View {
    @StateObject private var loginViewModel = LoginViewModel()
    
    var body: some View {
        MainTabView()
            .environmentObject(loginViewModel)
    }
}

#Preview {
    ContentView()
}
