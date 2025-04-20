//
//  MyBookingDetailView.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 20.04.25.
//

import SwiftUI

struct MyBookingDetailView: View {
    @EnvironmentObject var route: RouteObject
    @ObservedObject var viewModel: MyBookingDetailViewModel
    
    var body: some View {
        Text("BookingDetailView")
    }
}

#Preview {
    MyBookingDetailView(viewModel: .init(booking: .init()))
}
