# The Deeper Love of Go — Study & Exercises

Personal study repository following along with [*The Deeper Love of Go*](https://bitfieldconsulting.com/books/deeper) by John Arundel.

## 📁 Repository Layout

```text
.
├── bookstore/       # Active project: following along with Happy Fun Books
├── chapters/        # Point-in-time snapshots by chapter (ch1, ch2, ch3...)
└── chapterChecks/   # Custom challenges & checks testing concepts learned
```

### 1. `bookstore/`
The primary, living workspace where I follow along with the book's core bookstore project.
* Implements the `books` package containing core book models, formatting, and unit tests (`books_test.go`).
* Contains runnable command-line entry points under `cmd/` (`cmd/list`, `cmd/find`).

### 2. `chapters/`
Saved reference checkpoints preserving the state of the code at the end of each chapter:
* **`ch1/`**: First programs, packages, imports, and basic functions.
* **`ch2/`**: Variables, types, and struct definitions.
* **`ch3/`**: Structuring packages, setting up `go.mod`, and writing tests with the standard `testing` package.
* **`ch4/`**: Slices (`[]Book`), unique IDs, catalog lookups (`GetAllBooks`, `GetBook`), multi-value returns / comma-ok idiom (`(Book, bool)`), `slices.Equal`, and `t.Parallel()`.
* **`ch5/`**: Maps (`map[string]Book`), $O(1)$ lookups with the comma-ok idiom, inserting/mutating maps (`AddBook`), iterators via `maps.Values` & `slices.Collect`, deterministic test sorting with `slices.SortFunc` / `cmp.Compare`, and CLI queries using `os.Args` (`cmd/find`).

### 3. `chapterChecks/`
Targeted practice problems and challenges to verify understanding before moving to the next topic:
* **`label`**: A shelf-label generator (`label.go` and `label_test.go`) testing test-first design, early returns (guard clauses), variable scoping (`:=` vs `=`), and singular/plural formatting.

---

## 🧪 Running Tests

Because each area is managed as its own Go module, tests can be run by navigating to the corresponding directory:

```bash
# Test the active bookstore project
cd bookstore && go test -v

# Test the chapter checks
cd chapterChecks && go test -v

# Test a specific chapter checkpoint
cd chapters/ch5 && go test -v
```
