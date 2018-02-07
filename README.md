# Google Cloud Stackdriver Pipe Logger

`some_process >2&1 | gcloud_pipe_logger "some-process-logs"`

`GOOGLE_APPLICATION_CREDENTIALS` and `GCLOUD_PROJECT_ID` need to be in the ENV.

There are 2 main methods to logging on Google Cloud's Stackdriver; either with
Fluentd or in-app client libraries (eg; Ruby, Python, etc libs). If you only
have basic requirements then it can easily be overkill to use either of those
methods.

So just download one of the prebuilt static binaries from the [releases](https://github.com/tombh/gcloud_pipe_logger/releases) and away you go.
