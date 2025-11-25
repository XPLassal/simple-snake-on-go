# 🐍 Simple Snake Game in Go

A high-performance classic Snake game implementation written in Go using a **Clean Architecture** approach and a unique **Linked List via Map** optimization.

![Preview](preview.gif)

---

## ✨ Features

* **High Performance (True O(1)):** The game engine uses a custom **Linked List over Hash Map** data structure. Movement and collision detection are instant regardless of the snake's size—no array shifting or iteration required!
* **Cross-Platform:** Runs natively on **Windows**, **Linux**, and **macOS** (Intel & Apple Silicon).
* **Dynamic Speed:** Optional "Hard Mode" where the game speeds up as you progress.
* **Clean Code:** Written in pure Go with a modular structure (Logic separated from Rendering).

---

## 🎮 How to Play

You don't need to install Go to play. Just download the binary for your system!

### Option 1: Download & Run (Recommended)

1.  Go to the [**Releases**](https://github.com/XPLassal/simple-snake-on-go/releases) page.
2.  Download the file for your OS:
    * 🪟 **Windows:** `snake-windows-amd64.exe`
    * 🐧 **Linux:** `snake-linux-amd64`
    * 🍎 **macOS (M1/M2/M3):** `snake-macos-arm64`
    * 🍎 **macOS (Intel):** `snake-macos-intel`

3.  **Run it:**

    **On Windows:**
    Double-click the `.exe` file or run it via terminal.

    **On Linux / macOS:**
    Open terminal in the folder with the file and run:
    ```bash
    # Make it executable (only needed once)
    chmod +x snake-linux-amd64  # (replace with your file name)

    # Run
    ./snake-linux-amd64
    ```

---

## 🕹 Controls

| Key | Action |
| :---: | :--- |
| **W** | Move Up ⬆️ |
| **S** | Move Down ⬇️ |
| **A** | Move Left ⬅️ |
| **D** | Move Right ➡️ |
| **Q** | Quit Game |

---

## 🏗 Technical Details

For developers interested in the architecture:

* **Rendering:** Console-based rendering using ANSI escape codes for colors.
* **Data Structures:**
    * **Body Storage:** `map[Coordinates]Coordinates`.
    * **Logic:** This implements a **Singly Linked List** on top of a Hash Map.
        * **Key:** Current segment coordinate.
        * **Value:** Coordinate of the *next* segment (towards the head).
    * **Performance:** This structure ensures **O(1)** time complexity for:
        * **Collision checks:** Instant map lookup.
        * **Movement:** Simply updating the pointer of the old Head and removing the old Tail key. No memory shifting (`copy`/`append`) is performed, making it extremely efficient for very long snakes.
* **Architecture:** The project follows strict separation of concerns:
    * `structs`: Contains all game entities (Snake, Apple) and rules.
    * `render`: Handles drawing to the terminal.

---

## 🛠 Build from Source

If you want to modify the code or build it yourself:

1.  **Install Go** (1.23 or newer).
2.  **Clone the repository:**
    ```bash
    git clone [https://github.com/XPLassal/simple-snake-on-go.git](https://github.com/XPLassal/simple-snake-on-go.git)
    cd simple-snake-on-go
    ```
3.  **Run locally:**
    ```bash
    go run .
    ```
4.  **Build binaries (using script):**
    ```bash
    # Only on Linux/macOS/WSL
    ./build.sh
    ```

---

## 📄 License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
