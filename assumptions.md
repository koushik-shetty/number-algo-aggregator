- The challenge has response body with different keys `numbers` in few places and `number` in others  assuming `numbers` as the json key here.  
e.g    
    ```
    { "numbers": [ 1, 2, 3, 5, 8, 13 ] }
    ```
    ```
    <<< { "number": [ 2, 3, 5, 7, 11, 13 ] }
    ```


