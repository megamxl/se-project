//
//  FindCarDetailViewModel.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 18.04.25.
//

import Foundation
import MapKit
import OpenAPIClient
import _MapKit_SwiftUI
import SwiftUI

class FindCarDetailViewModel: ObservableObject {
    @AppStorage("selectedMapProvider") private var selectedMapProviderRaw = MapProvider.apple.rawValue

    @Published var showAlert = false

    var alertMessage: String = ""
    let car: OpenAPIClientAPI.CarListInner
    let currency: OpenAPIClientAPI.Currency
    let from: Date
    let to: Date
    
    public init(car: OpenAPIClientAPI.CarListInner,
                currency: OpenAPIClientAPI.Currency,
                from: Date,
                to: Date) {
        self.car = car
        self.currency = currency
        self.from = from
        self.to = to
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

    func bookCar() {
        let request = OpenAPIClientAPI.BookCarRequest(VIN: car.VIN, currency: currency, startTime: from, endTime: to)
        
        OpenAPIClientAPI.BookingAPI.bookCar(bookCarRequest: request) { [weak self] (booking, error) in
            guard let self = self else { return }
            
            if let error {
                alertMessage = error.localizedDescription
            } else if booking != nil {
                alertMessage = "Booking was sucessfully!"
            }
            showAlert = true
        }
    }
}
