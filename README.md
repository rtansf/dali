To run gazelle:
bazel run //:gazelle

To run api:
bazel run //packages/dali:api


To run test at api level
bazel test //packages/dali:go_default_test --test_env=APP_DB_PORT=5432 --test_env=APP_DB_HOST=localhost --test_env=APP_DB_USERNAME=rtan --test_env=APP_DB_PASSWORD= --test_env=APP_DB_NAME=rtan

To run test at models level
bazel test //packages/dali/models:go_default_test --test_env=APP_DB_PORT=5432 --test_env=APP_DB_HOST=localhost --test_env=APP_DB_USERNAME=rtan --test_env=APP_DB_PASSWORD= --test_env=APP_DB_NAME=rtan

To run all tests:
bazel test //packages/dali/... --test_env=APP_DB_PORT=5432 --test_env=APP_DB_HOST=localhost --test_env=APP_DB_USERNAME=rtan --test_env=APP_DB_PASSWORD= --test_env=APP_DB_NAME=rtan