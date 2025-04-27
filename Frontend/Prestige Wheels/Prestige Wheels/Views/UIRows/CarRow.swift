//
//  CarRow.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 25.03.25.
//

import SwiftUI
import OpenAPIClient

struct CarRow: View {
    let car: OpenAPIClientAPI.CarListInner
    let currency: OpenAPIClientAPI.Currency
    
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
                        .frame(width: 120, height: 80)
                        .clipShape(RoundedRectangle(cornerRadius: 12))
                        .clipped()
                } else {
                    Image(systemName: "car.fill")
                        .resizable()
                        .scaledToFill()
                        .frame(width: 120, height: 80)
                        .clipShape(RoundedRectangle(cornerRadius: 12))
                        .foregroundColor(.gray)
                        .background(Color(.systemGray5))
                        .clipped()
                }
            }

            VStack(alignment: .leading, spacing: 0) {
                Text(car.model ?? "")
                    .font(.headline)
                Text(car.brand ?? "")
                    .font(.subheadline)
            }
            Spacer()
            VStack(alignment: .trailing, spacing: 4){
                if let priceOverAll = car.priceOverAll {
                    Text("\(priceOverAll.formatted(.currency(code: currency.rawValue)))")
                        .foregroundStyle(.primary)
                        .font(.callout)
                        .fontDesign(.rounded)
                }
                if let pricePerDay = car.pricePerDay {
                    Text("\(pricePerDay.formatted(.currency(code: currency.rawValue))) /day")
                        .foregroundStyle(.gray)
                        .font(.caption2)
                        .fontDesign(.rounded)
                }
            }
        }
        .hAlign(.leading)
        .padding(.horizontal, .spacingXS)
        .padding(.vertical, .spacingXS)
        .background(
            RoundedRectangle(cornerRadius: 12)
                .stroke(Color.gray.opacity(0.4), lineWidth: 1)
        )
    }
}

#Preview(traits: .sizeThatFitsLayout) {
    CarRow(car: OpenAPIClientAPI.CarListInner(VIN: "WVWDA7AJ3AW410109", model: "Golf", brand: "Volkswagen", imageURL: "https://groupcms-services-api.porsche-holding.com/dam/images/8576a348442ddf167ea302fac1d6d27ee9fcd4c3/702d47bb1c21f79f7d02f8356bdd9d5e/50a4f40a-012e-426e-9f6b-933d4325702e/crop:SMART/resize:3840:1920/702d47bb1c21f79f7d02f8356bdd9d5e", pricePerDay: 33.00), currency: .eur)
}

#Preview() {
    NavigationStack {
        List(1..<10) { _ in
            CarRow(car: OpenAPIClientAPI.CarListInner(VIN: "WVWDA7AJ3AW410109", model: "Golf", brand: "Volkswagen", imageURL: "https://groupcms-services-api.porsche-holding.com/dam/images/8576a348442ddf167ea302fac1d6d27ee9fcd4c3/702d47bb1c21f79f7d02f8356bdd9d5e/50a4f40a-012e-426e-9f6b-933d4325702e/crop:SMART/resize:3840:1920/702d47bb1c21f79f7d02f8356bdd9d5e", pricePerDay: 33.00), currency: .eur)
                .listRowSeparator(.hidden)
                .listRowInsets(.init())
                .padding(.vertical, .spacingXS)
                .padding(.horizontal)
        }
        .listStyle(.plain)
    }
}
