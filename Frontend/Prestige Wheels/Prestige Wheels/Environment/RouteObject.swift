//
//  RouteObject.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 18.04.25.
//

import Foundation

class RouteObject: ObservableObject {
    @Published var pathFindCar: [RouteFindCar] = []
    @Published var pathMyBookings: [RouteMyBookings] = []
    @Published var pathAdmin: [RouteAdmin] = []
}
