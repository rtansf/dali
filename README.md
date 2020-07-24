<h1>Dali - Example go/bazel project</h1>

__To run gazelle:__
bazel run //:gazelle

__To run api:__
bazel run //packages/dali:api

__To run test at api level:__
bazel test //packages/dali:go_default_test --test_env=APP_DB_PORT=5432 --test_env=APP_DB_HOST=localhost --test_env=APP_DB_USERNAME=rtan --test_env=APP_DB_PASSWORD= --test_env=APP_DB_NAME=rtan

__To run test at models level:__
bazel test //packages/dali/models:go_default_test --test_env=APP_DB_PORT=5432 --test_env=APP_DB_HOST=localhost --test_env=APP_DB_USERNAME=rtan --test_env=APP_DB_PASSWORD= --test_env=APP_DB_NAME=rtan

__To run all tests:__
bazel test //packages/dali/... --test_env=APP_DB_PORT=5432 --test_env=APP_DB_HOST=localhost --test_env=APP_DB_USERNAME=rtan --test_env=APP_DB_PASSWORD= --test_env=APP_DB_NAME=rtan