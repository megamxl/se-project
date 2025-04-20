//
//  BookingRow.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 20.04.25.
//

import SwiftUI
import OpenAPIClient

struct BookingRow: View {
    let booking: OpenAPIClientAPI.Booking
    
    var body: some View {
        HStack(spacing: .spacingL) {
            AsyncImage(
                url: URL(string: ""),
                transaction: Transaction(animation: .default)
            ) { phase in
                if let image = phase.image {
                    image
                        .resizable()
                        .scaledToFill()
                } else {
                    // While fetching, show placeholder.
                    Image(systemName: "car.fill")
                        .resizable()
                        .scaledToFit()
                }
            }
            .frame(width: 80, height: 80)
            VStack(alignment: .leading, spacing: 0) {
                Text("Placeholder")
                    .font(.headline)
                Text("Placeholder")
                    .font(.subheadline)
                
                // Start- und Enddatum
                /*if let startDate = booking.startDate, let endDate = booking.endDate {
                    Text("\(startDate) - \(endDate)")
                        .font(.subheadline)
                        .foregroundColor(.secondary)
                }*/

                Text("Status: \(booking.status ?? "Unbekannt")")
                    .font(.caption)
                    .foregroundColor(.gray)
            }
            
            Spacer()
            VStack(alignment: .trailing, spacing: 4){
                //if let priceOverAll = car.priceOverAll {
                    //Text("\(priceOverAll.formatted(.currency(code: currency.rawValue)))")
                    Text("placeholder")
                        .foregroundStyle(.primary)
                        .font(.callout)
                        .fontDesign(.rounded)
                //}
                //if let pricePerDay = car.pricePerDay {
                    //Text("\(pricePerDay.formatted(.currency(code: currency.rawValue))) /day")
                    Text("placeholder")
                        .foregroundStyle(.gray)
                        .font(.caption2)
                        .fontDesign(.rounded)
                //}
            }
        }
        .hAlign(.leading)
        .padding(.horizontal, .spacingL)
        .padding(.vertical, .spacingXS)
        .background(
            RoundedRectangle(cornerRadius: 12)
                .stroke(Color.gray.opacity(0.4), lineWidth: 1)
        )
    }
}

#Preview(traits: .sizeThatFitsLayout) {
    BookingRow(booking: .init(bookingId: "", userId: "", VIN: "", status: ""))
}
