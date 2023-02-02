# UVicAPI [WIP]

UVic API is a thin wrapper around all the fetching functions for UVic Banner services, an interface
where students can use to manage registeration, class lookups, etc..

While this package cannot perform any action for the student, it abstracts a lot of the fetching
functions that would be required to get data from UVic Banner.

## banner.uvic.ca

- Before any fetching can be performed, the term has to be set. This is done by calling the `NewAPI`

  ```go
  u, err := uvicapi.NewAPI("202301") // winter term 2023
  ```

  - This function captured the session cookies set by Banner, it is important to note that
    all subsequent calls to the Banner server will be in the scope of the set term.

- A lot of the response are "paginated" with Banner, each "page" contains a maximum of 500 entries,
  an example:

  ```go
  // GetSection
  res, err := u.GetSection(uvicapi.UVicQueryParams{
      Subject:      "CSC",
      CourseNumber: "225",
      Max:          500, // max entries
      Offset:       1, // page num
  })
  ```

## API

Examples can be found in [example](./examples/)
