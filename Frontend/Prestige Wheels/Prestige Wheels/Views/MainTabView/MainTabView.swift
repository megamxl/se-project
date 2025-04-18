//
//  MainTabView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 24.03.25.
//

import SwiftUI

struct MainTabView: View {
    @EnvironmentObject var loginViewModel: AuthenticationViewModel
    @StateObject var route = RouteObject()
    
    var body: some View {
        TabView {
            Tab("Find a Car", systemImage: "car.fill") {
                NavigationStack(path: $route.path) {
                    FindCarView()
                        .navigationDestination(for: Route.self) { route in view(for: route) }
                }
                .environmentObject(route)
            }
            Tab("Bookings", systemImage: "book.pages.fill") {
                MyBookingsView()
            }
            Tab("Profile", systemImage: "person.crop.circle.fill") {
                ProfileView()
            }
        }
        .fullScreenCover(isPresented: Binding(
            get: { !loginViewModel.isLoggedIn },
            set: { loginViewModel.isLoggedIn = !$0 }
        )) {
            AuthenticationView()
        }
    }


    @ViewBuilder
    func view(for route: Route) -> some View {
        switch route {
        case .findCarDetailView(let car, let currency, let from, let to):
            let viewModel = FindCarDetailViewModel(car: car, currency: currency, from: from, to: to)
            FindCarDetailView(viewModel: viewModel)
        }
    }
}

#Preview {
    MainTabView()
        .environmentObject(AuthenticationViewModel())
}
