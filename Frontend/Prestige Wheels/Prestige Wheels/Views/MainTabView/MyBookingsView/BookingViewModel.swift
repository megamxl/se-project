//
//  BookingViewModel.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 10.04.25.
//

import Foundation
import OpenAPIClient
import OSLog

class BookingViewModel: ObservableObject {
    @Published var bookings: [OpenAPIClientAPI.Booking]?
    
    func getBookingsForUser() {
        OpenAPIClientAPI.BookingAPI.getBookings(apiResponseQueue: DispatchQueue.main) { [weak self] (bookings, error) in
            guard let self = self else { return }
            
            if let error = error {
                Logger.backgroundProcessing.error("\(error.localizedDescription)")
            } else if let bookings = bookings {
                self.bookings = bookings
            }
        }
    }
    
    func bookCar(request: OpenAPIClientAPI.BookCarRequest) {
        OpenAPIClientAPI.BookingAPI.bookCar(bookCarRequest: request, apiResponseQueue: .main) { [weak self] (_, error) in
            guard let self = self else { return }

            if let error = error {
                Logger.backgroundProcessing.error("Booking failed: \(error.localizedDescription)")
            } else {
                Logger.backgroundProcessing.info("Booking successful (no error returned)")
            }
        }
    }
    
    func deleteBooking() {

    }
    
    func getBookingById() {
        
    }
    
    func updateBooking() {
        
    }
}
