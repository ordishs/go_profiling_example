# Go profiling example

This is a very simple project to help demonstrate how to profile an application in Go.

1. Run the application

2. While it is running go to http://localhost:9000/debug/pprof

3. Click on the "profile" link and wait for the profile to complete (and download)

4. Run ```go tool pprof -http localhost:8000 ~/Downloads/profile ```