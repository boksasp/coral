# Coral CLI example

This example demonstraits using the Coral module in a CLI application, matching a value on a specific json path.

Match JSON input by specifying `key` and `value`
```bash
> go run main.go -json '{\"person\":{\"id\":1755, \"name\":\"alexander\", \"family\":{\"spouse\": \"eliza\"}}}' -key user.family.spouse -value eliza
🟢 match
```
```bash
> go run main.go -json '{\"person\":{\"id\":1755, \"name\":\"alexander\", \"family\":{\"spouse\": \"eliza\"}}}' -key user.family.spouse -value angelica
🔴 no match
```
```bash
> go run main.go -json '{\"user\":{\"id\":1755,\"name\":\"alexander\",\"family\":{\"spouse\":\"eliza\",\"children\":[{\"id\":1782,\"name\":\"philip\"},{\"id\":1784,\"name\":\"angelica\"},{\"id\":1786,\"name\":\"alexander\"}]}}}' -key user.family.children.id -value 1782
🟢 match
```
```
> go run main.go -help
Usage of main:
  -blankVal
        Set if the filter should match on blank values
  -nullVal
        Set if the filter should match on null values
  -json string
        JSON input string
  -key string
        Key to filter on. Format: node.subnode.subsubnode
  -print
        Print json to output if match is found.
  -value string
        Key should match this value to be filtered to output.
        Use '-blankVal' if matching on empty string values.
```