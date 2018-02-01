Feature: processes images
  As a user I would like to be sure
  that the connection between my 
  function and the machinebox api 
  is correct.

Scenario: succesfully outputs image
    Given I have a valid image
    When I call my function
    Then I expect it to return a valid image
