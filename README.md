# Test Driven Development

### Usage

```go
go test

go test -cover
go test -coverprofile=result.out
go tool cover -html=result.out -o result.html

```

### Good practices

Do this...

1. Follow TDD steps: write test, write code, improve test, improve code and iter!
2. Use interfaces.
3. Use Table Driven Tests (TDT).
4. Use go native libraries.
5. Build tests of differents packages.
6. Do not change the behaviour of your tests (do not change your tests) when refactor the code.
7. Do not forget integration tests.

### Antipatterns

Do not do this...

1. Giving confusing information when a test fails. Solution: be explicit and give context when tests fails.
2. If your code is difficult to test, your code is difficult to use. Solution: make easy and clean code.
3. Not following the typical structure of a test. Solution: write setup and preconditions, then actions, follow with assertions and finally clean and tier down. Do not mix this steps!
4. Prefer DRY (dont repeat yourself) to DAMP (descriptive and meaningful phrases). Solution: repeat code if you need. Make a very smart test could make it complicated to understand.
5. Include error cases in Table Driven Tests. Solution: create tests cases for testing errors.
6. Violate encapsulation. Solution: do not make public methods, attributes, constants, etc only cause you need it on tests. Build a simple interfaces and methods.
