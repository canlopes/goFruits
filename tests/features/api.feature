Feature: API testing
  Scenario: test addFruit api
    Given I don't have fruits
    When I add a fruit with "POST" request to the endpoint "/fruit" with the following payload '{"fruit": "banana", "calories": 200, "price": 2}'
    Then the HTTP-response code should be "201"
    And the response content should be '{"fruit": "banana", "calories": 200, "price": 2.00 }'
    Then I retrieve fruit info with "GET" request to the endpoint "/fruit/banana"
    Then the HTTP-response code should be "200"

  Scenario: test updateFruit api
    Given I don't have fruits
    When I add a fruit with "POST" request to the endpoint "/fruit" with the following payload '{"fruit": "banana", "calories": 200, "price": 2.4}'
    Then the HTTP-response code should be "201"
    And I update a fruit with "PUT" request to the endpoint "/fruit/banana" with the following payload '{"fruit": "banana2", "calories": 300, "price": 2.3}'
    Then the HTTP-response code should be "202"
    Then I retrieve fruit info with "GET" request to the endpoint "/fruit/banana2"
    And the response content should be '{"fruit": "banana2", "calories": 300, "price": 2.30 }'

  Scenario: test deleteFruit api
    Given I don't have fruits
    When I add a fruit with "POST" request to the endpoint "/fruit" with the following payload '{"fruit": "banana", "calories": 200, "price": 2.4}'
    Then the HTTP-response code should be "201"
    And I delete a fruit with "DELETE" request to the endpoint "/fruit/banana"
    Then the HTTP-response code should be "200"
    Then I retrieve fruit info with "GET" request to the endpoint "/fruit/banana2"
    And the response content should be '{"message":"fruit not found banana2"}'
