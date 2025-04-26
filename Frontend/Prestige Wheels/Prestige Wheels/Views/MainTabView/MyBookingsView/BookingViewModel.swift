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
    @Published var bookings: [OpenAPIClientAPI.Booking] = []
    @Published var carsByBookingId: [String: OpenAPIClientAPI.Car] = [:]
        
    func getBookingsForUser() {
        OpenAPIClientAPI.BookingAPI.getBookings(apiResponseQueue: DispatchQueue.main) { [weak self] (bookings, error) in
            guard let self = self else { return }
            
            if let error {
                Logger.backgroundProcessing.error("\(error.localizedDescription)")
            } else if let bookings = bookings {
                self.bookings = bookings
                for booking in bookings {
                    if let vin = booking.VIN, let bookingId = booking.bookingId {
                        self.getCarByVin(vin: vin, bookingId: bookingId)
                    }
                }
            }
        }
    }
    
    func getCarByVin(vin: String, bookingId: String) {
        OpenAPIClientAPI.CarsAPI.getCarByVin(VIN: vin, apiResponseQueue: DispatchQueue.main) { [weak self] (car, error) in
            guard let self = self else { return }
            
            if let error {
                Logger.backgroundProcessing.error("\(error.localizedDescription)")
            } else if let car = car {
                self.carsByBookingId[bookingId] = car
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
