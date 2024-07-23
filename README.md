# MP3 Player with System Shutdown

Um to troll your friends? But I am not responsible this is not a virus just a little program to sleep your computer after the outro music ends.

## Features

- Embeds an MP3 file as a byte slice.
- Decodes and plays the MP3 file, configuring audio settings.
- Shuts down the computer when the music finishes.
- Optionally shuts down the computer during the last 5 seconds of playback.

## Requirements

- Go 1.16 or newer.
- `faiface/beep` and `faiface/beep/mp3` packages.

## Installation

1. **Install Go and Dependencies**: Ensure you have Go and the required packages installed. To install the necessary packages, run:

    ```sh
    go get github.com/faiface/beep
    go get github.com/faiface/beep/mp3
    ```

2. **Add MP3 File**: Place an MP3 file named `music.mp3` in the project directory. This file will be played by the program.

3. **Build the Code**: In the project directory, compile the Go code by running:

    ```sh
    go build -o mp3player main.go
    ```

4. **Run the Program**: Execute the compiled program:

    ```sh
    ./mp3player
    ```

## How It Works

- The program embeds the `music.mp3` file as a byte slice and reads it using the `faiface/beep` library.
- It initializes the audio output and starts playing the MP3 file.
- When the music finishes or during the last 5 seconds of playback, the program executes a command to shut down the computer.
- A timer is used to perform the shutdown operation during the last 5 seconds of playback.

## Platform Compatibility

This code has been tested on Windows and uses Windows-specific commands for shutdown. You may need to adjust the shutdown command for different operating systems.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

