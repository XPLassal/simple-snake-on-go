# SimpleSnakeGameWithGo 🐍

A simple Snake game that runs directly in your terminal, written entirely in Go.

![preview.gif](preview.gif)

---

## Features

* **Classic Gameplay:** Eat the apple (🍎) to grow your snake. Don't hit the walls or yourself!
* **Cross-Platform:** Runs in any standard terminal on Windows, macOS, and Linux.
* **Dynamic Speed:** Choose to have the game speed up as your score increases for an extra challenge.
* **Pure Go:** Built with no external game engine dependencies, just the standard library and a keyboard listener.

---

## How to Play

You have two options to run the game: download a pre-compiled version (easiest) or build it from the source code.

### Option 1: Download and Run (Recommended)

1.  Go to the [**Releases**](https://github.com/XPLassal/SimpleSnakeGameWithGo/releases) page of this repository.
2.  Download the appropriate file for your operating system:
    * For Windows: `SimpleSnakeGameWithGo.exe`
    * For Linux: `SimpleSnakeGameWithGo`
3.  Open your terminal or command prompt, navigate to the folder where you downloaded the file, and run it.

    **On Windows:**
    ```cmd
    SimpleSnakeGameWithGo.exe
    ```

    **On Linux:**
    (You may need to make the file executable first)
    ```bash
    chmod +x SimpleSnakeGameWithGo
    ./SimpleSnakeGameWithGo
    ```

### Option 2: Build from Source

This requires you to have [Go installed](https://go.dev/doc/install) on your system.

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/XPLassal/SimpleSnakeGameWithGo.git](https://github.com/XPLassal/SimpleSnakeGameWithGo.git)
    ```

2.  **Navigate into the project directory:**
    ```bash
    cd SimpleSnakeGameWithGo
    ```

3.  **Run the game:**
    ```bash
    go run .
    ```
    Go will automatically download the necessary dependencies and start the game.

---

## 🎮 Controls

* **W** - Move Up
* **A** - Move Left
* **S** - Move Down
* **D** - Move Right
* **Q** - Quit Game

---

## 📄 License

This project is licensed under the MIT License. See the `LICENSE` file for more details.