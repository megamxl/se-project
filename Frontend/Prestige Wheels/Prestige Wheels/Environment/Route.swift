//
//  Route.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 18.04.25.
//

import Foundation
import OpenAPIClient

enum RouteFindCar: Hashable {
    case findCarDetailView(car: OpenAPIClientAPI.CarListInner,
                           currency: OpenAPIClientAPI.Currency,
                           from: Date,
                           to: Date)
}

enum RouteMyBookings: Hashable {
    case bookingDetailView(booking: OpenAPIClientAPI.Booking)
}
