# ASCII Art Align (Go)

## Overview

This program prints ASCII art from a string and allows you to align the output in the terminal.

---

## Usage

```bash
go run . [OPTION] [STRING] [BANNER]
```

---

## Examples & Expected Output

### 1. Center Alignment

```bash
go run . --align=center "hello" standard
```

**Output (centered):**

 ```
                                                        _                _    _
                                                        | |              | |  | |
                                                        | |__      ___   | |  | |    ___
                                                        |  _ \    / _ \  | |  | |   / _ \
                                                        | | | |  |  __/  | |  | |  | (_) |
                                                        |_| |_|   \___|  |_|  |_|   \___/
 ```

---

### 2. Left Alignment

```bash
go run . --align=left "Hello" standard
```

**Output (left):**

```
 _    _           _    _                 
| |  | |         | |  | |              
| |__| |   ___   | |  | |    ___       
|  __  |  / _ \  | |  | |   / _ \      
| |  | | |  __/  | |  | |  | (_) |        
|_|  |_|  \___|  |_|  |_|   \___/         
```

---

### 3. Right Alignment

```bash
go run . --align=right "hello" shadow
```

**Output (right):**

```
                                                                                                                            _|                _| _|
                                                                                                                            _|_|_|     _|_|   _| _|   _|_|
                                                                                                                            _|    _| _|_|_|_| _| _| _|    _|
                                                                                                                            _|    _| _|       _| _| _|    _|
                                                                                                                            _|    _|   _|_|_| _| _|   _|_|
```

---


## Error Example

### Invalid Flag Format

```bash
go run . --align right "hello" standard
```

**Output:**

```
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
```

---

### Missing Arguments

```bash
go run .
```

**Output:**

```
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
```

---

## Note

* The program adjusts output based on terminal width

```bash
go run . "hello"
```
