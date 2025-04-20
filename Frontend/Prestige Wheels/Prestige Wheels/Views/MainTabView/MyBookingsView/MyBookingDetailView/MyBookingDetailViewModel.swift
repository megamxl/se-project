//
//  MyBookingDetailViewModel.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 20.04.25.
//

import Foundation
import OpenAPIClient

class MyBookingDetailViewModel: ObservableObject {
    let booking: OpenAPIClientAPI.Booking
    
    public init(booking: OpenAPIClientAPI.Booking) {
        self.booking = booking
    }
}
