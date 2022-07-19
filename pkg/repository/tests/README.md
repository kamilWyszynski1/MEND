## Integration tests

This directory could contain integration tests for all UserRepository implementations.

If I had to implement this I would use package Suite to do so. It is handy to create
integration tests like this because of methods like `BeforeTest` or `SetupTestSuite`.

Additionally I would use `docker` go client for setuping db containers for those tests.