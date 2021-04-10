# Patient-Feedback

## Quick Start

1. Place sample .json file in the root directory.
2. Rename `sample-config.yaml` to `config.yaml` and provide the database connection info.
3. Execute `structure.sql` against the postgres database
4. `go run feedback` to prompt for user input
5. `go run feedback analyze` to display some statistics about the collected data

Optional:
6. Execute `sample-data.sql` to load dummy data into the database
