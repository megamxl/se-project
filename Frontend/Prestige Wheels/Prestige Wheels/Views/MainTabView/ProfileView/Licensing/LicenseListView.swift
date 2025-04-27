//
//  LicenseListView.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 22.04.25.
//

import SwiftUI

// MARK: - Model for a dependency's license info
struct DependencyLicense: Identifiable, Codable {
    // Auto-generated ID, not part of JSON
    let id = UUID()
    let name: String
    let version: String
    let licenseName: String
    let licenseURL: URL

    enum CodingKeys: String, CodingKey {
        case name, version, licenseName, licenseURL
    }
}

// MARK: - View to display all dependency licenses
struct LicenseListView: View {
    @State private var licenses: [DependencyLicense]
    @State private var expandedIDs: Set<UUID> = []

    /// Allows injecting sample data for previews
    init(licenses: [DependencyLicense] = []) {
        _licenses = State(initialValue: licenses)
    }

    var body: some View {
        List(licenses) { dependency in
            DisclosureGroup(
                isExpanded: Binding(
                    get: { expandedIDs.contains(dependency.id) },
                    set: { isOn in
                        if isOn {
                            expandedIDs.insert(dependency.id)
                        } else {
                            expandedIDs.remove(dependency.id)
                        }
                    }
                ),
                content: {
                    VStack(alignment: .leading, spacing: 8) {
                        Link("View \(dependency.licenseName) License", destination: dependency.licenseURL)
                            .font(.subheadline)
                            .bold()
                    }
                    .padding(.vertical, 4)
                },
                label: {
                    HStack {
                        VStack(alignment: .leading) {
                            Text(dependency.name)
                                .font(.headline)
                            Text("v\(dependency.version)")
                                .font(.subheadline)
                                .foregroundColor(.secondary)
                        }
                        Spacer()
                        Text(dependency.licenseName)
                            .font(.footnote)
                            .foregroundColor(.blue)
                    }
                }
            )
            .padding(.vertical, 4)
        }
        .listStyle(InsetGroupedListStyle())
        .navigationTitle("Open Source Licenses")
        .onAppear(perform: loadLicenses)
    }

    // MARK: - JSON Loading
    private func loadLicenses() {
        guard let url = Bundle.main.url(forResource: "licenses", withExtension: "json") else {
            print("❌ licenses.json not found in bundle")
            return
        }

        do {
            let data = try Data(contentsOf: url)
            let decoder = JSONDecoder()
            let items = try decoder.decode([DependencyLicense].self, from: data)
            licenses = items
        } catch {
            print("❌ Failed to decode licenses.json: \(error)")
        }
    }
}
