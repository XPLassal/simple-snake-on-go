# 🐍 Simple Snake Game in Go

A high-performance classic Snake game implementation written in Go using a **Clean Architecture** approach and **Map + Queue** data structure optimization.

![Preview](preview.gif)

---

## ✨ Features

* **High Performance:** Uses `Map` for O(1) collision checks and `Queue` (Slice) for movement logic. The game remains fast regardless of the snake's length!
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
    * **Body Map:** `map[Coordinates]struct{}` for instant collision detection **O(1)**.
    * **Body Queue:** `[]Coordinates` slice to track the order of segments for movement.
* **Architecture:** The project is split into `structs` (Game Logic) and `render` (Presentation Layer).

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
