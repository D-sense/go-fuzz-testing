test:
	go test ./... -v -count=1

# Just as the -run flag takes a regular expression to
# specify which tests to run, the -fuzz flag does the same.
# In this case, we’re using the regular expression “.”,
# which matches all fuzz tests.
fuzz:
	go test -fuzz .
