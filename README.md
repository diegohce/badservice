# badservice
Helps you test how resilient you are generating bad HTTP responses.

## Status Code
`/badservice/status/:code`
Will response a `code` HTTP status code.

e.g.:
`/badservice/status/503`
Will response a `503` HTTP status code.

## Connection Delay
`/badservice/delay/:delay`
Will response a 200 status code after `delay` milliseconds.

e.g.:
`/badservice/delay/10000`
Will wait 10 seconds.

## Connection Drop

`/badservice/drop`
Will accept your request, but also will drop the TCP connection.

