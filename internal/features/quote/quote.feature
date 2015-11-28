Feature: Stock quote example
  This confirms we can obtain a quote via a git on /quote/{symbol}

  Scenario: Basic quote
    Given a symbbol
    And a request for a quote for that symbol
    Then the current price is returned