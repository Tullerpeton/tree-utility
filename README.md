# tree-utility
Tree utility: displays a tree of directories and files

## Run tests
Command:
```bash
go test -v
```

Output:
```bash
=== RUN   TestTreeFull
--- PASS: TestTreeFull (0.00s)
=== RUN   TestTreeDir
--- PASS: TestTreeDir (0.00s)
PASS
ok      coursera/homework/tree     0.127s
```

## Run program
Command:
```bash
go run main.go . -f
```

Output:
```bash
├───main.go (1881b)
├───main_test.go (1318b)
└───testdata
├───project
│   ├───file.txt (19b)
│	└───gopher.png (70372b)
├───static
│	├───css
│	│	└───body.css (28b)
│	├───html
│	│	└───index.html (57b)
│	└───js
│		└───site.js (10b)
├───zline
│	└───empty.txt (empty)
└───zzfile.txt (empty)
go run main.go .
└───testdata
├───project
├───static
│	├───css
│	├───html
│	└───js
└───zline
```