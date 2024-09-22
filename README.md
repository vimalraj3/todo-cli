# CLI To Do App with GoLang 📝

**CLI To Do App** is a simple and efficient command-line tool built with GoLang for managing your tasks directly from the terminal. Whether you're organizing your daily chores or tracking project milestones, this tool provides a straightforward way to add, update, complete, and delete your to-dos with ease.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Add a New Todo](#add-a-new-todo)
  - [Toggle Done Status](#toggle-done-status)
  - [List All Todos](#list-all-todos)
  - [Update a Todo](#update-a-todo)
  - [Delete a Todo](#delete-a-todo)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

## Features

- **Add Todos**: Quickly add new tasks with descriptions.
- **Toggle Completion Status**: Mark tasks as done or not done.
- **List Todos**: View all your tasks with their current statuses.
- **Update Todos**: Modify existing tasks as needed.
- **Delete Todos**: Remove tasks that are no longer relevant.
- **Persistent Storage**: All tasks are saved locally for easy access.

## Installation

### Prerequisites

- **GoLang**: Ensure you have Go installed on your machine. Download it from [golang.org](https://golang.org/dl/).

### Steps

1. **Clone the Repository**

    ```bash
    git clone https://github.com/yourusername/cli-todo-app.git
    cd cli-todo-app
    ```

2. **Build the Application**

    ```bash
    go build -o todo
    ```

3. **Move to Your PATH**

    ```bash
    mv todo /usr/local/bin/
    ```

    _Alternatively, you can use Go's install command:_

    ```bash
    go install
    ```

    Ensure that your `$GOPATH/bin` is added to your system's `PATH` to use the `todo` command globally.

## Usage

Once installed, you can start managing your todos using the following commands:

### Add a New Todo

Add a new task with a description.

```bash
todo --add "Buy a coke"
```

**Example:**

```bash
todo --add "Finish writing the report"
```

### Toggle Done Status

Mark a task as completed or revert it back to pending using its ID.

```bash
todo --done <id>
```

**Example:**

```bash
todo --done 3
```

### List All Todos

Display all your tasks with their statuses.

```bash
todo --list
```

**Example Output:**

```
ID | Description           | Status
---|-----------------------|--------
1  | Buy a coke            | Pending
2  | Finish writing report | Done
3  | Call the plumber      | Pending
```

### Update a Todo

Modify the description of an existing task using its ID.

```bash
todo --update <id> "New description"
```

**Example:**

```bash
todo --update 2 "Complete the annual report"
```

### Delete a Todo

Remove a task from your list using its ID.

```bash
todo --delete <id>
```

**Example:**

```bash
todo --delete 1
```

## Contributing

Contributions are welcome! Whether it's fixing bugs, adding new features, or improving documentation, your help is appreciated.

1. **Fork the Repository**
2. **Create a Feature Branch**

    ```bash
    git checkout -b feature/YourFeature
    ```

3. **Commit Your Changes**

    ```bash
    git commit -m "Add your message"
    ```

4. **Push to the Branch**

    ```bash
    git push origin feature/YourFeature
    ```

5. **Open a Pull Request**

Please make sure to update tests as appropriate and adhere to the [code of conduct](LICENSE).

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

- Inspired by various CLI to-do applications.
- Built with ❤️ using [GoLang](https://golang.org/). 
