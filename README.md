# Green Light



##  migration 
Create
`migrate create -seq -ext=.sql -dir=./migrations create_movies_table`

Execute
`migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up`


- Third-party staticcheck tool to carry out some additional static analysis checks.
`go install honnef.co/go/tools/cmd/staticcheck@latest`



## Setup



TODO
[ ] Creating an end-to-end test for the GET /v1/healthcheck endpoint to verify that the headers and response body are what you expect.
[ ] Creating a unit-test for the rateLimit() middleware to confirm that it sends a 429 Too Many Requests response after a certain number of requests.
[ ] Creating an end-to-end integration test, using a test database instance, which confirms that the authenticate() and requirePermission() middleware work together correctly to allow or disallow access to specific endpoints.




DD/MM/YYYY
--------
- 01/10/2024 - Started 
- 01/10/2024 - Finished Chapter #2: Getting started
- 05/10/2024 - Finished till page 52 of `Chapter #3 Sending JSON responses` 
- 06/10/2024 - Finished till page 67 of `Chapter #3 Sending JSON responses` 
- 07/10/2024 - `Missing`
- 08/10/2024 - Finished `Chapter #3 Sending JSON responses` (page no 73)
- 09/10/2024 - Finished till page 100 of `Chapter #4 Parsing JSON Request`
- 10/10/2024 - Finished `Chapter #4 parsing JSON request`
- 11/10/2024 - Finished `Chapter #5 Database setup and configuration`
- 12/10/2024 - Finished `Chapter #6 SQL Migration`
- 12/10/2024 - Finished till page 159 `Chapter #7 SQL Migration: Update Movie`
- 13/10/2024 - `Missing`
- 14/10/2024 - `Missing`
- 15/10/2024 - Finished `Chapter 7 CRUD Operations` till page 163
- 15/10/2024 - Finished till page 177 in  `Chapter #8 Advance CRUD Operations`
- 16/10/2024 - Finished `Chapter #8 Advance CRUD Operations`
- 17/10/2024 - Finished till page 199 in `Chapter #9 Filtering, Sorting, and Pagination` 
- 18/10/2024 - Finished till page 220 in `Chapter #9 Filtering, Sorting, and Pagination` 
- 19/10/2024 - Finished `Chapter #9 Filtering, Sorting, and Pagination` page 233 
- 20/10/2024 - Finished `Chapter #10 Structure logging` page no. 234 
- 21/10/2024 - `Missing`
- 22/10/2024 - `Missing`
- 23/10/2024 - Finished  `Chapter #11 Rate Limiting` page no. 261
- 24/10/2024 - `Missing`
- 25/10/2024 - `Missing`
- 26/10/2024 - `Missing`
- 27/10/2024 - `Missing`
- 28/10/2024 - `Missing`
- 29/10/2024 - `Missing`
- 30/10/2024 - `Missing`
- 31/10/2024 - `Missing`
- 01/11/2024 - `Completed Chapter #12 Graceful Shutdown` page no. 274
- 02/11/2024 - `Completed Chapter #13 User Model Setup and Registration` finished till page no. 287
- 03/11/2024 - Finished `Completed Chapter #13 User Model Setup and Registration`  page no. 292
- 04/11/2024 - `Missing`
- 05/11/2024 - `Missing`
- 06/11/2024 - `Missing`
- 07/11/2024 - `Missing`
- 08/11/2024 - `Missing`
- 09/11/2024 - `Missing`
- 10/11/2024 - Finished `Chapter #14 Sending Emails` page no 321
- 10/11/2024 - Finished `Chapter #14 Sending Emails` page no 321
- 10/12/2024 - Completed till page 327 of `Chapter #15 User Activation` 
- 10/13/2024 - Completed till page 336 of `Chapter #15 User Activation` 
- 10/13/2024 - `Missing`
- 10/14/2024 - `Missing`
- 10/16/2024 - Completed till page 340 of `Chapter #15 User Activation` 
- 10/17/2024 - Finished `Chapter #15 User Activation` page no 346
- 10/18/2024 - `Missing`
- 10/19/2024 - `Missing`
- 10/20/2024 - Finished `Chapter #16 Authentication` page no 369
- 10/28/2024 - Complete till page 379 of  `Chapter #17 Permission-based Authorization` 
- 10/29/2024 - Complete till page 389 of  `Chapter #17 Permission-based Authorization` 
- 10/30/2024 - `Finished Chapter #17 Permission-based Authorization` 
- 12/03/2024 - `Finished Chapter #18 Cross origin request` 
- 12/05/2024 - `Finished Chapter #20 Metrics`
- 12/19/2024 - `Finished Chapter #21  Building, Versioning and Quality Control`
- Next: `Chapter #22 Deployment and Hosting` page no. 500