//
//  MyBookingsView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.25.
//

import SwiftUI

struct MyBookingsView: View {
    
    @EnvironmentObject var bookingViewModel: BookingViewModel
    
    var body: some View {
        NavigationStack {
            ContentUnavailableView("View not implemented", systemImage: "car.2.fill", description: Text("You haven't booked any cars yet."))
                .navigationTitle("My Bookings")
        }
        .onAppear {
            bookingViewModel.getBookingsForUser()
        }
    }
}

#Preview {
    MyBookingsView()
        .environmentObject(BookingViewModel())
}
