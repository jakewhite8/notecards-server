package controller

import (
  "context"
  "fmt"
  "io"
  "os"
  "time"
  "text/tabwriter"

  "github.com/apache/arrow/go/v13/arrow"
  "github.com/InfluxCommunity/influxdb3-go/influxdb3"
)

func Query() error {

   // INFLUX_TOKEN is an environment variable you created
   // for your API read token.
  token := os.Getenv("INFLUXDB_TOKEN")
  database := "get-started"

  // Instantiate the client.
  client, err := influxdb3.New(influxdb3.ClientConfig{
    Host:     "https://us-east-1-1.aws.cloud2.influxdata.com/",
    Token:    token,
    Database: database,
  })

  // Close the client when the function returns.
  defer func(client *influxdb3.Client) {
    err := client.Close()
    if err != nil {
      panic(err)
    }
  }(client)

  // Define the query.
  query := `SELECT *
    FROM home`

  // `SELECT *
  // FROM home
  // WHERE time >= '2024-01-06T08:00:00Z'
  // AND time <= '2024-01-06T20:00:00Z';`

  // Execute the query.
  iterator, err := client.Query(context.Background(), query)

  if err != nil {
    panic(err)
  }

  w := tabwriter.NewWriter(io.Discard, 4, 4, 1, ' ', 0)
  w.Init(os.Stdout, 0, 8, 0, '\t', 0)
  fmt.Fprintln(w, "time\troom\ttemp\thum\tco")

  // Iterate over rows and prints column values in table format.
  for iterator.Next() {
    row := iterator.Value()
    // Use Go arrow and time packages to format unix timestamp
    // as a time with timezone layout (RFC3339).
    time := (row["time"].(arrow.Timestamp)).
      ToTime(arrow.TimeUnit(arrow.Nanosecond)).
      Format(time.RFC3339)
    fmt.Fprintf(w, "%s\t%s\t%d\t%.1f\t%.1f\n",
      time, row["room"], row["co"], row["hum"], row["temp"])
  }

  w.Flush()
  return nil
}
