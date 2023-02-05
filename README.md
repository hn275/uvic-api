# UVicAPI [WIP]

UVic API is a thin wrapper around all the fetching functions for UVic Banner services, an interface
where students can use to manage registeration, class lookups, etc..

While this package cannot perform any action for the student, it abstracts a lot of the fetching
functions that would be required to get data from UVic Banner.

## Quick note

Special thanks to [Aomi](https://github.com/aomi), he saved me a lot of time and headache with his
infinite Banner wisdom.

## banner.uvic.ca

- Before any fetching can be performed, the term has to be set. This is done by calling the `NewAPI`.

  ```go
  uvic, err := uvicapi.NewAPI("202301") // winter term 2023
  ```

  - This function captured the session cookies set by Banner, it is important to note that
    all subsequent calls to the Banner server will be in the scope of the set term.

- A lot of the response are "paginated" with Banner, each "page" contains a maximum of 500 entries,
  an example:

  ```go
  // GetSection
  res, err := uvic.GetSection(uvicapi.UVicQueryParams{
      Subject:      "CSC",
      CourseNumber: "225",
      Max:          500, // max entries
      Offset:       1, // page num
  })
  ```

## API References

[Implementations](./examples/main.go).

### ToC

[type UVicAPI](#uvicapi)

- [func NewAPI](#newapi)
- [func (c \*UVicAPI) GetTerms](#getterms)
- [func (c \*UVicAPI) GetSection](#getsection)
- [func (c \*UVicAPI) GetAllSections](#getallsections)
- [func (c \*UVicAPI) GetAllCourses](#getallcourses)

[type UVicQueryParams](#uvicqueryparams)

#### UVicAPI

```go
type UVicAPI struct{
  http.Client
  Term string
}
```

#### UVicQueryParams

```go
type UVicQueryParams struct {
  Subject      string
  CourseNumber string
  Offset       int
  Max          int
}
```

- Note that the term is not set here, see `NewAPI`.

#### NewAPI

```go
func NewAPI(term string) (*UVicAPI, error)
```

- Returns an instance of `UVicAPI`.
- This method also capture the cookie set by Banner for the querying term.
  All the subsequent queries will be in the scope of the set term.

#### GetTerms

```go
func (u *UVicAPI) GetTerms() ([]byte, error)
```

- Returns the **json encoding** fetched from banner. See the [json](./examples/data/GetTerms.json).

#### GetSection

```go
func (u *UVicAPI) GetSection(queryParams UVicQueryParams) ([]byte, error)
```

- Returns the **json encoding** fetched from banner. See the [json](./examples/data/GetSection.json).

#### GetCourseDesc (wip)

#### GetAllSections

```go
func (u *UVicAPI) GetAllSections(offset int) ([]byte, error)
```

- Returns the **json encoding** fetched from banner. See the [json](./examples/data/GetAllSections.json).
- Banner "paginates" the response for this, each page can contain up to 500 entries, think of the `offset` param
  like page number.

#### GetAllCourses

```go
func (u *UVicAPI) GetAllCourses() ([]byte, error)
```

- Returns the **json encoding** fetched from banner. See [json](./examples/data/GetAllCourses.json).
