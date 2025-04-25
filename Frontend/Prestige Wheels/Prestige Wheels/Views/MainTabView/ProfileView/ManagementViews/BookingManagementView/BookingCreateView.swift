//
//  BookingCreateView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 22.04.25.
//

import SwiftUI
import OpenAPIClient

struct BookingCreateView: View {
    @Environment(\.dismiss) private var dismiss
    @ObservedObject var viewModel: BookingManagementViewModel
    @State private var VIN = ""
    @State private var startTime = Date()
    @State private var endTime = Date()

    var body: some View {
        NavigationStack {
            Form {
                TextField("Car VIN", text: $VIN)
                DatePicker("Start Time", selection: $startTime)
                DatePicker("End Time", selection: $endTime)
            }
            .navigationTitle("New Booking")
            .toolbar {
                ToolbarItem(placement: .cancellationAction) {
                    Button("Cancel") { dismiss() }
                }
                ToolbarItem(placement: .confirmationAction) {
                    Button("Book") {
                        let req = OpenAPIClientAPI.BookCarRequest(
                            VIN: VIN,
                            currency: .eur,
                            startTime: startTime,
                            endTime: endTime
                        )
                        viewModel.add(request: req) { dismiss() }
                    }
                }
            }
        }
    }
}

struct BookingCreateView_Previews: PreviewProvider {
    static var previews: some View {
        BookingCreateView(viewModel: BookingManagementViewModel())
    }
}
