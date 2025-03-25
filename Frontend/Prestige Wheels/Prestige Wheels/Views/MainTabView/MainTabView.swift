//
//  MainTabView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 24.03.25.
//

import SwiftUI

struct MainTabView: View {
    @EnvironmentObject var loginViewModel: LoginViewModel
    
    var body: some View {
        TabView {
            Tab("Find a Car", systemImage: "car.fill") {
                FindCarView()
            }
            Tab("Bookings", systemImage: "book.pages.fill") {
                MyBookingsView()
            }
            Tab("Profile", systemImage: "person.crop.circle.fill") {
                ProfileView()
            }
        }
//        .fullScreenCover(isPresented: Binding(
//            get: { !loginViewModel.isLoggedIn },
//            set: { loginViewModel.isLoggedIn = !$0 }
//        )) {
//            LoginView()
//        }
    }
}

#Preview {
    MainTabView()
        .environmentObject(LoginViewModel())
}
