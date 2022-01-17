Feature: Balling Score Calculator

Calculator for taking system input of user's game results

Scenario: Calculate player's score of each frame after he/she finished playing 10 frames
    Given system has provided the user restuls of throws of all frames
    And the input data has 10 frames and is valid
    When the input as an array is keyed in as a whole
    Then accumulated scores would be provided for the user

Scenario: Show error if user didn't fully complete the 10 frames of 1 game
    Given system has provided incomplete results of throws
    When the input as an array is keyed in as a whole  
    Then error of "invalid input" would be shown 

Scenario: Show error if somehow system input has more than 10 frames of 1 game
    Given system has provided over-complete results or more than 10 frames of a game
    When the input as an array is keyed in as a whole 
    Then error of "invalid input" would be shown

Scenario: Show error if system input is invalid for some of user's throw results
    Given system has provided invalid inputs although there is no incomplete or overcomplete of game
    When the input as an array is keyed in as a whole
    Then error of "invalid input" would be shown
    Examples:
        |        input                 |           output          |
        | [[1,2], [], [3] ......       |      "invalid input"      |
        | [10,10] ..........           |      "invalid input" .    |