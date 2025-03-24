//
//  MainTabView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 24.03.25.
//

import SwiftUI

struct MainTabView: View {
    var body: some View {
        TabView {
            Tab("Find a Car", systemImage: "car.fill") {
                FindCarView()
            }
            Tab("Bookings", systemImage: "book.pages.fill") {
                EmptyView()
            }
            Tab("Profile", systemImage: "person.crop.circle.fill") {
                EmptyView()
            }
        }
    }
}

#Preview {
    MainTabView()
}
