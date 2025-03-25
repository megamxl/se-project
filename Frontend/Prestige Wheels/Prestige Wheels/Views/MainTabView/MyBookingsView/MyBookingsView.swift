//
//  MyBookingsView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.25.
//

import SwiftUI

struct MyBookingsView: View {
    var body: some View {
        NavigationStack {
            ContentUnavailableView("No bookings yet", systemImage: "car.2.fill", description: Text("You haven't booked any cars yet."))
                .navigationTitle("My Bookings")
        }
    }
}

#Preview {
    MyBookingsView()
}
