### Progress Util
A simple (and very rough) utility to process JUnit progress files

There are three options:

    progress hang <path to progress file>

Will show which tests have not completed yet. Helpful when diagnosing test hangs.

    progress times <path to progress file>

Will show how long each test class took and how many tests were in the class.

    progress unique <path to progress file>

Will give you the list of classes in the order run. This is helpful if you're trying to figure out if prior tests have dirtied the environment and you're wanting to run some of them as part of a suite.

### Building

    go build progress


