# Sails SFTP Client

## 🚀 Overview

Sails SFTP is a modern, intuitive, and blazingly fast SFTP desktop client crafted for developers, system administrators, and power users. It offers a streamlined experience for managing files over SFTP, simplifying connections to remote servers, transferring files, and managing directories with ease.

Built with the powerful [Wails](https://wails.io/) framework, Sails SFTP leverages the speed and efficiency of Go for its backend logic, ensuring a native feel and cross-platform compatibility. The robust SFTP functionality is powered by the battle-tested `go-sftp` library, guaranteeing secure and efficient file transfers. For the frontend, Sails SFTP utilizes [Vue 3](https://vuejs.org/) with [TypeScript](https://www.typescriptlang.org/), bundled by [Vite](https://vitejs.dev/) for a fast, responsive, and delightful user interface.

Sails SFTP is designed to be lightweight and efficient, making it an ideal choice for both casual users and professionals demanding a reliable and feature-rich SFTP client.

## ✨ Key Features

Sails SFTP is built from the ground up with a focus on performance, stability, and an exceptional user experience. It supports the latest SFTP protocols and offers a suite of powerful features to enhance your workflow:

  * **Multi-Tab Interface:** Effortlessly manage multiple concurrent connections and file transfers in a tabbed environment.
  * **Intuitive Drag-and-Drop:** Seamlessly transfer files and directories between your local machine and remote servers.
  * **Session & Directory Bookmarking:** Save frequently accessed remote directories and connection sessions for quick and easy access.
  * **Integrated Search:** Swiftly locate files within remote directories using the built-in search functionality.
  * **Connection Profiles:** Securely store and manage connection settings for various SFTP servers, streamlining your workflow.
  * **Profile Import/Export:** Easily share your connection profiles with team members or back them up for future use.
  * **Configurable Logging:** Gain insights into application activity and troubleshoot issues with flexible logging options.
  * **Extensible Architecture:** Designed with modularity in mind, making it easier to add new features and integrations.

## 📦 Installation

Sails SFTP is a cross-platform application. You can either download pre-built binaries or build it from source.

### Download Pre-built Binaries (Coming Soon\!)

Pre-built binaries for Windows, macOS, and Linux will be available on the [releases page](https://youtu.be/dQw4w9WgXcQ).

### Build from Source

To build Sails SFTP from its source code, you'll need to have the necessary development tools installed.

#### Prerequisites

  * **Go:** Version 1.20 or newer. Download from [go.dev/dl](https://go.dev/dl/).
  * **Node.js & npm/yarn/bun:** Version 16 or newer. Download from [nodejs.org](https://nodejs.org/).
  * **Wails CLI:** Install the Wails CLI globally:
    ```bash
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    ```
  * **System Dependencies:**
      * **Windows:** No specific additional dependencies.
      * **macOS:** Xcode Command Line Tools (`xcode-select --install`).
      * **Linux (Ubuntu/Debian-based):**
        ```bash
        sudo apt install libgtk-3-dev libwebkit2gtk-4.0-dev
        ```
      * **Linux (Fedora/RHEL-based):**
        ```bash
        sudo dnf install webkit2gtk3-devel gtk3-devel
        ```
      * **Linux (Arch-based):**
        ```bash
        sudo pacman -S webkit2gtk gtk3
        ```

#### Build Steps

1.  **Clone the Repository:**

    ```bash
    git clone https://github.com/YentlHendrickx/sails-sftp.git
    cd sails-sftp
    ```

2.  **Install Frontend Dependencies:**
    Navigate to the `frontend` directory and install the Node.js packages:

    ```bash
    cd frontend
    bun install # or npm/yarn install
    cd ..
    ```

3.  **Build the Application:**
    Use the Wails CLI to build the application. This will compile the Go backend and bundle the Vue.js frontend.

    ```bash
    wails build
    ```

    The executable will be generated in the `build/bin` directory within your project root.

4.  **Run in Development Mode (Optional):**
    For development and testing, you can run the application in hot-reload mode:

    ```bash
    wails dev
    ```

    This will automatically reload the frontend when you make changes to your Vue.js code.

## 🤝 Contributing

We welcome contributions to Sails SFTP\! If you're interested in improving the application, please consider:

  * Reporting bugs or suggesting features via [GitHub Issues](https://github.com/YentlHendrickx/sails-sftp/issues).
  * Submitting Pull Requests with bug fixes, new features, or improvements.
  * Helping with documentation or translations.

Please refer to our [CONTRIBUTING.md](https://github.com/YentlHendrickx/sails-sftp/blob/main/CONTRIBUTING.md) for detailed guidelines on how to contribute.

## 📜 License

Sails SFTP is released under the [MIT License](https://github.com/YentlHendrickx/sails-sftp/blob/main/LICENSE). See the `LICENSE` file for more details.

## 🙏 Acknowledgements

  * [Wails](https://wails.io/) - The incredible framework making cross-platform desktop apps with Go a breeze.
  * [go-sftp](https://github.com/pkg/sftp) - For robust and secure SFTP client functionality.
  * [Vue.js](https://vuejs.org/) - The progressive JavaScript framework for building user interfaces.
  * [Vite](https://vitejs.dev/) - A blazing fast frontend development tool.
  * And all other open-source libraries and communities that make this project possible\!
