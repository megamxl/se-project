//
//  Bundle+Extensions.swift
//  CampusPlus
//
//  Created by Michael Luegmayer on 10.02.2025.
//

import Foundation

extension Bundle {
    /// A computed property that returns the release version number of the bundle.
    ///
    /// This property retrieves the value associated with the `CFBundleShortVersionString` key
    /// from the bundle's `infoDictionary`. This is typically used to represent the user-facing
    /// version of an application.
    ///
    /// - Returns: A `String` representing the release version number if available; otherwise, `nil`.
    ///
    /// ### Usage:
    /// ```swift
    /// Text("\(Bundle.main.releaseVersionNumber ?? ""))
    /// ```
    var releaseVersionNumber: String? {
        return infoDictionary?["CFBundleShortVersionString"] as? String
    }

    /// A computed property that returns the build version number of the bundle.
    ///
    /// This property retrieves the value associated with the `CFBundleVersion` key from the
    /// bundle's `infoDictionary`. This value is generally used for internal version tracking.
    ///
    /// - Returns: A `String` representing the build version number if available; otherwise, `nil`.
    ///
    /// ### Usage:
    /// ```swift
    /// Text("\(Bundle.main.buildVersionNumber ?? ""))
    /// ```
    var buildVersionNumber: String? {
        return infoDictionary?["CFBundleVersion"] as? String
    }
}
