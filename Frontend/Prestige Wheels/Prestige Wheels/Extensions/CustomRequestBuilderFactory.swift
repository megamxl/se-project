//
//  CustomRequestBuilderFactory.swift
//  Prestige Wheels
//
//  Created by Michael Luegmayer on 06.04.25.
//

import Foundation
import OpenAPIClient

public class CustomRequestBuilderFactory: RequestBuilderFactory {
    private let session: URLSession

    public init(session: URLSession) {
        self.session = session
    }

    public func getBuilder<T: Decodable>() -> RequestBuilder<T>.Type {
        return URLSessionDecodableRequestBuilder<T>.self
    }

    public func getNonDecodableBuilder<T>() -> RequestBuilder<T>.Type {
        return URLSessionRequestBuilder<T>.self
    }
}
