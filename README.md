# Purchase Manager Utility

To run the application use any one of the following command
```go
go run cmd/main.go --csv="./test1.csv" --appId=374        // Ouput: 1
go run cmd/main.go --csv="./test2.csv" --appId=374        // Ouput: 3
go run cmd/main.go --csv="./test3.csv" --appId=374        // Ouput: 2
go run cmd/main.go --csv="./sample-small.csv" --appId=374 // Ouput: 190
go run cmd/main.go --csv="./sample-large.csv" --appId=374 // Ouput: 15336
```

Feel free to update the csv file path with desired sample data set