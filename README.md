# API for conversion  units of measurement

## Methods(CRUD)
- POST
  ```shell
  curl -X POST -H "Content-Type: application/json" -d '{"type":"_type_", "units":"_units_", "value":_value_, "converted_units":"_new_"}' http://localhost:8080 
  ```
  
- GET
  ```shell
  curl -X GET http://localhost:8080
  ```

- PATCH
  ```shell
  curl -X PATCH -H "Content-Type: application/json" -d '{"type":"_type_", "units":"_units_", "value":_value_, "converted_units":"_new_"}' "http://localhost:8080/conversions/_id_"
  ```

- DELETE
  ```shell
  curl -X DELETE http://localhost:8080/conversions/_id_
  ```
***

## Supported units of measurement
- Length
  - millimeter
  - centimetre
  - meter
  - kilometer
  - inch
  - feet
  - yard
  - mile
    
- Temperature
  - celsius
  - fahrenheit
  - kelvin
    
- Time
  - second
  - minute
  - hour
  - day
    
- Weight
  - milligram
  - gram
  - kilogram
  - pound
  - ounce

***

## Technology used
- Golang
- Web-framework Echo 
- PostgreSQL

    
