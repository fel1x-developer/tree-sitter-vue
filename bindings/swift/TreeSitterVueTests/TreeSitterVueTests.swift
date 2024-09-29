import XCTest
import SwiftTreeSitter
import TreeSitterVue

final class TreeSitterVueTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_vue())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Vue grammar")
    }
}
