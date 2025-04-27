//
//  UserRow.swift
//  Prestige Wheels
//
//  Created by Heinz Schweitzer on 27.04.25.
//

import SwiftUI
import OpenAPIClient

struct UserRow: View {
    let user: OpenAPIClientAPI.User

    var body: some View {
        HStack(spacing: .spacingL) {
            Image(systemName: "person.crop.circle.fill")
                .resizable()
                .scaledToFit()
                .frame(width: 30, height: 30)
                .foregroundColor(.gray)
            
            VStack(alignment: .leading, spacing: 4) {
                Text(user.username ?? "Unknown User")
                    .font(.headline)
                Text(user.email ?? "No Email")
                    .font(.subheadline)
                    .foregroundColor(.secondary)
            }
            .frame(maxWidth: .infinity, alignment: .leading)
        }
    }
}

#Preview {
    UserRow(user: OpenAPIClientAPI.User(username: "michael", email: "michael@example.com"))
}
