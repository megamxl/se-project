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
    let car: OpenAPIClientAPI.Car
    
    var body: some View {
        HStack(spacing: .spacingL) {
            AsyncImage(
                url: URL(string: car.imageURL ?? ""),
                transaction: Transaction(animation: .default)
            ) { phase in
                if let image = phase.image {
                    image
                        .resizable()
                        .scaledToFill()
                        .clipShape(RoundedRectangle(cornerRadius: 12))
                } else {
                    // While fetching, show placeholder.
                    Image(systemName: "car.fill")
                        .resizable()
                        .scaledToFit()
                }
            }
            .frame(width: 80, height: 80)

            VStack(alignment: .leading, spacing: 0) {
                Text(car.model ?? "")
                    .font(.headline)
                Text(car.brand ?? "")
                    .font(.subheadline)
                    .padding(.bottom, 4)
                
                
                Text("Status: \(booking.status ?? "-")")
                    .font(.system(size: 11))
                    .foregroundColor(Color(uiColor: .secondaryLabel))
                
                Text("\(booking.startDate ?? "") - \(booking.endDate ?? "")")
                    .font(.system(size: 10))
                    .foregroundColor(Color(uiColor: .secondaryLabel))
            }
            .frame(maxWidth: .infinity, alignment: .leading)
            
            VStack(alignment: .trailing, spacing: 4){
                if let paidAmount = booking.paidAmount {
                    Text("\(paidAmount.formatted(.currency(code: booking.currency?.rawValue ?? "-")))")
                        .foregroundStyle(.primary)
                        .font(.callout)
                        .fontDesign(.rounded)
                }
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
    BookingRow(booking: .init(bookingId: "", userId: "", VIN: "", status: ""), car: .init())
}
