//
//  FindCarDetailViewModel.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 18.04.25.
//

import Foundation
import OpenAPIClient

class FindCarDetailViewModel: ObservableObject {
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
    
    func bookCar() {
        let request = OpenAPIClientAPI.BookCarRequest(VIN: car.VIN, currency: currency, startTime: from, endTime: to)
        
        OpenAPIClientAPI.BookingAPI.bookCar(bookCarRequest: request) { [weak self] (booking, error) in
            guard let self = self else { return }
            
            if let error {
                alertMessage = error.localizedDescription
            } else if let booking {
                alertMessage = "Booking was sucessfully!"
            }
            showAlert = true
        }
    }
}
