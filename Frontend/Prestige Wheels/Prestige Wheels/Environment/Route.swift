//
//  Route.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 18.04.25.
//

import Foundation
import OpenAPIClient

enum Route: Hashable {
    case findCarDetailView(car: OpenAPIClientAPI.CarListInner)
}
