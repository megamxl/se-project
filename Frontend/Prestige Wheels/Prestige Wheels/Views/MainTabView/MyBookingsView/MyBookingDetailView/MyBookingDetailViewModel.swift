//
//  MyBookingDetailViewModel.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 20.04.25.
//

import Foundation
import OpenAPIClient
import _MapKit_SwiftUI
import SwiftUI

class MyBookingDetailViewModel: ObservableObject {
    @AppStorage("selectedMapProvider") private var selectedMapProviderRaw = MapProvider.apple.rawValue
    @Published var showAlert = false
    
    let booking: OpenAPIClientAPI.Booking
    let car: OpenAPIClientAPI.Car
    var alertMessage: String = ""
    
    public init(booking: OpenAPIClientAPI.Booking, car: OpenAPIClientAPI.Car) {
        self.booking = booking
        self.car = car
    }
    
    var position = MapCameraPosition.region(
        MKCoordinateRegion(
            center: CLLocationCoordinate2D(latitude: 48.157975, longitude: 16.381778),
            span: MKCoordinateSpan(latitudeDelta: 0.005, longitudeDelta: 0.005)
        )
    )
    
    var selectedMapProvider: MapProvider {
        MapProvider(rawValue: selectedMapProviderRaw) ?? .apple
    }
    
    func cancelBooking() {
        guard let bookingId = booking.bookingId else { return }
        
        OpenAPIClientAPI.BookingAPI.deleteBooking(bookingId: bookingId) { [weak self] result, error  in
            guard let self = self else { return }
            
            if let error = error {
                alertMessage = "Cancel failed: \(error.localizedDescription)"
            } else {
                alertMessage = "Cancel successfully"
            }
            showAlert = true
            
        }
    }
}
