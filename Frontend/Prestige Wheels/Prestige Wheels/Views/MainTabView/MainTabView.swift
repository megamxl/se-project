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
                NavigationStack(path: $route.pathFindCar) {
                    FindCarView()
                        .navigationDestination(for: RouteFindCar.self) { route in view(for: route) }
                }
            }
            Tab("Bookings", systemImage: "book.pages.fill") {
                NavigationStack(path: $route.pathMyBookings) {
                    MyBookingsView()
                        .navigationDestination(for: RouteMyBookings.self) { route in view(for: route) }
                }
            }
            Tab("Profile", systemImage: "person.crop.circle.fill") {
                NavigationStack(path: $route.pathAdmin) {
                    ProfileView()
                }
            }
        }
        .environmentObject(route)
        .fullScreenCover(isPresented: Binding(
            get: { !loginViewModel.isLoggedIn },
            set: { loginViewModel.isLoggedIn = !$0 }
        )) {
            AuthenticationView()
        }
    }


    @ViewBuilder
    func view(for route: RouteFindCar) -> some View {
        switch route {
        case .findCarDetailView(let car, let currency, let from, let to):
            let viewModel = FindCarDetailViewModel(car: car, currency: currency, from: from, to: to)
            FindCarDetailView(viewModel: viewModel)
        }
    }
    
    @ViewBuilder
    func view(for route: RouteMyBookings) -> some View {
        switch route {
        case .bookingDetailView(booking: let booking, car: let car):
            let viewModel = MyBookingDetailViewModel(booking: booking, car: car)
            MyBookingDetailView(viewModel: viewModel)
        }
    }
}

#Preview {
    MainTabView()
        .environmentObject(AuthenticationViewModel())
}
