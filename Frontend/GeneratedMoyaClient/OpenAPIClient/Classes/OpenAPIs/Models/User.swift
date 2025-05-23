//
// User.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation
#if canImport(AnyCodable)
import AnyCodable
#endif

@available(*, deprecated, renamed: "OpenAPIClientAPI.User")
public typealias User = OpenAPIClientAPI.User

extension OpenAPIClientAPI {

public struct User: Codable, JSONEncodable, Hashable {

    public var id: String?
    public var username: String?
    public var email: String?

    public init(id: String? = nil, username: String? = nil, email: String? = nil) {
        self.id = id
        self.username = username
        self.email = email
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case id
        case username
        case email
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(id, forKey: .id)
        try container.encodeIfPresent(username, forKey: .username)
        try container.encodeIfPresent(email, forKey: .email)
    }
}

}

@available(iOS 13, tvOS 13, watchOS 6, macOS 10.15, *)
extension OpenAPIClientAPI.User: Identifiable {}
