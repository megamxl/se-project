//
//  MyBookingsView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.25.
//

import SwiftUI

struct MyBookingsView: View {
    @EnvironmentObject var bookingViewModel: BookingViewModel
    @EnvironmentObject var route: RouteObject
    
    var body: some View {
        VStack {
            if bookingViewModel.bookings.isEmpty {
                noBookingsFound
            } else {
                bookingList
            }
        }
        .navigationTitle("My Bookings")
        .onAppear {
            bookingViewModel.getBookingsForUser()
        }
    }
    
    // MARK: - No BookingsFound
    
    var noBookingsFound: some View {
        ContentUnavailableView("No Result", systemImage: "car.2.fill", description: Text("You haven't booked any cars yet."))
            .navigationTitle("My Bookings")
    }
    
    // MARK: - Booking List
    
    var bookingList: some View {
        List(bookingViewModel.bookings, id: \.bookingId) { booking in
            Button {
                route.pathMyBookings.append(.bookingDetailView(booking: booking))
            } label: {
                BookingRow(booking: booking)
                    .listRowSeparator(.hidden)
                    .listRowInsets(.init())
                    .padding(.vertical, .spacingXS)
            }
        }
        .listStyle(.plain)
        .contentMargins(.top, 0, for: .scrollContent)
    }
}

#Preview {
    MyBookingsView()
        .environmentObject(BookingViewModel())
}
