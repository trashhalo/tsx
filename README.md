# TSX

Experiment at creating a data structure for terminal component hierachy. Inspired by JSX/h

## Example

A functional component

```golang
func hello() tsx.T {
	return tsx.T{
        	// name of the "tag" for styling purposes
        	Name: "hello",
        	// text to render in that tag
		Text: "Hello ",
	}
}
```

A more complete example can be seen in <./examples/hello/hello.go>
